package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	DataBase *mongo.Database
}

func NewStore(c context.Context, DbName string) (*MongoStore, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	// TODO Check context timeout db connexion best practice
	go func() {
		for {
			select {
			case <-c.Done():
				cancel()
			case <-ctx.Done():
				break
			}
		}
	}()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database(DbName)

	return &MongoStore{
		DataBase: db,
	}, nil
}

// Mongo
func (ms MongoStore) GetIntegration(integration string) string {
	return "mongo Get"
}

func (ms MongoStore) PutIntegration(integration string) string {
	return "mongo Put"
}

func (ms MongoStore) DeleteIntegration(integration string) string {
	return "mongo Delete"
}

// func (s Store) Put(k string, d string) (string, error) {
// 	s.Client.KV.Put(ctx context.Context, key string, val string, opts ...clientv3.OpOption)
// 	s.Data[k] = d
// 	return s.Data[k], nil
// }

// func (s Store) Get(k string) (string, error) {
// 	return s.Data[k], nil
// }

// func (s Store) GetByIntegration(i string) []string {
// 	var resp []string
// 	for k, v := range s.Data {
// 		regStr := fmt.Sprintf("^/%s/", i)
// 		var regLight = regexp.MustCompile(regStr)
// 		if regLight.MatchString(k) {
// 			resp = append(resp, v)
// 		}
// 	}
// 	return resp
// }
