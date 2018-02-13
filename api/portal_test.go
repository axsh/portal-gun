package api

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestNewPortalClient(t *testing.T) {
	assert := assert.New(t)

	p, ctx := NewPortalClient("0.0.0.0", "8002", true, "cert-file", "token",)

	assert.NotNil(p)
	assert.Equal("0.0.0.0", p.hostIp)
	assert.Equal("8002", p.hostPort)
	assert.Equal(true, p.insecure)
	assert.Equal("cert-file", p.certFile)

	md, ok := metadata.FromOutgoingContext(ctx)
	assert.True(ok)
	assert.Equal("token", strings.Join(md["auth_token"], ""))
}
