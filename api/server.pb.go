// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	CreateVpnRequest
	DestroyVpnRequest
	CreateVpnReply
	DestroyVpnReply
	RegisterNicRequest
	DeregisterNicRequest
	RegisterNicReply
	DeregisterNicReply
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "github.com/axsh/portal-gun/model"

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

type CreateVpnRequest struct {
	VpnDriver *model.VpnDriver `protobuf:"bytes,1,opt,name=vpn_driver,json=vpnDriver" json:"vpn_driver,omitempty"`
}

func (m *CreateVpnRequest) Reset()                    { *m = CreateVpnRequest{} }
func (m *CreateVpnRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVpnRequest) ProtoMessage()               {}
func (*CreateVpnRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateVpnRequest) GetVpnDriver() *model.VpnDriver {
	if m != nil {
		return m.VpnDriver
	}
	return nil
}

type DestroyVpnRequest struct {
	VpnDriver *model.VpnDriver `protobuf:"bytes,1,opt,name=vpn_driver,json=vpnDriver" json:"vpn_driver,omitempty"`
}

func (m *DestroyVpnRequest) Reset()                    { *m = DestroyVpnRequest{} }
func (m *DestroyVpnRequest) String() string            { return proto.CompactTextString(m) }
func (*DestroyVpnRequest) ProtoMessage()               {}
func (*DestroyVpnRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DestroyVpnRequest) GetVpnDriver() *model.VpnDriver {
	if m != nil {
		return m.VpnDriver
	}
	return nil
}

type CreateVpnReply struct {
	VpnName string `protobuf:"bytes,1,opt,name=vpn_name,json=vpnName" json:"vpn_name,omitempty"`
}

func (m *CreateVpnReply) Reset()                    { *m = CreateVpnReply{} }
func (m *CreateVpnReply) String() string            { return proto.CompactTextString(m) }
func (*CreateVpnReply) ProtoMessage()               {}
func (*CreateVpnReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateVpnReply) GetVpnName() string {
	if m != nil {
		return m.VpnName
	}
	return ""
}

type DestroyVpnReply struct {
	VpnName string `protobuf:"bytes,1,opt,name=vpn_name,json=vpnName" json:"vpn_name,omitempty"`
}

func (m *DestroyVpnReply) Reset()                    { *m = DestroyVpnReply{} }
func (m *DestroyVpnReply) String() string            { return proto.CompactTextString(m) }
func (*DestroyVpnReply) ProtoMessage()               {}
func (*DestroyVpnReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DestroyVpnReply) GetVpnName() string {
	if m != nil {
		return m.VpnName
	}
	return ""
}

type RegisterNicRequest struct {
	NetworkDriver *model.NetworkDriver `protobuf:"bytes,1,opt,name=network_driver,json=networkDriver" json:"network_driver,omitempty"`
}

func (m *RegisterNicRequest) Reset()                    { *m = RegisterNicRequest{} }
func (m *RegisterNicRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterNicRequest) ProtoMessage()               {}
func (*RegisterNicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterNicRequest) GetNetworkDriver() *model.NetworkDriver {
	if m != nil {
		return m.NetworkDriver
	}
	return nil
}

type DeregisterNicRequest struct {
	NetworkDriver *model.NetworkDriver `protobuf:"bytes,1,opt,name=network_driver,json=networkDriver" json:"network_driver,omitempty"`
}

func (m *DeregisterNicRequest) Reset()                    { *m = DeregisterNicRequest{} }
func (m *DeregisterNicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeregisterNicRequest) ProtoMessage()               {}
func (*DeregisterNicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeregisterNicRequest) GetNetworkDriver() *model.NetworkDriver {
	if m != nil {
		return m.NetworkDriver
	}
	return nil
}

type RegisterNicReply struct {
	InterfaceId string `protobuf:"bytes,1,opt,name=interface_id,json=interfaceId" json:"interface_id,omitempty"`
}

func (m *RegisterNicReply) Reset()                    { *m = RegisterNicReply{} }
func (m *RegisterNicReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterNicReply) ProtoMessage()               {}
func (*RegisterNicReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RegisterNicReply) GetInterfaceId() string {
	if m != nil {
		return m.InterfaceId
	}
	return ""
}

type DeregisterNicReply struct {
	InterfaceId string `protobuf:"bytes,1,opt,name=interface_id,json=interfaceId" json:"interface_id,omitempty"`
}

