package mongo

import (
	"context"
	"log"

	"github.com/sundae-party/api-server/pkg/apis/core/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//
//
// Light store function
//
//

// PutLight create or update a light in the store.
func (ms MongoStore) PutLight(ctx context.Context, light *types.Light) (*types.Light, error) {

	// Ensure kind is set to Light
	light.Mutation = "light"

	// Put the entity in the store
	res, err := ms.UpdateLightState(ctx, light)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// The key must be formated with "intergrationName/lightName" and can't be empty
func (ms MongoStore) GetLightByName(ctx context.Context, key string) (*types.Light, error) {

	res, err := ms.getEntityByName(ctx, key)
	if err != nil {
		return nil, err
	}

	// Convert bson result to light object
	var light types.Light
	err = res.Decode(&light)
	if err != nil {
		return nil, err
	}

	return &light, nil
}

// GetAllLight return all light stored in the store
func (ms MongoStore) GetAllLight(c context.Context) ([]types.Light, error) {
	cursor, err := ms.getAllEntities(c, "light", "")
	if err != nil {
		return nil, err
	}

	var lights []types.Light
	err = cursor.Decode(lights)
	if err != nil {
		return nil, err
	}

	return lights, nil
}

// GetLightByIntegration return a list of light
func (ms MongoStore) GetLightByIntegration(c context.Context, key string) ([]types.Light, error) {
	integrationName, _, err := decodeKey(key)
	if err != nil {
		return nil, err
	}

	cursor, err := ms.getAllEntities(c, "light", integrationName)
	if err != nil {
		return nil, err
	}

	var lights []types.Light
	err = cursor.Decode(lights)
	if err != nil {
		return nil, err
	}

	return lights, nil
}

// Delete a light in the store
func (ms MongoStore) DeleteLight(ctx context.Context, light *types.Light) (*types.Light, error) {
	res, err := ms.deleteEntity(ctx, light.Integration.Name, light.Name)
	if err != nil {
		return nil, err
	}

	var deleted types.Light
	err = res.Decode(&deleted)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// UpdateLightState will try to update the light state in the store, if the light is not present in the store it will be created
func (ms MongoStore) UpdateLightState(ctx context.Context, light *types.Light) (*types.Light, error) {

	// Convert light state to bson object
	bsonLightState := bson.M{"state": light.State}

	// Try to update light
	res, err := ms.updateEntityState(ctx, light.Integration.Name, light.Name, bsonLightState)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {

			// This light dosn't exist, create it
			log.Println("light dosn't exist") /// TODO: debug

			// Convert light to bson object
			bsonLight, err := bson.Marshal(light)
			if err != nil {
				return nil, err
			}

			// Try to create the new light
			newLight, err := ms.putEntity(ctx, light.Integration.Name, light.Name, bsonLight)
			if err != nil {
				return nil, err
			}

			// Convert bson result to light object
			var updated types.Light
			err = newLight.Decode(&updated)
			if err != nil {
				return nil, err
			}
			log.Println("new light created") /// TODO: debug
			return &updated, nil
		}
		return nil, err
	}

	// Light updated
	var updated types.Light
	err = res.Decode(&updated)
	if err != nil {
		return nil, err
	}
	log.Println("light state update") /// TODO: debug
	return &updated, nil
}
