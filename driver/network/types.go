package network

import (
	"golang.org/x/net/context"
	"github.com/axsh/vpnhub/driver"
)

type NetworkDriver interface {
	RegisterNic(i *driver.Nic) (string, error)
	DeregisterNic(i *driver.Nic) (string, error)
}

func NewNetworkDriver(ctx context.Context) (NetworkDriver, error) {
	return nil, nil
}
