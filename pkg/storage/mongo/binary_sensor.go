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

func buildBinarySensortKey(binarySensor *types.BinarySensor) (key string, err error) {
	if binarySensor.Integration != nil {
		if binarySensor.Integration.Name != "" && binarySensor.Name != "" {
			return fmt.Sprintf("%s/%s", binarySensor.Integration.Name, binarySensor.Name), nil
		}
		return "", errors.New("Invalid binary sensor format, the integration name or binary sensor name is missing.")
	}
	return "", errors.New("Invalid binary sensor format, the integration infos is empty.")
}

//
//
// binary sensor store function
//
//

// PutBinarySensor create or update a BinarySensor in the store.
func (ms MongoStore) PutBinarySensor(ctx context.Context, BinarySensor *types.BinarySensor) (*types.BinarySensor, error) {

	key, err := buildBinarySensortKey(BinarySensor)
	if err != nil {
		return nil, err
	}

	// Ensure kind is set to binarySensor
	BinarySensor.Mutation = binarySensorKind

	// Convert BinarySensor to bson object
	bsonBinarySensor, err := bson.Marshal(BinarySensor)

	// Put the BinarySensor in the store
	res, err := ms.putEntity(ctx, key, bsonBinarySensor)
	if err != nil {
		return nil, err
	}

	var newBinarySensor types.BinarySensor
	err = res.Decode(&newBinarySensor)
	if err != nil {
		return nil, err
	}

	return &newBinarySensor, nil
}

// The key must be formated with "intergrationName/BinarySensorName" and can't be empty
func (ms MongoStore) GetBinarySensorByName(ctx context.Context, key string) (*types.BinarySensor, error) {

	res, err := ms.getEntityByName(ctx, key)
	if err != nil {
		return nil, err
	}

	// Convert bson result to BinarySensor object
	var BinarySensor types.BinarySensor
	err = res.Decode(&BinarySensor)
	if err != nil {
		return nil, err
	}

	return &BinarySensor, nil
}

// GetAllBinarySensor return all BinarySensor stored in the store
func (ms MongoStore) GetAllBinarySensor(c context.Context) ([]types.BinarySensor, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	cursor, err := ms.getAllEntities(c, binarySensorKind, "")
	if err != nil {
		return nil, err
	}

	var BinarySensors []types.BinarySensor
	err = cursor.All(ctx, &BinarySensors)
	if err != nil {
		return nil, err
	}
	return BinarySensors, nil
}

// GetBinarySensorByIntegration return a list of BinarySensor
func (ms MongoStore) GetBinarySensorByIntegration(c context.Context, integrationName string) ([]types.BinarySensor, error) {
	cursor, err := ms.getAllEntities(c, binarySensorKind, integrationName)
	if err != nil {
		return nil, err
	}

	var BinarySensors []types.BinarySensor
	err = cursor.Decode(BinarySensors)
	if err != nil {
		return nil, err
	}

	return BinarySensors, nil
}

// Delete a BinarySensor in the store
func (ms MongoStore) DeleteBinarySensor(ctx context.Context, BinarySensor *types.BinarySensor) (*types.BinarySensor, error) {
	key := fmt.Sprintf("%s/%s", BinarySensor.Integration.Name, BinarySensor.Name)

	res, err := ms.deleteEntity(ctx, key)
	if err != nil {
		return nil, err
	}

	var deleted types.BinarySensor
	err = res.Decode(&deleted)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// UpdateBinarySensorState will try to update the BinarySensor state in the store, if the BinarySensor is not present in the store it will be created
func (ms MongoStore) UpdateBinarySensorState(ctx context.Context, BinarySensor *types.BinarySensor) (*types.BinarySensor, error) {

	// Convert BinarySensor state to bson object
	bsonBinarySensorState := bson.M{"state": BinarySensor.State}

	key := fmt.Sprintf("%s/%s", BinarySensor.Integration.Name, BinarySensor.Name)

	// Try to update BinarySensor
	res, err := ms.updateEntityState(ctx, key, bsonBinarySensorState)
	if err != nil {
		return nil, err
	}

	// BinarySensor updated
	var updated types.BinarySensor
	err = res.Decode(&updated)
	if err != nil {
		return nil, err
	}
	log.Println("BinarySensor state update") /// TODO: debug
	return &updated, nil
}
