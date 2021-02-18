package sun

import (
	context "context"
	"log"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/storage"
	store_type "github.com/sundae-party/api-server/pkg/storage/types"
)

type SunHandler struct {
	types.UnimplementedSunHandlerServer
	Store storage.Store
}

func (sh SunHandler) Get(ctx context.Context, sr *types.SunRequest) (*types.Sun, error) {
	return sh.Store.GetSun(ctx)
}
func (sh SunHandler) Create(ctx context.Context, sun *types.Sun) (*types.Sun, error) {
	return sh.Store.PutSun(ctx, sun.GetState())
}
func (sh SunHandler) Delete(ctx context.Context, sr *types.SunRequest) (*types.Sun, error) {
	return sh.Store.DeleteSun(ctx)
}
func (sh SunHandler) Watch(sr *types.SunRequest, stream types.SunHandler_WatchServer) error {
	cs, err := sh.Store.GetEntityEvent(stream.Context(), "sun")
	if err != nil {
		return err
	}

	for {
		select {
		// if gRPC stream is closed, close mongo change stream
		case <-stream.Context().Done():
			cs.Close(stream.Context())
			return nil
		default:
			if cs.Next(stream.Context()) {
				var event store_type.SunEvent
				err := cs.Decode(&event)
				if err != nil {
					log.Println("Mongo store -- decode mongo change stream sun event error:")
					log.Println(err)
				}
				err = stream.Send(&event.FullDocument)
				if err != nil {
					log.Printf("Sun event emit error: %s", err)
				}
			}
		}
	}
}
func (sh SunHandler) SetState(ctx context.Context, state *types.SunState) (*types.Sun, error) {
	return sh.Store.UpdateSunState(ctx, state)
}
