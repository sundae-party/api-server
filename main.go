package main

import (
	"context"
	"log"
	"net"

	"github.com/sundae-party/api-server/pkg/apis/core/integration"
	"github.com/sundae-party/api-server/pkg/apis/core/light"
	"github.com/sundae-party/api-server/pkg/apis/core/sensor"
	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/server"

	"github.com/sundae-party/api-server/pkg/storage"

	"google.golang.org/grpc"
)

func main() {
	// Create new mongo store
	mongo := &storage.StoreOption{
		Type:     "mongo",
		Address:  []string{"gogs.connan.pro:27018"},
		User:     "sundae",
		Password: "pass",
		DbName:   "sundae",
		RsName:   "rs0",
	}
	ctx := context.Background()
	mongoStore, err := storage.NewStore(ctx, mongo)
	if err != nil {
		log.Fatalln(err)
	}

	// Create new grpc TCP listener
	tcpLis, err := net.Listen("tcp", "0.0.0.0:8443")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	// Create new grpc LINUX socket listener
	// sockLis, err := net.Listen("unix", "/var/sundae/api-server.sock")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Create Integration handler object
	ih := &integration.IntegrationHandler{
		Store: mongoStore,
	}
	// Create Light handler
	lh := &light.LightHandler{
		Store: mongoStore,
	}
	// Create Sensor handler
	sh := &sensor.SensorHandler{
		Store: mongoStore,
	}

	// Add handlers to the server
	types.RegisterIntegrationHandlerServer(grpcServer, ih)
	types.RegisterLightHandlerServer(grpcServer, lh)
	types.RegisterSensorHandlerServer(grpcServer, sh)

	// GRPC servers listen
	go grpcServer.Serve(tcpLis)
	// go grpcServer.Serve(sockLis)

	// Http server, ui, api, ws
	tlsConf := server.ServerConfig{
		ServerMode:          server.HTTPSMode,
		HTTPSAddr:           ":443",
		HTTPAddr:            ":80",
		EnableHTTPSredirect: true,
		KeyPath:             "ssl/sundae.key",
		CertPath:            "ssl/sundae.pem",
		EnableMTLS:          true,
		ClientCAsPath:       []string{"ssl/ca.pem"},
	}
	server.Serve(tlsConf, &mongoStore)
}
