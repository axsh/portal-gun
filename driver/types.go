package driver
import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/axsh/portal-gun/model"
	"golang.org/x/net/context"
)

type Driver interface {
	IsDriver()
	DriverType() model.PortalDriver
}

type DriverCreator func() (Driver, error)

var drivers map[string]DriverCreator = make(map[string]DriverCreator)

func Register(d model.PortalDriver, creator DriverCreator) {
	if _, exists := drivers[d.String()]; exists {
		fmt.Print("Driver %s already regitered, skipping", d.String())
		return
	}

	drivers[d.String()] = creator
	fmt.Println("Registered driver:", d.String())
}

func createDriver(name string) (Driver, error) {
	creator, exists := drivers[name]
	if !exists {
		knownDrivers := make([]string, len(drivers))
		for d, _ := range drivers {
			knownDrivers = append(knownDrivers, d)
		}
		d := strings.Join(knownDrivers, ",")
		return nil, errors.Errorf("Unknonwn driver, must be one of %s", d)
	}

	return creator()
}

type Vpn interface {
	GenerateConfig(p model.VpnParam) error
	StartVpn() error
	StopVpn() error
	RemoveVpn() error
}

func NewVpnDriver(ctx context.Context, t model.VpnDriver_Type) (Vpn, error) {
	if t == model.VpnDriver_NONE {
		return nil, errors.Errorf("vpn driver not set")
	}

	driver, err := createDriver(t.String())
	if err != nil {
		return nil, errors.Errorf("Unable to create the vpn driver", err)
	}

	return driver.(Vpn), nil
}

type Network interface {
	RegisterNic(nic model.NicParam) (string, error)
	DeregisterNic(nic model.NicParam) (string, error)
}

func NewNetworkDriver(ctx context.Context, t model.NetworkDriver_Type) (Network, error) {
	if t == model.NetworkDriver_NONE {
		return nil, errors.Errorf("network driver not set")
	}

	driver, err := createDriver(t.String())
	if err != nil {
		return nil, errors.Errorf("Unable to create the network driver", err)
	}

	return driver.(Network), nil
}
