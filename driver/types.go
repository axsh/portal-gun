package driver

import (
	"fmt"

	"github.com/axsh/vpnhub/model"
	"golang.org/x/net/context"
)

type Driver interface {
	IsDriver()
}

type DriverCreator func() (Driver, error)

var drivers map[string]DriverCreator = make(map[string]DriverCreator)

func Register(name string, creator DriverCreator) {
	drivers[name] = creator
	fmt.Println("register", name, creator)
}

func GetCreator(name string) DriverCreator {
	return nil
}

type Vpn interface {
	GenerateConfig(p *model.VpnServer) error
	StartVpn() error
	StopVpn() error
	RemoveVpn() error
}

func NewVpnDriver(ctx context.Context, t model.VpnServer_Type) (Vpn, error) {
	if t == model.VpnServer_NONE {
		return nil, fmt.Errorf("vpn driver not set")
	}

	d, err := drivers[t.String()]()
	if err != nil {
		return nil, err
	}
	return d.(Vpn), nil
}

type Network interface {
	RegisterNic(nic *model.Nic) (string, error)
	DeregisterNic(nic *model.Nic) (string, error)
}

func NewNetworkDriver(ctx context.Context, t model.Nic_Type) (Network, error) {
	if t == model.Nic_NONE {
		return nil, fmt.Errorf("network driver not set")
	}

	d, err := drivers[t.String()]()
	if err != nil {
		return nil, err
	}
	return d.(Network), nil
}
