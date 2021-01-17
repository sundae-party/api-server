package integration

import (
	"context"
	"log"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	store_type "github.com/sundae-party/api-server/pkg/storage/types"
)

type IntegrationHandler struct {
	types.UnimplementedIntegrationHandlerServer
	Store storage.Store
}

func (ih IntegrationHandler) Create(ctx context.Context, i *types.Integration) (*types.Integration, error) {
	return ih.Store.PutIntegration(ctx, i)
}
func (ih IntegrationHandler) Get(ctx context.Context, ir *types.IntegrationServerRequest) (*types.Integration, error) {
	return ih.Store.GetIntegration(ctx, ir.IntegrationName)
}
func (ih IntegrationHandler) Delete(ctx context.Context, id *types.Integration) (*types.Integration, error) {
	resp, err := ih.Store.DeleteIntegration(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (ih IntegrationHandler) SetState(ctx context.Context, sr *types.SetIntegrationStateRequest) (*types.Integration, error) {
	intergration := &types.Integration{
		Name:  sr.IntegrationName,
		State: sr.State,
	}
	return ih.Store.UpdateIntegrationState(ctx, intergration)
}

func (ih IntegrationHandler) SetDesiredState(ctx context.Context, sr *types.SetIntegrationStateRequest) (*types.Integration, error) {
	intergration := &types.Integration{
		Name:         sr.IntegrationName,
		DesiredState: sr.State,
	}
	return ih.Store.UpdateIntegrationDesiredState(ctx, intergration)
}

func (ih IntegrationHandler) SubscribeEvents(integration *types.Integration, stream types.IntegrationHandler_SubscribeEventsServer) error {

	cs, err := ih.Store.GetIntegrationEvent(stream.Context())
	if err != nil {
		return err
	}

	_, err = ih.Store.PutIntegration(stream.Context(), integration)
	if err != nil {
		return err
	}

	for {
		select {
		case <-stream.Context().Done():
			cs.Close(stream.Context())
			return nil
		default:
			if cs.Next(stream.Context()) {
				var event store_type.IntegrationEvent
				err := cs.Decode(&event)
				if err != nil {
					log.Println("Mongo store -- decode mongo change stream integration event error:")
					log.Println(err)
				}
				err = stream.Send(&event.FullDocument)
				if err != nil {
					log.Printf("Integration event emit error: %s", err)
				}
			}
		}
	}

}

// TODO
func (ih IntegrationHandler) StorePut(ctx context.Context, req *types.IntegrationStoreRequest) (*types.IntegrationStoreRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorePut not implemented")
}

// TODO
func (ih IntegrationHandler) StoreGet(ctx context.Context, req *types.IntegrationStoreRequest) (*types.IntegrationStoreRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreGet not implemented")
}
