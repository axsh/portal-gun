package softether

import (
	"fmt"
	"github.com/axsh/vpnhub/driver"
)

type SoftEtherDriver struct {}

func (d *SoftEtherDriver) GenerateConfig(p *driver.VpnParams) error {
	params := p.(driver.SoftEtherParams)
	fmt.Println(p)
}

func (d *SoftEtherDriver) CreateVpn() error {
	return nil
}

func (d *SoftEtherDriver) DestroyVpn() error {
	return nil
}
