package openvnet

import (
	"fmt"
	"github.com/axsh/vpnhub/driver"
	"github.com/axsh/vpnhub/model"
)

type OpenVNetDriver struct {
	// params *driver.OpenVNetInterface
}

func init() {
	d := model.Nic_OPENVNET
	driver.Register(d.String(), func() (driver.Driver, error) {
		return &OpenVNetDriver{}, nil
	})
}

func(d *OpenVNetDriver) RegisterNic(nic *model.Nic) (string, error) {
	// p := nic.GetInterfaceParams()
	fmt.Println("openvnet regitster")
	return "", nil
}

func(d *OpenVNetDriver) DeregisterNic(nic *model.Nic) (string, error) {
	return "", nil
}

func (*OpenVNetDriver) IsDriver() {}
