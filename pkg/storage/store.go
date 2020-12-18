package storage

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"

	"sundae-party/api-server/pkg/apis/core/integration"

	mongo_store "sundae-party/api-server/pkg/storage/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var availableStoreType = map[string]string{
	"mongo": "mongo",
	"etcd3": "etcd3",
}

type Store interface {
	PutIntegration(context.Context, *integration.Integration) (*integration.Integration, error)
	GetIntegration(context.Context, string) (*integration.Integration, error)
	DeleteIntegration(context.Context, *integration.Integration) (string, error)
}

type StoreOption struct {
	// Type of store mongo | etcd3
	Type string
	// Endpoints list of db, for mongo only first item will be used
	Address []string
	// Port used to connect to the db
	Port string
	// User used with authentication procedure user:password
	User string
	// Password used with authentication procedure user:password
	Password string
	// TLS config object contenning client certificate for cert authentication procedure
	TLSConfig *tls.Config
	// DbName to use, availbale for mongo only
	DbName string
}

func NewStore(ctx context.Context, ops *StoreOption) (Store, error) {

	// If store type is mongo
	if ops.Type == availableStoreType["mongo"] {
		err := checkMongoOptions(ops)
		if err != nil {
			return nil, err
		}
		hosts := []string{fmt.Sprintf("%s:%s", ops.Address, ops.Port)}
		creds := options.Credential{AuthSource: "sundae", Username: ops.User, Password: ops.Password}
		// Create new MongoStore object with provided opetions
		ms, err := mongo_store.NewStore(ctx, ops.DbName, hosts, creds)
		if err != nil {
			fmt.Println(err)
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

func checkMongoOptions(ops *StoreOption) error {
	// Check if db name provided
	if ops.DbName == "" {
		return errors.New("No database name provided for mongo store.")
	}

	// Check endpoint address
	if len(ops.Address) == 0 || ops.Address[0] == "" {
		return errors.New("No mongo address provided.")
	}

	// Check endpoint port
	if ops.Port == "" {
		return errors.New("No mongo port provided.")
	}

	// Check if user and password or x509 provided
	if (ops.User != "" && ops.Password != "") || ops.TLSConfig != nil {
		return nil
	}
	return errors.New("Authentication options missing.")
}
