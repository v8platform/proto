// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package messagesv1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	v1 "github.com/v8platform/protos/gen/v8platform/serialize/v1"
	io "io"
)

func (x *GetConnectionsRequest) GetMessageType() MessageType {
	return MessageType_GET_CONNECTIONS_SHORT_REQUEST
}

func (x *GetConnectionsRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetConnectionsRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetConnectionsResponse) GetMessageType() MessageType {
	return MessageType_GET_CONNECTIONS_SHORT_RESPONSE
}

func (x *GetConnectionsResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Connections opts: order:1
	var size_Connections int
	if err := codec256.ParseSize(reader, &size_Connections); err != nil {
		return err
	}
	for i := 0; i < size_Connections; i++ {
		val := &v1.ConnectionInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Connections = append(x.Connections, val)
	}
	return nil
}
func (x *GetConnectionsResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Connections opts: order:1
	if err := codec256.FormatSize(writer, len(x.Connections)); err != nil {
		return err
	}
	for i := 0; i < len(x.Connections); i++ {
		if err := x.Connections[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *GetInfobaseConnectionsRequest) GetMessageType() MessageType {
	return MessageType_GET_INFOBASE_CONNECTIONS_SHORT_REQUEST
}

func (x *GetInfobaseConnectionsRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.InfobaseId opts: encoder:"uuid"  order:2
	if err := codec256.ParseUUID(reader, &x.InfobaseId); err != nil {
		return err
	}
	return nil
}
func (x *GetInfobaseConnectionsRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.InfobaseId opts: encoder:"uuid"  order:2
	if err := codec256.FormatUuid(writer, x.InfobaseId); err != nil {
		return err
	}
	return nil
}
func (x *GetInfobaseConnectionsResponse) GetMessageType() MessageType {
	return MessageType_GET_INFOBASE_SESSIONS_RESPONSE
}

func (x *GetInfobaseConnectionsResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Connections opts: order:1
	var size_Connections int
	if err := codec256.ParseSize(reader, &size_Connections); err != nil {
		return err
	}
	for i := 0; i < size_Connections; i++ {
		val := &v1.ConnectionInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Connections = append(x.Connections, val)
	}
	return nil
}
func (x *GetInfobaseConnectionsResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Connections opts: order:1
	if err := codec256.FormatSize(writer, len(x.Connections)); err != nil {
		return err
	}
	for i := 0; i < len(x.Connections); i++ {
		if err := x.Connections[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *DisconnectConnectionRequest) GetMessageType() MessageType {
	return MessageType_DISCONNECT_REQUEST
}

func (x *DisconnectConnectionRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.ProcessId opts: encoder:"uuid"  order:2
	if err := codec256.ParseUUID(reader, &x.ProcessId); err != nil {
		return err
	}
	// decode x.ConnectionId opts: encoder:"uuid"  order:3
	if err := codec256.ParseUUID(reader, &x.ConnectionId); err != nil {
		return err
	}
	return nil
}
func (x *DisconnectConnectionRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.ProcessId opts: encoder:"uuid"  order:2
	if err := codec256.FormatUuid(writer, x.ProcessId); err != nil {
		return err
	}
	// decode x.ConnectionId opts: encoder:"uuid"  order:3
	if err := codec256.FormatUuid(writer, x.ConnectionId); err != nil {
		return err
	}
	return nil
}
