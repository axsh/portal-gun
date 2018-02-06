package driver

import (
	"testing"

	"github.com/axsh/portal-gun/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

type testvpn struct{}

func (*testvpn) IsDriver()                             {}
func (*testvpn) DriverType() model.PortalDriver        { return model.VpnDriver_NULL_VPN }
func (*testvpn) GenerateConfig(p model.VpnParam) error { return nil }
func (*testvpn) StartVpn() error                       { return nil }
func (*testvpn) StopVpn() error                        { return nil }
func (*testvpn) RemoveVpn() error                      { return nil }

func TestNewVpnDriver(t *testing.T) {
	assert := assert.New(t)

	Register(model.VpnDriver_NULL_VPN, func() (Driver, error) {
		return &testvpn{}, nil
	})

	ctx := context.Background()
	d, e := NewVpnDriver(ctx, model.VpnDriver_NULL_VPN)

	assert.NoError(e)
	assert.NotNil(d)
	assert.Equal(d.(Driver).DriverType(), model.VpnDriver_NULL_VPN)
}

type testnetwork struct{}

func (*testnetwork) IsDriver()                                      {}
func (*testnetwork) DriverType() model.PortalDriver                 { return model.NetworkDriver_NULL_NET }
func (*testnetwork) RegisterNic(p model.NicParam) (string, error)   { return "", nil }
func (*testnetwork) DeregisterNic(p model.NicParam) (string, error) { return "", nil }

func TestNewNetworkDriver(t *testing.T) {
	assert := assert.New(t)

	Register(model.NetworkDriver_NULL_NET, func() (Driver, error) {
		return &testnetwork{}, nil
	})

	ctx := context.Background()
	d, e := NewNetworkDriver(ctx, model.NetworkDriver_NULL_NET)

	assert.NoError(e)
	assert.NotNil(d)
	assert.Equal(d.(Driver).DriverType(), model.NetworkDriver_NULL_NET)
}
