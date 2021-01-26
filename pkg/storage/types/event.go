package types

import (
	"github.com/sundae-party/api-server/pkg/apis/core/types"
)

type StoreEvent struct {
	OperationType string
	FullDocument  interface{}
	Ns            struct {
		Db   string
		Coll string
	}
}

type IntegrationEvent struct {
	OperationType string
	FullDocument  types.Integration
	Ns            struct {
		Db   string
		Coll string
	}
}

type LightEvent struct {
	OperationType string
	FullDocument  types.Light
	Ns            struct {
		Db   string
		Coll string
	}
}

type SunEvent struct {
	OperationType string
	FullDocument  types.Sun
	Ns            struct {
		Db   string
		Coll string
	}
}
