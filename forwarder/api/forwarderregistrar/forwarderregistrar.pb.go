// Code generated by protoc-gen-go. DO NOT EDIT.
// source: forwarderregistrar.proto

package forwarderregistrar

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ForwarderRegistrationRequest is sent by the forwarder to NSM
// to advertise itself and inform NSM about the location of the forwarder socket
// and its initially supported parameters.
type ForwarderRegistrationRequest struct {
	ForwarderName        string   `protobuf:"bytes,1,opt,name=forwarder_name,json=forwarderName,proto3" json:"forwarder_name,omitempty"`
	ForwarderSocket      string   `protobuf:"bytes,2,opt,name=forwarder_socket,json=forwarderSocket,proto3" json:"forwarder_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwarderRegistrationRequest) Reset()         { *m = ForwarderRegistrationRequest{} }
func (m *ForwarderRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*ForwarderRegistrationRequest) ProtoMessage()    {}
func (*ForwarderRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf2c0f4975ef21fe, []int{0}
}

func (m *ForwarderRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwarderRegistrationRequest.Unmarshal(m, b)
}
func (m *ForwarderRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwarderRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *ForwarderRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwarderRegistrationRequest.Merge(m, src)
}
func (m *ForwarderRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_ForwarderRegistrationRequest.Size(m)
}
func (m *ForwarderRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwarderRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForwarderRegistrationRequest proto.InternalMessageInfo

func (m *ForwarderRegistrationRequest) GetForwarderName() string {
	if m != nil {
		return m.ForwarderName
	}
	return ""
}

func (m *ForwarderRegistrationRequest) GetForwarderSocket() string {
	if m != nil {
		return m.ForwarderSocket
	}
	return ""
}

type ForwarderRegistrationReply struct {
	Registered           bool     `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwarderRegistrationReply) Reset()         { *m = ForwarderRegistrationReply{} }
func (m *ForwarderRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*ForwarderRegistrationReply) ProtoMessage()    {}
func (*ForwarderRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf2c0f4975ef21fe, []int{1}
}

