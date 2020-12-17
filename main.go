package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sundae-party/api-server/pkg/server"

	"sundae-party/api-server/pkg/storage"

	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()

	mongo_ops := storage.StoreOption{
		Type: "mongo",
	}

	etcd_ops := storage.StoreOption{
		Type: "etcd3",
	}

	ms, _ := storage.NewStore(ctx, mongo_ops)
	es, _ := storage.NewStore(ctx, etcd_ops)

	fmt.Println(ms.GetIntegration("Hue"))
	fmt.Println(es.GetIntegration("Hue"))

	lis, err := net.Listen("tcp", "0.0.0.0:8443")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// light.RegisterLightHandlerServer(grpcServer, &light.LightHandler{Store: s})
	go grpcServer.Serve(lis)

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

	server.Serve(tlsConf)
}
