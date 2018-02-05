// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model.proto

It has these top-level messages:
	VpnDriver
	SoftEtherParam
	NetworkDriver
	OpenVnetParam
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VpnDriver_Type int32

const (
	VpnDriver_NONE          VpnDriver_Type = 0
	VpnDriver_SOFTETHER_VPN VpnDriver_Type = 1
)

var VpnDriver_Type_name = map[int32]string{
	0: "NONE",
	1: "SOFTETHER_VPN",
}
var VpnDriver_Type_value = map[string]int32{
	"NONE":          0,
	"SOFTETHER_VPN": 1,
}

func (x VpnDriver_Type) String() string {
	return proto.EnumName(VpnDriver_Type_name, int32(x))
}
func (VpnDriver_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type NetworkDriver_Type int32

const (
	NetworkDriver_NONE     NetworkDriver_Type = 0
	NetworkDriver_OPENVNET NetworkDriver_Type = 1
)

var NetworkDriver_Type_name = map[int32]string{
	0: "NONE",
	1: "OPENVNET",
}
var NetworkDriver_Type_value = map[string]int32{
	"NONE":     0,
	"OPENVNET": 1,
}

func (x NetworkDriver_Type) String() string {
	return proto.EnumName(NetworkDriver_Type_name, int32(x))
}
func (NetworkDriver_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type VpnDriver struct {
	DriverType VpnDriver_Type `protobuf:"varint,1,opt,name=driver_type,json=driverType,enum=model.VpnDriver_Type" json:"driver_type,omitempty"`
	// Types that are valid to be assigned to ServerParams:
	//	*VpnDriver_SoftetherParams
	ServerParams isVpnDriver_ServerParams `protobuf_oneof:"ServerParams"`
}

func (m *VpnDriver) Reset()                    { *m = VpnDriver{} }
func (m *VpnDriver) String() string            { return proto.CompactTextString(m) }
func (*VpnDriver) ProtoMessage()               {}
func (*VpnDriver) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isVpnDriver_ServerParams interface {
	isVpnDriver_ServerParams()
}

type VpnDriver_SoftetherParams struct {
	SoftetherParams *SoftEtherParam `protobuf:"bytes,100,opt,name=softether_params,json=softetherParams,oneof"`
}

func (*VpnDriver_SoftetherParams) isVpnDriver_ServerParams() {}

func (m *VpnDriver) GetServerParams() isVpnDriver_ServerParams {
	if m != nil {
		return m.ServerParams
	}
	return nil
}

func (m *VpnDriver) GetDriverType() VpnDriver_Type {
	if m != nil {
		return m.DriverType
	}
	return VpnDriver_NONE
}

func (m *VpnDriver) GetSoftetherParams() *SoftEtherParam {
	if x, ok := m.GetServerParams().(*VpnDriver_SoftetherParams); ok {
		return x.SoftetherParams
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*VpnDriver) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _VpnDriver_OneofMarshaler, _VpnDriver_OneofUnmarshaler, _VpnDriver_OneofSizer, []interface{}{
		(*VpnDriver_SoftetherParams)(nil),
	}
}

func _VpnDriver_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*VpnDriver)
	// ServerParams
	switch x := m.ServerParams.(type) {
	case *VpnDriver_SoftetherParams:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SoftetherParams); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("VpnDriver.ServerParams has unexpected type %T", x)
	}
	return nil
}

func _VpnDriver_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*VpnDriver)
	switch tag {
	case 100: // ServerParams.softether_params
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SoftEtherParam)
		err := b.DecodeMessage(msg)
		m.ServerParams = &VpnDriver_SoftetherParams{msg}
		return true, err
	default:
		return false, nil
	}
}

