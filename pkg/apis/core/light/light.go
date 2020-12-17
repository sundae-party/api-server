package light

import (
	context "context"
	"encoding/json"
	"fmt"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"sundae-party/api-server/pkg/apis/core/integration"
	"sundae-party/api-server/pkg/storage/etcd3"
)

type LightHandler struct {
	UnimplementedLightHandlerServer
	Store *etcd3.Store
}

func (lh *LightHandler) SetDesiredState(context.Context, *Light) (*Light, error) {
	// Update light desiredState in ETCD
	// Send request to concerned Integration
	// Bradcast to the websocket clients
	return nil, status.Errorf(codes.Unimplemented, "method SetDesiredState not implemented")
}
func (lh *LightHandler) SetState(context.Context, *Light) (*Light, error) {
	// Update light state in ETCD
	// Bradcast to the websocket clients
	return nil, status.Errorf(codes.Unimplemented, "method SetState not implemented")
}
func (lh *LightHandler) Create(c context.Context, l *Light) (*Light, error) {
	// Add light in ETCD
	json, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("/%s/%s", l.Integration.Name, l.Name)
	lh.Store.Put(key, string(json))
	return l, nil
	//return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (lh *LightHandler) GetByDevice(*integration.Integration, LightHandler_GetByDeviceServer) error {
	// Get Light from ETCD / ? Integration ? filtered by a
	return status.Errorf(codes.Unimplemented, "method GetByDevice not implemented")
}
func (lh *LightHandler) GetByIntegration(i *integration.Integration, h LightHandler_GetByIntegrationServer) error {
	//return status.Errorf(codes.Unimplemented, "method GetByIntegration not implemented")
	lights := lh.Store.GetByIntegration(i.Name)
	for _, jsonLight := range lights {
		lightObj := &Light{}
		err := json.Unmarshal([]byte(jsonLight), lightObj)
		if err != nil {
			return err
		}
		h.Send(lightObj)
	}
	return nil
}
