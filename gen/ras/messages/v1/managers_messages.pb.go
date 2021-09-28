// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ras/messages/v1/managers_messages.proto

package messagesv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/v8platform/protoc-gen-go-ras/plugin/ras/encoding"
	v1 "github.com/v8platform/protos/gen/v8platform/serialize/v1"
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

type GetClusterManagersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *GetClusterManagersRequest) Reset() {
	*x = GetClusterManagersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_managers_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterManagersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterManagersRequest) ProtoMessage() {}

func (x *GetClusterManagersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_managers_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterManagersRequest.ProtoReflect.Descriptor instead.
func (*GetClusterManagersRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_managers_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GetClusterManagersRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type GetClusterManagersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Managers []*v1.ManagerInfo `protobuf:"bytes,1,rep,name=managers,proto3" json:"managers,omitempty"`
}

func (x *GetClusterManagersResponse) Reset() {
	*x = GetClusterManagersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_managers_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterManagersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterManagersResponse) ProtoMessage() {}

func (x *GetClusterManagersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_managers_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterManagersResponse.ProtoReflect.Descriptor instead.
func (*GetClusterManagersResponse) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_managers_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetClusterManagersResponse) GetManagers() []*v1.ManagerInfo {
	if x != nil {
		return x.Managers
	}
	return nil
}

type GetClusterManagerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ManagerId string `protobuf:"bytes,2,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
}

func (x *GetClusterManagerInfoRequest) Reset() {
	*x = GetClusterManagerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_managers_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterManagerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterManagerInfoRequest) ProtoMessage() {}

func (x *GetClusterManagerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_managers_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterManagerInfoRequest.ProtoReflect.Descriptor instead.
func (*GetClusterManagerInfoRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_managers_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetClusterManagerInfoRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetClusterManagerInfoRequest) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

type GetClusterManagerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *v1.ManagerInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetClusterManagerInfoResponse) Reset() {
	*x = GetClusterManagerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_managers_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterManagerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterManagerInfoResponse) ProtoMessage() {}

func (x *GetClusterManagerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_managers_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterManagerInfoResponse.ProtoReflect.Descriptor instead.
func (*GetClusterManagerInfoResponse) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_managers_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetClusterManagerInfoResponse) GetInfo() *v1.ManagerInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_ras_messages_v1_managers_messages_proto protoreflect.FileDescriptor

var file_ras_messages_v1_managers_messages_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x61, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72, 0x61, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x26, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x03, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xd8, 0x01, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb8, 0x01,
	0x92, 0x41, 0xa6, 0x01, 0x32, 0x55, 0xd0, 0xa3, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0,
	0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd1, 0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xb8, 0xd0, 0xb4, 0xd0,
	0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb8, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd1,
	0x82, 0xd0, 0xbe, 0xd1, 0x80, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb,
	0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0,
	0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x8a, 0x01, 0x45, 0x5e, 0x5b,
	0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41,
	0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d,
	0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d,
	0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31,
	0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x01, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x3a, 0xaf, 0x01, 0x92, 0x41, 0x87, 0x01, 0x0a, 0x84, 0x01, 0x2a, 0x19, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x5a, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb,
	0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd1, 0x81, 0xd0,
	0xbf, 0xd0, 0xb8, 0xd1, 0x81, 0xd0, 0xba, 0xd0, 0xb0, 0x20, 0xd0, 0xbc, 0xd0, 0xb5, 0xd0, 0xbd,
	0xd0, 0xb5, 0xd0, 0xb4, 0xd0, 0xb6, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xbe, 0xd0, 0xb2, 0x20, 0xd0,
	0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0,
	0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5,
	0xd1, 0x80, 0xd0, 0xb0, 0xd2, 0x01, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x1e, 0x3a, 0x1c, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4c, 0x55,
	0x53, 0x54, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x53, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x22, 0x8f, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x82, 0xf5,
	0xea, 0x94, 0x0e, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73,
	0x3a, 0x25, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x1f, 0x3a, 0x1d, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4c,
	0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x53, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x22, 0xb5, 0x05, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xd8, 0x01, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb8, 0x01,
	0x92, 0x41, 0xa6, 0x01, 0x32, 0x55, 0xd0, 0xa3, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0,
	0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd1, 0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xb8, 0xd0, 0xb4, 0xd0,
	0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb8, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd1,
	0x82, 0xd0, 0xbe, 0xd1, 0x80, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb,
	0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0,
	0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x8a, 0x01, 0x45, 0x5e, 0x5b,
	0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41,
	0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d,
	0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d,
	0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31,
	0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x01, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0xeb, 0x01, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0xcb, 0x01, 0x92, 0x41, 0xb9, 0x01, 0x32,
	0x68, 0xd0, 0xa3, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0,
	0xbd, 0xd1, 0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xb8, 0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82,
	0xd0, 0xb8, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80,
	0x20, 0xd0, 0xbc, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xb4, 0xd0, 0xb6, 0xd0, 0xb5, 0xd1,
	0x80, 0xd0, 0xb0, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c,
	0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb0, 0xd1,
	0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x8a, 0x01, 0x45, 0x5e, 0x5b, 0x30, 0x2d,
	0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46,
	0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33,
	0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d,
	0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d,
	0x24, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x10, 0x02, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49,
	0x64, 0x3a, 0xcb, 0x01, 0x92, 0x41, 0x9f, 0x01, 0x0a, 0x9c, 0x01, 0x2a, 0x1b, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x63, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb,
	0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd0, 0xb8, 0xd0,
	0xbd, 0xd1, 0x84, 0xd0, 0xbe, 0xd1, 0x80, 0xd0, 0xbc, 0xd0, 0xb0, 0xd1, 0x86, 0xd0, 0xb8, 0xd0,
	0xb8, 0x20, 0xd0, 0xbe, 0x20, 0xd0, 0xbc, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xb4, 0xd0,
	0xb6, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb5, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0,
	0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0,
	0xbb, 0xd0, 0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0xd2, 0x01, 0x0a,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x0a, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x22, 0x3a, 0x20, 0x47,
	0x45, 0x54, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47,
	0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x22,
	0x8e, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x3a, 0x29, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x23, 0x3a, 0x21, 0x47,
	0x45, 0x54, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47,
	0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x42, 0xc7, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x52, 0x4d, 0x58, 0xaa, 0x02, 0x0f, 0x52, 0x61, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x52, 0x61, 0x73, 0x5c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x52, 0x61, 0x73, 0x5c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x52, 0x61, 0x73, 0x3a, 0x3a, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ras_messages_v1_managers_messages_proto_rawDescOnce sync.Once
	file_ras_messages_v1_managers_messages_proto_rawDescData = file_ras_messages_v1_managers_messages_proto_rawDesc
)

func file_ras_messages_v1_managers_messages_proto_rawDescGZIP() []byte {
	file_ras_messages_v1_managers_messages_proto_rawDescOnce.Do(func() {
		file_ras_messages_v1_managers_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_messages_v1_managers_messages_proto_rawDescData)
	})
	return file_ras_messages_v1_managers_messages_proto_rawDescData
}

var file_ras_messages_v1_managers_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ras_messages_v1_managers_messages_proto_goTypes = []interface{}{
	(*GetClusterManagersRequest)(nil),     // 0: ras.messages.v1.GetClusterManagersRequest
	(*GetClusterManagersResponse)(nil),    // 1: ras.messages.v1.GetClusterManagersResponse
	(*GetClusterManagerInfoRequest)(nil),  // 2: ras.messages.v1.GetClusterManagerInfoRequest
	(*GetClusterManagerInfoResponse)(nil), // 3: ras.messages.v1.GetClusterManagerInfoResponse
	(*v1.ManagerInfo)(nil),                // 4: v8platform.serialize.v1.ManagerInfo
}
var file_ras_messages_v1_managers_messages_proto_depIdxs = []int32{
	4, // 0: ras.messages.v1.GetClusterManagersResponse.managers:type_name -> v8platform.serialize.v1.ManagerInfo
	4, // 1: ras.messages.v1.GetClusterManagerInfoResponse.info:type_name -> v8platform.serialize.v1.ManagerInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ras_messages_v1_managers_messages_proto_init() }
func file_ras_messages_v1_managers_messages_proto_init() {
	if File_ras_messages_v1_managers_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ras_messages_v1_managers_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterManagersRequest); i {
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
		file_ras_messages_v1_managers_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterManagersResponse); i {
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
		file_ras_messages_v1_managers_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterManagerInfoRequest); i {
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
		file_ras_messages_v1_managers_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterManagerInfoResponse); i {
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
			RawDescriptor: file_ras_messages_v1_managers_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_messages_v1_managers_messages_proto_goTypes,
		DependencyIndexes: file_ras_messages_v1_managers_messages_proto_depIdxs,
		MessageInfos:      file_ras_messages_v1_managers_messages_proto_msgTypes,
	}.Build()
	File_ras_messages_v1_managers_messages_proto = out.File
	file_ras_messages_v1_managers_messages_proto_rawDesc = nil
	file_ras_messages_v1_managers_messages_proto_goTypes = nil
	file_ras_messages_v1_managers_messages_proto_depIdxs = nil
}
