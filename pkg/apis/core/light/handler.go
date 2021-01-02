package light

import (
	"context"
	"sundae-party/api-server/pkg/apis/core/entity"
	"sundae-party/api-server/pkg/apis/core/types"

	"sundae-party/api-server/pkg/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LightHandler struct {
	types.EntityHandlerServer
	Store storage.Store
}

func (lh LightHandler) Create(ctx context.Context, l *types.Light) (*types.Light, error) {
	lh.Store.PutEntity(ctx, l, types.Light)
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (lh LightHandler) Update(ctx context.Context, l *types.Light) (*types.Light, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (lh LightHandler) Delete(ctx context.Context, l *types.Light) (*types.Light, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (lh LightHandler) Watch(l *types.Light, s entity.EntityHandler_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (lh LightHandler) GetByIntegration(i *types.Integration, s entity.EntityHandler_GetByIntegrationServer) error {
	return status.Errorf(codes.Unimplemented, "method GetByIntegration not implemented")
}
func (lh LightHandler) Get(r *entity.GetEntitiesRequest, s entity.EntityHandler_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
