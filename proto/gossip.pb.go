// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/gossip.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GossipRequest struct {
	Payload              int32    `protobuf:"varint,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Hops                 int32    `protobuf:"varint,2,opt,name=hops,proto3" json:"hops,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipRequest) Reset()         { *m = GossipRequest{} }
func (m *GossipRequest) String() string { return proto.CompactTextString(m) }
func (*GossipRequest) ProtoMessage()    {}
func (*GossipRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_b4df064583c7efe3, []int{0}
}
func (m *GossipRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipRequest.Unmarshal(m, b)
}
func (m *GossipRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipRequest.Marshal(b, m, deterministic)
}
func (dst *GossipRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipRequest.Merge(dst, src)
}
func (m *GossipRequest) XXX_Size() int {
	return xxx_messageInfo_GossipRequest.Size(m)
}
func (m *GossipRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GossipRequest proto.InternalMessageInfo

func (m *GossipRequest) GetPayload() int32 {
	if m != nil {
		return m.Payload
	}
	return 0
}

func (m *GossipRequest) GetHops() int32 {
	if m != nil {
		return m.Hops
	}
	return 0
}

type GossipResponse struct {
	Hot                  bool     `protobuf:"varint,1,opt,name=hot,proto3" json:"hot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipResponse) Reset()         { *m = GossipResponse{} }
func (m *GossipResponse) String() string { return proto.CompactTextString(m) }
func (*GossipResponse) ProtoMessage()    {}
func (*GossipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_b4df064583c7efe3, []int{1}
}
func (m *GossipResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipResponse.Unmarshal(m, b)
}
func (m *GossipResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipResponse.Marshal(b, m, deterministic)
}
func (dst *GossipResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipResponse.Merge(dst, src)
}
func (m *GossipResponse) XXX_Size() int {
	return xxx_messageInfo_GossipResponse.Size(m)
}
func (m *GossipResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GossipResponse proto.InternalMessageInfo

func (m *GossipResponse) GetHot() bool {
	if m != nil {
		return m.Hot
	}
	return false
}

func init() {
	proto.RegisterType((*GossipRequest)(nil), "hashicorp.hyparview.example.gossip.GossipRequest")
	proto.RegisterType((*GossipResponse)(nil), "hashicorp.hyparview.example.gossip.GossipResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GossipClient is the client API for Gossip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GossipClient interface {
	Gossip(ctx context.Context, in *GossipRequest, opts ...grpc.CallOption) (*GossipResponse, error)
}

type gossipClient struct {
	cc *grpc.ClientConn
}

func NewGossipClient(cc *grpc.ClientConn) GossipClient {
	return &gossipClient{cc}
}

func (c *gossipClient) Gossip(ctx context.Context, in *GossipRequest, opts ...grpc.CallOption) (*GossipResponse, error) {
	out := new(GossipResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.hyparview.example.gossip.Gossip/Gossip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GossipServer is the server API for Gossip service.
type GossipServer interface {
	Gossip(context.Context, *GossipRequest) (*GossipResponse, error)
}

func RegisterGossipServer(s *grpc.Server, srv GossipServer) {
	s.RegisterService(&_Gossip_serviceDesc, srv)
}

func _Gossip_Gossip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GossipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).Gossip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.hyparview.example.gossip.Gossip/Gossip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).Gossip(ctx, req.(*GossipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gossip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.hyparview.example.gossip.Gossip",
	HandlerType: (*GossipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Gossip",
			Handler:    _Gossip_Gossip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gossip.proto",
}

func init() { proto.RegisterFile("proto/gossip.proto", fileDescriptor_gossip_b4df064583c7efe3) }

var fileDescriptor_gossip_b4df064583c7efe3 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcf, 0x2f, 0x2e, 0xce, 0x2c, 0xd0, 0x03, 0x73, 0x84, 0x94, 0x32, 0x12, 0x8b,
	0x33, 0x32, 0x93, 0xf3, 0x8b, 0x0a, 0xf4, 0x32, 0x2a, 0x0b, 0x12, 0x8b, 0xca, 0x32, 0x53, 0xcb,
	0xf5, 0x52, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0xf5, 0x20, 0x2a, 0x95, 0x6c, 0xb9, 0x78, 0xdd,
	0xc1, 0xac, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca,
	0x9c, 0xfc, 0xc4, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x18, 0x57, 0x48, 0x88, 0x8b,
	0x25, 0x23, 0xbf, 0xa0, 0x58, 0x82, 0x09, 0x2c, 0x0c, 0x66, 0x2b, 0x29, 0x71, 0xf1, 0xc1, 0xb4,
	0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x09, 0x70, 0x31, 0x67, 0xe4, 0x97, 0x80, 0xf5, 0x72,
	0x04, 0x81, 0x98, 0x46, 0xd5, 0x5c, 0x6c, 0x10, 0x35, 0x42, 0x85, 0x70, 0x96, 0xa1, 0x1e, 0x61,
	0xb7, 0xe9, 0xa1, 0x38, 0x4c, 0xca, 0x88, 0x14, 0x2d, 0x10, 0xc7, 0x28, 0x31, 0x38, 0xb1, 0x47,
	0xb1, 0x82, 0x03, 0x23, 0x89, 0x0d, 0x4c, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x34, 0xe2,
	0x91, 0x9e, 0x29, 0x01, 0x00, 0x00,
}
