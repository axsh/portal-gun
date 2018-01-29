package openvnet

import (
	"github.com/axsh/vpnhub/driver/network"
	"github.com/axsh/vpnhub/model"
)

type OpenVNetDriver {}

func RegisterInterface(i *model.Interface) (string, error) {
	return "", nil
}

func DeregisterInterface(id string) error {
	return nil
}

func GetInterface(id string) (*model.Interface, error) {
	i := &model.Interface{}

	return i, error
}
