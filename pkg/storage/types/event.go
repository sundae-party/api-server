package types

import (
	"github.com/sundae-party/api-server/pkg/apis/core/types"
)

type StoreEvent struct {
	OperationType string
	FullDocument  interface{}
}

type IntegrationEvent struct {
	OperationType string
	FullDocument  types.Integration
}

type LightEvent struct {
	OperationType string
	FullDocument  types.Light
}

type SensorEvent struct {
	OperationType string
	FullDocument  types.Sensor
}

type SunEvent struct {
	OperationType string
	FullDocument  types.Sun
}
