package types

type Entity struct {
	Kind     string
	Metadata EntityMeta
	Spec     interface{}
	Uuid     string
}

type EntityMeta struct {
	Name        string
	Integration Integration
}

// func (e *Entity) Decode(val []byte, ) (error) {
// 	switch e.Kind {
// 	case "Light":
// 		return e.Spec.(*Light), nil
// 	// case "Sensor":
// 	// 	return e.Spec.(*Sensor), nil
// 	default:
// 		return nil, errors.New("Object type not reconized")
// 	}
// }
