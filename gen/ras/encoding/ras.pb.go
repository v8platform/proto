// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: ras/encoding/ras.proto

package encoding

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EncodingMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenerateMarshaller bool `protobuf:"varint,1,opt,name=generate_marshaller,json=generateMarshaller,proto3" json:"generate_marshaller,omitempty"`
	GenerateUnmarshal  bool `protobuf:"varint,2,opt,name=generate_unmarshal,json=generateUnmarshal,proto3" json:"generate_unmarshal,omitempty"`
	GenerateTags       bool `protobuf:"varint,3,opt,name=generate_tags,json=generateTags,proto3" json:"generate_tags,omitempty"`
}

func (x *EncodingMessageOptions) Reset() {
	*x = EncodingMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_encoding_ras_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodingMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodingMessageOptions) ProtoMessage() {}

func (x *EncodingMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ras_encoding_ras_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodingMessageOptions.ProtoReflect.Descriptor instead.
func (*EncodingMessageOptions) Descriptor() ([]byte, []int) {
	return file_ras_encoding_ras_proto_rawDescGZIP(), []int{0}
}

func (x *EncodingMessageOptions) GetGenerateMarshaller() bool {
	if x != nil {
		return x.GenerateMarshaller
	}
	return false
}

func (x *EncodingMessageOptions) GetGenerateUnmarshal() bool {
	if x != nil {
		return x.GenerateUnmarshal
	}
	return false
}

func (x *EncodingMessageOptions) GetGenerateTags() bool {
	if x != nil {
		return x.GenerateTags
	}
	return false
}

type EncodingFileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenerateMarshaller bool `protobuf:"varint,1,opt,name=generate_marshaller,json=generateMarshaller,proto3" json:"generate_marshaller,omitempty"`
	GenerateUnmarshal  bool `protobuf:"varint,2,opt,name=generate_unmarshal,json=generateUnmarshal,proto3" json:"generate_unmarshal,omitempty"`
	GenerateTags       bool `protobuf:"varint,3,opt,name=generate_tags,json=generateTags,proto3" json:"generate_tags,omitempty"`
}

func (x *EncodingFileOptions) Reset() {
	*x = EncodingFileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_encoding_ras_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodingFileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodingFileOptions) ProtoMessage() {}

func (x *EncodingFileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ras_encoding_ras_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodingFileOptions.ProtoReflect.Descriptor instead.
func (*EncodingFileOptions) Descriptor() ([]byte, []int) {
	return file_ras_encoding_ras_proto_rawDescGZIP(), []int{1}
}

func (x *EncodingFileOptions) GetGenerateMarshaller() bool {
	if x != nil {
		return x.GenerateMarshaller
	}
	return false
}

func (x *EncodingFileOptions) GetGenerateUnmarshal() bool {
	if x != nil {
		return x.GenerateUnmarshal
	}
	return false
}

func (x *EncodingFileOptions) GetGenerateTags() bool {
	if x != nil {
		return x.GenerateTags
	}
	return false
}

type EncodingFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encoder     *string           `protobuf:"bytes,1,opt,name=encoder,proto3,oneof" json:"encoder,omitempty"`
	Order       *int32            `protobuf:"varint,2,opt,name=order,proto3,oneof" json:"order,omitempty"`
	Version     *int32            `protobuf:"varint,3,opt,name=version,proto3,oneof" json:"version,omitempty"`
	Opts        map[string]string `protobuf:"bytes,4,rep,name=opts,proto3" json:"opts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NoMarshal   bool              `protobuf:"varint,5,opt,name=no_marshal,json=noMarshal,proto3" json:"no_marshal,omitempty"`
	NoUnmarshal bool              `protobuf:"varint,6,opt,name=no_unmarshal,json=noUnmarshal,proto3" json:"no_unmarshal,omitempty"`
	TypeField   *int32            `protobuf:"varint,7,opt,name=type_field,json=typeField,proto3,oneof" json:"type_field,omitempty"`
	TypeUrl     string            `protobuf:"bytes,8,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Ignore      bool              `protobuf:"varint,9,opt,name=ignore,proto3" json:"ignore,omitempty"`
	SizeField   *int32            `protobuf:"varint,10,opt,name=size_field,json=sizeField,proto3,oneof" json:"size_field,omitempty"`
}

func (x *EncodingFieldOptions) Reset() {
	*x = EncodingFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_encoding_ras_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodingFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodingFieldOptions) ProtoMessage() {}

