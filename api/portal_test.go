package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPortal(t *testing.T) {
	assert := assert.New(t)

	p := NewPortal("0.0.0.0", "8002")
	assert.NotNil(p)
	assert.Equal("0.0.0.0", p.hostIp)
	assert.Equal("8002", p.hostPort)
}