func _VpnDriver_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*VpnDriver)
	// ServerParams
	switch x := m.ServerParams.(type) {
	case *VpnDriver_SoftetherParams:
		s := proto.Size(x.SoftetherParams)
		n += proto.SizeVarint(100<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SoftEtherParam struct {
	VirtualHub       string `protobuf:"bytes,1,opt,name=virtual_hub,json=virtualHub" json:"virtual_hub,omitempty"`
	Ipv4Network      string `protobuf:"bytes,2,opt,name=ipv4_network,json=ipv4Network" json:"ipv4_network,omitempty"`
	Ipv4BeginAddress uint32 `protobuf:"varint,3,opt,name=ipv4_begin_address,json=ipv4BeginAddress" json:"ipv4_begin_address,omitempty"`
	Ipv4EndAddress   uint32 `protobuf:"varint,4,opt,name=ipv4_end_address,json=ipv4EndAddress" json:"ipv4_end_address,omitempty"`
	TapId            string `protobuf:"bytes,5,opt,name=tap_id,json=tapId" json:"tap_id,omitempty"`
}

func (m *SoftEtherParam) Reset()                    { *m = SoftEtherParam{} }
func (m *SoftEtherParam) String() string            { return proto.CompactTextString(m) }
func (*SoftEtherParam) ProtoMessage()               {}
func (*SoftEtherParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SoftEtherParam) GetVirtualHub() string {
	if m != nil {
		return m.VirtualHub
	}
	return ""
}

func (m *SoftEtherParam) GetIpv4Network() string {
	if m != nil {
		return m.Ipv4Network
	}
	return ""
}

func (m *SoftEtherParam) GetIpv4BeginAddress() uint32 {
	if m != nil {
		return m.Ipv4BeginAddress
	}
	return 0
}

func (m *SoftEtherParam) GetIpv4EndAddress() uint32 {
	if m != nil {
		return m.Ipv4EndAddress
	}
	return 0
}

func (m *SoftEtherParam) GetTapId() string {
	if m != nil {
		return m.TapId
	}
	return ""
}

type NetworkDriver struct {
	DriverType NetworkDriver_Type `protobuf:"varint,1,opt,name=driver_type,json=driverType,enum=model.NetworkDriver_Type" json:"driver_type,omitempty"`
	// Types that are valid to be assigned to InterfaceParams:
	//	*NetworkDriver_OpenvnetParams
	InterfaceParams isNetworkDriver_InterfaceParams `protobuf_oneof:"InterfaceParams"`
}

func (m *NetworkDriver) Reset()                    { *m = NetworkDriver{} }
func (m *NetworkDriver) String() string            { return proto.CompactTextString(m) }
func (*NetworkDriver) ProtoMessage()               {}
func (*NetworkDriver) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isNetworkDriver_InterfaceParams interface {
	isNetworkDriver_InterfaceParams()
}

type NetworkDriver_OpenvnetParams struct {
	OpenvnetParams *OpenVnetParam `protobuf:"bytes,100,opt,name=openvnet_params,json=openvnetParams,oneof"`
}

func (*NetworkDriver_OpenvnetParams) isNetworkDriver_InterfaceParams() {}

func (m *NetworkDriver) GetInterfaceParams() isNetworkDriver_InterfaceParams {
	if m != nil {
		return m.InterfaceParams
	}
	return nil
}

func (m *NetworkDriver) GetDriverType() NetworkDriver_Type {
	if m != nil {
		return m.DriverType
	}
	return NetworkDriver_NONE
}

func (m *NetworkDriver) GetOpenvnetParams() *OpenVnetParam {
	if x, ok := m.GetInterfaceParams().(*NetworkDriver_OpenvnetParams); ok {
		return x.OpenvnetParams
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NetworkDriver) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NetworkDriver_OneofMarshaler, _NetworkDriver_OneofUnmarshaler, _NetworkDriver_OneofSizer, []interface{}{
		(*NetworkDriver_OpenvnetParams)(nil),
	}
}

func _NetworkDriver_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NetworkDriver)
	// InterfaceParams
	switch x := m.InterfaceParams.(type) {
	case *NetworkDriver_OpenvnetParams:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OpenvnetParams); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NetworkDriver.InterfaceParams has unexpected type %T", x)
	}
	return nil
}

