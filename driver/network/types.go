package network

import (
	"golang.org/x/net/context"
	"github.com/axsh/vpnhub/model"
)

type NetworkDriver interface {
	RegisterNic(i *model.Nic) (string, error)
	DeregisterNic(i *model.Nic) (string, error)
}

func NewNetworkDriver(ctx context.Context) (NetworkDriver, error) {
	return nil, nil
}
