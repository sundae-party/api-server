package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"sundae-party/api-server/pkg/apis/core/integration"
	"sundae-party/api-server/pkg/apis/core/types"
	"sundae-party/api-server/pkg/server"
	"time"

	"sundae-party/api-server/pkg/storage"

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
		Store:        mongoStore,
		ServiceEvent: make(chan *types.CallIntegrationServiceRequest),
	}

	// TODO: test to be removed
	go func() {
		min := 1
		max := 600
		for {
			val := rand.Intn(max-min) + min
			time.Sleep(time.Duration(val) * time.Second)
			ih.ServiceEvent <- &types.CallIntegrationServiceRequest{
				IntegrationName: "Zwave",
				Service: &types.IntegrationService{
					Name: "refresh_entities",
				},
			}
		}
	}()

	// Add handlers to the server
	// light.RegisterLightHandlerServer(grpcServer, &light.LightHandler{Store: s})
	types.RegisterIntegrationHandlerServer(grpcServer, ih)

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
