// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package serializev1

import (
	codec256 "github.com/v8platform/encoder/ras/codec256"
	io "io"
)

func (x *ManagerInfo) Parse(reader io.Reader, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Uuid opts: encoder:"uuid"  order:1
	if err := codec256.ParseUUID(reader, &x.Uuid); err != nil {
		return err
	}
	// decode x.Descr opts: order:2
	if err := codec256.ParseString(reader, &x.Descr); err != nil {
		return err
	}
	// decode x.Host opts: order:3
	if err := codec256.ParseString(reader, &x.Host); err != nil {
		return err
	}
	// decode x.MainManager opts: order:4
	if err := codec256.ParseInt(reader, &x.MainManager); err != nil {
		return err
	}
	// decode x.Port opts: encoder:"short"  order:5
	if err := codec256.ParseShort(reader, &x.Port); err != nil {
		return err
	}
	// decode x.Pid opts: order:6
	if err := codec256.ParseString(reader, &x.Pid); err != nil {
		return err
	}
	return nil
}
func (x *ManagerInfo) Formatter(writer io.Writer, version int32) error {
	if x == nil {
		return nil
	}
	// decode x.Uuid opts: encoder:"uuid"  order:1
	if err := codec256.FormatUuid(writer, x.Uuid); err != nil {
		return err
	}
	// decode x.Descr opts: order:2
	if err := codec256.FormatString(writer, x.Descr); err != nil {
		return err
	}
	// decode x.Host opts: order:3
	if err := codec256.FormatString(writer, x.Host); err != nil {
		return err
	}
	// decode x.MainManager opts: order:4
	if err := codec256.FormatInt(writer, x.MainManager); err != nil {
		return err
	}
	// decode x.Port opts: encoder:"short"  order:5
	if err := codec256.FormatShort(writer, x.Port); err != nil {
		return err
	}
	// decode x.Pid opts: order:6
	if err := codec256.FormatString(writer, x.Pid); err != nil {
		return err
	}
	return nil
}
