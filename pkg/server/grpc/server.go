package grpc

import (
	"log"
	"net"

	"github.com/sundae-party/api-server/pkg/apis/core/integration"
	"github.com/sundae-party/api-server/pkg/apis/core/light"
	"github.com/sundae-party/api-server/pkg/apis/core/sensor"
	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/server/utils"
	"github.com/sundae-party/api-server/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type service struct {
	Store storage.Store
}

func CreateServer(clientCAsPath []string, certPath string, ceyPath string, store storage.Store) {

	tlsConfig, err := utils.BuildServerTlsConf(clientCAsPath, certPath, ceyPath)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		//grpc.UnaryInterceptor(getIntegrationFromCert),
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
