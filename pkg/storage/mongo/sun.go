package mongo

import (
	"context"
	"log"

	"github.com/sundae-party/api-server/pkg/apis/core/types"

	"go.mongodb.org/mongo-driver/bson"
)

const sunKey = "sun/sun"

//
//
// Sun store function
//
//

// PutSun create or update the Sun sensor in the store for the home.
func (ms MongoStore) PutSun(ctx context.Context, sunState *types.SunState, integration *types.Integration) (*types.Sun, error) {

	sun := &types.Sun{
		Name:          "sun",
		Integration:   integration,
		DisplayedName: "sun",
		State:         sunState,
		Mutation:      sunKind,
	}

	// Convert Sun to bson object
	bsonSun, err := bson.Marshal(sun)

	// Put the Sun in the store
	res, err := ms.putEntity(ctx, sunKey, bsonSun)
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

// GetSun return the home Sun sensor stored in the store
func (ms MongoStore) GetSun(c context.Context) (*types.Sun, error) {

	res, err := ms.getEntityByName(c, sunKey)
	if err != nil {
		return nil, err
	}

	var sun *types.Sun
	err = res.Decode(&sun)
	if err != nil {
		return nil, err
	}
	return sun, nil
}

// Delete the home Sun in the store
func (ms MongoStore) DeleteSun(ctx context.Context) (*types.Sun, error) {

	res, err := ms.deleteEntity(ctx, sunKey)
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

// UpdateSunState will try to update the Sun state in the store
func (ms MongoStore) UpdateSunState(ctx context.Context, state *types.SunState) (*types.Sun, error) {

	// Convert Sun value to bson object
	bsonSunState := bson.M{"state": state}

	// Try to update Sun
	res, err := ms.updateEntityState(ctx, sunKey, bsonSunState)
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
