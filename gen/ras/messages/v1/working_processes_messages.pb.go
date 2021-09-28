// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ras/messages/v1/working_processes_messages.proto

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

type GetWorkingProcessesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *GetWorkingProcessesRequest) Reset() {
	*x = GetWorkingProcessesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkingProcessesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkingProcessesRequest) ProtoMessage() {}

func (x *GetWorkingProcessesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkingProcessesRequest.ProtoReflect.Descriptor instead.
func (*GetWorkingProcessesRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_working_processes_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GetWorkingProcessesRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type GetWorkingProcessesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes []*v1.ProcessInfo `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
}

func (x *GetWorkingProcessesResponse) Reset() {
	*x = GetWorkingProcessesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkingProcessesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkingProcessesResponse) ProtoMessage() {}

func (x *GetWorkingProcessesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkingProcessesResponse.ProtoReflect.Descriptor instead.
func (*GetWorkingProcessesResponse) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_working_processes_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetWorkingProcessesResponse) GetProcesses() []*v1.ProcessInfo {
	if x != nil {
		return x.Processes
	}
	return nil
}

type GetWorkingProcessInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ProcessId string `protobuf:"bytes,2,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
}

func (x *GetWorkingProcessInfoRequest) Reset() {
	*x = GetWorkingProcessInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkingProcessInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkingProcessInfoRequest) ProtoMessage() {}

func (x *GetWorkingProcessInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkingProcessInfoRequest.ProtoReflect.Descriptor instead.
func (*GetWorkingProcessInfoRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_working_processes_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetWorkingProcessInfoRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetWorkingProcessInfoRequest) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

type GetWorkingProcessInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *v1.ProcessInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetWorkingProcessInfoResponse) Reset() {
	*x = GetWorkingProcessInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkingProcessInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkingProcessInfoResponse) ProtoMessage() {}

func (x *GetWorkingProcessInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkingProcessInfoResponse.ProtoReflect.Descriptor instead.
func (*GetWorkingProcessInfoResponse) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_working_processes_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetWorkingProcessInfoResponse) GetInfo() *v1.ProcessInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetServerWorkingProcessesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ServerId  string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *GetServerWorkingProcessesRequest) Reset() {
	*x = GetServerWorkingProcessesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerWorkingProcessesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerWorkingProcessesRequest) ProtoMessage() {}

