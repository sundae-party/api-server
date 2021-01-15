package integration

import (
	"context"
	"log"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IntegrationHandler struct {
	types.UnimplementedIntegrationHandlerServer
	Store storage.Store
}

func (ih IntegrationHandler) Create(ctx context.Context, i *types.Integration) (*types.Integration, error) {
	ni, err := ih.Store.PutIntegration(ctx, i)
	if err != nil {
		return nil, err
	}
	return ni, nil
}
func (ih IntegrationHandler) Get(ctx context.Context, ir *types.IntegrationServerRequest) (*types.Integration, error) {
	i, err := ih.Store.GetIntegration(ctx, ir.IntegrationName)
	if err != nil {
		return nil, err
	}
	return i, nil
}
func (ih IntegrationHandler) Delete(ctx context.Context, id *types.Integration) (*types.Integration, error) {
	resp, err := ih.Store.DeleteIntegration(ctx, id)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (ih IntegrationHandler) SetState(ctx context.Context, sr *types.SetIntegrationStateRequest) (*types.Integration, error) {
	i, err := ih.Store.GetIntegration(ctx, sr.IntegrationName)
	if err != nil {
		return nil, err
	}
	ui, err := ih.Store.PutIntegration(ctx, i)
	if err != nil {
		return nil, err
	}

	return ui, nil
}

func (ih IntegrationHandler) SetDesiredState(ctx context.Context, sr *types.SetIntegrationStateRequest) (*types.Integration, error) {
	i, err := ih.Store.GetIntegration(ctx, sr.IntegrationName)
	if err != nil {
		return nil, err
	}
	ui, err := ih.Store.PutIntegration(ctx, i)
	if err != nil {
		return nil, err
	}

	return ui, nil
}

func (ih IntegrationHandler) SubscribeEvents(integration *types.Integration, stream types.IntegrationHandler_SubscribeEventsServer) error {
	select {
	case event := <-ih.Store.GetIntegrationEvent():
		err := stream.Send(&event)
		if err != nil {
			log.Printf("Integration event emit error: %s", err)
		}
	}
	return nil
}

// TODO
func (ih IntegrationHandler) StorePut(ctx context.Context, req *types.IntegrationStoreRequest) (*types.IntegrationStoreRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorePut not implemented")
}

// TODO
func (ih IntegrationHandler) StoreGet(ctx context.Context, req *types.IntegrationStoreRequest) (*types.IntegrationStoreRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreGet not implemented")
}
