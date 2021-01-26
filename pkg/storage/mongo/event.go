package mongo

import (
	"context"
	"errors"

	store_type "github.com/sundae-party/api-server/pkg/storage/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (ms MongoStore) GetAllEvent() chan store_type.StoreEvent {
	return ms.Event
}

// TODO: do not return mongo change stream but event chan IntegrationEvent and chan close
// chan IntegrationEvent use to get entity/integration event
// chan close use to close mong change stream when gRPC/socket/... is closed

// GetIntegrationEvent return a mongo Change Stream filtered on integration
func (ms MongoStore) GetIntegrationEvent(ctx context.Context) (*mongo.ChangeStream, error) {

	cs, err := ms.DataBase.Collection("integrations").Watch(ctx, mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		return nil, err
	}

	return cs, nil

}

// GetEntityEvent return a mongo Change Stream filtered on kind
func (ms MongoStore) GetEntityEvent(ctx context.Context, kind string) (*mongo.ChangeStream, error) {

	// TODO kind empty = change stream on all entity for json object (ws)
	if kind == "" {
		return nil, errors.New("Kind can't be empty")
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "fullDocument.mutation", Value: kind}}}},
	}

	cs, err := ms.DataBase.Collection("entities").Watch(ctx, pipeline, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		return nil, err
	}

	return cs, nil

}
