package mongo

import (
	"context"

	store_type "github.com/sundae-party/api-server/pkg/storage/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (ms MongoStore) GetAllEvent() chan store_type.StoreEvent {
	return ms.Event
}

// TODO: GetIntegrationEvent
// Should watch for store event
// for each of them define with the mutation field in the full document fild the object type (Integration, light, sensor, ...)
func (ms MongoStore) GetIntegrationEvent(ctx context.Context) (*mongo.ChangeStream, error) {

	cs, err := ms.DataBase.Collection("integrations").Watch(ctx, mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		return nil, err
	}

	return cs, nil

}
