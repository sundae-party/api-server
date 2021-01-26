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

func buildSuntKey(sun *types.Sun) (key string, err error) {
	if sun.Integration != nil {
		if sun.Integration.Name != "" && sun.Name != "" {
			return fmt.Sprintf("%s/%s", sun.Integration.Name, sun.Name), nil
		}
		return "", errors.New("Invalid sun format, the integration name or sun name is missing.")
	}
	return "", errors.New("Invalid sun format, the integration infos is empty.")
}

//
//
// Sun store function
//
//

// PutSun create or update a Sun in the store.
func (ms MongoStore) PutSun(ctx context.Context, Sun *types.Sun) (*types.Sun, error) {

	key, err := buildSuntKey(Sun)
	if err != nil {
		return nil, err
	}

	// Ensure kind is set to sun
	Sun.Mutation = sunKind

	// Convert Sun to bson object
	bsonSun, err := bson.Marshal(Sun)

	// Put the Sun in the store
	res, err := ms.putEntity(ctx, key, bsonSun)
	if err != nil {
		return nil, err
	}

	var newSun types.Sun
	err = res.Decode(&newSun)
	if err != nil {
		return nil, err
	}

	return &newSun, nil
}

// GetAllSun return all Sun stored in the store
func (ms MongoStore) GetAllSun(c context.Context) ([]types.Sun, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	cursor, err := ms.getAllEntities(c, sunKind, "")
	if err != nil {
		return nil, err
	}

	var Suns []types.Sun
	err = cursor.All(ctx, &Suns)
	if err != nil {
		return nil, err
	}
	return Suns, nil
}

// Delete a Sun in the store
func (ms MongoStore) DeleteSun(ctx context.Context, Sun *types.Sun) (*types.Sun, error) {
	key := fmt.Sprintf("%s/%s", Sun.Integration.Name, Sun.Name)

	res, err := ms.deleteEntity(ctx, key)
	if err != nil {
		return nil, err
	}

	var deleted types.Sun
	err = res.Decode(&deleted)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// UpdateSunState will try to update the Sun state in the store, if the Sun is not present in the store it will be created
func (ms MongoStore) UpdateSunValue(ctx context.Context, Sun *types.Sun) (*types.Sun, error) {

	// Convert Sun value to bson object
	bsonSunState := bson.M{"state": Sun.State}

	key := fmt.Sprintf("%s/%s", Sun.Integration.Name, Sun.Name)

	// Try to update Sun
	res, err := ms.updateEntityState(ctx, key, bsonSunState)
	if err != nil {
		return nil, err
	}

	// Sun updated
	var updated types.Sun
	err = res.Decode(&updated)
	if err != nil {
		return nil, err
	}
	log.Println("Sun state update") /// TODO: debug
	return &updated, nil
}
