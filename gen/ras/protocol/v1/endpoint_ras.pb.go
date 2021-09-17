// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package protocolv1

import (
	bytes "bytes"
	fmt "fmt"
	codec256 "github.com/v8platform/encoder/ras/codec256"
	v1 "github.com/v8platform/protos/gen/ras/messages/v1"
	io "io"
)

func (x *EndpointOpen) GetPacketType() PacketType {
	return PacketType_PACKET_TYPE_ENDPOINT_OPEN
}

func (x *EndpointOpen) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Service opts: order:1
	if err := codec256.ParseString(reader, &x.Service); err != nil {
		return err
	}
	// decode x.Version opts: encoder:"string" order:3
	if err := codec256.ParseString(reader, &x.Version); err != nil {
		return err
	}
	// decode x.Params opts: order:4
	// TODO parse map
	return nil
}
func (x *EndpointOpen) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Service opts: order:1
	if err := codec256.FormatString(writer, x.Service); err != nil {
		return err
	}
	// decode x.Version opts: encoder:"string" order:3
	if err := codec256.FormatString(writer, x.Version); err != nil {
		return err
	}
	// decode x.Params opts: order:4
	if err := codec256.FormatSize(writer, len(x.Params)); err != nil {
		return err
	}
	return nil
}
func (x *EndpointOpenAck) GetPacketType() PacketType {
	return PacketType_PACKET_TYPE_ENDPOINT_OPEN_ACK
}

func (x *EndpointOpenAck) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Service opts: order:1
	if err := codec256.ParseString(reader, &x.Service); err != nil {
		return err
	}
	// decode x.Version opts: encoder:"string" order:2
	if err := codec256.ParseString(reader, &x.Version); err != nil {
		return err
	}
	// decode x.EndpointId opts: encoder:"nullable" order:3
	if err := codec256.ParseNullable(reader, &x.EndpointId); err != nil {
		return err
	}
	// decode x.Params opts: order:4
	// TODO parse map
	return nil
}
func (x *EndpointOpenAck) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Service opts: order:1
	if err := codec256.FormatString(writer, x.Service); err != nil {
		return err
	}
	// decode x.Version opts: encoder:"string" order:2
	if err := codec256.FormatString(writer, x.Version); err != nil {
		return err
	}
	// decode x.EndpointId opts: encoder:"nullable" order:3
	if err := codec256.FormatNullable(writer, x.EndpointId); err != nil {
		return err
	}
	// decode x.Params opts: order:4
	if err := codec256.FormatSize(writer, len(x.Params)); err != nil {
		return err
	}
	return nil
}

// Helpers generated by github.com/v8platform/protoc-gen-go-ras. DO NOT EDIT

type EndpointImpl interface {
	GetVersion() int32
	GetId() int32
	GetService() string
	GetFormat() int32
	NewMessage(data interface{}) (*EndpointMessage, error)
	UnpackMessage(data interface{}, into EndpointMessageParser) error
}

func NewEndpoint(id int32, version int32) EndpointImpl {
	return &Endpoint{
		Service: "v8.service.Admin.Cluster",
		Version: version,
		Id:      id,
		Format:  codec256.Version()}
}

func (x *Endpoint) NewMessage(data interface{}) (*EndpointMessage, error) {
	switch typed := data.(type) {
	case io.Reader:
		packet, err := NewPacket(data)
		if err != nil {
			return nil, err
		}
		var message EndpointMessage
		if err := packet.Unpack(&message); err != nil {
			return nil, err
		}
		return &message, nil
	case EndpointMessageFormatter:
		return NewEndpointMessage(x, typed)
	default:
		return nil, fmt.Errorf("unknown type <%T> to create new message", typed)
	}
}

func (x *Endpoint) UnpackMessage(data interface{}, into EndpointMessageParser) error {
	switch typed := data.(type) {
	case io.Reader:
		packet, err := NewPacket(data)
		if err != nil {
			return err
		}
		var message EndpointMessage
		if err := packet.Unpack(&message); err != nil {
			return err
		}
		return message.Unpack(x, into)
	case Packet:
		var message EndpointMessage
		if err := typed.Unpack(&message); err != nil {
			return err
		}
		return message.Unpack(x, into)
	case *EndpointMessage:
		return typed.Unpack(x, into)
	default:
		return fmt.Errorf("unknown type <%T> to create unpack message", typed)
	}
}

func (x *EndpointMessage) GetPacketType() PacketType {
	return PacketType_PACKET_TYPE_ENDPOINT_MESSAGE
}

// Helpers generated by github.com/v8platform/protoc-r-go-ras. DO NOT EDIT

type EndpointMessageFormatter interface {
	GetMessageType() v1.MessageType
	Formatter(writer io.Writer, version int32) error
}

