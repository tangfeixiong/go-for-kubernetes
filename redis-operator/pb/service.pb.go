// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/service.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/service.proto
	pb/fake.proto

It has these top-level messages:
	FakeReqResp
	FakeMessage
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FakeReqResp struct {
	FakeMessage  *FakeMessage `protobuf:"bytes,1,opt,name=fake_message,json=fakeMessage" json:"fake_message,omitempty"`
	StateCode    int32        `protobuf:"varint,2,opt,name=state_code,json=stateCode,proto3" json:"state_code,omitempty"`
	StateMessage string       `protobuf:"bytes,3,opt,name=state_message,json=stateMessage,proto3" json:"state_message,omitempty"`
}

func (m *FakeReqResp) Reset()                    { *m = FakeReqResp{} }
func (m *FakeReqResp) String() string            { return proto.CompactTextString(m) }
func (*FakeReqResp) ProtoMessage()               {}
func (*FakeReqResp) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{0} }

func (m *FakeReqResp) GetFakeMessage() *FakeMessage {
	if m != nil {
		return m.FakeMessage
	}
	return nil
}

func (m *FakeReqResp) GetStateCode() int32 {
	if m != nil {
		return m.StateCode
	}
	return 0
}

func (m *FakeReqResp) GetStateMessage() string {
	if m != nil {
		return m.StateMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*FakeReqResp)(nil), "pb.FakeReqResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RedisOpService service

type RedisOpServiceClient interface {
	Hello(ctx context.Context, in *FakeReqResp, opts ...grpc.CallOption) (*FakeReqResp, error)
}

type redisOpServiceClient struct {
	cc *grpc.ClientConn
}

func NewRedisOpServiceClient(cc *grpc.ClientConn) RedisOpServiceClient {
	return &redisOpServiceClient{cc}
}

func (c *redisOpServiceClient) Hello(ctx context.Context, in *FakeReqResp, opts ...grpc.CallOption) (*FakeReqResp, error) {
	out := new(FakeReqResp)
	err := grpc.Invoke(ctx, "/pb.RedisOpService/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RedisOpService service

type RedisOpServiceServer interface {
	Hello(context.Context, *FakeReqResp) (*FakeReqResp, error)
}

func RegisterRedisOpServiceServer(s *grpc.Server, srv RedisOpServiceServer) {
	s.RegisterService(&_RedisOpService_serviceDesc, srv)
}

func _RedisOpService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FakeReqResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedisOpServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RedisOpService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedisOpServiceServer).Hello(ctx, req.(*FakeReqResp))
	}
	return interceptor(ctx, in, info, handler)
}

var _RedisOpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RedisOpService",
	HandlerType: (*RedisOpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _RedisOpService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}

func init() { proto.RegisterFile("pb/service.proto", fileDescriptorService) }

var fileDescriptorService = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0x2f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x92, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb,
	0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x90, 0xe2, 0x2d, 0x48, 0xd2, 0x4f,
	0x4b, 0xcc, 0x86, 0x6a, 0x50, 0x6a, 0x65, 0xe4, 0xe2, 0x76, 0x4b, 0xcc, 0x4e, 0x0d, 0x4a, 0x2d,
	0x0c, 0x4a, 0x2d, 0x2e, 0x10, 0x32, 0xe2, 0xe2, 0x01, 0xc9, 0xc6, 0xe7, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xf1, 0xeb, 0x15, 0x24, 0xe9, 0x81, 0x94,
	0xf9, 0x42, 0x84, 0x83, 0xb8, 0xd3, 0x10, 0x1c, 0x21, 0x59, 0x2e, 0xae, 0xe2, 0x92, 0xc4, 0x92,
	0xd4, 0xf8, 0xe4, 0xfc, 0x94, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x4e, 0xb0, 0x88,
	0x73, 0x7e, 0x4a, 0xaa, 0x90, 0x32, 0x17, 0x2f, 0x44, 0x1a, 0x66, 0x26, 0xb3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x0f, 0x58, 0x10, 0x6a, 0x86, 0x51, 0x34, 0x17, 0x5f, 0x50, 0x6a, 0x4a, 0x66, 0xb1,
	0x7f, 0x41, 0x30, 0xc4, 0x43, 0x42, 0x9e, 0x5c, 0xac, 0x1e, 0xa9, 0x39, 0x39, 0xf9, 0x42, 0x70,
	0xcb, 0xa1, 0x6e, 0x94, 0x42, 0x17, 0x50, 0x92, 0x6b, 0xba, 0xfc, 0x64, 0x32, 0x93, 0x84, 0x92,
	0xb0, 0x7e, 0x11, 0xc8, 0x8c, 0xfc, 0x02, 0x90, 0xd7, 0xcb, 0x0c, 0xc1, 0xfe, 0xb4, 0x62, 0xd4,
	0x4a, 0x62, 0x03, 0xfb, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x63, 0x03, 0x2d, 0x30,
	0x01, 0x00, 0x00,
}
