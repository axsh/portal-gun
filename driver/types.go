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

func Register(d model.PortalDriver, creator DriverCreator) {
	drivers[d.String()] = creator
	fmt.Println("register", d.String(), creator)
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
	GenerateConfig(p model.VpnParam) error
	StartVpn() error
	StopVpn() error
	RemoveVpn() error
}

func NewVpnDriver(ctx context.Context, t model.VpnDriver_Type) (Vpn, error) {
	if t == model.VpnDriver_NONE {
		return nil, fmt.Errorf("vpn driver not set")
	}

	driver, err := createDriver(t.String())
	if err != nil {
		return nil, err
	}

	return driver.(Vpn), err
}

type Network interface {
	RegisterNic(nic model.NicParam) (string, error)
	DeregisterNic(nic model.NicParam) (string, error)
}

// TODO: rename
func NewNetworkDriver(ctx context.Context, t model.NetworkDriver_Type) (Network, error) {
	if t == model.NetworkDriver_NONE {
		return nil, fmt.Errorf("network driver not set")
	}

	driver, err := createDriver(t.String())
	if err != nil {
		return nil, err
	}

	return driver.(Network), err
}
