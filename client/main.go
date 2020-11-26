package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := &http.Client{
		Transport: transport(),
	}

	response, issue := client.Get("https://localhost:5000")
	if issue != nil {
		log.Fatal(issue.Error())
	}

	binaryResponse, issue := ioutil.ReadAll(response.Body)
	if issue != nil {
		log.Fatal(issue.Error())
	}

	response.Body.Close()
	fmt.Printf("Code: %d\nProtocol: %s\n", response.StatusCode, binaryResponse)
}

func transport() *http2.Transport {
	return &http2.Transport{
		TLSClientConfig:    tlsConfig(),
		DisableCompression: true,
		AllowHTTP:          false,
	}
}

func tlsConfig() *tls.Config {
	// Copy server/certificates/certificate.pem file to certificates directory
	certificate, issue := ioutil.ReadFile("certificates/certificate.pem")
	if issue != nil {
		log.Fatal(issue.Error())
	}

	rootCertificateAuthority := x509.NewCertPool()
	rootCertificateAuthority.AppendCertsFromPEM(certificate)

	return &tls.Config{
		RootCAs:            rootCertificateAuthority,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}
