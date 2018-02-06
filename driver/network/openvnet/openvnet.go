package openvnet

import (
	"fmt"
	"github.com/axsh/portal-gun/driver"
	"github.com/axsh/portal-gun/model"
)

type OpenVnetDriver struct {
	params *model.OpenVnetParam
}

func init() {
	driver.Register(model.NetworkDriver_OPENVNET, func() (driver.Driver, error) {
		return &OpenVnetDriver{}, nil
	})
}

func (d *OpenVnetDriver) RegisterNic(nic model.NicParam) (string, error) {
	d.params = nic.(*model.OpenVnetParam)

	fmt.Println(d.params.GetIpLease())
	fmt.Println("openvnet regitster")
	return "", nil
}

func (d *OpenVnetDriver) DeregisterNic(nic model.NicParam) (string, error) {
	return "", nil
}

func (*OpenVnetDriver) IsDriver()                      {}
func (*OpenVnetDriver) DriverType() model.PortalDriver { return model.NetworkDriver_OPENVNET }
