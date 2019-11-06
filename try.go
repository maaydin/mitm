// +build ignore

package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"path"
	"fmt"
	"time"

	"../mitm"
)

var (
	dir      = path.Join(os.Getenv("HOME"), ".mitm")
	keyFile  = path.Join(dir, "ca-key.pem")
	certFile = path.Join(dir, "ca-cert.pem")
)

func main() {
	ca, err := loadCA()
	if err != nil {
		log.Fatal(err)
	}
	p := &mitm.Proxy{
		CA: &ca,
		TLSServerConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			//CipherSuites: []uint16{tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA},
		},
		Analyze: analyze,
	}
	
	log.Println("Proxy starting on: http://localhost:3128 https://localhost:3128")
	
	if err := http.ListenAndServe(":3128", p); err != nil {
		log.Fatal("Proxy failed to start up: ", err)
	}
}

func loadCA() (cert tls.Certificate, err error) {
	// TODO(kr): check file permissions
	cert, err = tls.LoadX509KeyPair(certFile, keyFile)
	if os.IsNotExist(err) {
		log.Fatal("CA Certificate not found on path: ", dir)
	}
	if err == nil {
		cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	}
	return
}

func analyze(rr mitm.RequestStat) {
	fmt.Printf("Method: %s\n", rr.Method)
	fmt.Printf("Scheme: %s\n", rr.Scheme)
	fmt.Printf("Host: %s\n", rr.Host)
	fmt.Printf("Path: %s\n", rr.Path)
	fmt.Printf("StatusCode: %d\n", rr.StatusCode)
	fmt.Println("Start time: " + rr.StartTime.Format(time.RFC3339Nano))
	fmt.Println("End time: " + rr.EndTime.Format(time.RFC3339Nano))
	fmt.Printf("Elapsed Time: %dms\n", rr.ElapsedTime)
}