package network

import (
	"github.com/axsh/vpnhub/model"
)

type NetworkDriver interface {
	RegisterInterface(i *model.Interface) (string, error)
	DeregisterInterface(id string) (string, error)
	GetInterface(id string) (*model.Interface, error)
}
