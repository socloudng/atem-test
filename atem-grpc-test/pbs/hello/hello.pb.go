// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RequestMsg struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMsg) Reset()         { *m = RequestMsg{} }
func (m *RequestMsg) String() string { return proto.CompactTextString(m) }
func (*RequestMsg) ProtoMessage()    {}
func (*RequestMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *RequestMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMsg.Unmarshal(m, b)
}
func (m *RequestMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMsg.Marshal(b, m, deterministic)
}
func (m *RequestMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMsg.Merge(m, src)
}
func (m *RequestMsg) XXX_Size() int {
	return xxx_messageInfo_RequestMsg.Size(m)
}
func (m *RequestMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMsg proto.InternalMessageInfo

func (m *RequestMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ResponseMsg struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMsg) Reset()         { *m = ResponseMsg{} }
func (m *ResponseMsg) String() string { return proto.CompactTextString(m) }
func (*ResponseMsg) ProtoMessage()    {}
func (*ResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *ResponseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMsg.Unmarshal(m, b)
}
func (m *ResponseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMsg.Marshal(b, m, deterministic)
}
func (m *ResponseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMsg.Merge(m, src)
}
func (m *ResponseMsg) XXX_Size() int {
	return xxx_messageInfo_ResponseMsg.Size(m)
}
func (m *ResponseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMsg proto.InternalMessageInfo

func (m *ResponseMsg) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestMsg)(nil), "hello.RequestMsg")
	proto.RegisterType((*ResponseMsg)(nil), "hello.ResponseMsg")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x14, 0xb8, 0xb8, 0x82,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x7c, 0x8b, 0xd3, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x75, 0x2e, 0xee, 0xa0, 0xd4,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0x12, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4,
	0x74, 0x98, 0x2a, 0x18, 0xd7, 0xc8, 0x8d, 0x8b, 0xdf, 0xb7, 0x38, 0xdd, 0x03, 0x64, 0x6c, 0x70,
	0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x31, 0x17, 0x47, 0x70, 0x62, 0x25, 0x58, 0x48, 0x48,
	0x50, 0x0f, 0x62, 0x3d, 0xc2, 0x3a, 0x29, 0x21, 0xb8, 0x10, 0xdc, 0x7c, 0x25, 0x06, 0x27, 0xfe,
	0x28, 0x5e, 0x3d, 0x7d, 0xb0, 0x84, 0x35, 0x98, 0x4c, 0x62, 0x03, 0xbb, 0xd8, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xcd, 0x39, 0xfc, 0xe4, 0xc0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgHelloServiceClient is the client API for MsgHelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgHelloServiceClient interface {
	SayHello(ctx context.Context, in *RequestMsg, opts ...grpc.CallOption) (*ResponseMsg, error)
}

type msgHelloServiceClient struct {
	cc *grpc.ClientConn
}

func NewMsgHelloServiceClient(cc *grpc.ClientConn) MsgHelloServiceClient {
	return &msgHelloServiceClient{cc}
}

func (c *msgHelloServiceClient) SayHello(ctx context.Context, in *RequestMsg, opts ...grpc.CallOption) (*ResponseMsg, error) {
	out := new(ResponseMsg)
	err := c.cc.Invoke(ctx, "/hello.MsgHelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgHelloServiceServer is the server API for MsgHelloService service.
type MsgHelloServiceServer interface {
	SayHello(context.Context, *RequestMsg) (*ResponseMsg, error)
}

// UnimplementedMsgHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgHelloServiceServer struct {
}

func (*UnimplementedMsgHelloServiceServer) SayHello(ctx context.Context, req *RequestMsg) (*ResponseMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterMsgHelloServiceServer(s *grpc.Server, srv interface{}) {
	s.RegisterService(&_MsgHelloService_serviceDesc, srv)
}

func _MsgHelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgHelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.MsgHelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgHelloServiceServer).SayHello(ctx, req.(*RequestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgHelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.MsgHelloService",
	HandlerType: (*MsgHelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _MsgHelloService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}