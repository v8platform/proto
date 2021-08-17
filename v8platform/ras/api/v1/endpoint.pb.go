// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: v8platform/ras/api/v1/endpoint.proto

package v1

import (
	_ "github.com/v8platform/protos/v8platform/ras"
	_ "github.com/v8platform/protos/v8platform/ras/serialize"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ParamType int32

const (
	ParamType_PARAM_UNKNOWN_TYPE  ParamType = 0
	ParamType_PARAM_BOOLEAN       ParamType = 1
	ParamType_PARAM_BYTE          ParamType = 2
	ParamType_PARAM_SHORT         ParamType = 3
	ParamType_PARAM_INT           ParamType = 4
	ParamType_PARAM_LONG          ParamType = 5
	ParamType_PARAM_FLOAT         ParamType = 6
	ParamType_PARAM_DOUBLE        ParamType = 7
	ParamType_PARAM_SIZE          ParamType = 8
	ParamType_PARAM_NULLABLE_SIZE ParamType = 9
	ParamType_PARAM_STRING        ParamType = 10
	ParamType_PARAM_UUID          ParamType = 11
	ParamType_PARAM_TYPE          ParamType = 12
	ParamType_PARAM_ENDPOINT_ID   ParamType = 13
)

// Enum value maps for ParamType.
var (
	ParamType_name = map[int32]string{
		0:  "PARAM_UNKNOWN_TYPE",
		1:  "PARAM_BOOLEAN",
		2:  "PARAM_BYTE",
		3:  "PARAM_SHORT",
		4:  "PARAM_INT",
		5:  "PARAM_LONG",
		6:  "PARAM_FLOAT",
		7:  "PARAM_DOUBLE",
		8:  "PARAM_SIZE",
		9:  "PARAM_NULLABLE_SIZE",
		10: "PARAM_STRING",
		11: "PARAM_UUID",
		12: "PARAM_TYPE",
		13: "PARAM_ENDPOINT_ID",
	}
	ParamType_value = map[string]int32{
		"PARAM_UNKNOWN_TYPE":  0,
		"PARAM_BOOLEAN":       1,
		"PARAM_BYTE":          2,
		"PARAM_SHORT":         3,
		"PARAM_INT":           4,
		"PARAM_LONG":          5,
		"PARAM_FLOAT":         6,
		"PARAM_DOUBLE":        7,
		"PARAM_SIZE":          8,
		"PARAM_NULLABLE_SIZE": 9,
		"PARAM_STRING":        10,
		"PARAM_UUID":          11,
		"PARAM_TYPE":          12,
		"PARAM_ENDPOINT_ID":   13,
	}
)

func (x ParamType) Enum() *ParamType {
	p := new(ParamType)
	*p = x
	return p
}

func (x ParamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParamType) Descriptor() protoreflect.EnumDescriptor {
	return file_v8platform_ras_api_v1_endpoint_proto_enumTypes[0].Descriptor()
}

func (ParamType) Type() protoreflect.EnumType {
	return &file_v8platform_ras_api_v1_endpoint_proto_enumTypes[0]
}

func (x ParamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParamType.Descriptor instead.
func (ParamType) EnumDescriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP(), []int{0}
}

type EndpointServiceType int32

const (
	EndpointServiceType_UNKNOWN_SERVICE EndpointServiceType = 0
	EndpointServiceType_ADMIN_CLUSTER   EndpointServiceType = 1 // "v8.service.Admin.Cluster"
)

// Enum value maps for EndpointServiceType.
var (
	EndpointServiceType_name = map[int32]string{
		0: "UNKNOWN_SERVICE",
		1: "ADMIN_CLUSTER",
	}
	EndpointServiceType_value = map[string]int32{
		"UNKNOWN_SERVICE": 0,
		"ADMIN_CLUSTER":   1,
	}
)

func (x EndpointServiceType) Enum() *EndpointServiceType {
	p := new(EndpointServiceType)
	*p = x
	return p
}

func (x EndpointServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndpointServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_v8platform_ras_api_v1_endpoint_proto_enumTypes[1].Descriptor()
}

func (EndpointServiceType) Type() protoreflect.EnumType {
	return &file_v8platform_ras_api_v1_endpoint_proto_enumTypes[1]
}

func (x EndpointServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EndpointServiceType.Descriptor instead.
func (EndpointServiceType) EnumDescriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP(), []int{1}
}

type Format int32

const (
	Format_FORMAT_UNKNOWN Format = 0
	Format_FORMAT_255     Format = 1
)

// Enum value maps for Format.
var (
	Format_name = map[int32]string{
		0: "FORMAT_UNKNOWN",
		1: "FORMAT_255",
	}
	Format_value = map[string]int32{
		"FORMAT_UNKNOWN": 0,
		"FORMAT_255":     1,
	}
)

func (x Format) Enum() *Format {
	p := new(Format)
	*p = x
	return p
}

func (x Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Format) Descriptor() protoreflect.EnumDescriptor {
	return file_v8platform_ras_api_v1_endpoint_proto_enumTypes[2].Descriptor()
}

func (Format) Type() protoreflect.EnumType {
	return &file_v8platform_ras_api_v1_endpoint_proto_enumTypes[2]
}

func (x Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Format.Descriptor instead.
func (Format) EnumDescriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP(), []int{2}
}

type EndpointOpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service  EndpointServiceType `protobuf:"varint,1,opt,name=service,proto3,enum=ras.api.v1.EndpointServiceType" json:"service,omitempty"`
	Encoding string              `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding,omitempty"`
	Version  string              `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
	Params   map[string]*Param   `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EndpointOpenRequest) Reset() {
	*x = EndpointOpenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointOpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointOpenRequest) ProtoMessage() {}

