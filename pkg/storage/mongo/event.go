package mongo

import (
	"github.com/sundae-party/api-server/pkg/apis/core/types"
	store_type "github.com/sundae-party/api-server/pkg/storage/types"
)

func (ms MongoStore) GetAllEvent() chan store_type.StoreEvent {
	return ms.Event
}

// TODO: GetIntegrationEvent
// Should watch for store event
// for each of them define with the mutation field in the full document fild the object type (Integration, light, sensor, ...)
func (ms MongoStore) GetIntegrationEvent() chan types.Integration {
	integrationChan := make(chan types.Integration)
	go func() {
		select {
		case event := <-ms.GetAllEvent():
			if event.Ns.Coll == integrationCollection {
				integrationChan <- event.FullDocument.(types.Integration)
			}
		}
	}()
	return integrationChan
}
