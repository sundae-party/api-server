package storage

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"

	mongo_store "sundae-party/api-server/pkg/storage/mongo"
)

var availableStoreType = map[string]string{
	"mongo": "mongo",
	"etcd3": "etcd3",
}

type Store interface {
	GetIntegration(string) string
	PutIntegration(string) string
	DeleteIntegration(string) string
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

type etcdStore struct {
	cli string
}

func NewStore(ctx context.Context, ops StoreOption) (Store, error) {

	log.Printf("Creating %s", ops.Type)
	// Check if type exist
	ok := false
	for _, v := range availableStoreType {
		if v == ops.Type {
			ok = true
			break
		}
	}
	if !ok {
		return nil, errors.New("This store type in not available.")
	}

	// If store type is mongo
	if ops.Type == availableStoreType["mongo"] {
		// Create new MongoStore object with provided opetions

		// TODO Check if ops.DbName != nil
		// TODO Check user, password, cert and provide it.
		ms, err := mongo_store.NewStore(ctx, ops.DbName)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Fail connecting to mongo.")
		}
		return ms, nil

	} else if ops.Type == availableStoreType["etcd3"] {
		return &etcdStore{cli: "etcd3"}, nil
	}
	// TODO improve check ...
	log.Fatal("Error creating store")
	return nil, nil
}

// ETCD
func (es etcdStore) GetIntegration(integration string) string {
	return "etcd Get"
}

func (es etcdStore) PutIntegration(integration string) string {
	return "etcd Put"
}

func (ems etcdStore) DeleteIntegration(integration string) string {
	return "etcd Delete"
}
