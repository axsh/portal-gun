package softether

import (
	"fmt"
	"github.com/axsh/portal-gun/driver"
	"github.com/axsh/portal-gun/model"
)

type SoftEtherDriver struct {
}

func init() {
	driver.Register(model.VpnDriver_SOFTETHER_VPN, func() (driver.Driver, error) {
		return &SoftEtherDriver{}, nil
	})
}

func (d *SoftEtherDriver) GenerateConfig(p model.VpnParam) error {
	fmt.Println("softether config")
	return nil
}

func (d *SoftEtherDriver) StartVpn() error {
	return nil
}

func (d *SoftEtherDriver) StopVpn() error {
	return nil
}

func (d *SoftEtherDriver) RemoveVpn() error {

	return nil
}

func (*SoftEtherDriver) IsDriver() {}
func (*SoftEtherDriver) DriverType() model.PortalDriver { return model.VpnDriver_SOFTETHER_VPN }
