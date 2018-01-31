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
import model "github.com/axsh/vpnhub/model"

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
	VpnServer *model.VpnServer `protobuf:"bytes,1,opt,name=vpn_server,json=vpnServer" json:"vpn_server,omitempty"`
}

func (m *CreateVpnRequest) Reset()                    { *m = CreateVpnRequest{} }
func (m *CreateVpnRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVpnRequest) ProtoMessage()               {}
func (*CreateVpnRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateVpnRequest) GetVpnServer() *model.VpnServer {
	if m != nil {
		return m.VpnServer
	}
	return nil
}

type DestroyVpnRequest struct {
	VpnName string `protobuf:"bytes,1,opt,name=vpn_name,json=vpnName" json:"vpn_name,omitempty"`
}

func (m *DestroyVpnRequest) Reset()                    { *m = DestroyVpnRequest{} }
func (m *DestroyVpnRequest) String() string            { return proto.CompactTextString(m) }
func (*DestroyVpnRequest) ProtoMessage()               {}
func (*DestroyVpnRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DestroyVpnRequest) GetVpnName() string {
	if m != nil {
		return m.VpnName
	}
	return ""
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
	Nic *model.Nic `protobuf:"bytes,1,opt,name=nic" json:"nic,omitempty"`
}

func (m *RegisterNicRequest) Reset()                    { *m = RegisterNicRequest{} }
func (m *RegisterNicRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterNicRequest) ProtoMessage()               {}
func (*RegisterNicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterNicRequest) GetNic() *model.Nic {
	if m != nil {
		return m.Nic
	}
	return nil
}

type DeregisterNicRequest struct {
	Nic *model.Nic `protobuf:"bytes,1,opt,name=nic" json:"nic,omitempty"`
}

func (m *DeregisterNicRequest) Reset()                    { *m = DeregisterNicRequest{} }
func (m *DeregisterNicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeregisterNicRequest) ProtoMessage()               {}
func (*DeregisterNicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeregisterNicRequest) GetNic() *model.Nic {
	if m != nil {
		return m.Nic
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
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x95, 0x90, 0xf0, 0x31, 0x10, 0xc5, 0x15, 0x44, 0x1a, 0x0e, 0xda, 0x93, 0x89, 0xa6, 0x4d,
	0xea, 0x57, 0x4c, 0x3c, 0x01, 0x17, 0x2f, 0x1c, 0x6a, 0xc2, 0xc1, 0x0b, 0x59, 0xca, 0x08, 0x9b,
	0xd0, 0xed, 0xba, 0x5d, 0x1a, 0x39, 0xf9, 0x0f, 0xfc, 0xcd, 0xa6, 0xdb, 0x56, 0xb1, 0x35, 0x7e,
	0xdc, 0x36, 0x6f, 0xe6, 0xbd, 0x79, 0xf3, 0x66, 0xa1, 0x19, 0xa2, 0x8c, 0x50, 0x5a, 0x42, 0x06,
	0x2a, 0x20, 0x65, 0x2a, 0x98, 0xd1, 0xf0, 0x83, 0x39, 0xae, 0x12, 0xc4, 0x1c, 0x42, 0x6b, 0x28,
	0x91, 0x2a, 0x9c, 0x08, 0xee, 0xe2, 0xf3, 0x1a, 0x43, 0x45, 0x6c, 0x80, 0x48, 0xf0, 0x69, 0xc2,
	0x3c, 0x2a, 0x1d, 0x97, 0x4e, 0x1b, 0x4e, 0xcb, 0x4a, 0x58, 0x13, 0xc1, 0x1f, 0x34, 0xee, 0xd6,
	0xa3, 0xec, 0x69, 0x5a, 0xb0, 0x3f, 0xc2, 0x50, 0xc9, 0x60, 0xb3, 0xa5, 0xd2, 0x83, 0x5a, 0xac,
	0xc2, 0xa9, 0x8f, 0x5a, 0xa3, 0xee, 0x56, 0x23, 0xc1, 0xc7, 0xd4, 0x47, 0xf3, 0x0c, 0x76, 0xb7,
	0x86, 0x8a, 0xd5, 0xe6, 0xa7, 0xe6, 0x73, 0xd8, 0xdb, 0x16, 0xff, 0xa5, 0xdb, 0x01, 0xe2, 0xe2,
	0x82, 0x85, 0x0a, 0xe5, 0x98, 0x79, 0x99, 0x97, 0x3e, 0x94, 0x39, 0xf3, 0xd2, 0x55, 0x20, 0x5d,
	0x25, 0xae, 0xc7, 0xb0, 0x79, 0x09, 0xed, 0x11, 0xca, 0xff, 0xb2, 0xae, 0xa0, 0xf5, 0x65, 0x52,
	0x6c, 0xec, 0x04, 0x9a, 0x8c, 0x2b, 0x94, 0x4f, 0xd4, 0xc3, 0x29, 0x9b, 0xa7, 0xe6, 0x1a, 0x1f,
	0xd8, 0xfd, 0xdc, 0xbc, 0x01, 0x92, 0x1b, 0xf6, 0x37, 0xa2, 0xf3, 0x0a, 0x90, 0x86, 0xcf, 0x3c,
	0x24, 0xd7, 0x50, 0x49, 0x22, 0x24, 0x1d, 0x8b, 0x0a, 0x66, 0xe5, 0x8f, 0x68, 0x1c, 0xe4, 0x61,
	0xb1, 0xda, 0x98, 0x3b, 0xe4, 0x16, 0xaa, 0x69, 0x9a, 0xe4, 0x50, 0x77, 0x14, 0x0e, 0x67, 0xb4,
	0x0b, 0xb8, 0xa6, 0x3a, 0x6f, 0x25, 0x80, 0x31, 0xf3, 0x32, 0x07, 0x77, 0x50, 0xcb, 0xf6, 0x27,
	0x5d, 0x4d, 0x29, 0x06, 0x6f, 0x74, 0x8a, 0x85, 0xc4, 0xc7, 0x00, 0xe0, 0x33, 0x06, 0xd2, 0x4b,
	0x47, 0x16, 0x8f, 0x60, 0x74, 0xbf, 0x2b, 0x69, 0x8d, 0x41, 0xff, 0xd1, 0x58, 0x30, 0xb5, 0x5c,
	0xcf, 0x2c, 0x2f, 0xf0, 0x6d, 0xfa, 0x12, 0x2e, 0xed, 0x48, 0xf0, 0xe5, 0x7a, 0x66, 0x53, 0xc1,
	0x66, 0x15, 0xfd, 0xc1, 0x2f, 0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0x89, 0xf9, 0xaf, 0xc9, 0x02,
	0x03, 0x00, 0x00,
}
