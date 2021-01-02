package integration

import (
	"context"
	"sundae-party/api-server/pkg/apis/core/types"
	"sundae-party/api-server/pkg/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IntegrationHandler struct {
	types.IntegrationHandlerServer
	Store storage.Store
}

func (ih IntegrationHandler) Create(ctx context.Context, i *types.Integration) (*types.Integration, error) {
	ni, err := ih.Store.PutIntegration(ctx, i)
	if err != nil {
		return nil, err
	}
	return ni, nil
}
func (ih IntegrationHandler) Get(ctx context.Context, ir *types.Integration) (*types.Integration, error) {
	i, err := ih.Store.GetIntegration(ctx, ir.Metadata.Name)
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
func (ih IntegrationHandler) Update(ctx context.Context, i *types.Integration) (*types.Integration, error) {
	ui, err := ih.Store.PutIntegration(ctx, i)
	if err != nil {
		return nil, err
	}

	return ui, nil
}
func (ih IntegrationHandler) Watch(i *types.Integration, stream types.IntegrationHandler_WatchServer) error {
	// Watch and send service call event only for given integration
	// for event := range ih.Store.GetEvent() {
	// 	if event.IntegrationName == i.Name {
	// 		stream.Send(event)
	// 	}
	// }
	// TODO: Check for stream connexion close
	return nil
}

func (ih IntegrationHandler) CallService(context.Context, *types.IntegrationServiceRequest) (*types.IntegrationServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallService not implemented")
}
func (ih IntegrationHandler) WatchService(*types.Integration, types.IntegrationHandler_WatchServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchService not implemented")
}
func (ih IntegrationHandler) StorePut(context.Context, *types.IntegrationStoreRequest) (*types.IntegrationStoreRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorePut not implemented")
}
func (ih IntegrationHandler) StoreGet(context.Context, *types.IntegrationStoreRequest) (*types.IntegrationStoreRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreGet not implemented")
}