func (x *EncodingFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_ras_encoding_ras_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodingFieldOptions.ProtoReflect.Descriptor instead.
func (*EncodingFieldOptions) Descriptor() ([]byte, []int) {
	return file_ras_encoding_ras_proto_rawDescGZIP(), []int{2}
}

func (x *EncodingFieldOptions) GetEncoder() string {
	if x != nil && x.Encoder != nil {
		return *x.Encoder
	}
	return ""
}

func (x *EncodingFieldOptions) GetOrder() int32 {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return 0
}

func (x *EncodingFieldOptions) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *EncodingFieldOptions) GetOpts() map[string]string {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *EncodingFieldOptions) GetNoMarshal() bool {
	if x != nil {
		return x.NoMarshal
	}
	return false
}

func (x *EncodingFieldOptions) GetNoUnmarshal() bool {
	if x != nil {
		return x.NoUnmarshal
	}
	return false
}

func (x *EncodingFieldOptions) GetTypeField() int32 {
	if x != nil && x.TypeField != nil {
		return *x.TypeField
	}
	return 0
}

func (x *EncodingFieldOptions) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *EncodingFieldOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *EncodingFieldOptions) GetSizeField() int32 {
	if x != nil && x.SizeField != nil {
		return *x.SizeField
	}
	return 0
}

var file_ras_encoding_ras_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*EncodingMessageOptions)(nil),
		Field:         475223888,
		Name:          "ras.encoding.message_options",
		Tag:           "bytes,475223888,opt,name=message_options",
		Filename:      "ras/encoding/ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*EncodingFileOptions)(nil),
		Field:         475223888,
		Name:          "ras.encoding.options",
		Tag:           "bytes,475223888,opt,name=options",
		Filename:      "ras/encoding/ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847939,
		Name:          "ras.encoding.tags",
		Tag:           "bytes,847939,opt,name=tags",
		Filename:      "ras/encoding/ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         847939,
		Name:          "ras.encoding.oneof_tags",
		Tag:           "bytes,847939,opt,name=oneof_tags",
		Filename:      "ras/encoding/ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*EncodingFieldOptions)(nil),
		Field:         475223888,
		Name:          "ras.encoding.field",
		Tag:           "bytes,475223888,opt,name=field",
		Filename:      "ras/encoding/ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*EncodingFieldOptions)(nil),
		Field:         475223888,
		Name:          "ras.encoding.oneof_field",
		Tag:           "bytes,475223888,opt,name=oneof_field",
		Filename:      "ras/encoding/ras.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         475223891,
		Name:          "ras.encoding.type_url",
		Tag:           "bytes,475223891,opt,name=type_url",
		Filename:      "ras/encoding/ras.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional ras.encoding.EncodingMessageOptions message_options = 475223888;
	E_MessageOptions = &file_ras_encoding_ras_proto_extTypes[0]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional ras.encoding.EncodingFileOptions options = 475223888;
	E_Options = &file_ras_encoding_ras_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// Multiple Tags can be spcified.
	//
	// optional string tags = 847939;
	E_Tags = &file_ras_encoding_ras_proto_extTypes[2]
	// Multiple Tags can be spcified.
	//
	// optional ras.encoding.EncodingFieldOptions field = 475223888;
	E_Field = &file_ras_encoding_ras_proto_extTypes[4]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// Multiple Tags can be spcified.
	//
	// optional string oneof_tags = 847939;
	E_OneofTags = &file_ras_encoding_ras_proto_extTypes[3]
	// optional ras.encoding.EncodingFieldOptions oneof_field = 475223888;
	E_OneofField = &file_ras_encoding_ras_proto_extTypes[5]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// Ссылка на тип связанного сообщения
	//
	// optional string type_url = 475223891;
	E_TypeUrl = &file_ras_encoding_ras_proto_extTypes[6]
)

var File_ras_encoding_ras_proto protoreflect.FileDescriptor

