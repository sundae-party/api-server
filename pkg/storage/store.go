package storage

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"strings"

	"sundae-party/api-server/pkg/apis/core/types"
	store_type "sundae-party/api-server/pkg/storage/types"

	mongo_store "sundae-party/api-server/pkg/storage/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var availableStoreType = map[string]string{
	"mongo": "mongo",
	"etcd3": "etcd3",
}

type Store interface {
	GetEvent() chan store_type.StoreEvent

	PutIntegration(context.Context, *types.Integration) (*types.Integration, error)
	GetIntegration(context.Context, string) (*types.Integration, error)
	DeleteIntegration(context.Context, *types.Integration) (*types.Integration, error)

	// Genric entity
	GetAllEntities(ctx context.Context, kind string, integrationName string) ([]byte, error)

	// Light store
	PutLight(context.Context, *types.Light) (*types.Light, error)
	GetLightByName(context.Context, string) (*types.Light, error)
	GetAllLight(context.Context) ([]types.Light, error)
	GetLightByIntegration(context.Context, string) ([]types.Light, error)
	DeleteLight(context.Context, *types.Light) (*types.Light, error)
}

// TODO: Create a genrique watchEvent type struct
type EventMessage struct {
	OperationType string
	FullDocument  interface{}
	Ns            ns
}
type ns struct {
	Db   string
	Coll string
}

type StoreOption struct {
	// Type of store mongo | etcd3
	Type string
	// Endpoints list of db in style ADDR:PORT
	Address []string
	// User used with authentication procedure user:password
	User string
	// Password used with authentication procedure user:password
	Password string
	// TLS config object contenning client certificate for cert authentication procedure
	TLSConfig *tls.Config
	// DbName to use, availbale for mongo only
	DbName string
	// RsName to connect in replicat set mode, availbale for mongo only
	RsName string
}

func NewStore(ctx context.Context, ops *StoreOption) (Store, error) {

	// If store type is mongo
	if ops.Type == availableStoreType["mongo"] {
		// Check mongo ops
		err := checkMongoOptions(ops)
		if err != nil {
			return nil, err
		}
		// Build mongo uri
		uri, err := buildMongoUri(ops)
		if err != nil {
			return nil, err
		}
		// Build creds ops
		creds := options.Credential{AuthSource: ops.DbName, Username: ops.User, Password: ops.Password}
		// Create new MongoStore object with provided opetions
		ms, err := mongo_store.NewStore(ctx, ops.DbName, uri, creds)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Fail connecting to mongo.")
		}
		return ms, nil
	}

	// TODO If store type is ETCD
	// if ops.Type == availableStoreType["etcd3"] {
	// 	return &etcdStore{cli: "etcd3"}, nil
	// }

	return nil, errors.New("Store type unavailable.")
}

func buildMongoUri(ops *StoreOption) (string, error) {
	var uri strings.Builder
	_, err := uri.WriteString("mongodb://")
	if err != nil {
		return "", err
	}
	for _, addr := range ops.Address {
		_, err = uri.WriteString(addr)
		if err != nil {
			return "", err
		}
	}
	_, err = uri.WriteString("/?replicaSet=")
	if err != nil {
		return "", err
	}
	_, err = uri.WriteString(ops.RsName)
	if err != nil {
		return "", err
	}
	return uri.String(), nil
}

func checkMongoOptions(ops *StoreOption) error {
	// Check if db name provided
	if ops.DbName == "" {
		return errors.New("No database name provided for mongo store.")
	}

	// Check endpoint address
	if len(ops.Address) == 0 || ops.Address[0] == "" {
		return errors.New("No mongo address provided.")
	}

	// Check if replicat set name is provided
	if ops.RsName == "" {
		return errors.New("Mongo replicat set name not provided.")
	}

	// Check if user and password or x509 provided
	if (ops.User != "" && ops.Password != "") || ops.TLSConfig != nil {
		return nil
	}
	return errors.New("Authentication options missing.")
}
