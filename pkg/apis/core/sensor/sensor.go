package sensor

import (
	context "context"
	"fmt"
	"log"

	"github.com/sundae-party/api-server/pkg/apis/core/types"
	"github.com/sundae-party/api-server/pkg/storage"

	store_type "github.com/sundae-party/api-server/pkg/storage/types"
)

type SensorHandler struct {
	types.UnimplementedSensorHandlerServer
	Store storage.Store
}

func (sh SensorHandler) Get(ctx context.Context, s *types.Sensor) (*types.Sensor, error) {
	key := fmt.Sprintf("%s/%s", s.Integration.Name, s.Name)
	return sh.Store.GetSensorByName(ctx, key)
}
func (sh SensorHandler) Create(ctx context.Context, s *types.Sensor) (*types.Sensor, error) {
	return sh.Store.PutSensor(ctx, s)
}
func (sh SensorHandler) Update(ctx context.Context, s *types.Sensor) (*types.Sensor, error) {
	return sh.Store.PutSensor(ctx, s)
}
func (sh SensorHandler) Delete(ctx context.Context, s *types.Sensor) (*types.Sensor, error) {
	return sh.Store.DeleteSensor(ctx, s)
}
func (sh SensorHandler) GetAll(r *types.GetAllRequest, stream types.SensorHandler_GetAllServer) error {
	sensors, err := sh.Store.GetAllSensor(stream.Context())
	if err != nil {
		return err
	}

	for _, s := range sensors {
		err := stream.Send(&s)
		if err != nil {
			fmt.Printf("Error sending sensors to getAllSensor stream: %s\n", err)
		}
	}
	return nil
}

func (sh SensorHandler) WatchAll(s *types.GetAllRequest, stream types.SensorHandler_WatchAllServer) error {
	cs, err := sh.Store.GetEntityEvent(stream.Context(), "sensor")
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
				var event store_type.SensorEvent
				err := cs.Decode(&event)
				if err != nil {
					log.Println("Mongo store -- decode mongo change stream sensor event error:")
					log.Println(err)
				}
				err = stream.Send(&event.FullDocument)
				if err != nil {
					log.Printf("Sensor event emit error: %s", err)
				}
			}
		}
	}
}
func (sh SensorHandler) SetValue(ctx context.Context, r *types.SetSensorValueRequest) (*types.Sensor, error) {

	// update value
	sensor := &types.Sensor{
		Name: r.SensorName,
		Integration: &types.Integration{
			Name: r.IntegrationName,
		},
		Value: r.Value,
	}
	log.Printf("%s", sensor)
	return sh.Store.UpdateSensorValue(ctx, sensor)
}
