// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

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

// The request message containing the process name.
type CertificateRequest struct {
	Csr                  []byte   `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_62f8d9161514c8eb, []int{0}
}
func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}
func (dst *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(dst, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

// The response message containing the requested logs.
type CertificateResponse struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_62f8d9161514c8eb, []int{1}
}
func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}
func (dst *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(dst, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "proto.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "proto.CertificateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ROTDClient is the client API for ROTD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ROTDClient interface {
	Certificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type rOTDClient struct {
	cc *grpc.ClientConn
}

func NewROTDClient(cc *grpc.ClientConn) ROTDClient {
	return &rOTDClient{cc}
}

func (c *rOTDClient) Certificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/proto.ROTD/Certificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ROTDServer is the server API for ROTD service.
type ROTDServer interface {
	Certificate(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

func RegisterROTDServer(s *grpc.Server, srv ROTDServer) {
	s.RegisterService(&_ROTD_serviceDesc, srv)
}

func _ROTD_Certificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ROTDServer).Certificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ROTD/Certificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ROTDServer).Certificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ROTD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ROTD",
	HandlerType: (*ROTDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certificate",
			Handler:    _ROTD_Certificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_62f8d9161514c8eb) }

var fileDescriptor_api_62f8d9161514c8eb = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6a, 0x5c, 0x42, 0xce, 0xa9, 0x45,
	0x25, 0x99, 0x69, 0x99, 0xc9, 0x89, 0x25, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x02, 0x5c, 0xcc, 0xc9, 0xc5, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x20, 0xa6, 0x92,
	0x36, 0x97, 0x30, 0x8a, 0xba, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0xa4,
	0xca, 0x92, 0xd4, 0x62, 0xa8, 0x52, 0x08, 0xc7, 0xc8, 0x8f, 0x8b, 0x25, 0xc8, 0x3f, 0xc4, 0x45,
	0xc8, 0x8d, 0x8b, 0x1b, 0x49, 0x93, 0x90, 0x24, 0xc4, 0x6a, 0x3d, 0x4c, 0x0b, 0xa5, 0xa4, 0xb0,
	0x49, 0x41, 0xec, 0x50, 0x62, 0x48, 0x62, 0x03, 0x4b, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb4, 0xaa, 0x0a, 0x1e, 0xbf, 0x00, 0x00, 0x00,
}
