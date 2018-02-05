package openvnet

import (
	"fmt"
	"github.com/axsh/portal-gun/driver"
	"github.com/axsh/portal-gun/model"
)

type OpenVNetDriver struct {
	// params *driver.OpenVNetInterface
}

func init() {
	driver.Register(model.NetworkDriver_OPENVNET, func() (driver.Driver, error) {
		return &OpenVNetDriver{}, nil
	})
}

func (d *OpenVNetDriver) RegisterNic(nic model.NicParam) (string, error) {
	p := nic.(model.OpenVNetParam)
	fmt.Println(p)
	fmt.Println("openvnet regitster")
	return "", nil
}

func (d *OpenVNetDriver) DeregisterNic(nic model.NicParam) (string, error) {
	return "", nil
}

func (*OpenVNetDriver) IsDriver() {}
