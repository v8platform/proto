// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package serializev1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
)

func (x *ProcessInfo) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Uuid opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.Uuid); err != nil {
		return err
	}
	// decode x.AvgBackCallTime opts: order:2
	if err := codec256.ParseDouble(reader, &x.AvgBackCallTime); err != nil {
		return err
	}
	// decode x.AvgCallTime opts: order:3
	if err := codec256.ParseDouble(reader, &x.AvgCallTime); err != nil {
		return err
	}
	// decode x.AvgDbCallTime opts: order:4
	if err := codec256.ParseDouble(reader, &x.AvgDbCallTime); err != nil {
		return err
	}
	// decode x.AvgLockCallTime opts: order:5
	if err := codec256.ParseDouble(reader, &x.AvgLockCallTime); err != nil {
		return err
	}
	// decode x.AvgServerCallTime opts: order:6
	if err := codec256.ParseDouble(reader, &x.AvgServerCallTime); err != nil {
		return err
	}
	// decode x.AvgThreads opts: order:7
	if err := codec256.ParseDouble(reader, &x.AvgThreads); err != nil {
		return err
	}
	// decode x.Capacity opts: order:8
	if err := codec256.ParseInt(reader, &x.Capacity); err != nil {
		return err
	}
	// decode x.Connections opts: order:9
	if err := codec256.ParseInt(reader, &x.Connections); err != nil {
		return err
	}
	// decode x.Host opts: order:10
	if err := codec256.ParseString(reader, &x.Host); err != nil {
		return err
	}
	// decode x.Enable opts: order:11
	if err := codec256.ParseBool(reader, &x.Enable); err != nil {
		return err
	}
	// decode x.Licenses opts: order:12
	var size_Licenses int
	if err := codec256.ParseSize(reader, &size_Licenses); err != nil {
		return err
	}
	for i := 0; i < size_Licenses; i++ {
		val := &LicenseInfo{}
		if err := val.Parse(reader, version); err != nil {
			return err
		}

		x.Licenses = append(x.Licenses, val)
	}
	// decode x.Port opts: encoder:"short"  order:13
	if err := codec256.ParseShort(reader, &x.Port); err != nil {
		return err
	}
	// decode x.MemoryExcessTime opts: order:14
	if err := codec256.ParseInt(reader, &x.MemoryExcessTime); err != nil {
		return err
	}
	// decode x.MemorySize opts: order:15
	if err := codec256.ParseInt(reader, &x.MemorySize); err != nil {
		return err
	}
	// decode x.Pid opts: order:16
	if err := codec256.ParseString(reader, &x.Pid); err != nil {
		return err
	}
	// decode x.Running opts: encoder:"int"  order:17
	if err := codec256.ParseInt(reader, &x.Running); err != nil {
		return err
	}
	// decode x.SelectionSize opts: order:18
	if err := codec256.ParseInt(reader, &x.SelectionSize); err != nil {
		return err
	}
	// decode x.StartedAt opts: encoder:"time"  order:18
	x.StartedAt = &timestamppb.Timestamp{}
	if err := codec256.ParseTime(reader, x.StartedAt); err != nil {
		return err
	}
	// decode x.Use opts: encoder:"int"  order:20
	if err := codec256.ParseInt(reader, &x.Use); err != nil {
		return err
	}
	// decode x.AvailablePerfomance opts: order:21
	if err := codec256.ParseInt(reader, &x.AvailablePerfomance); err != nil {
		return err
	}
	if version >= 9 {
		// decode x.Reverse opts: order:22  version:9
		if err := codec256.ParseBool(reader, &x.Reverse); err != nil {
			return err
		}
	}
	return nil
}
func (x *ProcessInfo) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Uuid opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.Uuid); err != nil {
		return err
	}
	// decode x.AvgBackCallTime opts: order:2
	if err := codec256.FormatDouble(writer, x.AvgBackCallTime); err != nil {
		return err
	}
	// decode x.AvgCallTime opts: order:3
	if err := codec256.FormatDouble(writer, x.AvgCallTime); err != nil {
		return err
	}
	// decode x.AvgDbCallTime opts: order:4
	if err := codec256.FormatDouble(writer, x.AvgDbCallTime); err != nil {
		return err
	}
	// decode x.AvgLockCallTime opts: order:5
	if err := codec256.FormatDouble(writer, x.AvgLockCallTime); err != nil {
		return err
	}
	// decode x.AvgServerCallTime opts: order:6
	if err := codec256.FormatDouble(writer, x.AvgServerCallTime); err != nil {
		return err
	}
	// decode x.AvgThreads opts: order:7
	if err := codec256.FormatDouble(writer, x.AvgThreads); err != nil {
		return err
	}
	// decode x.Capacity opts: order:8
	if err := codec256.FormatInt(writer, x.Capacity); err != nil {
		return err
	}
	// decode x.Connections opts: order:9
	if err := codec256.FormatInt(writer, x.Connections); err != nil {
		return err
	}
	// decode x.Host opts: order:10
	if err := codec256.FormatString(writer, x.Host); err != nil {
		return err
	}
	// decode x.Enable opts: order:11
	if err := codec256.FormatBool(writer, x.Enable); err != nil {
		return err
	}
	// decode x.Licenses opts: order:12
	if err := codec256.FormatSize(writer, len(x.Licenses)); err != nil {
		return err
	}
	for i := 0; i < len(x.Licenses); i++ {
		if err := x.Licenses[i].Formatter(writer, version); err != nil {
			return err
		}
	}
	// decode x.Port opts: encoder:"short"  order:13
	if err := codec256.FormatShort(writer, x.Port); err != nil {
		return err
	}
	// decode x.MemoryExcessTime opts: order:14
	if err := codec256.FormatInt(writer, x.MemoryExcessTime); err != nil {
		return err
	}
	// decode x.MemorySize opts: order:15
	if err := codec256.FormatInt(writer, x.MemorySize); err != nil {
		return err
	}
	// decode x.Pid opts: order:16
	if err := codec256.FormatString(writer, x.Pid); err != nil {
		return err
	}
	// decode x.Running opts: encoder:"int"  order:17
	if err := codec256.FormatInt(writer, x.Running); err != nil {
		return err
	}
	// decode x.SelectionSize opts: order:18
	if err := codec256.FormatInt(writer, x.SelectionSize); err != nil {
		return err
	}
	// decode x.StartedAt opts: encoder:"time"  order:18
	// TODO check nil
	if err := codec256.FormatTime(writer, x.GetStartedAt().AsTime()); err != nil {
		return err
	}
	// decode x.Use opts: encoder:"int"  order:20
	if err := codec256.FormatInt(writer, x.Use); err != nil {
		return err
	}
	// decode x.AvailablePerfomance opts: order:21
	if err := codec256.FormatInt(writer, x.AvailablePerfomance); err != nil {
		return err
	}
	if version >= 9 {
		// decode x.Reverse opts: order:22  version:9
		if err := codec256.FormatBool(writer, x.Reverse); err != nil {
			return err
		}
	}
	return nil
}