var file_ras_encoding_ras_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x72,
	0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x16, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73,
	0x68, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x13, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2f, 0x0a, 0x13, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x72,
	0x73, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x6e,
	0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x73, 0x22, 0xe7, 0x03, 0x0a, 0x14, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d,
	0x0a, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x5f,
	0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6e,
	0x6f, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x75,
	0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x6e, 0x6f, 0x55, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0a, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x03, 0x52, 0x09, 0x74, 0x79, 0x70, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x4f, 0x70, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a,
	0x72, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd0, 0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x5d, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0xae, 0xcd,
	0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x3a, 0x33, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc3, 0xe0, 0x33, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x3e, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc3, 0xe0, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x54, 0x61, 0x67, 0x73, 0x3a, 0x5b, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd0, 0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x61, 0x73, 0x2e,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x3a, 0x66, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xd0, 0xae, 0xcd, 0xe2, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72,
	0x61, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x40, 0x0a, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0xae, 0xcd, 0xe2,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x9c,
	0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x08, 0x52, 0x61, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0xa2, 0x02,
	0x03, 0x52, 0x45, 0x58, 0xaa, 0x02, 0x0c, 0x52, 0x61, 0x73, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0xca, 0x02, 0x0c, 0x52, 0x61, 0x73, 0x5c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0xe2, 0x02, 0x18, 0x52, 0x61, 0x73, 0x5c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d,
	0x52, 0x61, 0x73, 0x3a, 0x3a, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_encoding_ras_proto_rawDescOnce sync.Once
	file_ras_encoding_ras_proto_rawDescData = file_ras_encoding_ras_proto_rawDesc
)

func file_ras_encoding_ras_proto_rawDescGZIP() []byte {
	file_ras_encoding_ras_proto_rawDescOnce.Do(func() {
		file_ras_encoding_ras_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_encoding_ras_proto_rawDescData)
	})
	return file_ras_encoding_ras_proto_rawDescData
}

var file_ras_encoding_ras_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ras_encoding_ras_proto_goTypes = []interface{}{
	(*EncodingMessageOptions)(nil),        // 0: ras.encoding.EncodingMessageOptions
	(*EncodingFileOptions)(nil),           // 1: ras.encoding.EncodingFileOptions
	(*EncodingFieldOptions)(nil),          // 2: ras.encoding.EncodingFieldOptions
	nil,                                   // 3: ras.encoding.EncodingFieldOptions.OptsEntry
	(*descriptorpb.MessageOptions)(nil),   // 4: google.protobuf.MessageOptions
	(*descriptorpb.FileOptions)(nil),      // 5: google.protobuf.FileOptions
	(*descriptorpb.FieldOptions)(nil),     // 6: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil),     // 7: google.protobuf.OneofOptions
	(*descriptorpb.EnumValueOptions)(nil), // 8: google.protobuf.EnumValueOptions
}
var file_ras_encoding_ras_proto_depIdxs = []int32{
	3,  // 0: ras.encoding.EncodingFieldOptions.opts:type_name -> ras.encoding.EncodingFieldOptions.OptsEntry
	4,  // 1: ras.encoding.message_options:extendee -> google.protobuf.MessageOptions
	5,  // 2: ras.encoding.options:extendee -> google.protobuf.FileOptions
	6,  // 3: ras.encoding.tags:extendee -> google.protobuf.FieldOptions
	7,  // 4: ras.encoding.oneof_tags:extendee -> google.protobuf.OneofOptions
	6,  // 5: ras.encoding.field:extendee -> google.protobuf.FieldOptions
	7,  // 6: ras.encoding.oneof_field:extendee -> google.protobuf.OneofOptions
	8,  // 7: ras.encoding.type_url:extendee -> google.protobuf.EnumValueOptions
	0,  // 8: ras.encoding.message_options:type_name -> ras.encoding.EncodingMessageOptions
	1,  // 9: ras.encoding.options:type_name -> ras.encoding.EncodingFileOptions
	2,  // 10: ras.encoding.field:type_name -> ras.encoding.EncodingFieldOptions
	2,  // 11: ras.encoding.oneof_field:type_name -> ras.encoding.EncodingFieldOptions
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	8,  // [8:12] is the sub-list for extension type_name
	1,  // [1:8] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_ras_encoding_ras_proto_init() }
func file_ras_encoding_ras_proto_init() {
	if File_ras_encoding_ras_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ras_encoding_ras_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodingMessageOptions); i {
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
		file_ras_encoding_ras_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodingFileOptions); i {
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
		file_ras_encoding_ras_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodingFieldOptions); i {
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
	file_ras_encoding_ras_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ras_encoding_ras_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 7,
			NumServices:   0,
		},
		GoTypes:           file_ras_encoding_ras_proto_goTypes,
		DependencyIndexes: file_ras_encoding_ras_proto_depIdxs,
		MessageInfos:      file_ras_encoding_ras_proto_msgTypes,
		ExtensionInfos:    file_ras_encoding_ras_proto_extTypes,
	}.Build()
	File_ras_encoding_ras_proto = out.File
	file_ras_encoding_ras_proto_rawDesc = nil
	file_ras_encoding_ras_proto_goTypes = nil
	file_ras_encoding_ras_proto_depIdxs = nil
}
