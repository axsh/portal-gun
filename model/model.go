package model

//go:generate protoc -I../proto --go_out=plugins=grpc:${GOPATH}/src ../proto/model.proto

type PortalDriver interface {
	isPortalDriver()
	String() string
}

func (VpnDriver_Type) isPortalDriver() {}
func (NetworkDriver_Type) isPortalDriver() {}

type VpnParam interface {
	isVpnType()
}
func (SoftEtherParam) isVpnType() {}

func (d *VpnDriver) GetVpnParams() VpnParam {
	return d.GetServerParams().(VpnParam)
}

type NicParam interface {
	isNicParam()
}

func (OpenVNetParam) isNicParam() {}

func (d *NetworkDriver) GetNicParams() NicParam {
	return d.GetInterfaceParams().(NicParam)
}