func (m *ForwarderRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwarderRegistrationReply.Unmarshal(m, b)
}
func (m *ForwarderRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwarderRegistrationReply.Marshal(b, m, deterministic)
}
func (m *ForwarderRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwarderRegistrationReply.Merge(m, src)
}
func (m *ForwarderRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_ForwarderRegistrationReply.Size(m)
}
func (m *ForwarderRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwarderRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_ForwarderRegistrationReply proto.InternalMessageInfo

func (m *ForwarderRegistrationReply) GetRegistered() bool {
	if m != nil {
		return m.Registered
	}
	return false
}

// ForwarderUnRegistrationRequest is sent by the forwarder to NSM
// to remove itself from the list of available forwarders.
type ForwarderUnRegistrationRequest struct {
	ForwarderName        string   `protobuf:"bytes,1,opt,name=forwarder_name,json=forwarderName,proto3" json:"forwarder_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwarderUnRegistrationRequest) Reset()         { *m = ForwarderUnRegistrationRequest{} }
func (m *ForwarderUnRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*ForwarderUnRegistrationRequest) ProtoMessage()    {}
func (*ForwarderUnRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf2c0f4975ef21fe, []int{2}
}

func (m *ForwarderUnRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwarderUnRegistrationRequest.Unmarshal(m, b)
}
func (m *ForwarderUnRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwarderUnRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *ForwarderUnRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwarderUnRegistrationRequest.Merge(m, src)
}
func (m *ForwarderUnRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_ForwarderUnRegistrationRequest.Size(m)
}
func (m *ForwarderUnRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwarderUnRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForwarderUnRegistrationRequest proto.InternalMessageInfo

func (m *ForwarderUnRegistrationRequest) GetForwarderName() string {
	if m != nil {
		return m.ForwarderName
	}
	return ""
}

type ForwarderUnRegistrationReply struct {
	UnRegistered         bool     `protobuf:"varint,1,opt,name=un_registered,json=unRegistered,proto3" json:"un_registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwarderUnRegistrationReply) Reset()         { *m = ForwarderUnRegistrationReply{} }
func (m *ForwarderUnRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*ForwarderUnRegistrationReply) ProtoMessage()    {}
func (*ForwarderUnRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf2c0f4975ef21fe, []int{3}
}

func (m *ForwarderUnRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwarderUnRegistrationReply.Unmarshal(m, b)
}
func (m *ForwarderUnRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwarderUnRegistrationReply.Marshal(b, m, deterministic)
}
func (m *ForwarderUnRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwarderUnRegistrationReply.Merge(m, src)
}
func (m *ForwarderUnRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_ForwarderUnRegistrationReply.Size(m)
}
func (m *ForwarderUnRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwarderUnRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_ForwarderUnRegistrationReply proto.InternalMessageInfo

func (m *ForwarderUnRegistrationReply) GetUnRegistered() bool {
	if m != nil {
		return m.UnRegistered
	}
	return false
}

func init() {
	proto.RegisterType((*ForwarderRegistrationRequest)(nil), "forwarderregistrar.ForwarderRegistrationRequest")
	proto.RegisterType((*ForwarderRegistrationReply)(nil), "forwarderregistrar.ForwarderRegistrationReply")
	proto.RegisterType((*ForwarderUnRegistrationRequest)(nil), "forwarderregistrar.ForwarderUnRegistrationRequest")
	proto.RegisterType((*ForwarderUnRegistrationReply)(nil), "forwarderregistrar.ForwarderUnRegistrationReply")
}

func init() { proto.RegisterFile("forwarderregistrar.proto", fileDescriptor_bf2c0f4975ef21fe) }

var fileDescriptor_bf2c0f4975ef21fe = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x18, 0xc4, 0x59, 0x0f, 0xa2, 0x1f, 0xd6, 0xca, 0x82, 0x1a, 0x42, 0x09, 0x12, 0x11, 0xea, 0x65,
	0x0d, 0xf1, 0xea, 0x4d, 0xaa, 0x17, 0xf1, 0x10, 0xf1, 0x5c, 0x52, 0xfb, 0x35, 0x04, 0x93, 0xdd,
	0x75, 0x77, 0xa3, 0xe4, 0xe6, 0xc9, 0x07, 0xf1, 0xcd, 0x7c, 0x13, 0xc9, 0x1f, 0x63, 0x49, 0x93,
	0x42, 0x7b, 0xc9, 0x61, 0xf2, 0x9b, 0xdd, 0x99, 0x4c, 0xc0, 0x5a, 0x08, 0xf5, 0x11, 0xaa, 0x39,
	0x2a, 0x85, 0x51, 0xac, 0x8d, 0x0a, 0x15, 0x93, 0x4a, 0x18, 0x41, 0xe9, 0xea, 0x1b, 0xdb, 0x92,
	0x26, 0x97, 0xa8, 0xaf, 0x30, 0x95, 0x26, 0xaf, 0x9e, 0x15, 0xed, 0x4a, 0x18, 0xdd, 0xfd, 0xf1,
	0x41, 0xcd, 0x9b, 0x58, 0xf0, 0x00, 0xdf, 0x32, 0xd4, 0x86, 0x5e, 0xc0, 0x61, 0x73, 0xde, 0x94,
	0x87, 0x29, 0x5a, 0xe4, 0x8c, 0x8c, 0xf7, 0x83, 0x41, 0xa3, 0x3e, 0x86, 0x29, 0xd2, 0x4b, 0x38,
	0xfa, 0xc7, 0xb4, 0x78, 0x79, 0x45, 0x63, 0xed, 0x94, 0xe0, 0xb0, 0xd1, 0x9f, 0x4a, 0xd9, 0xbd,
	0x01, 0xbb, 0xe7, 0x46, 0x99, 0xe4, 0xd4, 0x01, 0xa8, 0x62, 0xa3, 0xc2, 0x79, 0x79, 0xd7, 0x5e,
	0xb0, 0xa4, 0xb8, 0xf7, 0xe0, 0x34, 0xee, 0x67, 0xbe, 0x7d, 0x62, 0xf7, 0x76, 0xa9, 0x78, 0xfb,
	0xa0, 0x22, 0xc8, 0x39, 0x0c, 0x32, 0x3e, 0x5d, 0xc9, 0x72, 0x90, 0xd5, 0x6c, 0xa1, 0xf9, 0x3f,
	0x04, 0x8e, 0x3b, 0xcb, 0xd0, 0x4f, 0x02, 0xa3, 0x3a, 0x51, 0x37, 0xe0, 0xb1, 0x8e, 0x05, 0xd7,
	0x4d, 0x61, 0xb3, 0x0d, 0x1c, 0x45, 0x83, 0x09, 0x0c, 0x6b, 0xeb, 0x43, 0xfc, 0x8e, 0x1c, 0xb5,
	0xa6, 0x27, 0x2c, 0x12, 0x22, 0x4a, 0xb0, 0x1a, 0x7f, 0x96, 0x2d, 0xd8, 0xa4, 0xf8, 0x17, 0xec,
	0x1e, 0x7d, 0x4c, 0x3c, 0xe2, 0x7f, 0x13, 0x38, 0xed, 0xf9, 0x52, 0xf4, 0x8b, 0x80, 0xd3, 0x6e,
	0xd9, 0x42, 0xfc, 0xb5, 0xa9, 0x3b, 0x27, 0xb4, 0xbd, 0x8d, 0x3c, 0x32, 0xc9, 0x67, 0xbb, 0x65,
	0xf0, 0xeb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xaf, 0x67, 0xb2, 0x17, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ForwarderRegistrationClient is the client API for ForwarderRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForwarderRegistrationClient interface {
	RequestForwarderRegistration(ctx context.Context, in *ForwarderRegistrationRequest, opts ...grpc.CallOption) (*ForwarderRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the forwarder that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the forwarder needs to start re-registration logic.
	RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (ForwarderRegistration_RequestLivenessClient, error)
}

type forwarderRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewForwarderRegistrationClient(cc *grpc.ClientConn) ForwarderRegistrationClient {
	return &forwarderRegistrationClient{cc}
}

func (c *forwarderRegistrationClient) RequestForwarderRegistration(ctx context.Context, in *ForwarderRegistrationRequest, opts ...grpc.CallOption) (*ForwarderRegistrationReply, error) {
	out := new(ForwarderRegistrationReply)
	err := c.cc.Invoke(ctx, "/forwarderregistrar.ForwarderRegistration/RequestForwarderRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwarderRegistrationClient) RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (ForwarderRegistration_RequestLivenessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ForwarderRegistration_serviceDesc.Streams[0], "/forwarderregistrar.ForwarderRegistration/RequestLiveness", opts...)
	if err != nil {
		return nil, err
	}
	x := &forwarderRegistrationRequestLivenessClient{stream}
	return x, nil
}

type ForwarderRegistration_RequestLivenessClient interface {
	Send(*empty.Empty) error
	Recv() (*empty.Empty, error)
	grpc.ClientStream
}

type forwarderRegistrationRequestLivenessClient struct {
	grpc.ClientStream
}

func (x *forwarderRegistrationRequestLivenessClient) Send(m *empty.Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *forwarderRegistrationRequestLivenessClient) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ForwarderRegistrationServer is the server API for ForwarderRegistration service.
type ForwarderRegistrationServer interface {
	RequestForwarderRegistration(context.Context, *ForwarderRegistrationRequest) (*ForwarderRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the forwarder that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the forwarder needs to start re-registration logic.
	RequestLiveness(ForwarderRegistration_RequestLivenessServer) error
}

// UnimplementedForwarderRegistrationServer can be embedded to have forward compatible implementations.
type UnimplementedForwarderRegistrationServer struct {
}

func (*UnimplementedForwarderRegistrationServer) RequestForwarderRegistration(ctx context.Context, req *ForwarderRegistrationRequest) (*ForwarderRegistrationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestForwarderRegistration not implemented")
}
func (*UnimplementedForwarderRegistrationServer) RequestLiveness(srv ForwarderRegistration_RequestLivenessServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestLiveness not implemented")
}

func RegisterForwarderRegistrationServer(s *grpc.Server, srv ForwarderRegistrationServer) {
	s.RegisterService(&_ForwarderRegistration_serviceDesc, srv)
}

func _ForwarderRegistration_RequestForwarderRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwarderRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwarderRegistrationServer).RequestForwarderRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarderregistrar.ForwarderRegistration/RequestForwarderRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwarderRegistrationServer).RequestForwarderRegistration(ctx, req.(*ForwarderRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ForwarderRegistration_RequestLiveness_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ForwarderRegistrationServer).RequestLiveness(&forwarderRegistrationRequestLivenessServer{stream})
}

type ForwarderRegistration_RequestLivenessServer interface {
	Send(*empty.Empty) error
	Recv() (*empty.Empty, error)
	grpc.ServerStream
}

type forwarderRegistrationRequestLivenessServer struct {
	grpc.ServerStream
}

func (x *forwarderRegistrationRequestLivenessServer) Send(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *forwarderRegistrationRequestLivenessServer) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ForwarderRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "forwarderregistrar.ForwarderRegistration",
	HandlerType: (*ForwarderRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestForwarderRegistration",
			Handler:    _ForwarderRegistration_RequestForwarderRegistration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestLiveness",
			Handler:       _ForwarderRegistration_RequestLiveness_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "forwarderregistrar.proto",
}

// ForwarderUnRegistrationClient is the client API for ForwarderUnRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForwarderUnRegistrationClient interface {
	RequestForwarderUnRegistration(ctx context.Context, in *ForwarderUnRegistrationRequest, opts ...grpc.CallOption) (*ForwarderUnRegistrationReply, error)
}

type forwarderUnRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewForwarderUnRegistrationClient(cc *grpc.ClientConn) ForwarderUnRegistrationClient {
	return &forwarderUnRegistrationClient{cc}
}

func (c *forwarderUnRegistrationClient) RequestForwarderUnRegistration(ctx context.Context, in *ForwarderUnRegistrationRequest, opts ...grpc.CallOption) (*ForwarderUnRegistrationReply, error) {
	out := new(ForwarderUnRegistrationReply)
	err := c.cc.Invoke(ctx, "/forwarderregistrar.ForwarderUnRegistration/RequestForwarderUnRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForwarderUnRegistrationServer is the server API for ForwarderUnRegistration service.
type ForwarderUnRegistrationServer interface {
	RequestForwarderUnRegistration(context.Context, *ForwarderUnRegistrationRequest) (*ForwarderUnRegistrationReply, error)
}

// UnimplementedForwarderUnRegistrationServer can be embedded to have forward compatible implementations.
type UnimplementedForwarderUnRegistrationServer struct {
}

func (*UnimplementedForwarderUnRegistrationServer) RequestForwarderUnRegistration(ctx context.Context, req *ForwarderUnRegistrationRequest) (*ForwarderUnRegistrationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestForwarderUnRegistration not implemented")
}

func RegisterForwarderUnRegistrationServer(s *grpc.Server, srv ForwarderUnRegistrationServer) {
	s.RegisterService(&_ForwarderUnRegistration_serviceDesc, srv)
}

func _ForwarderUnRegistration_RequestForwarderUnRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwarderUnRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwarderUnRegistrationServer).RequestForwarderUnRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarderregistrar.ForwarderUnRegistration/RequestForwarderUnRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwarderUnRegistrationServer).RequestForwarderUnRegistration(ctx, req.(*ForwarderUnRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ForwarderUnRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "forwarderregistrar.ForwarderUnRegistration",
	HandlerType: (*ForwarderUnRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestForwarderUnRegistration",
			Handler:    _ForwarderUnRegistration_RequestForwarderUnRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forwarderregistrar.proto",
}