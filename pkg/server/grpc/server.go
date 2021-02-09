package grpc

import (
	"context"
	"errors"
	"log"
	"net"
	"strings"

	"github.com/sundae-party/api-server/pkg/apis/core/integration"
	"github.com/sundae-party/api-server/pkg/apis/core/light"
	"github.com/sundae-party/api-server/pkg/apis/core/sensor"
	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/server/utils"
	"github.com/sundae-party/api-server/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type service struct {
	Store storage.Store
}

func CreateServer(clientCAsPath []string, certPath string, ceyPath string, store storage.Store) {

	tlsConfig, err := utils.BuildTlsConf(clientCAsPath, certPath, ceyPath)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		grpc.UnaryInterceptor(getIntegrationFromCert),
	)

	// Create Integration handler object
	ih := &integration.IntegrationHandler{
		Store: store,
	}
	// Create Light handler
	lh := &light.LightHandler{
		Store: store,
	}
	// Create Sensor handler
	sh := &sensor.SensorHandler{
		Store: store,
	}

	// Add handlers to the server
	types.RegisterIntegrationHandlerServer(grpcServer, ih)
	types.RegisterLightHandlerServer(grpcServer, lh)
	types.RegisterSensorHandlerServer(grpcServer, sh)

	// GRPC servers listen
	go func() {

		// Create new grpc TCP listener
		tcpLis, err := net.Listen("tcp", "0.0.0.0:8443")
		if err != nil {
			log.Fatalf("%s\n", err)
		}

		// TCP sock
		err = grpcServer.Serve(tcpLis)
		if err != nil {
			log.Fatal(err)
		}

		// Create new grpc LINUX socket listener
		// sockLis, err := net.Listen("unix", "/var/sundae/api-server.sock")

		// Unix sock
		//grpcServer.Serve(sockLis)
	}()

}

func getIntegrationFromCert(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// Test getting integration name from client cert
	peer, ok := peer.FromContext(ctx)
	if ok {
		mtls, ok := peer.AuthInfo.(credentials.TLSInfo)
		if ok {
			// get cn from mtls cert
			cn := mtls.State.PeerCertificates[0].Subject.CommonName
			// get integration name from cn
			cliInfos := strings.Split(cn, ":")
			if len(cliInfos) != 2 && cliInfos[0] != "integration" {
				return nil, errors.New("Invalid integration CN format in the cli mtls cert.")
			}
			// info.Server return server handler containing the store to use to get the integration object
			// store := info.Server.(service)
			// integration, err := store.GetIntegration(ctx, cliInfos[1])
			// if err != nil {
			// 	return nil, err
			// }
			log.Printf("%s \n", cliInfos[1])
		}
	}

	//metadata.AppendToOutgoingContext(ctx, )
	return handler(ctx, req)
}
