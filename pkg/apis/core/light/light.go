package light

import (
	context "context"
	"log"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/storage"
	store_type "github.com/sundae-party/api-server/pkg/storage/types"
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
func (lh LightHandler) GetAll(_ *types.GetAllRequest, stream types.LightHandler_GetAllServer) error {
	lights, err := lh.Store.GetAllLight(stream.Context())
	if err != nil {
		return nil
	}
	for _, light := range lights {
		stream.Send(&light)
	}
	return nil
}
func (lh LightHandler) WatchAll(r *types.GetAllRequest, stream types.LightHandler_WatchAllServer) error {
	cs, err := lh.Store.GetEntityEvent(stream.Context(), "light")
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
				var event store_type.LightEvent
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

func (lh LightHandler) SetDesiredState(ctx context.Context, lsr *types.SetLightStateRequest) (*types.Light, error) {
	lightRequest := &types.Light{
		Name:            lsr.LightName,
		IntegrationName: lsr.IntegrationName,
		DesiredState:    lsr.State,
	}
	return lh.Store.UpdateLightStateDesiredState(ctx, lightRequest)
}
func (lh LightHandler) SetState(ctx context.Context, lsr *types.SetLightStateRequest) (*types.Light, error) {
	lightRequest := &types.Light{
		Name:            lsr.LightName,
		IntegrationName: lsr.IntegrationName,
		DesiredState:    lsr.State,
	}
	return lh.Store.UpdateLightState(ctx, lightRequest)

}
