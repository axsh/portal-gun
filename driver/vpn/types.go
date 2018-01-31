package vpn

import (
	"golang.org/x/net/context"
	"github.com/axsh/vpnhub/model"
)

type VpnDriver interface {
	GenerateConfig(p *model.VpnServer) error
	StartVpn() error
	StopVpn() error
	RemoveVpn() error
}

func NewVpnDriver(ctx context.Context) (VpnDriver, error) {
	return nil, nil
}
