package softether

import (
	"testing"

	"github.com/axsh/portal-gun/driver"
	"github.com/axsh/portal-gun/model"
	"github.com/stretchr/testify/assert"
)

func TestSoftEtherDriver(t *testing.T) {
	assert := assert.New(t)
	d := new(SoftEtherDriver)

	assert.Implements((*driver.Driver)(nil), d)
	assert.Implements((*driver.Vpn)(nil), d)
	assert.Equal(model.VpnDriver_SOFTETHER_VPN, d.DriverType())
}
