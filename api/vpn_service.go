package api

import (
	"golang.org/x/net/context"
	"github.com/axsh/vpnhub/driver"
)

type VpnService struct {
	api    *VpnHubAPIServer
}

func(a *VpnService) Create(ctx context.Context, req *CreateVpnRequest) (*CreateVpnReply, error) {
	vpn := req.GetVpnServer()
	d, err := driver.NewVpnDriver(ctx, vpn.GetDriverType())
	if err != nil {
		return nil, err
	}

	if err := d.GenerateConfig(vpn); err != nil {
		return nil, err
	}
	if err := d.StartVpn(); err != nil {
		return nil, err
	}
	return &CreateVpnReply{}, nil
}

func(a *VpnService) Destroy(ctx context.Context, req *DestroyVpnRequest) (*DestroyVpnReply, error) {
	vpn := req.GetVpnServer()
	d, err := driver.NewVpnDriver(ctx, vpn.GetDriverType())
	if err != nil {
		return nil, err
	}

	d.StopVpn()
	d.RemoveVpn()
	return &DestroyVpnReply{}, nil
}
