package model

import (
	"reflect"
	"strings"
)

//go:generate protoc -I../proto --go_out=plugins=grpc:${GOPATH}/src ../proto/model.proto

type PortalDriver interface {
	isPortalDriver()
	String() string
}

func (VpnDriver_Type) isPortalDriver()     {}
func (NetworkDriver_Type) isPortalDriver() {}

type VpnParam interface {
	isVpnType()
}

func (SoftEtherParam) isVpnType() {}

func (d *VpnDriver) GetVpnParams() VpnParam {
	if d == nil {
		return nil
	}

	t := reflect.TypeOf(d.ServerParams)
	driverType := strings.TrimPrefix(t.String(), "*model.VpnDriver_")
	method := reflect.ValueOf(d).MethodByName(strings.Join([]string{"Get", driverType}, ""))
	// might want a better solution, call() is slow
	resp := method.Call(nil)
	return resp[0].Interface().(VpnParam)
}

type NicParam interface {
	isNicType()
}

func (OpenVnetParam) isNicType() {}

func (d *NetworkDriver) GetNicParams() NicParam {
	if d == nil {
		return nil
	}

	t := reflect.TypeOf(d.InterfaceParams)
	driverType := strings.TrimPrefix(t.String(), "*model.NetworkDriver_")
	method := reflect.ValueOf(d).MethodByName(strings.Join([]string{"Get", driverType}, ""))
	// might want a better solution, call() is slow
	resp := method.Call(nil)
	return resp[0].Interface().(NicParam)
}
