// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: light.proto

package types

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LightState_FlashMode int32

const (
	LightState_SHORT LightState_FlashMode = 0
	LightState_LONG  LightState_FlashMode = 1
)

// Enum value maps for LightState_FlashMode.
var (
	LightState_FlashMode_name = map[int32]string{
		0: "SHORT",
		1: "LONG",
	}
	LightState_FlashMode_value = map[string]int32{
		"SHORT": 0,
		"LONG":  1,
	}
)

func (x LightState_FlashMode) Enum() *LightState_FlashMode {
	p := new(LightState_FlashMode)
	*p = x
	return p
}

func (x LightState_FlashMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LightState_FlashMode) Descriptor() protoreflect.EnumDescriptor {
	return file_light_proto_enumTypes[0].Descriptor()
}

func (LightState_FlashMode) Type() protoreflect.EnumType {
	return &file_light_proto_enumTypes[0]
}

func (x LightState_FlashMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LightState_FlashMode.Descriptor instead.
func (LightState_FlashMode) EnumDescriptor() ([]byte, []int) {
	return file_light_proto_rawDescGZIP(), []int{4, 0}
}

type SetLightStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LightName       string      `protobuf:"bytes,1,opt,name=lightName,proto3" json:"lightName,omitempty"`
	IntegrationName string      `protobuf:"bytes,2,opt,name=integrationName,proto3" json:"integrationName,omitempty"`
	State           *LightState `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SetLightStateRequest) Reset() {
	*x = SetLightStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLightStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLightStateRequest) ProtoMessage() {}

func (x *SetLightStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_light_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLightStateRequest.ProtoReflect.Descriptor instead.
func (*SetLightStateRequest) Descriptor() ([]byte, []int) {
	return file_light_proto_rawDescGZIP(), []int{0}
}

func (x *SetLightStateRequest) GetLightName() string {
	if x != nil {
		return x.LightName
	}
	return ""
}

func (x *SetLightStateRequest) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

func (x *SetLightStateRequest) GetState() *LightState {
	if x != nil {
		return x.State
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_light_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_light_proto_rawDescGZIP(), []int{1}
}

type Light struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Integration   *Integration `protobuf:"bytes,2,opt,name=integration,proto3" json:"integration,omitempty"`
	Device        string       `protobuf:"bytes,3,opt,name=Device,proto3" json:"Device,omitempty"`
	DisplayedName string       `protobuf:"bytes,4,opt,name=displayed_name,json=displayedName,proto3" json:"displayed_name,omitempty"`
	Room          string       `protobuf:"bytes,5,opt,name=room,proto3" json:"room,omitempty"`
	DesiredState  *LightState  `protobuf:"bytes,6,opt,name=desired_state,json=desiredState,proto3" json:"desired_state,omitempty"`
	State         *LightState  `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Mutation      string       `protobuf:"bytes,8,opt,name=mutation,proto3" json:"mutation,omitempty"`
}

func (x *Light) Reset() {
	*x = Light{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Light) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Light) ProtoMessage() {}

func (x *Light) ProtoReflect() protoreflect.Message {
	mi := &file_light_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Light.ProtoReflect.Descriptor instead.
func (*Light) Descriptor() ([]byte, []int) {
	return file_light_proto_rawDescGZIP(), []int{2}
}

func (x *Light) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Light) GetIntegration() *Integration {
	if x != nil {
		return x.Integration
	}
	return nil
}

func (x *Light) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Light) GetDisplayedName() string {
	if x != nil {
		return x.DisplayedName
	}
	return ""
}

func (x *Light) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *Light) GetDesiredState() *LightState {
	if x != nil {
		return x.DesiredState
	}
	return nil
}

func (x *Light) GetState() *LightState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Light) GetMutation() string {
	if x != nil {
		return x.Mutation
	}
	return ""
}

type LightColorRGB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Red   int32 `protobuf:"varint,1,opt,name=red,proto3" json:"red,omitempty"`
	Green int32 `protobuf:"varint,2,opt,name=green,proto3" json:"green,omitempty"`
	Blue  int32 `protobuf:"varint,3,opt,name=blue,proto3" json:"blue,omitempty"`
}

func (x *LightColorRGB) Reset() {
	*x = LightColorRGB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightColorRGB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightColorRGB) ProtoMessage() {}

func (x *LightColorRGB) ProtoReflect() protoreflect.Message {
	mi := &file_light_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightColorRGB.ProtoReflect.Descriptor instead.
func (*LightColorRGB) Descriptor() ([]byte, []int) {
	return file_light_proto_rawDescGZIP(), []int{3}
}

func (x *LightColorRGB) GetRed() int32 {
	if x != nil {
		return x.Red
	}
	return 0
}

func (x *LightColorRGB) GetGreen() int32 {
	if x != nil {
		return x.Green
	}
	return 0
}

func (x *LightColorRGB) GetBlue() int32 {
	if x != nil {
		return x.Blue
	}
	return 0
}

type LightState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brightness int32                `protobuf:"varint,1,opt,name=Brightness,proto3" json:"Brightness,omitempty"`
	ColorRGB   *LightColorRGB       `protobuf:"bytes,2,opt,name=colorRGB,proto3" json:"colorRGB,omitempty"`
	On         bool                 `protobuf:"varint,3,opt,name=on,proto3" json:"on,omitempty"`
	Kelvin     int32                `protobuf:"varint,4,opt,name=kelvin,proto3" json:"kelvin,omitempty"`
	FlashMode  LightState_FlashMode `protobuf:"varint,5,opt,name=flashMode,proto3,enum=types.LightState_FlashMode" json:"flashMode,omitempty"`
	Transition int32                `protobuf:"varint,6,opt,name=transition,proto3" json:"transition,omitempty"`
}

