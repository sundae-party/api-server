package utils

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func BuildTlsConf(CAPaths []string, certFile string, keyFile string) (*tls.Config, error) {
	caCertPool := x509.NewCertPool()

	// Create a CA certificate pool with all Client CAs listed in srvConf.ClientCAsPath
	for _, caPath := range CAPaths {
		caCert, err := ioutil.ReadFile(caPath)
		if err != nil {
			log.Fatal(err)
		}
		caCertPool.AppendCertsFromPEM(caCert)
	}

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	// If cert and key file are provided for server ssl termination
	if certFile != "" && keyFile != "" {
		serverCert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		tlsConfig.Certificates = []tls.Certificate{serverCert}
	}

	return tlsConfig, nil
}
