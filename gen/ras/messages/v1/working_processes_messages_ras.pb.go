// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package messagesv1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	v1 "github.com/v8platform/protos/gen/v8platform/serialize/v1"
	io "io"
)

func (x *GetWorkingProcessesRequest) GetMessageType() MessageType {
	return MessageType_GET_WORKING_PROCESSES_REQUEST
}

func (x *GetWorkingProcessesRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingProcessesRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingProcessesResponse) GetMessageType() MessageType {
	return MessageType_GET_WORKING_PROCESSES_RESPONSE
}

func (x *GetWorkingProcessesResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Processes opts: order:1
	var size_Processes int
	if err := codec256.ParseSize(reader, &size_Processes); err != nil {
		return err
	}
	for i := 0; i < size_Processes; i++ {
		val := &v1.ProcessInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Processes = append(x.Processes, val)
	}
	return nil
}
func (x *GetWorkingProcessesResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Processes opts: order:1
	if err := codec256.FormatSize(writer, len(x.Processes)); err != nil {
		return err
	}
	for i := 0; i < len(x.Processes); i++ {
		if err := x.Processes[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	return nil
}
func (x *GetWorkingProcessInfoRequest) GetMessageType() MessageType {
	return MessageType_GET_WORKING_PROCESS_INFO_REQUEST
}

func (x *GetWorkingProcessInfoRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.ProcessId opts: encoder:"uuid" order:2
	if err := codec256.ParseUUID(reader, &x.ProcessId); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingProcessInfoRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.ProcessId opts: encoder:"uuid" order:2
	if err := codec256.FormatUuid(writer, x.ProcessId); err != nil {
		return err
	}
	return nil
}
func (x *GetWorkingProcessInfoResponse) GetMessageType() MessageType {
	return MessageType_GET_WORKING_PROCESS_INFO_RESPONSE
}

func (x *GetWorkingProcessInfoResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Info opts: order:1
	x.Info = &v1.ProcessInfo{}
	if err := x.Info.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *GetWorkingProcessInfoResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Info opts: order:1
	if err := x.Info.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
func (x *GetServerWorkingProcessesRequest) GetMessageType() MessageType {
	return MessageType_GET_SERVER_WORKING_PROCESSES_REQUEST
}

func (x *GetServerWorkingProcessesRequest) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.ParseUUID(reader, &x.ClusterId); err != nil {
		return err
	}
	// decode x.ServerId opts: encoder:"uuid" order:2
	if err := codec256.ParseUUID(reader, &x.ServerId); err != nil {
		return err
	}
	return nil
}
func (x *GetServerWorkingProcessesRequest) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.ClusterId opts: encoder:"uuid" order:1
	if err := codec256.FormatUuid(writer, x.ClusterId); err != nil {
		return err
	}
	// decode x.ServerId opts: encoder:"uuid" order:2
	if err := codec256.FormatUuid(writer, x.ServerId); err != nil {
		return err
	}
	return nil
}
func (x *GetServerWorkingProcessesResponse) GetMessageType() MessageType {
	return MessageType_GET_SERVER_WORKING_PROCESSES_RESPONSE
}

func (x *GetServerWorkingProcessesResponse) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Processes opts: order:1
	x.Processes = &v1.ProcessInfo{}
	if err := x.Processes.Parse(reader, version); err != nil {
		return err
	}

	return nil
}
func (x *GetServerWorkingProcessesResponse) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Processes opts: order:1
	if err := x.Processes.Formatter(writer, version); err != nil {
		return err
	}
	return nil
}
