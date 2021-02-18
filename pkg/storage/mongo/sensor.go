package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sundae-party/api-server/pkg/apis/core/types"

	"go.mongodb.org/mongo-driver/bson"
)

func buildSensortKey(sensor *types.Sensor) (key string, err error) {
	if sensor.IntegrationName != "" && sensor.Name != "" {
		return fmt.Sprintf("%s/%s", sensor.IntegrationName, sensor.Name), nil
	}
	return "", errors.New("Invalid sensor format, the integration name or sensor name is missing.")
}

//
//
// Sensor store function
//
//

// PutSensor create or update a Sensor in the store.
func (ms MongoStore) PutSensor(ctx context.Context, Sensor *types.Sensor) (*types.Sensor, error) {

	key, err := buildSensortKey(Sensor)
	if err != nil {
		return nil, err
	}

	// Ensure kind is set to sensor
	Sensor.Mutation = sensorKind

	// Convert Sensor to bson object
	bsonSensor, err := bson.Marshal(Sensor)

	// Put the Sensor in the store
	res, err := ms.putEntity(ctx, key, bsonSensor)
	if err != nil {
		return nil, err
	}

	var newSensor types.Sensor
	err = res.Decode(&newSensor)
	if err != nil {
		return nil, err
	}

	return &newSensor, nil
}

// The key must be formated with "intergrationName/SensorName" and can't be empty
func (ms MongoStore) GetSensorByName(ctx context.Context, key string) (*types.Sensor, error) {

	res, err := ms.getEntityByName(ctx, key)
	if err != nil {
		return nil, err
	}

	// Convert bson result to Sensor object
	var Sensor types.Sensor
	err = res.Decode(&Sensor)
	if err != nil {
		return nil, err
	}

	return &Sensor, nil
}

// GetAllSensor return all Sensor stored in the store
func (ms MongoStore) GetAllSensor(c context.Context) ([]types.Sensor, error) {
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()

	cursor, err := ms.getAllEntities(c, sensorKind, "")
	if err != nil {
		return nil, err
	}

	var Sensors []types.Sensor
	err = cursor.All(ctx, &Sensors)
	if err != nil {
		return nil, err
	}
	return Sensors, nil
}

// GetSensorByIntegration return a list of Sensor
func (ms MongoStore) GetSensorByIntegration(c context.Context, integrationName string) ([]types.Sensor, error) {
	cursor, err := ms.getAllEntities(c, sensorKind, integrationName)
	if err != nil {
		return nil, err
	}

	var Sensors []types.Sensor
	err = cursor.Decode(Sensors)
	if err != nil {
		return nil, err
	}

	return Sensors, nil
}

// Delete a Sensor in the store
func (ms MongoStore) DeleteSensor(ctx context.Context, Sensor *types.Sensor) (*types.Sensor, error) {
	key := fmt.Sprintf("%s/%s", Sensor.IntegrationName, Sensor.Name)

	res, err := ms.deleteEntity(ctx, key)
	if err != nil {
		return nil, err
	}

	var deleted types.Sensor
	err = res.Decode(&deleted)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// UpdateSensorState will try to update the Sensor state in the store, if the Sensor is not present in the store it will be created
func (ms MongoStore) UpdateSensorValue(ctx context.Context, Sensor *types.Sensor) (*types.Sensor, error) {

	// Convert Sensor value to bson object
	bsonSensorState := bson.M{"value": Sensor.Value}

	key := fmt.Sprintf("%s/%s", Sensor.IntegrationName, Sensor.Name)

	// Try to update Sensor
	res, err := ms.updateEntityState(ctx, key, bsonSensorState)
	if err != nil {
		return nil, err
	}

	// Sensor updated
	var updated types.Sensor
	err = res.Decode(&updated)
	if err != nil {
		return nil, err
	}
	log.Println("Sensor value update") /// TODO: debug
	return &updated, nil
}
