package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/sundae-party/api-server/pkg/apis/core/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	newIntegration.Mutation = "integration"

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

func (ms MongoStore) UpdateIntegrationState(ctx context.Context, integration *types.Integration) (*types.Integration, error) {

	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection(integrationCollection)

	// Select the integration with it's unique name
	filter := bson.D{
		{Key: "name", Value: integration.Name},
	}

	// Return new val insted old val
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Convert integration state to bson object
	update := bson.M{"$set": bson.M{"state": integration.State}}

	res := collection.FindOneAndUpdate(c, filter, update, opts)
	if res.Err() != nil {
		if res.Err().Error() == mongo.ErrNoDocuments.Error() {
			// Try to create the new integration
			res, err := ms.PutIntegration(ctx, integration)
			if err != nil {
				return nil, err
			}
			return res, nil
		} else {
			return nil, res.Err()
		}
	}

	// integration updated
	var updated types.Integration
	err := res.Decode(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (ms MongoStore) UpdateIntegrationDesiredState(ctx context.Context, integration *types.Integration) (*types.Integration, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	collection := ms.DataBase.Collection(integrationCollection)

	// Select the integration with it's unique name
	filter := bson.D{
		{Key: "name", Value: integration.Name},
	}

	// Return new val insted old val
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Convert integration state to bson object
	update := bson.M{"$set": bson.M{"desiredstate": integration.DesiredState}}

	res := collection.FindOneAndUpdate(c, filter, update, opts)
	if res.Err() != nil {
		return nil, res.Err()
	}

	// integration updated
	var updated types.Integration
	err := res.Decode(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}
