package integration

import (
	"context"
	"sundae-party/api-server/pkg/apis/core/types"
	"sundae-party/api-server/pkg/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IntegrationHandler struct {
	types.UnimplementedIntegrationHandlerServer
	Store        storage.Store
	ServiceEvent chan *types.CallIntegrationServiceRequest
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
func (ih IntegrationHandler) SetState(ctx context.Context, sr *types.SetIntegrationStateRequest) (*types.IntegrationState, error) {
	i, err := ih.Store.GetIntegration(ctx, sr.IntegrationName)
	if err != nil {
		return nil, err
	}
	ui, err := ih.Store.PutIntegration(ctx, i)
	if err != nil {
		return nil, err
	}

	return ui.State, nil
}

// TODO: how to contact desired integration
func (ih IntegrationHandler) CallService(ctx context.Context, svc *types.CallIntegrationServiceRequest) (*types.CallIntegrationServiceResponse, error) {
	ih.ServiceEvent <- svc
	// TODO: how to get Call service response status
	return &types.CallIntegrationServiceResponse{Success: true}, nil
	// svc.IntegrationName
	// svc.Service.Name
	// svc.Service.Data
	// return nil, status.Errorf(codes.Unimplemented, "method CallService not implemented")
}

func (ih IntegrationHandler) Connect(i *types.Integration, stream types.IntegrationHandler_ConnectServer) error {
	// TODO: Need to create new context, no context in param
	ctx := context.Background()
	_, err := ih.Store.PutIntegration(ctx, i)
	if err != nil {
		return err
	}
	// Watch and send service call event only for given integration
	for event := range ih.ServiceEvent {
		if event.IntegrationName == i.Name {
			stream.Send(event)
		}
	}
	// TODO: Check for stream connexion close
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
