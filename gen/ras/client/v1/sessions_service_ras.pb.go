// Code generated by protoc-gen-go-ras. DO NOT EDIT.

// This is a compile-time assertion to ensure that this generated file
// is compatible with the v8platform/protoc-gen-go-ras ras it is being compiled against.

package clientv1

import (
	context "context"
	v1 "github.com/v8platform/protos/gen/ras/messages/v1"
	proto "google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type SessionsServiceImpl interface {
	GetSessions(ctx context.Context, req *v1.GetSessionsRequest) (*v1.GetSessionsResponse, error)
	GetInfobaseSessions(ctx context.Context, req *v1.GetInfobaseSessionsRequest) (*v1.GetInfobaseSessionsResponse, error)
	GetSession(ctx context.Context, req *v1.GetSessionInfoRequest) (*v1.GetSessionInfoResponse, error)
	TerminateSession(ctx context.Context, req *v1.TerminateSessionRequest) (*emptypb.Empty, error)
}

func NewSessionsService(endpointService EndpointServiceImpl) SessionsServiceImpl {
	return &SessionsService{
		endpointService,
	}
}

// SessionsService is the endpoint message service for RAS service.
type SessionsService struct {
	e EndpointServiceImpl
}

func (x *SessionsService) GetSessions(ctx context.Context, req *v1.GetSessionsRequest) (*v1.GetSessionsResponse, error) {

	var resp v1.GetSessionsResponse

	anyRequest, err := anypb.New(req)
	if err != nil {
		return nil, err
	}

	anyRespond, err := anypb.New(&resp)
	if err != nil {
		return nil, err
	}

	endpointRequest := &EndpointRequest{
		Request: anyRequest,
		Respond: anyRespond,
	}

	response, err := x.e.Request(ctx, endpointRequest)
	if err != nil {
		return nil, err
	}

	if err := anypb.UnmarshalTo(response, &resp, proto.UnmarshalOptions{}); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (x *SessionsService) GetInfobaseSessions(ctx context.Context, req *v1.GetInfobaseSessionsRequest) (*v1.GetInfobaseSessionsResponse, error) {

	var resp v1.GetInfobaseSessionsResponse

	anyRequest, err := anypb.New(req)
	if err != nil {
		return nil, err
	}

	anyRespond, err := anypb.New(&resp)
	if err != nil {
		return nil, err
	}

	endpointRequest := &EndpointRequest{
		Request: anyRequest,
		Respond: anyRespond,
	}

	response, err := x.e.Request(ctx, endpointRequest)
	if err != nil {
		return nil, err
	}

	if err := anypb.UnmarshalTo(response, &resp, proto.UnmarshalOptions{}); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (x *SessionsService) GetSession(ctx context.Context, req *v1.GetSessionInfoRequest) (*v1.GetSessionInfoResponse, error) {

	var resp v1.GetSessionInfoResponse

	anyRequest, err := anypb.New(req)
	if err != nil {
		return nil, err
	}

	anyRespond, err := anypb.New(&resp)
	if err != nil {
		return nil, err
	}

	endpointRequest := &EndpointRequest{
		Request: anyRequest,
		Respond: anyRespond,
	}

	response, err := x.e.Request(ctx, endpointRequest)
	if err != nil {
		return nil, err
	}

	if err := anypb.UnmarshalTo(response, &resp, proto.UnmarshalOptions{}); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (x *SessionsService) TerminateSession(ctx context.Context, req *v1.TerminateSessionRequest) (*emptypb.Empty, error) {

	var resp emptypb.Empty

	anyRequest, err := anypb.New(req)
	if err != nil {
		return nil, err
	}

	anyRespond, err := anypb.New(&resp)
	if err != nil {
		return nil, err
	}

	endpointRequest := &EndpointRequest{
		Request: anyRequest,
		Respond: anyRespond,
	}

	response, err := x.e.Request(ctx, endpointRequest)
	if err != nil {
		return nil, err
	}

	if err := anypb.UnmarshalTo(response, &resp, proto.UnmarshalOptions{}); err != nil {
		return nil, err
	}
	return &resp, nil
}