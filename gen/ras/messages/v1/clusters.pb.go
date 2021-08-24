// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: ras/messages/v1/clusters.proto

package messagesv1

import (
	_ "github.com/v8platform/protos/gen/ras/encoding"
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

type GetClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClustersRequest) Reset() {
	*x = GetClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_clusters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClustersRequest) ProtoMessage() {}

func (x *GetClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_clusters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClustersRequest.ProtoReflect.Descriptor instead.
func (*GetClustersRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_clusters_proto_rawDescGZIP(), []int{0}
}

type GetClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*v1.ClusterInfo `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *GetClustersResponse) Reset() {
	*x = GetClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_clusters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClustersResponse) ProtoMessage() {}

func (x *GetClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_clusters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClustersResponse.ProtoReflect.Descriptor instead.
func (*GetClustersResponse) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_clusters_proto_rawDescGZIP(), []int{1}
}

func (x *GetClustersResponse) GetClusters() []*v1.ClusterInfo {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type GetClusterInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  option (ras.encoding.endpoint_message_type) = GET_CLUSTER_INFO_REQUEST;
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *GetClusterInfoRequest) Reset() {
	*x = GetClusterInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_clusters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterInfoRequest) ProtoMessage() {}

func (x *GetClusterInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_clusters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterInfoRequest.ProtoReflect.Descriptor instead.
func (*GetClusterInfoRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_clusters_proto_rawDescGZIP(), []int{2}
}

func (x *GetClusterInfoRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type GetClusterInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  option (ras.encoding.endpoint_message_type) = GET_CLUSTER_INFO_RESPONSE;
	ClusterInfo *v1.ClusterInfo `protobuf:"bytes,1,opt,name=cluster_info,json=clusterInfo,proto3" json:"cluster_info,omitempty"`
}

func (x *GetClusterInfoResponse) Reset() {
	*x = GetClusterInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_clusters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterInfoResponse) ProtoMessage() {}

func (x *GetClusterInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_clusters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterInfoResponse.ProtoReflect.Descriptor instead.
func (*GetClusterInfoResponse) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_clusters_proto_rawDescGZIP(), []int{3}
}

func (x *GetClusterInfoResponse) GetClusterInfo() *v1.ClusterInfo {
	if x != nil {
		return x.ClusterInfo
	}
	return nil
}

type RegClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  option (ras.encoding.endpoint_message_type) = REG_CLUSTER_REQUEST;
	ClusterInfo *v1.ClusterInfo `protobuf:"bytes,1,opt,name=cluster_info,json=clusterInfo,proto3" json:"cluster_info,omitempty"`
}

func (x *RegClusterRequest) Reset() {
	*x = RegClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_clusters_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegClusterRequest) ProtoMessage() {}

func (x *RegClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_clusters_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegClusterRequest.ProtoReflect.Descriptor instead.
func (*RegClusterRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_clusters_proto_rawDescGZIP(), []int{4}
}

func (x *RegClusterRequest) GetClusterInfo() *v1.ClusterInfo {
	if x != nil {
		return x.ClusterInfo
	}
	return nil
}

type RegClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  option (ras.encoding.endpoint_message_type) = REG_CLUSTER_RESPONSE;
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *RegClusterResponse) Reset() {
	*x = RegClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_clusters_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegClusterResponse) ProtoMessage() {}

func (x *RegClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_clusters_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegClusterResponse.ProtoReflect.Descriptor instead.
func (*RegClusterResponse) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_clusters_proto_rawDescGZIP(), []int{5}
}

func (x *RegClusterResponse) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type UnregClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  option (ras.encoding.endpoint_message_type) = UNREG_CLUSTER_REQUEST;
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *UnregClusterRequest) Reset() {
	*x = UnregClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_messages_v1_clusters_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregClusterRequest) ProtoMessage() {}

func (x *UnregClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_messages_v1_clusters_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregClusterRequest.ProtoReflect.Descriptor instead.
func (*UnregClusterRequest) Descriptor() ([]byte, []int) {
	return file_ras_messages_v1_clusters_proto_rawDescGZIP(), []int{6}
}

func (x *UnregClusterRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

var File_ras_messages_v1_clusters_proto protoreflect.FileDescriptor

var file_ras_messages_v1_clusters_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x61, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x72, 0x61, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x61, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x72, 0x61, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x3a, 0x06, 0x98, 0xf5, 0xea, 0x94, 0x0e, 0x0b, 0x22, 0x69, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94,
	0x0e, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x3a, 0x06,
	0x98, 0xf5, 0xea, 0x94, 0x0e, 0x0c, 0x22, 0x4f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x3a, 0x22, 0x75, 0x75, 0x69, 0x64, 0x2c, 0x31, 0x22, 0x52, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5c, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x13, 0x9a,
	0x84, 0x9e, 0x03, 0x0e, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c,
	0x31, 0x22, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x71, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x5c, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x38, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x13, 0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x3a, 0x22, 0x2c, 0x31, 0x22, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84,
	0x9e, 0x03, 0x12, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x75, 0x75,
	0x69, 0x64, 0x2c, 0x31, 0x22, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4d, 0x0a, 0x13, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e,
	0x03, 0x12, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x75, 0x75, 0x69,
	0x64, 0x2c, 0x31, 0x22, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42,
	0xcb, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x4d, 0x58, 0xaa, 0x02, 0x0f, 0x52, 0x61,
	0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f,
	0x52, 0x61, 0x73, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1b, 0x52, 0x61, 0x73, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11,
	0x52, 0x61, 0x73, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x06, 0x08, 0x01, 0x10, 0x01, 0x18, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_messages_v1_clusters_proto_rawDescOnce sync.Once
	file_ras_messages_v1_clusters_proto_rawDescData = file_ras_messages_v1_clusters_proto_rawDesc
)

func file_ras_messages_v1_clusters_proto_rawDescGZIP() []byte {
	file_ras_messages_v1_clusters_proto_rawDescOnce.Do(func() {
		file_ras_messages_v1_clusters_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_messages_v1_clusters_proto_rawDescData)
	})
	return file_ras_messages_v1_clusters_proto_rawDescData
}

var file_ras_messages_v1_clusters_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ras_messages_v1_clusters_proto_goTypes = []interface{}{
	(*GetClustersRequest)(nil),     // 0: ras.messages.v1.GetClustersRequest
	(*GetClustersResponse)(nil),    // 1: ras.messages.v1.GetClustersResponse
	(*GetClusterInfoRequest)(nil),  // 2: ras.messages.v1.GetClusterInfoRequest
	(*GetClusterInfoResponse)(nil), // 3: ras.messages.v1.GetClusterInfoResponse
	(*RegClusterRequest)(nil),      // 4: ras.messages.v1.RegClusterRequest
	(*RegClusterResponse)(nil),     // 5: ras.messages.v1.RegClusterResponse
	(*UnregClusterRequest)(nil),    // 6: ras.messages.v1.UnregClusterRequest
	(*v1.ClusterInfo)(nil),         // 7: v8platform.serialize.v1.ClusterInfo
}
var file_ras_messages_v1_clusters_proto_depIdxs = []int32{
	7, // 0: ras.messages.v1.GetClustersResponse.clusters:type_name -> v8platform.serialize.v1.ClusterInfo
	7, // 1: ras.messages.v1.GetClusterInfoResponse.cluster_info:type_name -> v8platform.serialize.v1.ClusterInfo
	7, // 2: ras.messages.v1.RegClusterRequest.cluster_info:type_name -> v8platform.serialize.v1.ClusterInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ras_messages_v1_clusters_proto_init() }
func file_ras_messages_v1_clusters_proto_init() {
	if File_ras_messages_v1_clusters_proto != nil {
		return
	}
	file_ras_messages_v1_types_proto_init()
	file_ras_messages_v1_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ras_messages_v1_clusters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClustersRequest); i {
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
		file_ras_messages_v1_clusters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClustersResponse); i {
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
		file_ras_messages_v1_clusters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterInfoRequest); i {
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
		file_ras_messages_v1_clusters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterInfoResponse); i {
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
		file_ras_messages_v1_clusters_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegClusterRequest); i {
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
		file_ras_messages_v1_clusters_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegClusterResponse); i {
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
		file_ras_messages_v1_clusters_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregClusterRequest); i {
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
			RawDescriptor: file_ras_messages_v1_clusters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_messages_v1_clusters_proto_goTypes,
		DependencyIndexes: file_ras_messages_v1_clusters_proto_depIdxs,
		MessageInfos:      file_ras_messages_v1_clusters_proto_msgTypes,
	}.Build()
	File_ras_messages_v1_clusters_proto = out.File
	file_ras_messages_v1_clusters_proto_rawDesc = nil
	file_ras_messages_v1_clusters_proto_goTypes = nil
	file_ras_messages_v1_clusters_proto_depIdxs = nil
}