func (x *LightState) Reset() {
	*x = LightState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_light_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightState) ProtoMessage() {}

func (x *LightState) ProtoReflect() protoreflect.Message {
	mi := &file_light_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightState.ProtoReflect.Descriptor instead.
func (*LightState) Descriptor() ([]byte, []int) {
	return file_light_proto_rawDescGZIP(), []int{4}
}

func (x *LightState) GetBrightness() int32 {
	if x != nil {
		return x.Brightness
	}
	return 0
}

func (x *LightState) GetColorRGB() *LightColorRGB {
	if x != nil {
		return x.ColorRGB
	}
	return nil
}

func (x *LightState) GetOn() bool {
	if x != nil {
		return x.On
	}
	return false
}

func (x *LightState) GetKelvin() int32 {
	if x != nil {
		return x.Kelvin
	}
	return 0
}

func (x *LightState) GetFlashMode() LightState_FlashMode {
	if x != nil {
		return x.FlashMode
	}
	return LightState_SHORT
}

func (x *LightState) GetTransition() int32 {
	if x != nil {
		return x.Transition
	}
	return 0
}

var File_light_proto protoreflect.FileDescriptor

var file_light_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xa1, 0x02, 0x0a, 0x05, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x34, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x36, 0x0a, 0x0d, 0x64, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x0d, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x47, 0x42, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x65,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62,
	0x6c, 0x75, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x0a, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x42, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x47, 0x42, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67,
	0x68, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x47, 0x42, 0x52, 0x08, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x52, 0x47, 0x42, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x6c, 0x76, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6b, 0x65, 0x6c, 0x76, 0x69, 0x6e, 0x12, 0x39, 0x0a, 0x09,
	0x66, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x66, 0x6c,
	0x61, 0x73, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x09, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x48, 0x4f, 0x52, 0x54, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x01, 0x32, 0xfa, 0x02, 0x0a, 0x0c, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x1a,
	0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4c, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68,
	0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x2e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x30, 0x01, 0x12,
	0x30, 0x0a, 0x08, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x30,
	0x01, 0x12, 0x3c, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74,
	0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x35, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x65, 0x2d, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_light_proto_rawDescOnce sync.Once
	file_light_proto_rawDescData = file_light_proto_rawDesc
)

func file_light_proto_rawDescGZIP() []byte {
	file_light_proto_rawDescOnce.Do(func() {
		file_light_proto_rawDescData = protoimpl.X.CompressGZIP(file_light_proto_rawDescData)
	})
	return file_light_proto_rawDescData
}

var file_light_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_light_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_light_proto_goTypes = []interface{}{
	(LightState_FlashMode)(0),    // 0: types.LightState.FlashMode
	(*SetLightStateRequest)(nil), // 1: types.SetLightStateRequest
	(*GetAllRequest)(nil),        // 2: types.GetAllRequest
	(*Light)(nil),                // 3: types.Light
	(*LightColorRGB)(nil),        // 4: types.LightColorRGB
	(*LightState)(nil),           // 5: types.LightState
	(*Integration)(nil),          // 6: types.Integration
}
var file_light_proto_depIdxs = []int32{
	5,  // 0: types.SetLightStateRequest.state:type_name -> types.LightState
	6,  // 1: types.Light.integration:type_name -> types.Integration
	5,  // 2: types.Light.desired_state:type_name -> types.LightState
	5,  // 3: types.Light.state:type_name -> types.LightState
	4,  // 4: types.LightState.colorRGB:type_name -> types.LightColorRGB
	0,  // 5: types.LightState.flashMode:type_name -> types.LightState.FlashMode
	3,  // 6: types.LightHandler.Get:input_type -> types.Light
	3,  // 7: types.LightHandler.Create:input_type -> types.Light
	3,  // 8: types.LightHandler.Update:input_type -> types.Light
	3,  // 9: types.LightHandler.Delete:input_type -> types.Light
	2,  // 10: types.LightHandler.GetAll:input_type -> types.GetAllRequest
	2,  // 11: types.LightHandler.WatchAll:input_type -> types.GetAllRequest
	1,  // 12: types.LightHandler.SetDesiredState:input_type -> types.SetLightStateRequest
	1,  // 13: types.LightHandler.setState:input_type -> types.SetLightStateRequest
	3,  // 14: types.LightHandler.Get:output_type -> types.Light
	3,  // 15: types.LightHandler.Create:output_type -> types.Light
	3,  // 16: types.LightHandler.Update:output_type -> types.Light
	3,  // 17: types.LightHandler.Delete:output_type -> types.Light
	3,  // 18: types.LightHandler.GetAll:output_type -> types.Light
	3,  // 19: types.LightHandler.WatchAll:output_type -> types.Light
	3,  // 20: types.LightHandler.SetDesiredState:output_type -> types.Light
	3,  // 21: types.LightHandler.setState:output_type -> types.Light
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_light_proto_init() }
func file_light_proto_init() {
	if File_light_proto != nil {
		return
	}
	file_integration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_light_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLightStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_light_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_light_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Light); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_light_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightColorRGB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_light_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_light_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_light_proto_goTypes,
		DependencyIndexes: file_light_proto_depIdxs,
		EnumInfos:         file_light_proto_enumTypes,
		MessageInfos:      file_light_proto_msgTypes,
	}.Build()
	File_light_proto = out.File
	file_light_proto_rawDesc = nil
	file_light_proto_goTypes = nil
	file_light_proto_depIdxs = nil
}
