package vpn

import (
	"github.com/axsh/vpnhub/model"
)

type VpnDriver interface {
	LoadConfig(p model.VpnParams) error
	StartVpn() error
	StopVpn() error
}
