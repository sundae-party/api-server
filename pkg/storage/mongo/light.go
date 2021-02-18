package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sundae-party/api-server/pkg/apis/core/types"

	"go.mongodb.org/mongo-driver/bson"
)

func buildLightKey(light *types.Light) (key string, err error) {
	if light.IntegrationName != "" && light.Name != "" {
		return fmt.Sprintf("%s/%s", light.IntegrationName, light.Name), nil
	}
	return "", errors.New("Invalid light format, the integration name or light name is missing.")
}

//
//
// Light store function
//
//

// PutLight create or update a light in the store.
func (ms MongoStore) PutLight(ctx context.Context, light *types.Light) (*types.Light, error) {

	key, err := buildLightKey(light)
	if err != nil {
		return nil, err
	}

	// Ensure kind is set to Light
	light.Mutation = lightKind

	// Convert light to bson object
	bsonLight, err := bson.Marshal(light)

	// Put the light in the store
	res, err := ms.putEntity(ctx, key, bsonLight)
	if err != nil {
		return nil, err
	}

	newLight := &types.Light{}
	err = res.Decode(newLight)
	if err != nil {
		return nil, err
	}

	return newLight, nil
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
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	cursor, err := ms.getAllEntities(c, lightKind, "")
	if err != nil {
		return nil, err
	}

	var lights []types.Light
	err = cursor.All(ctx, &lights)
	if err != nil {
		return nil, err
	}
	return lights, nil
}

// GetLightByIntegration return a list of light
func (ms MongoStore) GetLightByIntegration(c context.Context, integrationName string) ([]types.Light, error) {
	cursor, err := ms.getAllEntities(c, lightKind, integrationName)
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
	key := fmt.Sprintf("%s/%s", light.IntegrationName, light.Name)

	res, err := ms.deleteEntity(ctx, key)
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

	key := fmt.Sprintf("%s/%s", light.IntegrationName, light.Name)

	// Try to update light
	res, err := ms.updateEntityState(ctx, key, bsonLightState)
	if err != nil {
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

func (ms MongoStore) UpdateLightStateDesiredState(ctx context.Context, light *types.Light) (*types.Light, error) {

	// Convert light state to bson object
	bsonLightState := bson.M{"desiredstate": light.DesiredState}

	key := fmt.Sprintf("%s/%s", light.IntegrationName, light.Name)

	// Try to update light
	res, err := ms.updateEntityState(ctx, key, bsonLightState)
	if err != nil {
		return nil, err
	}

	// Light updated
	var updated types.Light
	err = res.Decode(&updated)
	if err != nil {
		return nil, err
	}
	log.Println("light desiredstate update") /// TODO: debug
	return &updated, nil
}