func (x *GetServerWorkingProcessesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerWorkingProcessesRequest.ProtoReflect.Descriptor instead.
func (*GetServerWorkingProcessesRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_working_processes_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetServerWorkingProcessesRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetServerWorkingProcessesRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type GetServerWorkingProcessesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes *v1.ProcessInfo `protobuf:"bytes,1,opt,name=processes,proto3" json:"processes,omitempty"`
}

func (x *GetServerWorkingProcessesResponse) Reset() {
	*x = GetServerWorkingProcessesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerWorkingProcessesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerWorkingProcessesResponse) ProtoMessage() {}

func (x *GetServerWorkingProcessesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_working_processes_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerWorkingProcessesResponse.ProtoReflect.Descriptor instead.
func (*GetServerWorkingProcessesResponse) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_working_processes_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GetServerWorkingProcessesResponse) GetProcesses() *v1.ProcessInfo {
	if x != nil {
		return x.Processes
	}
	return nil
}

var File_ras_messages_v1_working_processes_messages_proto protoreflect.FileDescriptor

var file_ras_messages_v1_working_processes_messages_proto_rawDesc = []byte{
	0x0a, 0x30, 0x72, 0x61, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x72, 0x61, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x76, 0x38, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x03, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0xd8, 0x01, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb8, 0x01, 0x92, 0x41, 0xa6, 0x01, 0x32,
	0x55, 0xd0, 0xa3, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0,
	0xbd, 0xd1, 0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xb8, 0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82,
	0xd0, 0xb8, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80,
	0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0,
	0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb0, 0xd1, 0x81, 0xd1, 0x82,
	0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x8a, 0x01, 0x45, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d,
	0x46, 0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34,
	0x7d, 0x2d, 0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b,
	0x38, 0x39, 0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d,
	0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0xa2, 0x02,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x10, 0x01, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x3a, 0xc1,
	0x01, 0x92, 0x41, 0x98, 0x01, 0x0a, 0x95, 0x01, 0x2a, 0x1a, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0x6a, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87,
	0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xb8, 0xd1,
	0x81, 0xd0, 0xba, 0xd0, 0xb0, 0x20, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0, 0xbe, 0xd1, 0x87,
	0xd0, 0xb8, 0xd1, 0x85, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd1, 0x86, 0xd0, 0xb5, 0xd1,
	0x81, 0xd1, 0x81, 0xd0, 0xbe, 0xd0, 0xb2, 0x20, 0xd0, 0xbd, 0xd0, 0xb0, 0x20, 0xd0, 0xbb, 0xd0,
	0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xbc, 0x20,
	0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb5,
	0xd2, 0x01, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x8a, 0xf5, 0xea,
	0x94, 0x0e, 0x1f, 0x3a, 0x1d, 0x47, 0x45, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x49, 0x4e, 0x47,
	0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x22, 0x93, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x82, 0xf5, 0xea,
	0x94, 0x0e, 0x02, 0x10, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x3a, 0x26, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x20, 0x3a, 0x1e, 0x47, 0x45, 0x54, 0x5f, 0x57, 0x4f,
	0x52, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x53, 0x5f,
	0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x22, 0xd0, 0x05, 0x0a, 0x1c, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xd8, 0x01, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb8,
	0x01, 0x92, 0x41, 0xa6, 0x01, 0x32, 0x55, 0xd0, 0xa3, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba, 0xd0,
	0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd1, 0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xb8, 0xd0, 0xb4,
	0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb8, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0,
	0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0,
	0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb,
	0xd0, 0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x8a, 0x01, 0x45, 0x5e,
	0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39,
	0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46,
	0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41,
	0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b,
	0x31, 0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x01, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0xfa, 0x01, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0xda, 0x01, 0x92, 0x41, 0xc8, 0x01,
	0x32, 0x77, 0xd0, 0xa3, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c,
	0xd0, 0xbd, 0xd1, 0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xb8, 0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1,
	0x82, 0xd0, 0xb8, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd1,
	0x80, 0x20, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0, 0xbe, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xb3,
	0xd0, 0xbe, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd1, 0x86, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
	0x81, 0xd0, 0xb0, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c,
	0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb0, 0xd1,
	0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x8a, 0x01, 0x45, 0x5e, 0x5b, 0x30, 0x2d,
	0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46,
	0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33,
	0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d,
	0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d,
	0x24, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x10, 0x02, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x3a, 0xd7, 0x01, 0x92, 0x41, 0xab, 0x01, 0x0a, 0xa8, 0x01, 0x2a, 0x1b, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x70, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb,
	0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd0, 0xb8, 0xd0,
	0xbd, 0xd1, 0x84, 0xd0, 0xbe, 0xd1, 0x80, 0xd0, 0xbc, 0xd0, 0xb0, 0xd1, 0x86, 0xd0, 0xb8, 0xd0,
	0xb8, 0x20, 0xd0, 0xbe, 0x20, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0, 0xbe, 0xd1, 0x87, 0xd0,
	0xb5, 0xd0, 0xbc, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd1, 0x86, 0xd0, 0xb5, 0xd1, 0x81,
	0xd1, 0x81, 0xd0, 0xb5, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1,
	0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb0,
	0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0xd2, 0x01, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x22, 0x3a, 0x20, 0x47, 0x45, 0x54, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x22, 0x8e, 0x01, 0x0a, 0x1d,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x38,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x01, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x3a, 0x29, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x23, 0x3a, 0x21, 0x47, 0x45, 0x54, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x22, 0xf4, 0x05, 0x0a,
	0x20, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0xd8, 0x01, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb8, 0x01, 0x92, 0x41, 0xa6, 0x01, 0x32, 0x55, 0xd0,
	0xa3, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd1,
	0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xb8, 0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb8,
	0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80, 0x20, 0xd0,
	0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0,
	0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba, 0xd0, 0xbb, 0xd0, 0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5,
	0xd1, 0x80, 0xd0, 0xb0, 0x8a, 0x01, 0x45, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d,
	0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d,
	0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39,
	0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b,
	0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10,
	0x01, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0xf6, 0x01, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0xd8, 0x01, 0x92, 0x41, 0xc6, 0x01, 0x32, 0x75, 0xd0, 0xa3, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0,
	0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd1, 0x8b, 0xd0, 0xb9, 0x20, 0xd0, 0xb8,
	0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x82, 0xd0, 0xb8, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xba,
	0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80, 0x20, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0,
	0xbe, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd1, 0x81, 0xd0, 0xb5, 0xd1, 0x80,
	0xd0, 0xb2, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb5, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xba, 0xd0,
	0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd0, 0xba,
	0xd0, 0xbb, 0xd0, 0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x8a, 0x01,
	0x45, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30,
	0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41,
	0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d,
	0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46,
	0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0x82, 0xf5, 0xea,
	0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x02, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x3a, 0xfb, 0x01, 0x92, 0x41, 0xcb, 0x01, 0x0a, 0xc8, 0x01, 0x2a,
	0x20, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0x8a, 0x01, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb5,
	0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xb8, 0xd1, 0x81, 0xd0,
	0xba, 0xd0, 0xb0, 0x20, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0, 0xbe, 0xd1, 0x87, 0xd0, 0xb8,
	0xd1, 0x85, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd1, 0x86, 0xd0, 0xb5, 0xd1, 0x81, 0xd1,
	0x81, 0xd0, 0xbe, 0xd0, 0xb2, 0x20, 0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xb1, 0xd0, 0xbe, 0xd1, 0x87,
	0xd0, 0xb5, 0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd1, 0x81, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb2, 0xd0,
	0xb5, 0xd1, 0x80, 0xd0, 0xb0, 0x20, 0xd0, 0xbd, 0xd0, 0xb0, 0x20, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0,
	0xba, 0xd0, 0xb0, 0xd0, 0xbb, 0xd1, 0x8c, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xbc, 0x20, 0xd0, 0xba,
	0xd0, 0xbb, 0xd0, 0xb0, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb5, 0xd2, 0x01,
	0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x26, 0x3a, 0x24, 0x47,
	0x45, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x49, 0x4e,
	0x47, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x22, 0xa0, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76,
	0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x01, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x3a, 0x2d, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x27, 0x3a,
	0x25, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x57, 0x4f, 0x52, 0x4b,
	0x49, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x53, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x42, 0xcf, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x72,
	0x61, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1d,
	0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52,
	0x4d, 0x58, 0xaa, 0x02, 0x0f, 0x52, 0x61, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x52, 0x61, 0x73, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x52, 0x61, 0x73, 0x5c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x52, 0x61, 0x73, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_messages_v1_working_processes_messages_proto_rawDescOnce sync.Once
	file_ras_messages_v1_working_processes_messages_proto_rawDescData = file_ras_messages_v1_working_processes_messages_proto_rawDesc
)

func file_ras_messages_v1_working_processes_messages_proto_rawDescGZIP() []byte {
	file_ras_messages_v1_working_processes_messages_proto_rawDescOnce.Do(func() {
		file_ras_messages_v1_working_processes_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_messages_v1_working_processes_messages_proto_rawDescData)
	})
	return file_ras_messages_v1_working_processes_messages_proto_rawDescData
}

