package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
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

func GetClientAuthInformationFromCert(ctx context.Context) (clientType string, clientName string, err error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		mtls, ok := peer.AuthInfo.(credentials.TLSInfo)
		if ok {
			// get cn from mtls cert
			cn := mtls.State.PeerCertificates[0].Subject.CommonName
			// get integration name from cn
			cliInfos := strings.Split(cn, ":")
			if len(cliInfos) != 2 && cliInfos[0] != "integration" {
				return "", "", errors.New("Invalid integration CN format in the cli mtls cert.")
			}
			return cliInfos[0], cliInfos[1], nil
		}
		return "", "", errors.New("Error getting the peer authentitcation information.")
	}
	return "", "", errors.New("Error getting the peer information in ctx.")
}
