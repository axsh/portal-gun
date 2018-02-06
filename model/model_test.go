package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkerInterface(t *testing.T) {
	assert := assert.New(t)

	assert.Implements((*PortalDriver)(nil), new(VpnDriver_Type))
	assert.Implements((*PortalDriver)(nil), new(NetworkDriver_Type))

	assert.Implements((*VpnParam)(nil), new(SoftEtherParam))
	assert.Implements((*NicParam)(nil), new(OpenVnetParam))
}
