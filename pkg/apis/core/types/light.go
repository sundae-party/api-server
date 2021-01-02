package types

type Light struct {
	Entity
	State        LightState // Overide spec with light spec
	DesiredState LightState // Overide status with light status
}

type ColorRGB struct {
	Red   int32
	Green int32
	Blue  int32
}

type LightState struct {
	Brightness int32
	ColorRGB   ColorRGB
	On         bool
	Kelvin     int32
	FlashMode  int32
}
