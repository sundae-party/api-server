package mongo

import (
	"context"
	"errors"
	"log"
	"time"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PutEntity create or replace an entity in the store
func (ms MongoStore) putEntity(ctx context.Context, integrationName string, entityName string, entity []byte) (*mongo.SingleResult, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	if integrationName == "" || entityName == "" {
		return nil, errors.New("Cannot insert this entity, Integration name or entity name is empty")
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
		{Key: "integration.name", Value: integrationName},
		{Key: "name", Value: entityName},
	}

	// Create or replace entity
	res := collection.FindOneAndReplace(c, filter, entity, opts)
	if res.Err() != nil {
		return nil, res.Err()
	}
	return res, nil
}

// Get an entity with key from the store
func (ms MongoStore) getEntityByName(ctx context.Context, key string) (*mongo.SingleResult, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	integrationName, entityName, err := decodeKey(key)
	if err != nil {
		return nil, err
	}

	collection := ms.DataBase.Collection(entityCollection)

	// Select the entity with it's unique name
	filter := bson.D{
		{Key: "integration.name", Value: integrationName},
		{Key: "name", Value: entityName},
	}

	// Get light
	res := collection.FindOne(c, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

// getAllEntities return a mongo cursor object that should be decode with the corect type
func (ms MongoStore) getAllEntities(c context.Context, kind string, integrationName string) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection(entityCollection)

	filter := bson.D{}

	if kind != "" {
		filter = append(filter, bson.E{Key: "kind", Value: kind})
	}
	if integrationName != "" {
		filter = append(filter, bson.E{Key: "integration.name", Value: integrationName})
	}

	// Get all entity
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

// GetAllEntities return a json object with all entity stored in the store
func (ms MongoStore) GetAllEntities(c context.Context, kind string, integrationName string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection(entityCollection)

	filter := bson.D{}

	if kind != "" {
		filter = append(filter, bson.E{Key: "kind", Value: kind})
	}
	if integrationName != "" {
		filter = append(filter, bson.E{Key: "integration.name", Value: integrationName})
	}

	// Get all entity
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var bsonRes []bson.M
	err = cursor.All(ctx, &bsonRes)
	if err != nil {
		return nil, err
	}

	log.Printf("%s", bsonRes)

	res, err := json.Marshal(bsonRes)
	if err != nil {
		return nil, err
	}

	log.Printf("%s", res)

	return res, nil
}

// deleteEntity delete an entity in the store
func (ms MongoStore) deleteEntity(ctx context.Context, integrationName string, entityName string) (*mongo.SingleResult, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection(entityCollection)

	// Select the entity with it's unique name
	filter := bson.D{
		{Key: "integration.name", Value: integrationName},
		{Key: "name", Value: entityName},
	}

	res := collection.FindOneAndDelete(c, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}