func (x *EndpointOpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointOpenRequest.ProtoReflect.Descriptor instead.
func (*EndpointOpenRequest) Descriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointOpenRequest) GetService() EndpointServiceType {
	if x != nil {
		return x.Service
	}
	return EndpointServiceType_UNKNOWN_SERVICE
}

func (x *EndpointOpenRequest) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *EndpointOpenRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *EndpointOpenRequest) GetParams() map[string]*Param {
	if x != nil {
		return x.Params
	}
	return nil
}

type EndpointOpenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Format Format `protobuf:"varint,1,opt,name=format,proto3,enum=ras.api.v1.Format" json:"format,omitempty"`
}

func (x *EndpointOpenResponse) Reset() {
	*x = EndpointOpenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointOpenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointOpenResponse) ProtoMessage() {}

func (x *EndpointOpenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointOpenResponse.ProtoReflect.Descriptor instead.
func (*EndpointOpenResponse) Descriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointOpenResponse) GetFormat() Format {
	if x != nil {
		return x.Format
	}
	return Format_FORMAT_UNKNOWN
}

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  ParamType  `protobuf:"varint,1,opt,name=type,proto3,enum=ras.api.v1.ParamType" json:"type,omitempty"`
	Value *anypb.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *Param) GetType() ParamType {
	if x != nil {
		return x.Type
	}
	return ParamType_PARAM_UNKNOWN_TYPE
}

func (x *Param) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type EndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId int32 `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Format     int32 `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"` // TODO int16
}

func (x *EndpointRequest) Reset() {
	*x = EndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRequest) ProtoMessage() {}

func (x *EndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRequest.ProtoReflect.Descriptor instead.
func (*EndpointRequest) Descriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP(), []int{3}
}

func (x *EndpointRequest) GetEndpointId() int32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *EndpointRequest) GetFormat() int32 {
	if x != nil {
		return x.Format
	}
	return 0
}

type EndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId int32 `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Format     int32 `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"` // TODO int16
}

func (x *EndpointResponse) Reset() {
	*x = EndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointResponse) ProtoMessage() {}

func (x *EndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_ras_api_v1_endpoint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointResponse.ProtoReflect.Descriptor instead.
func (*EndpointResponse) Descriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP(), []int{4}
}

func (x *EndpointResponse) GetEndpointId() int32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *EndpointResponse) GetFormat() int32 {
	if x != nil {
		return x.Format
	}
	return 0
}

var File_v8platform_ras_api_v1_endpoint_proto protoreflect.FileDescriptor

var file_v8platform_ras_api_v1_endpoint_proto_rawDesc = []byte{
	0x0a, 0x24, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x61, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x1a, 0x27, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72,
	0x61, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x76, 0x38, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x13, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x4c, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x42, 0x0a, 0x14, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x22, 0x5e, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x29, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x4a, 0x0a, 0x0f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x22, 0x4b, 0x0a, 0x10, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2a, 0x8b, 0x02,
	0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x50,
	0x41, 0x52, 0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x42, 0x4f, 0x4f,
	0x4c, 0x45, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f,
	0x42, 0x59, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f,
	0x53, 0x48, 0x4f, 0x52, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x5f, 0x49, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f,
	0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f,
	0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x49, 0x5a, 0x45,
	0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x55, 0x55,
	0x49, 0x44, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x45, 0x4e,
	0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x49, 0x44, 0x10, 0x0d, 0x2a, 0x3d, 0x0a, 0x13, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x2a, 0x2c, 0x0a, 0x06, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d,
	0x41, 0x54, 0x5f, 0x32, 0x35, 0x35, 0x10, 0x01, 0x32, 0xa6, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04,
	0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x61, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v8platform_ras_api_v1_endpoint_proto_rawDescOnce sync.Once
	file_v8platform_ras_api_v1_endpoint_proto_rawDescData = file_v8platform_ras_api_v1_endpoint_proto_rawDesc
)

