package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", baseHandler)

	server := &http.Server{
		Addr:         ":5000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		TLSConfig:    tlsConfig(),
	}

	log.Println("Server listens to localhost:5000")

	issue := server.ListenAndServeTLS("", "")
	if issue != nil {
		log.Fatal(issue.Error())
	}
}

func tlsConfig() *tls.Config {
	// openssl genrsa -out certificates/private-key.pem 2048
	// openssl req -nodes -new -x509 -sha256 -days 365 -config certificates/configuration.conf -extensions 'req_ext' -key certificates/private-key.pem -out certificates/certificate.pem
	// openssl x509 -in certificates/certificate.pem -text
	key, issue := ioutil.ReadFile("certificates/private-key.pem")
	certificate, issue := ioutil.ReadFile("certificates/certificate.pem")
	if issue != nil {
		log.Fatal(issue.Error())
	}

	keyPair, issue := tls.X509KeyPair(certificate, key)
	if issue != nil {
		log.Fatal(issue.Error())
	}

	return &tls.Config{
		Certificates: []tls.Certificate{keyPair},
		ServerName:   "localhost",
	}
}

func baseHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte(request.Proto))
}
