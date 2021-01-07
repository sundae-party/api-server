package mongo

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"sundae-party/api-server/pkg/apis/core/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	integrationCollection = "integrations"
	entityCollection      = "entities"
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

	// Select and init the database
	db := client.Database(DbName)
	err = initDb(c, db)
	if err != nil {
		return nil, err
	}

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
			s.Event <- cs.Current.String()
			// TODO: Debug mode only
			log.Println(cs.Current.String())
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

	collection := ms.DataBase.Collection(integrationCollection)
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

	collection := ms.DataBase.Collection(integrationCollection)

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

	collection := ms.DataBase.Collection(integrationCollection)

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

//
//
// TODO: Use an interface to represent an entity
//
//

func (ms MongoStore) PutEntity(ctx context.Context, key string, entity []byte) ([]byte, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	// decode key in order to extract integration and entity name
	// ex /integration01/entityA
	entityInfos := strings.Split(key, "/")
	if len(entityInfos) != 2 {
		return nil, errors.New("Invalid key format.")
	}

	collection := ms.DataBase.Collection(entityCollection)
	// If no entity found create new one
	// If entity exist update it and return updated value
	opts := options.FindOneAndReplace().SetUpsert(true)
	opts = opts.SetReturnDocument(options.After)
	// Select the entity with it's unique name and integration pairs
	// ex:
	//	{"integration": int1, "name": "ent1"}
	//	{"integration": int1, "name": "ent1"} KO
	//	{"integration": int1, "name": "ent2"} OK
	//	{"integration": int2, "name": "ent1"} Ok
	//
	filter := bson.D{
		{Key: "integration.name", Value: entityInfos[0]},
		{Key: "name", Value: entityInfos[1]},
	}

	res := collection.FindOneAndReplace(c, filter, entity, opts)
	if res.Err() != nil {
		return nil, res.Err()
	}
	updated, err := res.DecodeBytes()
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (ms MongoStore) GetEntityByName(ctx context.Context, key string) ([]byte, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	// decode key in order to extract integration and entity name
	// ex /integration01/entityA
	entityInfos := strings.Split(key, "/")
	if len(entityInfos) != 2 {
		return nil, errors.New("Invalid key format.")
	}

	collection := ms.DataBase.Collection(entityCollection)

	// Select the entity with it's unique name
	filter := bson.D{
		{Key: "integration.name", Value: entityInfos[0]},
		{Key: "name", Value: entityInfos[1]},
	}

	res := collection.FindOne(c, filter)

	if res.Err() != nil {
		return nil, res.Err()
	}

	entity, err := res.DecodeBytes()
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (ms MongoStore) GetAllEntities(c context.Context) ([][]byte, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection(entityCollection)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var res [][]byte
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (ms MongoStore) GetEntitiesByIntegration(c context.Context, key string) ([][]byte, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	// decode key in order to extract integration and entity name
	// ex /integration01/entityA
	entityInfos := strings.Split(key, "/")
	if len(entityInfos) != 2 {
		return nil, errors.New("Invalid key format.")
	}

	collection := ms.DataBase.Collection(entityCollection)

	// Select the entity with it's unique name
	filter := bson.D{
		{Key: "integration.name", Value: entityInfos[0]},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var res [][]byte
	err = cursor.All(ctx, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ms MongoStore) DeleteEntity(ctx context.Context, key string, entity []byte) ([]byte, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	// decode key in order to extract integration and entity name
	// ex /integration01/entityA
	entityInfos := strings.Split(key, "/")
	if len(entityInfos) != 2 {
		return nil, errors.New("Invalid key format.")
	}

	collection := ms.DataBase.Collection(entityCollection)

	// Select the entity with it's unique name
	filter := bson.D{
		{Key: "integration.name", Value: entityInfos[0]},
		{Key: "name", Value: entityInfos[1]},
	}

	res := collection.FindOneAndDelete(c, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}
	deleted, err := res.DecodeBytes()
	if err != nil {
		return nil, err
	}
	return deleted, nil
}