var file_ras_messages_v1_working_processes_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ras_messages_v1_working_processes_messages_proto_goTypes = []interface{}{
	(*GetWorkingProcessesRequest)(nil),        // 0: ras.messages.v1.GetWorkingProcessesRequest
	(*GetWorkingProcessesResponse)(nil),       // 1: ras.messages.v1.GetWorkingProcessesResponse
	(*GetWorkingProcessInfoRequest)(nil),      // 2: ras.messages.v1.GetWorkingProcessInfoRequest
	(*GetWorkingProcessInfoResponse)(nil),     // 3: ras.messages.v1.GetWorkingProcessInfoResponse
	(*GetServerWorkingProcessesRequest)(nil),  // 4: ras.messages.v1.GetServerWorkingProcessesRequest
	(*GetServerWorkingProcessesResponse)(nil), // 5: ras.messages.v1.GetServerWorkingProcessesResponse
	(*v1.ProcessInfo)(nil),                    // 6: v8platform.serialize.v1.ProcessInfo
}
var file_ras_messages_v1_working_processes_messages_proto_depIdxs = []int32{
	6, // 0: ras.messages.v1.GetWorkingProcessesResponse.processes:type_name -> v8platform.serialize.v1.ProcessInfo
	6, // 1: ras.messages.v1.GetWorkingProcessInfoResponse.info:type_name -> v8platform.serialize.v1.ProcessInfo
	6, // 2: ras.messages.v1.GetServerWorkingProcessesResponse.processes:type_name -> v8platform.serialize.v1.ProcessInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ras_messages_v1_working_processes_messages_proto_init() }
func file_ras_messages_v1_working_processes_messages_proto_init() {
	if File_ras_messages_v1_working_processes_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ras_messages_v1_working_processes_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkingProcessesRequest); i {
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
		file_ras_messages_v1_working_processes_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkingProcessesResponse); i {
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
		file_ras_messages_v1_working_processes_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkingProcessInfoRequest); i {
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
		file_ras_messages_v1_working_processes_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkingProcessInfoResponse); i {
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
		file_ras_messages_v1_working_processes_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerWorkingProcessesRequest); i {
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
		file_ras_messages_v1_working_processes_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerWorkingProcessesResponse); i {
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
			RawDescriptor: file_ras_messages_v1_working_processes_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_messages_v1_working_processes_messages_proto_goTypes,
		DependencyIndexes: file_ras_messages_v1_working_processes_messages_proto_depIdxs,
		MessageInfos:      file_ras_messages_v1_working_processes_messages_proto_msgTypes,
	}.Build()
	File_ras_messages_v1_working_processes_messages_proto = out.File
	file_ras_messages_v1_working_processes_messages_proto_rawDesc = nil
	file_ras_messages_v1_working_processes_messages_proto_goTypes = nil
	file_ras_messages_v1_working_processes_messages_proto_depIdxs = nil
}
