package utils

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

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
