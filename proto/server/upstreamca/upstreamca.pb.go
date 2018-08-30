// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upstreamca.proto

package upstreamca

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import plugin "github.com/spiffe/spire/proto/common/plugin"

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

// ConfigureRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureRequest = plugin.ConfigureRequest

// GlobalConfig from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureRequest_GlobalConfig = plugin.ConfigureRequest_GlobalConfig

// ConfigureResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type ConfigureResponse = plugin.ConfigureResponse

// GetPluginInfoRequest from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoRequest = plugin.GetPluginInfoRequest

// GetPluginInfoResponse from public import github.com/spiffe/spire/proto/common/plugin/plugin.proto
type GetPluginInfoResponse = plugin.GetPluginInfoResponse

type SubmitCSRRequest struct {
	// * Certificate signing request.
	Csr                  []byte   `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitCSRRequest) Reset()         { *m = SubmitCSRRequest{} }
func (m *SubmitCSRRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitCSRRequest) ProtoMessage()    {}
func (*SubmitCSRRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_upstreamca_9a8b85c164a5df1d, []int{0}
}
func (m *SubmitCSRRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitCSRRequest.Unmarshal(m, b)
}
func (m *SubmitCSRRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitCSRRequest.Marshal(b, m, deterministic)
}
func (dst *SubmitCSRRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitCSRRequest.Merge(dst, src)
}
func (m *SubmitCSRRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitCSRRequest.Size(m)
}
func (m *SubmitCSRRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitCSRRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitCSRRequest proto.InternalMessageInfo

func (m *SubmitCSRRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type SubmitCSRResponse struct {
	// * Signed certificate
	Cert []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	// * Upstream trust bundle.
	UpstreamTrustBundle  []byte   `protobuf:"bytes,2,opt,name=upstreamTrustBundle,proto3" json:"upstreamTrustBundle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitCSRResponse) Reset()         { *m = SubmitCSRResponse{} }
func (m *SubmitCSRResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitCSRResponse) ProtoMessage()    {}
func (*SubmitCSRResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_upstreamca_9a8b85c164a5df1d, []int{1}
}
func (m *SubmitCSRResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitCSRResponse.Unmarshal(m, b)
}
func (m *SubmitCSRResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitCSRResponse.Marshal(b, m, deterministic)
}
func (dst *SubmitCSRResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitCSRResponse.Merge(dst, src)
}
func (m *SubmitCSRResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitCSRResponse.Size(m)
}
func (m *SubmitCSRResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitCSRResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitCSRResponse proto.InternalMessageInfo

func (m *SubmitCSRResponse) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *SubmitCSRResponse) GetUpstreamTrustBundle() []byte {
	if m != nil {
		return m.UpstreamTrustBundle
	}
	return nil
}

func init() {
	proto.RegisterType((*SubmitCSRRequest)(nil), "spire.server.upstreamca.SubmitCSRRequest")
	proto.RegisterType((*SubmitCSRResponse)(nil), "spire.server.upstreamca.SubmitCSRResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UpstreamCA service

type UpstreamCAClient interface {
	// * Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
	// * Will take in a CSR and submit it to the upstream CA for signing
	// (“upstream” CA can be local self-signed root in simple case).
	SubmitCSR(ctx context.Context, in *SubmitCSRRequest, opts ...grpc.CallOption) (*SubmitCSRResponse, error)
}

type upstreamCAClient struct {
	cc *grpc.ClientConn
}

func NewUpstreamCAClient(cc *grpc.ClientConn) UpstreamCAClient {
	return &upstreamCAClient{cc}
}

func (c *upstreamCAClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := grpc.Invoke(ctx, "/spire.server.upstreamca.UpstreamCA/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamCAClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/spire.server.upstreamca.UpstreamCA/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *upstreamCAClient) SubmitCSR(ctx context.Context, in *SubmitCSRRequest, opts ...grpc.CallOption) (*SubmitCSRResponse, error) {
	out := new(SubmitCSRResponse)
	err := grpc.Invoke(ctx, "/spire.server.upstreamca.UpstreamCA/SubmitCSR", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UpstreamCA service

type UpstreamCAServer interface {
	// * Responsible for configuration of the plugin.
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
	// * Will take in a CSR and submit it to the upstream CA for signing
	// (“upstream” CA can be local self-signed root in simple case).
	SubmitCSR(context.Context, *SubmitCSRRequest) (*SubmitCSRResponse, error)
}

func RegisterUpstreamCAServer(s *grpc.Server, srv UpstreamCAServer) {
	s.RegisterService(&_UpstreamCA_serviceDesc, srv)
}

func _UpstreamCA_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamca.UpstreamCA/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamCA_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamca.UpstreamCA/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpstreamCA_SubmitCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpstreamCAServer).SubmitCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.upstreamca.UpstreamCA/SubmitCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpstreamCAServer).SubmitCSR(ctx, req.(*SubmitCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UpstreamCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.server.upstreamca.UpstreamCA",
	HandlerType: (*UpstreamCAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _UpstreamCA_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _UpstreamCA_GetPluginInfo_Handler,
		},
		{
			MethodName: "SubmitCSR",
			Handler:    _UpstreamCA_SubmitCSR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upstreamca.proto",
}

func init() { proto.RegisterFile("upstreamca.proto", fileDescriptor_upstreamca_9a8b85c164a5df1d) }

var fileDescriptor_upstreamca_9a8b85c164a5df1d = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbb, 0x4e, 0x33, 0x31,
	0x10, 0x85, 0xff, 0xe4, 0x47, 0x48, 0x19, 0x05, 0x69, 0x31, 0x05, 0x51, 0x2a, 0x14, 0x01, 0x02,
	0x0a, 0x2f, 0x82, 0x86, 0x96, 0x6c, 0x81, 0xe8, 0xa2, 0x0d, 0x14, 0xa4, 0xcb, 0x9a, 0xf1, 0x62,
	0x29, 0xbe, 0xe0, 0x0b, 0x4f, 0xc6, 0x03, 0x22, 0xbc, 0xde, 0x0d, 0x97, 0x08, 0xa8, 0x3c, 0xf2,
	0xf9, 0xc6, 0xe7, 0x78, 0x06, 0xb2, 0x60, 0x9c, 0xb7, 0xb8, 0x94, 0x6c, 0x49, 0x8d, 0xd5, 0x5e,
	0x93, 0x7d, 0x67, 0x84, 0x45, 0xea, 0xd0, 0xbe, 0xa0, 0xa5, 0x6b, 0x79, 0x7c, 0x55, 0x0b, 0xff,
	0x14, 0x2a, 0xca, 0xb4, 0xcc, 0x9d, 0x11, 0x9c, 0x63, 0x1e, 0xd1, 0x3c, 0xf6, 0xe5, 0x4c, 0x4b,
	0xa9, 0x55, 0x6e, 0x56, 0xa1, 0x16, 0xed, 0xd1, 0x3c, 0x39, 0x39, 0x84, 0x6c, 0x1e, 0x2a, 0x29,
	0x7c, 0x31, 0x2f, 0x4b, 0x7c, 0x0e, 0xe8, 0x3c, 0xc9, 0xe0, 0x3f, 0x73, 0x76, 0xd4, 0x3b, 0xe8,
	0x9d, 0x0c, 0xcb, 0xf7, 0x72, 0xf2, 0x00, 0xbb, 0x1f, 0x28, 0x67, 0xb4, 0x72, 0x48, 0x08, 0x6c,
	0x31, 0xb4, 0x3e, 0x71, 0xb1, 0x26, 0xe7, 0xb0, 0xd7, 0xc6, 0xba, 0xb3, 0xc1, 0xf9, 0x69, 0x50,
	0x8f, 0x2b, 0x1c, 0xf5, 0x23, 0xb2, 0x49, 0xba, 0x78, 0xed, 0x03, 0xdc, 0xa7, 0xfb, 0xe2, 0x9a,
	0x2c, 0x60, 0x50, 0x68, 0xc5, 0x45, 0x1d, 0x2c, 0x92, 0x23, 0xda, 0x7c, 0xb8, 0xc9, 0x4f, 0x53,
	0xf0, 0x4e, 0x4f, 0x79, 0xc7, 0xc7, 0xbf, 0x61, 0x29, 0x30, 0x87, 0x9d, 0x1b, 0xf4, 0xb3, 0x28,
	0xdf, 0x2a, 0xae, 0xc9, 0xe9, 0xc6, 0xc6, 0x4f, 0x4c, 0xeb, 0x71, 0xf6, 0x17, 0x34, 0xf9, 0x54,
	0x30, 0xe8, 0xa6, 0xd5, 0x79, 0x7c, 0x5b, 0x1a, 0xfd, 0x3a, 0xf7, 0xce, 0xe3, 0x47, 0xb4, 0xf1,
	0x98, 0x0e, 0x17, 0xb0, 0xd6, 0x67, 0xff, 0xaa, 0xed, 0xb8, 0xce, 0xcb, 0xb7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4f, 0x59, 0x5b, 0x78, 0x35, 0x02, 0x00, 0x00,
}
