// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: integration.proto

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

type IntegrationStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationName string `protobuf:"bytes,1,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	Key             string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value           string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IntegrationStoreRequest) Reset() {
	*x = IntegrationStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationStoreRequest) ProtoMessage() {}

func (x *IntegrationStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationStoreRequest.ProtoReflect.Descriptor instead.
func (*IntegrationStoreRequest) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{0}
}

func (x *IntegrationStoreRequest) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

func (x *IntegrationStoreRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IntegrationStoreRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type IntegrationServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationName string `protobuf:"bytes,1,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
}

func (x *IntegrationServerRequest) Reset() {
	*x = IntegrationServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationServerRequest) ProtoMessage() {}

func (x *IntegrationServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationServerRequest.ProtoReflect.Descriptor instead.
func (*IntegrationServerRequest) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{1}
}

func (x *IntegrationServerRequest) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

type CallIntegrationServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationName string              `protobuf:"bytes,1,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	Service         *IntegrationService `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *CallIntegrationServiceRequest) Reset() {
	*x = CallIntegrationServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallIntegrationServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallIntegrationServiceRequest) ProtoMessage() {}

func (x *CallIntegrationServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallIntegrationServiceRequest.ProtoReflect.Descriptor instead.
func (*CallIntegrationServiceRequest) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{2}
}

func (x *CallIntegrationServiceRequest) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

func (x *CallIntegrationServiceRequest) GetService() *IntegrationService {
	if x != nil {
		return x.Service
	}
	return nil
}

type CallIntegrationServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CallIntegrationServiceResponse) Reset() {
	*x = CallIntegrationServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallIntegrationServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallIntegrationServiceResponse) ProtoMessage() {}

func (x *CallIntegrationServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallIntegrationServiceResponse.ProtoReflect.Descriptor instead.
func (*CallIntegrationServiceResponse) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{3}
}

