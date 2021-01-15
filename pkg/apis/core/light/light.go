package light

import (
	context "context"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/storage"
)

type LightHandler struct {
	types.UnimplementedLightHandlerServer
	Store storage.Store
}

func (lh LightHandler) Get(ctx context.Context, l *types.Light) (*types.Light, error) {
	return lh.Store.GetLightByName(ctx, l.Name)
}
func (lh LightHandler) Create(ctx context.Context, l *types.Light) (*types.Light, error) {
	return lh.Store.PutLight(ctx, l)
}
func (lh LightHandler) Update(ctx context.Context, l *types.Light) (*types.Light, error) {
	return lh.Store.UpdateLightState(ctx, l)
}
func (lh LightHandler) Delete(ctx context.Context, l *types.Light) (*types.Light, error) {
	return lh.Store.DeleteLight(ctx, l)
}
func (lh LightHandler) GetAll(*types.GetAllRequest, types.LightHandler_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (lh LightHandler) WatchAll(r *types.GetAllRequest, stream types.LightHandler_WatchAllServer) error {

	go func() {
		for {
			select {
			case event := <-lh.Store.GetAllEvent():
				log.Printf("%s on %s", event.OperationType, event.Ns.Coll)
				// stream.Send(event)
			}
		}
	}()
	return nil
}
