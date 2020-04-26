// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

/*
Package clusterpb is a generated protocol buffer package.

It is generated from these files:
	cluster.proto

It has these top-level messages:
	MemberInfo
	RegisterRequest
	RegisterResponse
	UnregisterRequest
	UnregisterResponse
	RequestMessage
	NotifyMessage
	ResponseMessage
	PushMessage
	MemberHandleResponse
	NewMemberRequest
	NewMemberResponse
	DelMemberRequest
	DelMemberResponse
	SessionClosedRequest
	SessionClosedResponse
	CloseSessionRequest
	CloseSessionResponse
*/
package clusterpb

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MemberInfo struct {
	Label       string   `protobuf:"bytes,1,opt,name=label" json:"label"`
	ServiceAddr string   `protobuf:"bytes,2,opt,name=serviceAddr" json:"serviceAddr"`
	Services    []string `protobuf:"bytes,3,rep,name=services" json:"services"`
}

func (m *MemberInfo) Reset()                    { *m = MemberInfo{} }
func (m *MemberInfo) String() string            { return proto.CompactTextString(m) }
func (*MemberInfo) ProtoMessage()               {}
func (*MemberInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MemberInfo) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *MemberInfo) GetServiceAddr() string {
	if m != nil {
		return m.ServiceAddr
	}
	return ""
}

func (m *MemberInfo) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

type RegisterRequest struct {
	MemberInfo *MemberInfo `protobuf:"bytes,1,opt,name=memberInfo" json:"memberInfo"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetMemberInfo() *MemberInfo {
	if m != nil {
		return m.MemberInfo
	}
	return nil
}

type RegisterResponse struct {
	Members []*MemberInfo `protobuf:"bytes,1,rep,name=members" json:"members"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterResponse) GetMembers() []*MemberInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

type UnregisterRequest struct {
	ServiceAddr string `protobuf:"bytes,1,opt,name=serviceAddr" json:"serviceAddr"`
}

func (m *UnregisterRequest) Reset()                    { *m = UnregisterRequest{} }
func (m *UnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*UnregisterRequest) ProtoMessage()               {}
func (*UnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UnregisterRequest) GetServiceAddr() string {
	if m != nil {
		return m.ServiceAddr
	}
	return ""
}

type UnregisterResponse struct {
}

func (m *UnregisterResponse) Reset()                    { *m = UnregisterResponse{} }
func (m *UnregisterResponse) String() string            { return proto.CompactTextString(m) }
func (*UnregisterResponse) ProtoMessage()               {}
func (*UnregisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RequestMessage struct {
	GateAddr  string `protobuf:"bytes,1,opt,name=gateAddr" json:"gateAddr"`
	SessionId int64  `protobuf:"varint,2,opt,name=sessionId" json:"sessionId"`
	Id        uint64 `protobuf:"varint,3,opt,name=id" json:"id"`
	Route     string `protobuf:"bytes,4,opt,name=route" json:"route"`
	Data      []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data"`
}

func (m *RequestMessage) Reset()                    { *m = RequestMessage{} }
func (m *RequestMessage) String() string            { return proto.CompactTextString(m) }
func (*RequestMessage) ProtoMessage()               {}
func (*RequestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RequestMessage) GetGateAddr() string {
	if m != nil {
		return m.GateAddr
	}
	return ""
}

func (m *RequestMessage) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *RequestMessage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RequestMessage) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *RequestMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type NotifyMessage struct {
	GateAddr  string `protobuf:"bytes,1,opt,name=gateAddr" json:"gateAddr"`
	SessionId int64  `protobuf:"varint,2,opt,name=sessionId" json:"sessionId"`
	Route     string `protobuf:"bytes,3,opt,name=route" json:"route"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
}

func (m *NotifyMessage) Reset()                    { *m = NotifyMessage{} }
func (m *NotifyMessage) String() string            { return proto.CompactTextString(m) }
func (*NotifyMessage) ProtoMessage()               {}
func (*NotifyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NotifyMessage) GetGateAddr() string {
	if m != nil {
		return m.GateAddr
	}
	return ""
}

func (m *NotifyMessage) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *NotifyMessage) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *NotifyMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ResponseMessage struct {
	SessionId int64  `protobuf:"varint,1,opt,name=sessionId" json:"sessionId"`
	Id        uint64 `protobuf:"varint,2,opt,name=id" json:"id"`
	Route     string `protobuf:"bytes,3,opt,name=route" json:"route"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data"`
}

