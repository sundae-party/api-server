package light

import (
	context "context"
	"encoding/json"
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
	light, err := lh.Store.GetLight(ctx, l.Name)
	if err != nil {
		return nil, err
	}
	return light, nil
}
func (lh LightHandler) Create(ctx context.Context, l *types.Light) (*types.Light, error) {
	light, err := lh.Store.PutLight(ctx, l)
	if err != nil {
		return nil, err
	}
	return light, nil
}
func (lh LightHandler) Update(ctx context.Context, l *types.Light) (*types.Light, error) {
	light, err := lh.Store.PutLight(ctx, l)
	if err != nil {
		return nil, err
	}
	return light, nil
}
func (lh LightHandler) Delete(ctx context.Context, l *types.Light) (*types.Light, error) {
	light, err := lh.Store.DeleteLight(ctx, l)
	if err != nil {
		return nil, err
	}
	return light, nil
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
				msg := &storage.EventMessage{}
				json.Unmarshal([]byte(event), msg)
				log.Printf("%s on %s", msg.OperationType, msg.Ns.Coll)
				// stream.Send(event)
			}
		}
	}()
	return nil
}

// func (lh *LightHandler) SetDesiredState(context.Context, *types.Light) (*types.Light, error) {
// 	// Update light desiredState in ETCD
// 	// Send request to concerned Integration
// 	// Bradcast to the websocket clients
// 	return nil, status.Errorf(codes.Unimplemented, "method SetDesiredState not implemented")
// }
// func (lh *LightHandler) SetState(context.Context, *types.Light) (*types.Light, error) {
// 	// Update light state in ETCD
// 	// Bradcast to the websocket clients
// 	return nil, status.Errorf(codes.Unimplemented, "method SetState not implemented")
// }
// func (lh *LightHandler) Create(c context.Context, l *types.Light) (*types.Light, error) {
// 	// Add light in ETCD
// 	// json, err := json.Marshal(l)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// key := fmt.Sprintf("/%s/%s", l.Integration.Name, l.Name)
// 	// lh.Store.Put(key, string(json))
// 	return l, nil
// 	//return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
// }
// func (lh *LightHandler) GetByDevice(*types.Integration, types.LightHandler_GetByDeviceServer) error {
// 	// Get Light from ETCD / ? Integration ? filtered by a
// 	return status.Errorf(codes.Unimplemented, "method GetByDevice not implemented")
// }
// func (lh *LightHandler) GetByIntegration(i *types.Integration, h types.LightHandler_GetByIntegrationServer) error {
// 	//return status.Errorf(codes.Unimplemented, "method GetByIntegration not implemented")
// 	// lights := lh.Store.GetByIntegration(i.Name)
// 	// for _, jsonLight := range lights {
// 	// 	lightObj := &Light{}
// 	// 	err := json.Unmarshal([]byte(jsonLight), lightObj)
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}
// 	// 	h.Send(lightObj)
// 	// }
// 	return nil
// }
