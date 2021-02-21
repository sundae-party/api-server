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

// BuildServerTlsConf create a tlsConfig object of type *tls.Config configured to be used in the server side.
// If one or many CA certificates are provided through CAPaths, the mTLS configuration will be enabled and this certificates will be used to validate the client certificates.
func BuildServerTlsConf(CAPaths []string, certPath string, keyPath string) (tlsConfig *tls.Config, err error) {
	if certPath == "" {
		return nil, errors.New("Missing informations, certificate file not provided.")
	}

	if keyPath == "" {
		return nil, errors.New("Missing informations, certificate key file not provided.")
	}

	// SSL server configuration
	serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}
	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}

	// mTLS configuration
	if len(CAPaths) > 0 {
		caCertPool := x509.NewCertPool()

		// Create a CA certificate pool with all CAs used to sign client certificates for the mTLS.
		for _, caPath := range CAPaths {
			caCert, err := ioutil.ReadFile(caPath)
			if err != nil {
				log.Fatal(err)
			}
			caCertPool.AppendCertsFromPEM(caCert)
		}

		// Create the server TLS Config with the CA pool and enable Client certificate validation.
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		tlsConfig.ClientCAs = caCertPool
	}

	return tlsConfig, nil
}

// LoadKeyPair create a tlsConfig object of type credentials.TransportCredentials configured to be used in the client side.
func LoadKeyPair(certPath string, keyPath string, caPath string) (clientTLSConfig credentials.TransportCredentials, err error) {
	if certPath == "" {
		return nil, errors.New("Missing informations, certificate file not provided.")
	}

	if keyPath == "" {
		return nil, errors.New("Missing informations, certificate key file not provided.")
	}

	// Load client cert & key
	certificate, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	// Load CA
	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	capool := x509.NewCertPool()
	if !capool.AppendCertsFromPEM(ca) {
		return nil, errors.New("Cann't add ca")
	}

	// Build TLS config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      capool,
	}

	return credentials.NewTLS(tlsConfig), nil
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
