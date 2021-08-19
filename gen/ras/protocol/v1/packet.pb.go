// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: ras/protocol/v1/packet.proto

package protocolv1

import (
	_ "github.com/v8platform/protos/gen/ras/encoding"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   PacketType `protobuf:"varint,1,opt,name=Type,proto3,enum=ras.protocol.v1.PacketType" json:"Type,omitempty" ras:"byte,1"`
	Length int32      `protobuf:"varint,2,opt,name=Length,proto3" json:"Length,omitempty" ras:"size,2"`
	Data   []byte     `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty" ras:"bytes,3"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_v1_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_v1_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_ras_protocol_v1_packet_proto_rawDescGZIP(), []int{0}
}

func (x *Packet) GetType() PacketType {
	if x != nil {
		return x.Type
	}
	return PacketType_PACKET_TYPE_NEGOTIATE
}

func (x *Packet) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Packet) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ras_protocol_v1_packet_proto protoreflect.FileDescriptor

var file_ras_protocol_v1_packet_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x61,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x61, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f,
	0x01, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x11, 0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x72, 0x61, 0x73, 0x3a, 0x22,
	0x62, 0x79, 0x74, 0x65, 0x2c, 0x31, 0x22, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a,
	0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x11, 0x9a,
	0x84, 0x9e, 0x03, 0x0c, 0x72, 0x61, 0x73, 0x3a, 0x22, 0x73, 0x69, 0x7a, 0x65, 0x2c, 0x32, 0x22,
	0x52, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d, 0x72, 0x61, 0x73,
	0x3a, 0x22, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2c, 0x33, 0x22, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x42, 0xc9, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x50, 0x58, 0xaa, 0x02, 0x0f, 0x52, 0x61, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x52,
	0x61, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x52, 0x61, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x52,
	0x61, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x3a, 0x56, 0x31,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x06, 0x08, 0x01, 0x10, 0x01, 0x18, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_protocol_v1_packet_proto_rawDescOnce sync.Once
	file_ras_protocol_v1_packet_proto_rawDescData = file_ras_protocol_v1_packet_proto_rawDesc
)

func file_ras_protocol_v1_packet_proto_rawDescGZIP() []byte {
	file_ras_protocol_v1_packet_proto_rawDescOnce.Do(func() {
		file_ras_protocol_v1_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_protocol_v1_packet_proto_rawDescData)
	})
	return file_ras_protocol_v1_packet_proto_rawDescData
}

var file_ras_protocol_v1_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ras_protocol_v1_packet_proto_goTypes = []interface{}{
	(*Packet)(nil),  // 0: ras.protocol.v1.Packet
	(PacketType)(0), // 1: ras.protocol.v1.PacketType
}
var file_ras_protocol_v1_packet_proto_depIdxs = []int32{
	1, // 0: ras.protocol.v1.Packet.Type:type_name -> ras.protocol.v1.PacketType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ras_protocol_v1_packet_proto_init() }
func file_ras_protocol_v1_packet_proto_init() {
	if File_ras_protocol_v1_packet_proto != nil {
		return
	}
	file_ras_protocol_v1_param_proto_init()
	file_ras_protocol_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ras_protocol_v1_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
			RawDescriptor: file_ras_protocol_v1_packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_protocol_v1_packet_proto_goTypes,
		DependencyIndexes: file_ras_protocol_v1_packet_proto_depIdxs,
		MessageInfos:      file_ras_protocol_v1_packet_proto_msgTypes,
	}.Build()
	File_ras_protocol_v1_packet_proto = out.File
	file_ras_protocol_v1_packet_proto_rawDesc = nil
	file_ras_protocol_v1_packet_proto_goTypes = nil
	file_ras_protocol_v1_packet_proto_depIdxs = nil
}