func (x *CallIntegrationServiceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CallIntegrationServiceResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SetIntegrationStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationName string            `protobuf:"bytes,1,opt,name=integration_name,json=integrationName,proto3" json:"integration_name,omitempty"`
	State           *IntegrationState `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SetIntegrationStateRequest) Reset() {
	*x = SetIntegrationStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetIntegrationStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIntegrationStateRequest) ProtoMessage() {}

func (x *SetIntegrationStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIntegrationStateRequest.ProtoReflect.Descriptor instead.
func (*SetIntegrationStateRequest) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{4}
}

func (x *SetIntegrationStateRequest) GetIntegrationName() string {
	if x != nil {
		return x.IntegrationName
	}
	return ""
}

func (x *SetIntegrationStateRequest) GetState() *IntegrationState {
	if x != nil {
		return x.State
	}
	return nil
}

type Integration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Documentation string                `protobuf:"bytes,2,opt,name=documentation,proto3" json:"documentation,omitempty"`
	Version       string                `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Url           string                `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	State         *IntegrationState     `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Services      []*IntegrationService `protobuf:"bytes,6,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *Integration) Reset() {
	*x = Integration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integration) ProtoMessage() {}

func (x *Integration) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integration.ProtoReflect.Descriptor instead.
func (*Integration) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{5}
}

func (x *Integration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Integration) GetDocumentation() string {
	if x != nil {
		return x.Documentation
	}
	return ""
}

func (x *Integration) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Integration) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Integration) GetState() *IntegrationState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Integration) GetServices() []*IntegrationService {
	if x != nil {
		return x.Services
	}
	return nil
}

type IntegrationState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connected bool `protobuf:"varint,1,opt,name=connected,proto3" json:"connected,omitempty"`
}

func (x *IntegrationState) Reset() {
	*x = IntegrationState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationState) ProtoMessage() {}

func (x *IntegrationState) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationState.ProtoReflect.Descriptor instead.
func (*IntegrationState) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{6}
}

func (x *IntegrationState) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

type IntegrationService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *IntegrationService) Reset() {
	*x = IntegrationService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationService) ProtoMessage() {}

func (x *IntegrationService) ProtoReflect() protoreflect.Message {
	mi := &file_integration_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationService.ProtoReflect.Descriptor instead.
func (*IntegrationService) Descriptor() ([]byte, []int) {
	return file_integration_proto_rawDescGZIP(), []int{7}
}

func (x *IntegrationService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IntegrationService) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_integration_proto protoreflect.FileDescriptor

var file_integration_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x17, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x45, 0x0a, 0x18, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x7f, 0x0a, 0x1d, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x50, 0x0a, 0x1e, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x76, 0x0a, 0x1a, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xd9, 0x01, 0x0a, 0x0b, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xb7, 0x04, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x30, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a,
	0x08, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x24, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x50, 0x75, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x65, 0x74,
	0x12, 0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x2d, 0x5a, 0x2b, 0x73, 0x75, 0x6e, 0x64, 0x61, 0x65, 0x2d, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integration_proto_rawDescOnce sync.Once
	file_integration_proto_rawDescData = file_integration_proto_rawDesc
)

func file_integration_proto_rawDescGZIP() []byte {
	file_integration_proto_rawDescOnce.Do(func() {
		file_integration_proto_rawDescData = protoimpl.X.CompressGZIP(file_integration_proto_rawDescData)
	})
	return file_integration_proto_rawDescData
}

var file_integration_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_integration_proto_goTypes = []interface{}{
	(*IntegrationStoreRequest)(nil),        // 0: types.IntegrationStoreRequest
	(*IntegrationServerRequest)(nil),       // 1: types.IntegrationServerRequest
	(*CallIntegrationServiceRequest)(nil),  // 2: types.CallIntegrationServiceRequest
	(*CallIntegrationServiceResponse)(nil), // 3: types.CallIntegrationServiceResponse
	(*SetIntegrationStateRequest)(nil),     // 4: types.SetIntegrationStateRequest
	(*Integration)(nil),                    // 5: types.Integration
	(*IntegrationState)(nil),               // 6: types.IntegrationState
	(*IntegrationService)(nil),             // 7: types.IntegrationService
}
var file_integration_proto_depIdxs = []int32{
	7,  // 0: types.CallIntegrationServiceRequest.service:type_name -> types.IntegrationService
	6,  // 1: types.SetIntegrationStateRequest.state:type_name -> types.IntegrationState
	6,  // 2: types.Integration.state:type_name -> types.IntegrationState
	7,  // 3: types.Integration.services:type_name -> types.IntegrationService
	5,  // 4: types.IntegrationHandler.Create:input_type -> types.Integration
	1,  // 5: types.IntegrationHandler.Get:input_type -> types.IntegrationServerRequest
	5,  // 6: types.IntegrationHandler.Delete:input_type -> types.Integration
	4,  // 7: types.IntegrationHandler.SetState:input_type -> types.SetIntegrationStateRequest
	2,  // 8: types.IntegrationHandler.CallService:input_type -> types.CallIntegrationServiceRequest
	5,  // 9: types.IntegrationHandler.Connect:input_type -> types.Integration
	0,  // 10: types.IntegrationHandler.StorePut:input_type -> types.IntegrationStoreRequest
	0,  // 11: types.IntegrationHandler.StoreGet:input_type -> types.IntegrationStoreRequest
	5,  // 12: types.IntegrationHandler.Create:output_type -> types.Integration
	5,  // 13: types.IntegrationHandler.Get:output_type -> types.Integration
	5,  // 14: types.IntegrationHandler.Delete:output_type -> types.Integration
	6,  // 15: types.IntegrationHandler.SetState:output_type -> types.IntegrationState
	3,  // 16: types.IntegrationHandler.CallService:output_type -> types.CallIntegrationServiceResponse
	2,  // 17: types.IntegrationHandler.Connect:output_type -> types.CallIntegrationServiceRequest
	0,  // 18: types.IntegrationHandler.StorePut:output_type -> types.IntegrationStoreRequest
	0,  // 19: types.IntegrationHandler.StoreGet:output_type -> types.IntegrationStoreRequest
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_integration_proto_init() }
func file_integration_proto_init() {
	if File_integration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationStoreRequest); i {
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
		file_integration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationServerRequest); i {
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
		file_integration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallIntegrationServiceRequest); i {
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
		file_integration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallIntegrationServiceResponse); i {
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
		file_integration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetIntegrationStateRequest); i {
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
		file_integration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integration); i {
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
		file_integration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationState); i {
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
		file_integration_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationService); i {
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
			RawDescriptor: file_integration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integration_proto_goTypes,
		DependencyIndexes: file_integration_proto_depIdxs,
		MessageInfos:      file_integration_proto_msgTypes,
	}.Build()
	File_integration_proto = out.File
	file_integration_proto_rawDesc = nil
	file_integration_proto_goTypes = nil
	file_integration_proto_depIdxs = nil
}
