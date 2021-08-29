// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ras/api/v1/infobases_request.proto

package apiv1

import (
	v1 "github.com/v8platform/protos/gen/v8platform/serialize/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InfobasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *InfobasesRequest) Reset() {
	*x = InfobasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_infobases_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfobasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfobasesRequest) ProtoMessage() {}

func (x *InfobasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_infobases_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfobasesRequest.ProtoReflect.Descriptor instead.
func (*InfobasesRequest) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_infobases_request_proto_rawDescGZIP(), []int{0}
}

func (x *InfobasesRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type InfobasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*v1.InfobaseSummaryInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *InfobasesResponse) Reset() {
	*x = InfobasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_infobases_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfobasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfobasesResponse) ProtoMessage() {}

func (x *InfobasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_infobases_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfobasesResponse.ProtoReflect.Descriptor instead.
func (*InfobasesResponse) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_infobases_request_proto_rawDescGZIP(), []int{1}
}

func (x *InfobasesResponse) GetItems() []*v1.InfobaseSummaryInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetInfobaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster    *Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	InfobaseId string   `protobuf:"bytes,2,opt,name=infobase_id,json=infobaseId,proto3" json:"infobase_id,omitempty"`
}

func (x *GetInfobaseRequest) Reset() {
	*x = GetInfobaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_infobases_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfobaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfobaseRequest) ProtoMessage() {}

func (x *GetInfobaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_infobases_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfobaseRequest.ProtoReflect.Descriptor instead.
func (*GetInfobaseRequest) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_infobases_request_proto_rawDescGZIP(), []int{2}
}

func (x *GetInfobaseRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *GetInfobaseRequest) GetInfobaseId() string {
	if x != nil {
		return x.InfobaseId
	}
	return ""
}

type LookupInfobaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *Cluster       `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Field   []*FieldFilter `protobuf:"bytes,2,rep,name=field,proto3" json:"field,omitempty"`
}

func (x *LookupInfobaseRequest) Reset() {
	*x = LookupInfobaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_infobases_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupInfobaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupInfobaseRequest) ProtoMessage() {}

func (x *LookupInfobaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_infobases_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupInfobaseRequest.ProtoReflect.Descriptor instead.
func (*LookupInfobaseRequest) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_infobases_request_proto_rawDescGZIP(), []int{3}
}

func (x *LookupInfobaseRequest) GetCluster() *Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *LookupInfobaseRequest) GetField() []*FieldFilter {
	if x != nil {
		return x.Field
	}
	return nil
}

type LookupInfobaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*v1.InfobaseSummaryInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *LookupInfobaseResponse) Reset() {
	*x = LookupInfobaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_infobases_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupInfobaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupInfobaseResponse) ProtoMessage() {}

func (x *LookupInfobaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_infobases_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupInfobaseResponse.ProtoReflect.Descriptor instead.
func (*LookupInfobaseResponse) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_infobases_request_proto_rawDescGZIP(), []int{4}
}

func (x *LookupInfobaseResponse) GetItems() []*v1.InfobaseSummaryInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type FieldFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameReqexp string     `protobuf:"bytes,1,opt,name=name_reqexp,json=nameReqexp,proto3" json:"name_reqexp,omitempty"`
	Value      *anypb.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FieldFilter) Reset() {
	*x = FieldFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_api_v1_infobases_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldFilter) ProtoMessage() {}

func (x *FieldFilter) ProtoReflect() protoreflect.Message {
	mi := &file_ras_api_v1_infobases_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldFilter.ProtoReflect.Descriptor instead.
func (*FieldFilter) Descriptor() ([]byte, []int) {
	return file_ras_api_v1_infobases_request_proto_rawDescGZIP(), []int{5}
}

func (x *FieldFilter) GetNameReqexp() string {
	if x != nil {
		return x.NameReqexp
	}
	return ""
}

func (x *FieldFilter) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_ras_api_v1_infobases_request_proto protoreflect.FileDescriptor

var file_ras_api_v1_infobases_request_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66,
	0x6f, 0x62, 0x61, 0x73, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x1a, 0x27, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x72, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10, 0x49, 0x6e, 0x66,
	0x6f, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x57, 0x0a, 0x11,
	0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x62,
	0x61, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x64, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72,
	0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x15, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x22, 0x5c, 0x0a, 0x16, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x76, 0x38,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x5a, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x65, 0x78, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x65, 0x78, 0x70,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x91, 0x02, 0x0a,
	0x10, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4a, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x1c,
	0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72,
	0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x21, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x42, 0xa4, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x42, 0x15, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72,
	0x61, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x52, 0x41, 0x58, 0xaa, 0x02, 0x0a, 0x52, 0x61, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0a, 0x52, 0x61, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x16, 0x52, 0x61, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x52, 0x61, 0x73, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_api_v1_infobases_request_proto_rawDescOnce sync.Once
	file_ras_api_v1_infobases_request_proto_rawDescData = file_ras_api_v1_infobases_request_proto_rawDesc
)

func file_ras_api_v1_infobases_request_proto_rawDescGZIP() []byte {
	file_ras_api_v1_infobases_request_proto_rawDescOnce.Do(func() {
		file_ras_api_v1_infobases_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_api_v1_infobases_request_proto_rawDescData)
	})
	return file_ras_api_v1_infobases_request_proto_rawDescData
}

var file_ras_api_v1_infobases_request_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ras_api_v1_infobases_request_proto_goTypes = []interface{}{
	(*InfobasesRequest)(nil),       // 0: ras.api.v1.InfobasesRequest
	(*InfobasesResponse)(nil),      // 1: ras.api.v1.InfobasesResponse
	(*GetInfobaseRequest)(nil),     // 2: ras.api.v1.GetInfobaseRequest
	(*LookupInfobaseRequest)(nil),  // 3: ras.api.v1.LookupInfobaseRequest
	(*LookupInfobaseResponse)(nil), // 4: ras.api.v1.LookupInfobaseResponse
	(*FieldFilter)(nil),            // 5: ras.api.v1.FieldFilter
	(*Cluster)(nil),                // 6: ras.api.v1.Cluster
	(*v1.InfobaseSummaryInfo)(nil), // 7: v8platform.serialize.v1.InfobaseSummaryInfo
	(*anypb.Any)(nil),              // 8: google.protobuf.Any
	(*v1.InfobaseInfo)(nil),        // 9: v8platform.serialize.v1.InfobaseInfo
}
var file_ras_api_v1_infobases_request_proto_depIdxs = []int32{
	6,  // 0: ras.api.v1.InfobasesRequest.cluster:type_name -> ras.api.v1.Cluster
	7,  // 1: ras.api.v1.InfobasesResponse.items:type_name -> v8platform.serialize.v1.InfobaseSummaryInfo
	6,  // 2: ras.api.v1.GetInfobaseRequest.cluster:type_name -> ras.api.v1.Cluster
	6,  // 3: ras.api.v1.LookupInfobaseRequest.cluster:type_name -> ras.api.v1.Cluster
	5,  // 4: ras.api.v1.LookupInfobaseRequest.field:type_name -> ras.api.v1.FieldFilter
	7,  // 5: ras.api.v1.LookupInfobaseResponse.items:type_name -> v8platform.serialize.v1.InfobaseSummaryInfo
	8,  // 6: ras.api.v1.FieldFilter.value:type_name -> google.protobuf.Any
	0,  // 7: ras.api.v1.InfobasesService.Infobases:input_type -> ras.api.v1.InfobasesRequest
	3,  // 8: ras.api.v1.InfobasesService.LookupInfobase:input_type -> ras.api.v1.LookupInfobaseRequest
	2,  // 9: ras.api.v1.InfobasesService.GetInfobase:input_type -> ras.api.v1.GetInfobaseRequest
	1,  // 10: ras.api.v1.InfobasesService.Infobases:output_type -> ras.api.v1.InfobasesResponse
	4,  // 11: ras.api.v1.InfobasesService.LookupInfobase:output_type -> ras.api.v1.LookupInfobaseResponse
	9,  // 12: ras.api.v1.InfobasesService.GetInfobase:output_type -> v8platform.serialize.v1.InfobaseInfo
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_ras_api_v1_infobases_request_proto_init() }
func file_ras_api_v1_infobases_request_proto_init() {
	if File_ras_api_v1_infobases_request_proto != nil {
		return
	}
	file_ras_api_v1_helpers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ras_api_v1_infobases_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfobasesRequest); i {
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
		file_ras_api_v1_infobases_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfobasesResponse); i {
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
		file_ras_api_v1_infobases_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfobaseRequest); i {
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
		file_ras_api_v1_infobases_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupInfobaseRequest); i {
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
		file_ras_api_v1_infobases_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupInfobaseResponse); i {
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
		file_ras_api_v1_infobases_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldFilter); i {
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
			RawDescriptor: file_ras_api_v1_infobases_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ras_api_v1_infobases_request_proto_goTypes,
		DependencyIndexes: file_ras_api_v1_infobases_request_proto_depIdxs,
		MessageInfos:      file_ras_api_v1_infobases_request_proto_msgTypes,
	}.Build()
	File_ras_api_v1_infobases_request_proto = out.File
	file_ras_api_v1_infobases_request_proto_rawDesc = nil
	file_ras_api_v1_infobases_request_proto_goTypes = nil
	file_ras_api_v1_infobases_request_proto_depIdxs = nil
}