func file_v8platform_ras_api_v1_endpoint_proto_rawDescGZIP() []byte {
	file_v8platform_ras_api_v1_endpoint_proto_rawDescOnce.Do(func() {
		file_v8platform_ras_api_v1_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_v8platform_ras_api_v1_endpoint_proto_rawDescData)
	})
	return file_v8platform_ras_api_v1_endpoint_proto_rawDescData
}

var file_v8platform_ras_api_v1_endpoint_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_v8platform_ras_api_v1_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v8platform_ras_api_v1_endpoint_proto_goTypes = []interface{}{
	(ParamType)(0),               // 0: ras.api.v1.ParamType
	(EndpointServiceType)(0),     // 1: ras.api.v1.EndpointServiceType
	(Format)(0),                  // 2: ras.api.v1.Format
	(*EndpointOpenRequest)(nil),  // 3: ras.api.v1.EndpointOpenRequest
	(*EndpointOpenResponse)(nil), // 4: ras.api.v1.EndpointOpenResponse
	(*Param)(nil),                // 5: ras.api.v1.Param
	(*EndpointRequest)(nil),      // 6: ras.api.v1.EndpointRequest
	(*EndpointResponse)(nil),     // 7: ras.api.v1.EndpointResponse
	nil,                          // 8: ras.api.v1.EndpointOpenRequest.ParamsEntry
	(*anypb.Any)(nil),            // 9: google.protobuf.Any
}
var file_v8platform_ras_api_v1_endpoint_proto_depIdxs = []int32{
	1, // 0: ras.api.v1.EndpointOpenRequest.service:type_name -> ras.api.v1.EndpointServiceType
	8, // 1: ras.api.v1.EndpointOpenRequest.params:type_name -> ras.api.v1.EndpointOpenRequest.ParamsEntry
	2, // 2: ras.api.v1.EndpointOpenResponse.format:type_name -> ras.api.v1.Format
	0, // 3: ras.api.v1.Param.type:type_name -> ras.api.v1.ParamType
	9, // 4: ras.api.v1.Param.value:type_name -> google.protobuf.Any
	5, // 5: ras.api.v1.EndpointOpenRequest.ParamsEntry.value:type_name -> ras.api.v1.Param
	3, // 6: ras.api.v1.EndpointService.Open:input_type -> ras.api.v1.EndpointOpenRequest
	6, // 7: ras.api.v1.EndpointService.Request:input_type -> ras.api.v1.EndpointRequest
	4, // 8: ras.api.v1.EndpointService.Open:output_type -> ras.api.v1.EndpointOpenResponse
	7, // 9: ras.api.v1.EndpointService.Request:output_type -> ras.api.v1.EndpointResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v8platform_ras_api_v1_endpoint_proto_init() }
func file_v8platform_ras_api_v1_endpoint_proto_init() {
	if File_v8platform_ras_api_v1_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v8platform_ras_api_v1_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointOpenRequest); i {
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
		file_v8platform_ras_api_v1_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointOpenResponse); i {
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
		file_v8platform_ras_api_v1_endpoint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_v8platform_ras_api_v1_endpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRequest); i {
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
		file_v8platform_ras_api_v1_endpoint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointResponse); i {
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
			RawDescriptor: file_v8platform_ras_api_v1_endpoint_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v8platform_ras_api_v1_endpoint_proto_goTypes,
		DependencyIndexes: file_v8platform_ras_api_v1_endpoint_proto_depIdxs,
		EnumInfos:         file_v8platform_ras_api_v1_endpoint_proto_enumTypes,
		MessageInfos:      file_v8platform_ras_api_v1_endpoint_proto_msgTypes,
	}.Build()
	File_v8platform_ras_api_v1_endpoint_proto = out.File
	file_v8platform_ras_api_v1_endpoint_proto_rawDesc = nil
	file_v8platform_ras_api_v1_endpoint_proto_goTypes = nil
	file_v8platform_ras_api_v1_endpoint_proto_depIdxs = nil
}
