package softether

import (
	"fmt"
	"github.com/axsh/vpnhub/driver"
	"github.com/axsh/vpnhub/model"
)

type SoftEtherDriver struct {
}

func init() {
	d := model.VpnServer_SOFTETHER_VPN

	driver.Register(d.String(), func() (driver.Driver, error) {
		return &SoftEtherDriver{}, nil
	})
}

func (d *SoftEtherDriver) GenerateConfig(p *model.VpnServer) error {
	// params := p.GetParam().(driver.SoftEtherParam)
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
