package softether

import (
	"fmt"
	"github.com/axsh/vpnhub"
)

type SoftEtherDriver struct {}

func (d *SoftEtherDriver) GenerateConfig(p *model.VpnParams) error {
	params := p.(model.SoftEtherParams)
	fmt.Println(p)
}

func (d *SoftEtherDriver) CreateVpn() error {
	return nil
}

func (d *SoftEtherDriver) DestroyVpn() error {
	return nil
}
