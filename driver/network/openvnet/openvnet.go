package openvnet

import (
	"github.com/axsh/vpnhub/driver/network"
	"github.com/axsh/vpnhub/driver"
)

type OpenVNetDriver {

}

func(d *OpenVNetDriver) RegisterNic(i *driver.Nic) (string, error) {
	return "", nil
}

func(d *OpenVNetDriver) DeregisterNic(id string) error {
	return nil
}
