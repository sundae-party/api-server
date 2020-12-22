package mongo

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"sundae-party/api-server/pkg/apis/core/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

type MongoStore struct {
	Client   *mongo.Client
	DataBase *mongo.Database
	Event    chan string
	Exit     chan os.Signal
}

func NewStore(c context.Context, DbName string, uri string, creds options.Credential) (*MongoStore, error) {

	// Create new mongo client with ops
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(creds))
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
		Event:    make(chan string),
		Exit:     make(chan os.Signal),
	}
	// signal.Notify(ms.Exit, syscall.SIGINT, syscall.SIGTERM)

	err = WatchEvent(c, ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

// Start mongo watch event on collections
func WatchEvent(ctx context.Context, s *MongoStore) error {
	csIntegration, err := s.DataBase.Collection("integrations").Watch(ctx, mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		return err
	}
	// TODO quit with cancel
	c, cancel := context.WithCancel(ctx)
	go func() {
		<-s.Exit
		log.Print("Stop prog: stop mongo watch")
		cancel()
	}()
	// Integration loop event
	log.Println("Start watch")
	go func() {
		for csIntegration.Next(c) {
			s.Event <- csIntegration.Current.String()
			log.Println(csIntegration.Current.String())
		}
	}()
	return nil
}

func (ms MongoStore) GetEvent() chan string {
	return ms.Event
}

// PutIntegration create or update an integration in mongo store.
func (ms MongoStore) PutIntegration(ctx context.Context, newIntegration *types.Integration) (*types.Integration, error) {
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
	updated := &types.Integration{}
	res.Decode(updated)
	return updated, nil
}

// GetIntegration find unique integration with given name and return it.
// If more one integration is find with this name an error is returned
func (ms MongoStore) GetIntegration(ctx context.Context, name string) (*types.Integration, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection("integrations")

	// Select an integration with it's unique name
	filter := bson.D{{Key: "name", Value: name}}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}

	var result []types.Integration
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

func (ms MongoStore) DeleteIntegration(ctx context.Context, deleteIntegration *types.Integration) (*types.Integration, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection("integrations")

	// Select an integration with it's unique name
	filter := bson.D{{Key: "name", Value: deleteIntegration.Name}}

	res := collection.FindOneAndDelete(c, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}
	deleted := &types.Integration{}
	res.Decode(deleted)

	return deleted, nil
}
