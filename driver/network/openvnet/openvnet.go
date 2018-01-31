package openvnet

import (
	"github.com/axsh/vpnhub/driver/network"
	"github.com/axsh/vpnhub/model"
)

type OpenVNetDriver {

}

func(d *OpenVNetDriver) RegisterNic(i *model.Nic) (string, error) {
	return "", nil
}

func(d *OpenVNetDriver) DeregisterNic(id string) error {
	return nil
}

func(d *OpenVNetDriver) GetNic(id string) (*model.Nic, error) {
	i := &model.Nic{}

	return i, error
}

func(d *OpenVNetDriver) GetDriver() {

}
