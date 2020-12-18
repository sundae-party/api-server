package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"sundae-party/api-server/pkg/apis/core/integration"
)

type MongoStore struct {
	Client   *mongo.Client
	DataBase *mongo.Database
}

func NewStore(c context.Context, DbName string, hosts []string, creds options.Credential) (*MongoStore, error) {

	// Create new mongo client with ops
	client, err := mongo.NewClient(options.Client().SetHosts(hosts).SetAuth(creds))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()
	// Connect client to mongo instance
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Select the database
	db := client.Database(DbName)

	ms := &MongoStore{
		Client:   client,
		DataBase: db,
	}

	return ms, nil
}

// PutIntegration create or update an integration in mongo store.
func (ms MongoStore) PutIntegration(ctx context.Context, newIntegration *integration.Integration) (*integration.Integration, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection("integrations")
	// If no integration found create new one and return updated value instead old value
	opts := options.FindOneAndReplace().SetUpsert(true)
	opts = opts.SetReturnDocument(options.After)
	// Select an integration with it's unique name
	filter := bson.D{{Key: "name", Value: newIntegration.Name}}

	// Convert intergration to bson
	replacment, err := bson.Marshal(newIntegration)
	if err != nil {
		return nil, err
	}
	res := collection.FindOneAndReplace(c, filter, replacment, opts)
	if res.Err() != nil {
		return nil, res.Err()
	}
	updated := &integration.Integration{}
	res.Decode(updated)
	return updated, nil
}

// GetIntegration find unique integration with given name and return it.
// If more one integration is find with this name an error is returned
func (ms MongoStore) GetIntegration(ctx context.Context, name string) (*integration.Integration, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection("integrations")

	// Select an integration with it's unique name
	filter := bson.D{{Key: "name", Value: name}}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}

	var result []integration.Integration
	err = cursor.All(c, &result)
	if err != nil {
		return nil, err
	}

	if len(result) > 1 {
		return nil, errors.New("This integration has been found more than once.")
	}
	if len(result) == 0 {
		return nil, errors.New("Integration not found.")
	}
	return &result[0], nil
}

func (ms MongoStore) DeleteIntegration(ctx context.Context, deleteIntegration *integration.Integration) (string, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection("integrations")

	// Select an integration with it's unique name
	filter := bson.D{{Key: "name", Value: deleteIntegration.Name}}

	res := collection.FindOneAndDelete(c, filter)
	if res.Err() != nil {
		return "", res.Err()
	}
	deleted := &integration.Integration{}
	res.Decode(deleted)

	return fmt.Sprintf("%s deleted.", deleted.Name), nil
}