func (m *ResponseMessage) Reset()                    { *m = ResponseMessage{} }
func (m *ResponseMessage) String() string            { return proto.CompactTextString(m) }
func (*ResponseMessage) ProtoMessage()               {}
func (*ResponseMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResponseMessage) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *ResponseMessage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ResponseMessage) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *ResponseMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PushMessage struct {
	SessionId int64  `protobuf:"varint,1,opt,name=sessionId" json:"sessionId"`
	Route     string `protobuf:"bytes,2,opt,name=route" json:"route"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (m *PushMessage) Reset()                    { *m = PushMessage{} }
func (m *PushMessage) String() string            { return proto.CompactTextString(m) }
func (*PushMessage) ProtoMessage()               {}
func (*PushMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PushMessage) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *PushMessage) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *PushMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type MemberHandleResponse struct {
}

func (m *MemberHandleResponse) Reset()                    { *m = MemberHandleResponse{} }
func (m *MemberHandleResponse) String() string            { return proto.CompactTextString(m) }
func (*MemberHandleResponse) ProtoMessage()               {}
func (*MemberHandleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type NewMemberRequest struct {
	MemberInfo *MemberInfo `protobuf:"bytes,1,opt,name=memberInfo" json:"memberInfo"`
}

func (m *NewMemberRequest) Reset()                    { *m = NewMemberRequest{} }
func (m *NewMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*NewMemberRequest) ProtoMessage()               {}
func (*NewMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *NewMemberRequest) GetMemberInfo() *MemberInfo {
	if m != nil {
		return m.MemberInfo
	}
	return nil
}

type NewMemberResponse struct {
}

func (m *NewMemberResponse) Reset()                    { *m = NewMemberResponse{} }
func (m *NewMemberResponse) String() string            { return proto.CompactTextString(m) }
func (*NewMemberResponse) ProtoMessage()               {}
func (*NewMemberResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type DelMemberRequest struct {
	ServiceAddr string `protobuf:"bytes,1,opt,name=serviceAddr" json:"serviceAddr"`
}

func (m *DelMemberRequest) Reset()                    { *m = DelMemberRequest{} }
func (m *DelMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*DelMemberRequest) ProtoMessage()               {}
func (*DelMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DelMemberRequest) GetServiceAddr() string {
	if m != nil {
		return m.ServiceAddr
	}
	return ""
}

type DelMemberResponse struct {
}

func (m *DelMemberResponse) Reset()                    { *m = DelMemberResponse{} }
func (m *DelMemberResponse) String() string            { return proto.CompactTextString(m) }
func (*DelMemberResponse) ProtoMessage()               {}
func (*DelMemberResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type SessionClosedRequest struct {
	SessionId int64 `protobuf:"varint,1,opt,name=sessionId" json:"sessionId"`
}

func (m *SessionClosedRequest) Reset()                    { *m = SessionClosedRequest{} }
func (m *SessionClosedRequest) String() string            { return proto.CompactTextString(m) }
func (*SessionClosedRequest) ProtoMessage()               {}
func (*SessionClosedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SessionClosedRequest) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type SessionClosedResponse struct {
}

func (m *SessionClosedResponse) Reset()                    { *m = SessionClosedResponse{} }
func (m *SessionClosedResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionClosedResponse) ProtoMessage()               {}
func (*SessionClosedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type CloseSessionRequest struct {
	SessionId int64 `protobuf:"varint,1,opt,name=sessionId" json:"sessionId"`
}

func (m *CloseSessionRequest) Reset()                    { *m = CloseSessionRequest{} }
func (m *CloseSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CloseSessionRequest) ProtoMessage()               {}
func (*CloseSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *CloseSessionRequest) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type CloseSessionResponse struct {
}

func (m *CloseSessionResponse) Reset()                    { *m = CloseSessionResponse{} }
func (m *CloseSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*CloseSessionResponse) ProtoMessage()               {}
func (*CloseSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func init() {
	proto.RegisterType((*MemberInfo)(nil), "clusterpb.MemberInfo")
	proto.RegisterType((*RegisterRequest)(nil), "clusterpb.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "clusterpb.RegisterResponse")
	proto.RegisterType((*UnregisterRequest)(nil), "clusterpb.UnregisterRequest")
	proto.RegisterType((*UnregisterResponse)(nil), "clusterpb.UnregisterResponse")
	proto.RegisterType((*RequestMessage)(nil), "clusterpb.RequestMessage")
	proto.RegisterType((*NotifyMessage)(nil), "clusterpb.NotifyMessage")
	proto.RegisterType((*ResponseMessage)(nil), "clusterpb.ResponseMessage")
	proto.RegisterType((*PushMessage)(nil), "clusterpb.PushMessage")
	proto.RegisterType((*MemberHandleResponse)(nil), "clusterpb.MemberHandleResponse")
	proto.RegisterType((*NewMemberRequest)(nil), "clusterpb.NewMemberRequest")
	proto.RegisterType((*NewMemberResponse)(nil), "clusterpb.NewMemberResponse")
	proto.RegisterType((*DelMemberRequest)(nil), "clusterpb.DelMemberRequest")
	proto.RegisterType((*DelMemberResponse)(nil), "clusterpb.DelMemberResponse")
	proto.RegisterType((*SessionClosedRequest)(nil), "clusterpb.SessionClosedRequest")
	proto.RegisterType((*SessionClosedResponse)(nil), "clusterpb.SessionClosedResponse")
	proto.RegisterType((*CloseSessionRequest)(nil), "clusterpb.CloseSessionRequest")
	proto.RegisterType((*CloseSessionResponse)(nil), "clusterpb.CloseSessionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Master service

type MasterClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Master/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error) {
	out := new(UnregisterResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Master/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Master service

type MasterServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Unregister(context.Context, *UnregisterRequest) (*UnregisterResponse, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Master/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Master/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Unregister(ctx, req.(*UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clusterpb.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Master_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _Master_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}

// Client API for Member service

type MemberClient interface {
	HandleRequest(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*MemberHandleResponse, error)
	HandleNotify(ctx context.Context, in *NotifyMessage, opts ...grpc.CallOption) (*MemberHandleResponse, error)
	HandlePush(ctx context.Context, in *PushMessage, opts ...grpc.CallOption) (*MemberHandleResponse, error)
	HandleResponse(ctx context.Context, in *ResponseMessage, opts ...grpc.CallOption) (*MemberHandleResponse, error)
	NewMember(ctx context.Context, in *NewMemberRequest, opts ...grpc.CallOption) (*NewMemberResponse, error)
	DelMember(ctx context.Context, in *DelMemberRequest, opts ...grpc.CallOption) (*DelMemberResponse, error)
	SessionClosed(ctx context.Context, in *SessionClosedRequest, opts ...grpc.CallOption) (*SessionClosedResponse, error)
	CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*CloseSessionResponse, error)
}

type memberClient struct {
	cc *grpc.ClientConn
}

func NewMemberClient(cc *grpc.ClientConn) MemberClient {
	return &memberClient{cc}
}

func (c *memberClient) HandleRequest(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*MemberHandleResponse, error) {
	out := new(MemberHandleResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Member/HandleRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) HandleNotify(ctx context.Context, in *NotifyMessage, opts ...grpc.CallOption) (*MemberHandleResponse, error) {
	out := new(MemberHandleResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Member/HandleNotify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) HandlePush(ctx context.Context, in *PushMessage, opts ...grpc.CallOption) (*MemberHandleResponse, error) {
	out := new(MemberHandleResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Member/HandlePush", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) HandleResponse(ctx context.Context, in *ResponseMessage, opts ...grpc.CallOption) (*MemberHandleResponse, error) {
	out := new(MemberHandleResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Member/HandleResponse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) NewMember(ctx context.Context, in *NewMemberRequest, opts ...grpc.CallOption) (*NewMemberResponse, error) {
	out := new(NewMemberResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Member/NewMember", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) DelMember(ctx context.Context, in *DelMemberRequest, opts ...grpc.CallOption) (*DelMemberResponse, error) {
	out := new(DelMemberResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Member/DelMember", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) SessionClosed(ctx context.Context, in *SessionClosedRequest, opts ...grpc.CallOption) (*SessionClosedResponse, error) {
	out := new(SessionClosedResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Member/SessionClosed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberClient) CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*CloseSessionResponse, error) {
	out := new(CloseSessionResponse)
	err := grpc.Invoke(ctx, "/clusterpb.Member/CloseSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Member service

type MemberServer interface {
	HandleRequest(context.Context, *RequestMessage) (*MemberHandleResponse, error)
	HandleNotify(context.Context, *NotifyMessage) (*MemberHandleResponse, error)
	HandlePush(context.Context, *PushMessage) (*MemberHandleResponse, error)
	HandleResponse(context.Context, *ResponseMessage) (*MemberHandleResponse, error)
	NewMember(context.Context, *NewMemberRequest) (*NewMemberResponse, error)
	DelMember(context.Context, *DelMemberRequest) (*DelMemberResponse, error)
	SessionClosed(context.Context, *SessionClosedRequest) (*SessionClosedResponse, error)
	CloseSession(context.Context, *CloseSessionRequest) (*CloseSessionResponse, error)
}

func RegisterMemberServer(s *grpc.Server, srv MemberServer) {
	s.RegisterService(&_Member_serviceDesc, srv)
}

func _Member_HandleRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).HandleRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Member/HandleRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).HandleRequest(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_HandleNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).HandleNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Member/HandleNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).HandleNotify(ctx, req.(*NotifyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_HandlePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).HandlePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Member/HandlePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).HandlePush(ctx, req.(*PushMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_HandleResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResponseMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).HandleResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Member/HandleResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).HandleResponse(ctx, req.(*ResponseMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_NewMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).NewMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Member/NewMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).NewMember(ctx, req.(*NewMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_DelMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).DelMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Member/DelMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).DelMember(ctx, req.(*DelMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_SessionClosed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionClosedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).SessionClosed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Member/SessionClosed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).SessionClosed(ctx, req.(*SessionClosedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Member_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusterpb.Member/CloseSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServer).CloseSession(ctx, req.(*CloseSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Member_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clusterpb.Member",
	HandlerType: (*MemberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleRequest",
			Handler:    _Member_HandleRequest_Handler,
		},
		{
			MethodName: "HandleNotify",
			Handler:    _Member_HandleNotify_Handler,
		},
		{
			MethodName: "HandlePush",
			Handler:    _Member_HandlePush_Handler,
		},
		{
			MethodName: "HandleResponse",
			Handler:    _Member_HandleResponse_Handler,
		},
		{
			MethodName: "NewMember",
			Handler:    _Member_NewMember_Handler,
		},
		{
			MethodName: "DelMember",
			Handler:    _Member_DelMember_Handler,
		},
		{
			MethodName: "SessionClosed",
			Handler:    _Member_SessionClosed_Handler,
		},
		{
			MethodName: "CloseSession",
			Handler:    _Member_CloseSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x5d, 0x73, 0xd2, 0x4c,
	0x14, 0xc7, 0x09, 0xa1, 0x3c, 0xe5, 0xf0, 0xf2, 0xc0, 0x81, 0x62, 0x8c, 0xa8, 0x99, 0x5c, 0x71,
	0x85, 0x33, 0xd4, 0x7e, 0x00, 0xa7, 0x3a, 0xc2, 0x38, 0x54, 0x4d, 0xed, 0xbd, 0xa1, 0xd9, 0x62,
	0x66, 0x52, 0x82, 0xd9, 0xa0, 0xe3, 0xbd, 0x1f, 0xc3, 0x6f, 0xea, 0x8d, 0x93, 0x6c, 0xb2, 0x9c,
	0x0d, 0x41, 0x99, 0xe9, 0x1d, 0xbb, 0x67, 0xcf, 0xef, 0xbc, 0xfe, 0x09, 0xb4, 0x6f, 0x83, 0x2d,
	0x8f, 0x59, 0x34, 0xd9, 0x44, 0x61, 0x1c, 0x62, 0x23, 0x3b, 0x6e, 0x96, 0xf6, 0x67, 0x80, 0x05,
	0xbb, 0x5f, 0xb2, 0x68, 0xbe, 0xbe, 0x0b, 0x71, 0x00, 0x27, 0x81, 0xbb, 0x64, 0x81, 0xa1, 0x59,
	0xda, 0xb8, 0xe1, 0x88, 0x03, 0x5a, 0xd0, 0xe4, 0x2c, 0xfa, 0xe6, 0xdf, 0xb2, 0x57, 0x9e, 0x17,
	0x19, 0xd5, 0xd4, 0x46, 0xaf, 0xd0, 0x84, 0xd3, 0xec, 0xc8, 0x0d, 0xdd, 0xd2, 0xc7, 0x0d, 0x47,
	0x9e, 0xed, 0x19, 0xfc, 0xef, 0xb0, 0x95, 0x9f, 0xc4, 0x73, 0xd8, 0xd7, 0x2d, 0xe3, 0x31, 0x5e,
	0x00, 0xdc, 0xcb, 0xa0, 0x69, 0xac, 0xe6, 0xf4, 0x6c, 0x22, 0x93, 0x9a, 0xec, 0x32, 0x72, 0xc8,
	0x43, 0xfb, 0x12, 0xba, 0x3b, 0x12, 0xdf, 0x84, 0x6b, 0xce, 0xf0, 0x05, 0xfc, 0x27, 0x5e, 0x70,
	0x43, 0xb3, 0xf4, 0xc3, 0x9c, 0xfc, 0x95, 0x7d, 0x01, 0xbd, 0x9b, 0x75, 0x54, 0x48, 0xa8, 0x50,
	0xa1, 0xb6, 0x57, 0xa1, 0x3d, 0x00, 0xa4, 0x6e, 0x22, 0xba, 0xfd, 0x53, 0x83, 0x4e, 0xc6, 0x58,
	0x30, 0xce, 0xdd, 0x15, 0x4b, 0x5a, 0xb1, 0x72, 0x63, 0xca, 0x91, 0x67, 0x1c, 0x41, 0x83, 0x33,
	0xce, 0xfd, 0x70, 0x3d, 0xf7, 0xd2, 0x36, 0xea, 0xce, 0xee, 0x02, 0x3b, 0x50, 0xf5, 0x3d, 0x43,
	0xb7, 0xb4, 0x71, 0xcd, 0xa9, 0xfa, 0x5e, 0x32, 0x8c, 0x28, 0xdc, 0xc6, 0xcc, 0xa8, 0x89, 0x61,
	0xa4, 0x07, 0x44, 0xa8, 0x79, 0x6e, 0xec, 0x1a, 0x27, 0x96, 0x36, 0x6e, 0x39, 0xe9, 0x6f, 0x9b,
	0x43, 0xfb, 0x2a, 0x8c, 0xfd, 0xbb, 0x1f, 0x0f, 0x4f, 0x42, 0x06, 0xd5, 0xcb, 0x82, 0xd6, 0x48,
	0xd0, 0xeb, 0x64, 0xae, 0xa2, 0x0f, 0x79, 0x58, 0x05, 0xad, 0x95, 0xd7, 0x57, 0x95, 0xf5, 0xe5,
	0x50, 0x9d, 0x40, 0x6f, 0xa0, 0xf9, 0x61, 0xcb, 0xbf, 0x1c, 0x07, 0x94, 0xb9, 0x56, 0xcb, 0x72,
	0xa5, 0xd8, 0x21, 0x0c, 0xc4, 0x2e, 0xcc, 0xdc, 0xb5, 0x17, 0x30, 0x39, 0xbf, 0x39, 0x74, 0xaf,
	0xd8, 0x77, 0x61, 0x7a, 0xe0, 0x72, 0xf6, 0xa1, 0x47, 0x50, 0x19, 0xff, 0x25, 0x74, 0x5f, 0xb3,
	0x40, 0xe5, 0xff, 0x7b, 0xd7, 0xfa, 0xd0, 0x23, 0x5e, 0x12, 0x35, 0xb8, 0x16, 0x95, 0x5f, 0x06,
	0x21, 0x67, 0x5e, 0x8e, 0xfb, 0x6b, 0x8b, 0xec, 0x47, 0x70, 0x56, 0xf0, 0xca, 0x70, 0xe7, 0xd0,
	0x4f, 0x6f, 0x32, 0xeb, 0x71, 0xb4, 0x21, 0x0c, 0x54, 0x27, 0x01, 0x9b, 0xfe, 0xd2, 0xa0, 0xbe,
	0x70, 0x93, 0xfe, 0xe0, 0x1b, 0x38, 0xcd, 0x35, 0x8a, 0x26, 0xe9, 0x5a, 0xe1, 0x2f, 0xc0, 0x7c,
	0x52, 0x6a, 0xcb, 0x92, 0xab, 0xe0, 0x3b, 0x80, 0x9d, 0xdc, 0x70, 0x44, 0x1e, 0xef, 0x89, 0xd7,
	0x7c, 0x7a, 0xc0, 0x9a, 0xc3, 0xa6, 0xbf, 0x6b, 0x50, 0x17, 0xdd, 0xc4, 0x05, 0xb4, 0xf3, 0x15,
	0x10, 0x05, 0x3f, 0x56, 0xf2, 0xa0, 0x4a, 0x36, 0x9f, 0xef, 0x0d, 0xbd, 0xb0, 0x3d, 0x49, 0x9a,
	0x2d, 0x71, 0x27, 0xe4, 0x87, 0x06, 0x71, 0x51, 0x14, 0x79, 0x0c, 0xec, 0x2d, 0x80, 0xb8, 0x4b,
	0x14, 0x80, 0x43, 0xe2, 0x40, 0x24, 0x71, 0x0c, 0xe8, 0x3d, 0x74, 0xd4, 0xbb, 0xc2, 0x24, 0x14,
	0xd1, 0x1e, 0x03, 0x9c, 0x41, 0x43, 0xee, 0x36, 0xd2, 0xc9, 0x15, 0xc5, 0x63, 0x8e, 0xca, 0x8d,
	0x94, 0x24, 0x57, 0x5b, 0x21, 0x15, 0x65, 0xa2, 0x90, 0xf6, 0xd5, 0x50, 0xc1, 0x4f, 0xd0, 0x56,
	0x36, 0x1b, 0x69, 0x1d, 0x65, 0x4a, 0x31, 0xad, 0xc3, 0x0f, 0x24, 0xf5, 0x23, 0xb4, 0xe8, 0x86,
	0xe3, 0x33, 0xe2, 0x53, 0xa2, 0x17, 0xa5, 0x79, 0x65, 0xd2, 0xb0, 0x2b, 0xcb, 0x7a, 0xfa, 0xcd,
	0x3d, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x80, 0x79, 0x99, 0xc0, 0x84, 0x07, 0x00, 0x00,
}
