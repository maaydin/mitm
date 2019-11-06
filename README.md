# mitm - mitm is a SSL-capable man-in-the-middle proxy for use with golang net/http.

It is heavily inspired by the mitmproxy project (https://mitmproxy.org/).

**Install**

	go get github.com/maaydin/mitm

**Docs**

http://godoc.org/github.com/maaydin/mitm

**Create CA Root Cert**
```
mkdir ~/.mitm
cd ~/.mitm
case `uname -s` in                                                                                                            
    Linux*)     sslConfig=/etc/ssl/openssl.cnf;;
    Darwin*)    sslConfig=/System/Library/OpenSSL/openssl.cnf;;
esac

openssl ecparam -genkey -name prime256v1 -noout -out ca-key.pem

openssl req \
    -new \
    -x509 \
    -key ca-key.pem \
    -new \
    -out ca-cert.pem \
    -subj /CN=$HOST \
    -reqexts SAN \
    -extensions SAN \
    -extensions v3_ca \
    -config <(cat $sslConfig \
        <(printf '[SAN]\nsubjectAltName=DNS:localhost\n\n[v3_ca]\nbasicConstraints = critical,CA:TRUE\nsubjectKeyIdentifier=hash\nauthorityKeyIdentifier=keyid:always,issuer:always')) \
    -sha256 \
    -days 3650
 ```

**Contributors**

* Keith Rarick (@kr)
* Blake Mizerany (@bmizerany)
* Mehmet Ali Aydin (@maaydin)
