package mongo

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	store_type "github.com/sundae-party/api-server/pkg/storage/types"
)

const (
	integrationCollection = "integrations"
	entityCollection      = "entities"
	lightKind             = "light"
	binarySensorKind      = "binarySensor"
)

type MongoStore struct {
	Client   *mongo.Client
	DataBase *mongo.Database
	// Events struct {
	//	allJsonEvents chan []byte
	//	integrationEvents chan Integration
	//  lightEvents chan Light
	// }
	Event            chan store_type.StoreEvent
	IntegrationEvent chan types.Integration
	Exit             chan os.Signal
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

	// Select and init the database
	db := client.Database(DbName)
	err = initDb(c, db)
	if err != nil {
		return nil, err
	}

	ms := &MongoStore{
		Client:           client,
		DataBase:         db,
		Event:            make(chan store_type.StoreEvent),
		IntegrationEvent: make(chan types.Integration),
		Exit:             make(chan os.Signal),
	}
	// signal.Notify(ms.Exit, syscall.SIGINT, syscall.SIGTERM)

	err = WatchEvent(c, ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func initDb(c context.Context, db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(c, time.Second*2)
	defer cancel()

	// Create integration index with name as unique value
	integrationCollection := db.Collection(integrationCollection)

	integrationIndex := mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	indexName, err := integrationCollection.Indexes().CreateOne(ctx, integrationIndex)
	if err != nil {
		return err
	}
	log.Printf("Index %s created\n", indexName)

	// Create entity index with integration name and entity name pair as unique value
	ientityCollection := db.Collection(entityCollection)
	entityIndex := mongo.IndexModel{
		Keys: bson.D{
			{Key: "integration.name", Value: 1},
			{Key: "name", Value: 2},
		},
		Options: options.Index().SetUnique(true),
	}

	indexName, err = ientityCollection.Indexes().CreateOne(ctx, entityIndex)
	if err != nil {
		return err
	}
	log.Printf("Index %s created\n", indexName)
	return nil
}

// Start mongo watch on all the db
func WatchEvent(ctx context.Context, s *MongoStore) error {

	cs, err := s.DataBase.Watch(ctx, mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		return err
	}

	// Db loop event
	log.Println("Start watch")
	go func() {
		for cs.Next(ctx) {
			var event store_type.StoreEvent
			err := cs.Decode(&event)
			if err != nil {
				log.Println("Mongo store -- decode mongo change stream event error")
				log.Println(err)
			}
			s.Event <- event
			// TODO: Debug mode only
			// log.Println(cs.Current.String())
		}
	}()

	return nil
}
