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
