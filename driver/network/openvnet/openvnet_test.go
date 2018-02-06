package openvnet

import (
	"testing"

	"github.com/axsh/portal-gun/driver"
	"github.com/axsh/portal-gun/model"
	"github.com/stretchr/testify/assert"
)

func TestOpenVnetDriver(t *testing.T) {
	assert := assert.New(t)
	d := new(OpenVnetDriver)

	assert.Implements((*driver.Driver)(nil), d)
	assert.Implements((*driver.Network)(nil), d)
	assert.Equal(model.NetworkDriver_OPENVNET, d.DriverType())
}