type EndpointMessageParser interface {
	GetMessageType() v1.MessageType
	Parse(reader io.Reader, version int32) error
}

func NewEndpointMessage(endpoint EndpointImpl, message EndpointMessageFormatter) (*EndpointMessage, error) {
	buf := &bytes.Buffer{}
	if err := message.Formatter(buf, endpoint.GetVersion()); err != nil {
		return nil, err
	}
	return &EndpointMessage{
		Type:       EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE,
		Format:     endpoint.GetFormat(),
		EndpointId: endpoint.GetId(),
		Data: &EndpointMessage_Message{
			Message: &EndpointDataMessage{
				Bytes: buf.Bytes(),
				Type:  message.GetMessageType(),
			},
		},
	}, nil
}

func (x *EndpointMessage) Unpack(endpoint EndpointImpl, into EndpointMessageParser) error {
	switch x.GetType() {
	case EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE:
		buf := bytes.NewBuffer(x.GetMessage().GetBytes())
		if err := into.Parse(buf, endpoint.GetVersion()); err != nil {
			return err
		}
		return nil
	case EndpointDataType_ENDPOINT_DATA_TYPE_VOID_MESSAGE:
		return nil
	case EndpointDataType_ENDPOINT_DATA_TYPE_EXCEPTION:
		return x.GetFailure()
	default:
		return fmt.Errorf("unknown message type <%s>", x.GetType())
	}
}

func (x *EndpointMessage) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.EndpointId opts: encoder:"nullable" order:1
	if err := codec256.ParseNullable(reader, &x.EndpointId); err != nil {
		return err
	}
	// decode x.Format opts: encoder:"short" order:2
	if err := codec256.ParseShort(reader, &x.Format); err != nil {
		return err
	}
	// decode x.Type opts: encoder:"byte" order:3
	var val_Type int32
	if err := codec256.ParseByte(reader, &val_Type); err != nil {
		return err
	}
	x.Type = EndpointDataType(val_Type)
	// decode x.VoidMessage opts: order:4 type_field:3
	if x.GetType() == EndpointDataType_ENDPOINT_DATA_TYPE_VOID_MESSAGE {
		void_message := &EndpointDataVoidMessage{}
		if err := void_message.Parse(reader, version); err != nil {
			return err
		}

		x.Data = &EndpointMessage_VoidMessage{
			VoidMessage: void_message,
		}
	}
	// decode x.Message opts: order:4 type_field:3
	if x.GetType() == EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE {
		message := &EndpointDataMessage{}
		if err := message.Parse(reader, version); err != nil {
			return err
		}

		x.Data = &EndpointMessage_Message{
			Message: message,
		}
	}
	// decode x.Failure opts: order:4 type_field:3
	if x.GetType() == EndpointDataType_ENDPOINT_DATA_TYPE_EXCEPTION {
		failure := &EndpointFailureMessage{}
		if err := failure.Parse(reader, version); err != nil {
			return err
		}

		x.Data = &EndpointMessage_Failure{
			Failure: failure,
		}
	}
	return nil
}
func (x *EndpointMessage) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.EndpointId opts: encoder:"nullable" order:1
	if err := codec256.FormatNullable(writer, x.EndpointId); err != nil {
		return err
	}
	// decode x.Format opts: encoder:"short" order:2
	if err := codec256.FormatShort(writer, x.Format); err != nil {
		return err
	}
	// decode x.Type opts: encoder:"byte" order:3
	if err := codec256.FormatByte(writer, int32(x.Type)); err != nil {
		return err
	}
	// decode x.VoidMessage opts: order:4 type_field:3
	if val := x.GetVoidMessage(); val != nil {
		if err := val.Formatter(writer, version); err != nil {
			return err
		}
	}
	// decode x.Message opts: order:4 type_field:3
	if val := x.GetMessage(); val != nil {
		if err := val.Formatter(writer, version); err != nil {
			return err
		}
	}
	// decode x.Failure opts: order:4 type_field:3
	if val := x.GetFailure(); val != nil {
		if err := val.Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *EndpointFailureAck) GetPacketType() PacketType {
	return PacketType_PACKET_TYPE_ENDPOINT_FAILURE
}

func (x *EndpointFailureAck) Error() string {
	return x.String()
}

