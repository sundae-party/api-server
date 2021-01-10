package types

type StoreEvent struct {
	OperationType string
	FullDocument  interface{}
	Ns            struct {
		Db   string
		Coll string
	}
}
