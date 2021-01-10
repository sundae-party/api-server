package light

import (
	context "context"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"sundae-party/api-server/pkg/apis/core/types"
	"sundae-party/api-server/pkg/storage"
)

type LightHandler struct {
	types.UnimplementedLightHandlerServer
	Store storage.Store
}

func (lh LightHandler) Get(ctx context.Context, l *types.Light) (*types.Light, error) {
	res, err := lh.Store.GetLightByName(ctx, l.Name)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (lh LightHandler) Create(ctx context.Context, l *types.Light) (*types.Light, error) {
	res, err := lh.Store.PutLight(ctx, l)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (lh LightHandler) Update(ctx context.Context, l *types.Light) (*types.Light, error) {
	return lh.Create(ctx, l)
}
func (lh LightHandler) Delete(ctx context.Context, l *types.Light) (*types.Light, error) {
	res, err := lh.Store.DeleteLight(ctx, l)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (lh LightHandler) GetAll(*types.GetAllRequest, types.LightHandler_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (lh LightHandler) WatchAll(r *types.GetAllRequest, stream types.LightHandler_WatchAllServer) error {
	lh.Store.GetEvent()

	go func() {
		for {
			select {
			case event := <-lh.Store.GetEvent():
				log.Printf("%s on %s", event.OperationType, event.Ns.Coll)
				// stream.Send(event)
			}
		}
	}()
	return nil
}