func (x *EndpointFailureAck) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ServiceId opts: order:1
	if err := codec256.ParseString(reader, &x.ServiceId); err != nil {
		return err
	}
	// decode x.Version opts: order:2
	if err := codec256.ParseString(reader, &x.Version); err != nil {
		return err
	}
	// decode x.EndpointId opts: encoder:"nullable" order:3
	if err := codec256.ParseNullable(reader, &x.EndpointId); err != nil {
		return err
	}
	// decode x.ClassCause opts: order:4
	if err := codec256.ParseString(reader, &x.ClassCause); err != nil {
		return err
	}
	// decode x.Message opts: order:5
	if err := codec256.ParseString(reader, &x.Message); err != nil {
		return err
	}
	// decode x.Trace opts: order:6
	var size_Trace int
	if err := codec256.ParseSize(reader, &size_Trace); err != nil {
		return err
	}
	for i := 0; i < size_Trace; i++ {
		var val string
		if err := codec256.ParseString(reader, &val); err != nil {
			return err
		}
		x.Trace = append(x.Trace, val)
	}
	// decode x.Cause opts: order:7
	x.Cause = &CauseError{}
	if err := x.Cause.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *EndpointFailureAck) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ServiceId opts: order:1
	if err := codec256.FormatString(writer, x.ServiceId); err != nil {
		return err
	}
	// decode x.Version opts: order:2
	if err := codec256.FormatString(writer, x.Version); err != nil {
		return err
	}
	// decode x.EndpointId opts: encoder:"nullable" order:3
	if err := codec256.FormatNullable(writer, x.EndpointId); err != nil {
		return err
	}
	// decode x.ClassCause opts: order:4
	if err := codec256.FormatString(writer, x.ClassCause); err != nil {
		return err
	}
	// decode x.Message opts: order:5
	if err := codec256.FormatString(writer, x.Message); err != nil {
		return err
	}
	// decode x.Trace opts: order:6
	if err := codec256.FormatSize(writer, len(x.Trace)); err != nil {
		return err
	}
	for i := 0; i < len(x.Trace); i++ {
		if err := codec256.FormatString(writer, x.Trace[i]); err != nil {
			return err
		}
	}
	// decode x.Cause opts: order:7
	if err := x.Cause.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *CauseError) Error() string {
	return x.String()
}

func (x *CauseError) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ServiceId opts: order:1
	if err := codec256.ParseString(reader, &x.ServiceId); err != nil {
		return err
	}
	// decode x.Message opts: order:2
	if err := codec256.ParseString(reader, &x.Message); err != nil {
		return err
	}
	return nil
}
func (x *CauseError) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ServiceId opts: order:1
	if err := codec256.FormatString(writer, x.ServiceId); err != nil {
		return err
	}
	// decode x.Message opts: order:2
	if err := codec256.FormatString(writer, x.Message); err != nil {
		return err
	}
	return nil
}
func (x *EndpointDataVoidMessage) GetEndpointDataType() EndpointDataType {
	return EndpointDataType_ENDPOINT_DATA_TYPE_VOID_MESSAGE
}

func (x *EndpointDataVoidMessage) Parse(reader io.Reader, version int32) error {
	return nil
}
func (x *EndpointDataVoidMessage) Formatter(writer io.Writer, version int32) error {
	return nil
}
func (x *EndpointDataMessage) GetEndpointDataType() EndpointDataType {
	return EndpointDataType_ENDPOINT_DATA_TYPE_MESSAGE
}

func (x *EndpointDataMessage) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Type opts: encoder:"byte" order:1
	var val_Type int32
	if err := codec256.ParseByte(reader, &val_Type); err != nil {
		return err
	}
	x.Type = v1.MessageType(val_Type)
	// decode x.Bytes opts: order:2
	var err error
	x.Bytes, err = io.ReadAll(reader)
	if err != nil {
		return err
	}
	return nil
}
func (x *EndpointDataMessage) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Type opts: encoder:"byte" order:1
	if err := codec256.FormatByte(writer, int32(x.Type)); err != nil {
		return err
	}
	// decode x.Bytes opts: order:2
	if err := codec256.FormatBytes(writer, x.Bytes); err != nil {
		return err
	}
	return nil
}
func (x *EndpointFailureMessage) GetEndpointDataType() EndpointDataType {
	return EndpointDataType_ENDPOINT_DATA_TYPE_EXCEPTION
}

func (x *EndpointFailureMessage) Error() string {
	return x.String()
}

func (x *EndpointFailureMessage) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ServiceId opts: order:1
	if err := codec256.ParseString(reader, &x.ServiceId); err != nil {
		return err
	}
	// decode x.Message opts: order:2
	if err := codec256.ParseString(reader, &x.Message); err != nil {
		return err
	}
	// decode x.Cause opts: order:3
	x.Cause = &CauseError{}
	if err := x.Cause.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *EndpointFailureMessage) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ServiceId opts: order:1
	if err := codec256.FormatString(writer, x.ServiceId); err != nil {
		return err
	}
	// decode x.Message opts: order:2
	if err := codec256.FormatString(writer, x.Message); err != nil {
		return err
	}
	// decode x.Cause opts: order:3
	if err := x.Cause.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