func (m *DeregisterNicReply) Reset()                    { *m = DeregisterNicReply{} }
func (m *DeregisterNicReply) String() string            { return proto.CompactTextString(m) }
func (*DeregisterNicReply) ProtoMessage()               {}
func (*DeregisterNicReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeregisterNicReply) GetInterfaceId() string {
	if m != nil {
		return m.InterfaceId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateVpnRequest)(nil), "api.CreateVpnRequest")
	proto.RegisterType((*DestroyVpnRequest)(nil), "api.DestroyVpnRequest")
	proto.RegisterType((*CreateVpnReply)(nil), "api.CreateVpnReply")
	proto.RegisterType((*DestroyVpnReply)(nil), "api.DestroyVpnReply")
	proto.RegisterType((*RegisterNicRequest)(nil), "api.RegisterNicRequest")
	proto.RegisterType((*DeregisterNicRequest)(nil), "api.DeregisterNicRequest")
	proto.RegisterType((*RegisterNicReply)(nil), "api.RegisterNicReply")
	proto.RegisterType((*DeregisterNicReply)(nil), "api.DeregisterNicReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VpnService service

type VpnServiceClient interface {
	Create(ctx context.Context, in *CreateVpnRequest, opts ...grpc.CallOption) (*CreateVpnReply, error)
	Destroy(ctx context.Context, in *DestroyVpnRequest, opts ...grpc.CallOption) (*DestroyVpnReply, error)
}

type vpnServiceClient struct {
	cc *grpc.ClientConn
}

func NewVpnServiceClient(cc *grpc.ClientConn) VpnServiceClient {
	return &vpnServiceClient{cc}
}

func (c *vpnServiceClient) Create(ctx context.Context, in *CreateVpnRequest, opts ...grpc.CallOption) (*CreateVpnReply, error) {
	out := new(CreateVpnReply)
	err := grpc.Invoke(ctx, "/api.VpnService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vpnServiceClient) Destroy(ctx context.Context, in *DestroyVpnRequest, opts ...grpc.CallOption) (*DestroyVpnReply, error) {
	out := new(DestroyVpnReply)
	err := grpc.Invoke(ctx, "/api.VpnService/Destroy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VpnService service

type VpnServiceServer interface {
	Create(context.Context, *CreateVpnRequest) (*CreateVpnReply, error)
	Destroy(context.Context, *DestroyVpnRequest) (*DestroyVpnReply, error)
}

func RegisterVpnServiceServer(s *grpc.Server, srv VpnServiceServer) {
	s.RegisterService(&_VpnService_serviceDesc, srv)
}

func _VpnService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVpnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VpnServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VpnService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VpnServiceServer).Create(ctx, req.(*CreateVpnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VpnService_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyVpnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VpnServiceServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VpnService/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VpnServiceServer).Destroy(ctx, req.(*DestroyVpnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VpnService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.VpnService",
	HandlerType: (*VpnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VpnService_Create_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _VpnService_Destroy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

// Client API for NicService service

type NicServiceClient interface {
	Register(ctx context.Context, in *RegisterNicRequest, opts ...grpc.CallOption) (*RegisterNicReply, error)
	Deregister(ctx context.Context, in *DeregisterNicRequest, opts ...grpc.CallOption) (*DeregisterNicReply, error)
}

type nicServiceClient struct {
	cc *grpc.ClientConn
}

func NewNicServiceClient(cc *grpc.ClientConn) NicServiceClient {
	return &nicServiceClient{cc}
}

func (c *nicServiceClient) Register(ctx context.Context, in *RegisterNicRequest, opts ...grpc.CallOption) (*RegisterNicReply, error) {
	out := new(RegisterNicReply)
	err := grpc.Invoke(ctx, "/api.NicService/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nicServiceClient) Deregister(ctx context.Context, in *DeregisterNicRequest, opts ...grpc.CallOption) (*DeregisterNicReply, error) {
	out := new(DeregisterNicReply)
	err := grpc.Invoke(ctx, "/api.NicService/Deregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NicService service

type NicServiceServer interface {
	Register(context.Context, *RegisterNicRequest) (*RegisterNicReply, error)
	Deregister(context.Context, *DeregisterNicRequest) (*DeregisterNicReply, error)
}

func RegisterNicServiceServer(s *grpc.Server, srv NicServiceServer) {
	s.RegisterService(&_NicService_serviceDesc, srv)
}

func _NicService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NicServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NicService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NicServiceServer).Register(ctx, req.(*RegisterNicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NicService_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterNicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NicServiceServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NicService/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NicServiceServer).Deregister(ctx, req.(*DeregisterNicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NicService",
	HandlerType: (*NicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _NicService_Register_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _NicService_Deregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x4f, 0xe2, 0x40,
	0x14, 0xc7, 0x97, 0x6c, 0xc2, 0x8f, 0x07, 0xcb, 0xb2, 0xb3, 0xb0, 0x2c, 0x3d, 0x6c, 0xd8, 0x9e,
	0x4c, 0xd4, 0x36, 0xc1, 0xa8, 0x31, 0x7a, 0x82, 0x5e, 0xbc, 0x34, 0xb1, 0x24, 0x1c, 0xbc, 0x90,
	0xa1, 0x7d, 0xc2, 0xc4, 0x76, 0x3a, 0x4e, 0x87, 0x2a, 0x27, 0xff, 0x03, 0xff, 0x66, 0x43, 0x7f,
	0x28, 0x50, 0x13, 0x4d, 0xf4, 0xd6, 0x7e, 0xdf, 0xbc, 0xef, 0x7c, 0x5e, 0xbf, 0xaf, 0xd0, 0x88,
	0x50, 0xc6, 0x28, 0x0d, 0x21, 0x43, 0x15, 0x92, 0xef, 0x54, 0x30, 0xad, 0x1e, 0x84, 0x1e, 0xfa,
	0xa9, 0xa2, 0x8f, 0xa0, 0x35, 0x92, 0x48, 0x15, 0x4e, 0x04, 0x77, 0xf0, 0x6e, 0x89, 0x91, 0x22,
	0x26, 0x40, 0x2c, 0xf8, 0xd4, 0x93, 0x2c, 0x46, 0xf9, 0xb7, 0xd4, 0x2f, 0xed, 0xd5, 0x07, 0x2d,
	0x23, 0xed, 0x9a, 0x08, 0x6e, 0x25, 0xba, 0x53, 0x8b, 0xf3, 0x47, 0xdd, 0x82, 0x5f, 0x16, 0x46,
	0x4a, 0x86, 0xab, 0xcf, 0xb8, 0xec, 0x43, 0x73, 0x03, 0x45, 0xf8, 0x2b, 0xd2, 0x83, 0xea, 0xda,
	0x82, 0xd3, 0x00, 0x13, 0x83, 0x9a, 0x53, 0x89, 0x05, 0xb7, 0x69, 0x80, 0xfa, 0x01, 0xfc, 0xdc,
	0xbc, 0xf2, 0x9d, 0xd3, 0x57, 0x40, 0x1c, 0x9c, 0xb3, 0x48, 0xa1, 0xb4, 0x99, 0x9b, 0x13, 0x9e,
	0x43, 0x93, 0xa3, 0xba, 0x0f, 0xe5, 0xed, 0x36, 0x65, 0x3b, 0xa3, 0xb4, 0xd3, 0x62, 0x46, 0xfa,
	0x83, 0x6f, 0xbe, 0xea, 0x63, 0x68, 0x5b, 0x28, 0xbf, 0xd8, 0xf4, 0x18, 0x5a, 0x5b, 0x9c, 0xeb,
	0xb1, 0xfe, 0x43, 0x83, 0x71, 0x85, 0xf2, 0x86, 0xba, 0x38, 0x65, 0x5e, 0x36, 0x5a, 0xfd, 0x45,
	0xbb, 0xf4, 0xf4, 0x53, 0x20, 0x3b, 0x2c, 0x1f, 0x6b, 0x1c, 0x3c, 0x02, 0x4c, 0x04, 0x1f, 0xa3,
	0x8c, 0x99, 0x8b, 0xe4, 0x04, 0xca, 0x69, 0x00, 0xa4, 0x63, 0x50, 0xc1, 0x8c, 0xdd, 0xc5, 0xd0,
	0x7e, 0xef, 0xca, 0xc2, 0x5f, 0xe9, 0xdf, 0xc8, 0x19, 0x54, 0xb2, 0x2c, 0xc8, 0x9f, 0xe4, 0x44,
	0x61, 0x19, 0xb4, 0x76, 0x41, 0x4f, 0x5a, 0x07, 0x4f, 0x25, 0x00, 0x9b, 0xb9, 0x39, 0xc1, 0x05,
	0x54, 0xf3, 0xf9, 0x49, 0x37, 0x69, 0x29, 0xc6, 0xa6, 0x75, 0x8a, 0x85, 0x94, 0x63, 0x08, 0xf0,
	0xfa, 0x19, 0x48, 0x2f, 0xbb, 0xb2, 0x98, 0x91, 0xd6, 0x7d, 0xab, 0x94, 0x78, 0x0c, 0xfb, 0xd7,
	0xff, 0xe6, 0x4c, 0x2d, 0x96, 0x33, 0xc3, 0x0d, 0x03, 0x93, 0x3e, 0x44, 0x0b, 0x53, 0x84, 0x52,
	0x51, 0xff, 0x70, 0xbe, 0xe4, 0x26, 0x15, 0x6c, 0x56, 0x4e, 0x7e, 0x9c, 0xa3, 0xe7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb9, 0xfa, 0xff, 0x55, 0x5a, 0x03, 0x00, 0x00,
}
