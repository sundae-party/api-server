package main

import (
	"context"
	"log"
	"os"

	"github.com/sundae-party/api-server/pkg/server/grpc"
	"github.com/sundae-party/api-server/pkg/server/mux"

	"github.com/sundae-party/api-server/pkg/storage"
)

func main() {
	// Create new mongo store
	mongo := &storage.StoreOption{
		Type:     "mongo",
		Address:  []string{os.Getenv("MONGO_ADDR")},
		User:     os.Getenv("MONGO_USR"),
		Password: os.Getenv("MONGO_PWD"),
		DbName:   os.Getenv("MONGO_DB"),
		RsName:   os.Getenv("MONGO_RS"),
	}
	ctx := context.Background()
	mongoStore, err := storage.NewStore(ctx, mongo)
	if err != nil {
		log.Fatalln(err)
	}

	// gRPC server
	grpc.CreateServer([]string{"ssl/ca.pem"}, "ssl/sundae-apiserver.pem", "ssl/sundae-apiserver.key", mongoStore)

	// Http server, ui, api, ws
	muxTlsConf := mux.ServerConfig{
		ServerMode:          mux.HTTPSMode,
		HTTPSAddr:           ":443",
		HTTPAddr:            ":80",
		EnableHTTPSredirect: true,
		KeyPath:             "ssl/sundae-apiserver.key",
		CertPath:            "ssl/sundae-apiserver.pem",
		EnableMTLS:          true,
		ClientCAsPath:       []string{"ssl/ca.pem"},
	}
	mux.Serve(muxTlsConf, &mongoStore)
}
