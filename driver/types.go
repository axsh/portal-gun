package driver

import (
	"fmt"
	"strings"

	"github.com/axsh/portal-gun/model"
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

func createDriver(name string) (Driver, error) {
	creator, exists := drivers[name]
	if !exists {
		knownDrivers := make([]string, len(drivers))
		for d, _ := range drivers {
			knownDrivers = append(knownDrivers, d)
		}
		d := strings.Join(knownDrivers, ",")
		return nil, fmt.Errorf("Unknonwn driver, must be one of %s", d)
	}

	return creator()
}

type Vpn interface {
	GenerateConfig(p *model.VpnServer) error
	StartVpn() error
	StopVpn() error
	RemoveVpn() error
}

func GetVpnDriver(ctx context.Context, t model.VpnServer_Type) (Vpn, error) {
	if t == model.VpnServer_NONE {
		return nil, fmt.Errorf("vpn driver not set")
	}

	driver, err := createDriver(t.String())
	if err != nil {
		return nil, err
	}
	return driver.(Vpn), err
}

type Network interface {
	RegisterNic(nic *model.Nic) (string, error)
	DeregisterNic(nic *model.Nic) (string, error)
}

func GetNetworkDriver(ctx context.Context, t model.Nic_Type) (Network, error) {
	if t == model.Nic_NONE {
		return nil, fmt.Errorf("network driver not set")
	}

	driver, err := createDriver(t.String())
	if err != nil {
		return nil, err
	}
	return driver.(Network), err
}