func _NetworkDriver_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NetworkDriver)
	switch tag {
	case 100: // InterfaceParams.openvnet_params
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OpenVnetParam)
		err := b.DecodeMessage(msg)
		m.InterfaceParams = &NetworkDriver_OpenvnetParams{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NetworkDriver_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NetworkDriver)
	// InterfaceParams
	switch x := m.InterfaceParams.(type) {
	case *NetworkDriver_OpenvnetParams:
		s := proto.Size(x.OpenvnetParams)
		n += proto.SizeVarint(100<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type OpenVnetParam struct {
	InterfaceId string                  `protobuf:"bytes,1,opt,name=interface_id,json=interfaceId" json:"interface_id,omitempty"`
	NetworkId   string                  `protobuf:"bytes,2,opt,name=network_id,json=networkId" json:"network_id,omitempty"`
	IpLease     *OpenVnetParam_IpLease  `protobuf:"bytes,4,opt,name=ip_lease,json=ipLease" json:"ip_lease,omitempty"`
	MacLease    *OpenVnetParam_MacLease `protobuf:"bytes,5,opt,name=mac_lease,json=macLease" json:"mac_lease,omitempty"`
}

func (m *OpenVnetParam) Reset()                    { *m = OpenVnetParam{} }
func (m *OpenVnetParam) String() string            { return proto.CompactTextString(m) }
func (*OpenVnetParam) ProtoMessage()               {}
func (*OpenVnetParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OpenVnetParam) GetInterfaceId() string {
	if m != nil {
		return m.InterfaceId
	}
	return ""
}

func (m *OpenVnetParam) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *OpenVnetParam) GetIpLease() *OpenVnetParam_IpLease {
	if m != nil {
		return m.IpLease
	}
	return nil
}

func (m *OpenVnetParam) GetMacLease() *OpenVnetParam_MacLease {
	if m != nil {
		return m.MacLease
	}
	return nil
}

type OpenVnetParam_IpLease struct {
	Ipv4Address string `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	MacLeaseId  string `protobuf:"bytes,2,opt,name=mac_lease_id,json=macLeaseId" json:"mac_lease_id,omitempty"`
	NetworkId   string `protobuf:"bytes,3,opt,name=network_id,json=networkId" json:"network_id,omitempty"`
}

func (m *OpenVnetParam_IpLease) Reset()                    { *m = OpenVnetParam_IpLease{} }
func (m *OpenVnetParam_IpLease) String() string            { return proto.CompactTextString(m) }
func (*OpenVnetParam_IpLease) ProtoMessage()               {}
func (*OpenVnetParam_IpLease) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *OpenVnetParam_IpLease) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *OpenVnetParam_IpLease) GetMacLeaseId() string {
	if m != nil {
		return m.MacLeaseId
	}
	return ""
}

func (m *OpenVnetParam_IpLease) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

type OpenVnetParam_MacLease struct {
	InterfaceId string `protobuf:"bytes,1,opt,name=interface_id,json=interfaceId" json:"interface_id,omitempty"`
	MacAddr     string `protobuf:"bytes,2,opt,name=mac_addr,json=macAddr" json:"mac_addr,omitempty"`
}

func (m *OpenVnetParam_MacLease) Reset()                    { *m = OpenVnetParam_MacLease{} }
func (m *OpenVnetParam_MacLease) String() string            { return proto.CompactTextString(m) }
func (*OpenVnetParam_MacLease) ProtoMessage()               {}
func (*OpenVnetParam_MacLease) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 1} }

func (m *OpenVnetParam_MacLease) GetInterfaceId() string {
	if m != nil {
		return m.InterfaceId
	}
	return ""
}

func (m *OpenVnetParam_MacLease) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*VpnDriver)(nil), "model.VpnDriver")
	proto.RegisterType((*SoftEtherParam)(nil), "model.SoftEtherParam")
	proto.RegisterType((*NetworkDriver)(nil), "model.NetworkDriver")
	proto.RegisterType((*OpenVnetParam)(nil), "model.OpenVnetParam")
	proto.RegisterType((*OpenVnetParam_IpLease)(nil), "model.OpenVnetParam.IpLease")
	proto.RegisterType((*OpenVnetParam_MacLease)(nil), "model.OpenVnetParam.MacLease")
	proto.RegisterEnum("model.VpnDriver_Type", VpnDriver_Type_name, VpnDriver_Type_value)
	proto.RegisterEnum("model.NetworkDriver_Type", NetworkDriver_Type_name, NetworkDriver_Type_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0xcd, 0x24, 0x21, 0xc0, 0x35, 0x7f, 0x19, 0x7d, 0x91, 0x48, 0xf4, 0xa5, 0xa5, 0x74, 0xc3,
	0xa2, 0x05, 0x29, 0xad, 0x5a, 0x29, 0x9b, 0xaa, 0xa8, 0xae, 0x40, 0x6a, 0x0d, 0x32, 0x88, 0x45,
	0x37, 0xd6, 0xc0, 0x0c, 0x60, 0x15, 0x8f, 0x47, 0xe3, 0x81, 0x36, 0xef, 0x56, 0xa9, 0x8b, 0x3e,
	0x58, 0xab, 0x19, 0x8f, 0x1d, 0x51, 0x58, 0x74, 0x77, 0x7d, 0xee, 0x39, 0xf7, 0xef, 0x8c, 0xc1,
	0x89, 0x62, 0xca, 0x36, 0x5d, 0x21, 0x63, 0x15, 0xe3, 0x82, 0xf9, 0x68, 0xff, 0x40, 0x50, 0x9e,
	0x09, 0xfe, 0x41, 0x86, 0x3b, 0x26, 0xf1, 0x1b, 0x70, 0xa8, 0x89, 0x02, 0xf5, 0x20, 0x58, 0x13,
	0xb5, 0x50, 0xa7, 0x76, 0x77, 0xd5, 0x4d, 0x75, 0x39, 0xad, 0x3b, 0x7d, 0x10, 0xcc, 0x87, 0x94,
	0xa9, 0x63, 0xdc, 0x87, 0x46, 0x12, 0x2f, 0x15, 0x53, 0x6b, 0x26, 0x03, 0x41, 0x24, 0x89, 0x92,
	0x26, 0x6d, 0xa1, 0x8e, 0x93, 0x8b, 0x27, 0xf1, 0x52, 0xb9, 0x3a, 0x3d, 0xd6, 0xd9, 0xc1, 0x89,
	0x5f, 0xcf, 0x05, 0x06, 0x49, 0xda, 0xcf, 0xe1, 0xdc, 0xd4, 0x2a, 0xc1, 0xb9, 0x37, 0xf2, 0xdc,
	0xc6, 0x09, 0xbe, 0x84, 0xea, 0x64, 0xf4, 0x71, 0xea, 0x4e, 0x07, 0xae, 0x1f, 0xcc, 0xc6, 0x5e,
	0x03, 0xf5, 0x6b, 0x50, 0x99, 0x30, 0xb9, 0xcb, 0x45, 0xbf, 0x10, 0xd4, 0xf6, 0x4b, 0xe3, 0xa7,
	0xe0, 0xec, 0x42, 0xa9, 0xb6, 0x64, 0x13, 0xac, 0xb7, 0x73, 0xb3, 0x43, 0xd9, 0x07, 0x0b, 0x0d,
	0xb6, 0x73, 0xfc, 0x0c, 0x2a, 0xa1, 0xd8, 0xbd, 0x0e, 0x38, 0x53, 0xdf, 0x62, 0xf9, 0xb5, 0x79,
	0x6a, 0x18, 0x8e, 0xc6, 0xbc, 0x14, 0xc2, 0x2f, 0x00, 0x1b, 0xca, 0x9c, 0xad, 0x42, 0x1e, 0x10,
	0x4a, 0x25, 0x4b, 0x92, 0xe6, 0x59, 0x0b, 0x75, 0xaa, 0x7e, 0x43, 0x67, 0xfa, 0x3a, 0xf1, 0x3e,
	0xc5, 0x71, 0x07, 0x0c, 0x16, 0x30, 0x4e, 0x73, 0xee, 0xb9, 0xe1, 0xd6, 0x34, 0xee, 0x72, 0x9a,
	0x31, 0xaf, 0xe0, 0x42, 0x11, 0x11, 0x84, 0xb4, 0x59, 0x30, 0x4d, 0x0b, 0x8a, 0x88, 0x21, 0x6d,
	0xff, 0x44, 0x50, 0xb5, 0xad, 0xad, 0x11, 0xf7, 0xc7, 0x8c, 0xb8, 0xb6, 0xb7, 0xdc, 0xa3, 0x1e,
	0x9a, 0xf1, 0x0e, 0xea, 0xb1, 0x60, 0x7c, 0xc7, 0x99, 0xda, 0xf7, 0xe2, 0x3f, 0xab, 0x1f, 0x09,
	0xc6, 0x67, 0x9c, 0xa9, 0xcc, 0x8a, 0x5a, 0x46, 0xb7, 0x47, 0x7d, 0x72, 0xe0, 0x44, 0x05, 0x4a,
	0xa3, 0xb1, 0xeb, 0xcd, 0x3c, 0x77, 0xda, 0x40, 0xfd, 0x4b, 0xa8, 0x0f, 0xb9, 0x62, 0x72, 0x49,
	0x16, 0xcc, 0x4a, 0x7e, 0x9f, 0x42, 0x75, 0xaf, 0xac, 0xb9, 0x72, 0x46, 0xd2, 0x0b, 0x23, 0x7b,
	0xe5, 0x0c, 0x1b, 0x52, 0x7c, 0x0b, 0x60, 0x3d, 0xd0, 0x84, 0xd4, 0x86, 0xb2, 0x45, 0x86, 0x14,
	0xbf, 0x85, 0x52, 0x28, 0x82, 0x0d, 0x23, 0x09, 0x33, 0xe7, 0x74, 0xee, 0xfe, 0x3f, 0xb6, 0x40,
	0x77, 0x28, 0x3e, 0x69, 0x8e, 0x5f, 0x0c, 0xd3, 0x00, 0xdf, 0x43, 0x39, 0x22, 0x0b, 0xab, 0x2c,
	0x18, 0xe5, 0xed, 0x51, 0xe5, 0x67, 0xb2, 0x48, 0xa5, 0xa5, 0xc8, 0x46, 0x37, 0x11, 0x14, 0x6d,
	0xbd, 0xfc, 0x9d, 0x64, 0x96, 0xa2, 0xc7, 0x77, 0x92, 0xf9, 0xd9, 0x82, 0x4a, 0xde, 0xe9, 0x71,
	0x07, 0xc8, 0xaa, 0x1d, 0xec, 0x78, 0xf6, 0xd7, 0x8e, 0x37, 0x03, 0x28, 0x65, 0x43, 0xfc, 0xcb,
	0xc5, 0xae, 0x41, 0x4f, 0x6a, 0x26, 0xb2, 0xbd, 0x8a, 0x11, 0x59, 0xe8, 0x69, 0xfa, 0xed, 0x2f,
	0xad, 0x55, 0xa8, 0xd6, 0xdb, 0x79, 0x77, 0x11, 0x47, 0x3d, 0xf2, 0x3d, 0x59, 0xf7, 0x44, 0x2c,
	0x15, 0xd9, 0xbc, 0x5c, 0x6d, 0x79, 0xcf, 0x6c, 0x3f, 0xbf, 0x30, 0xbf, 0xfe, 0xab, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0xde, 0x6b, 0x1a, 0x09, 0x04, 0x00, 0x00,
}
