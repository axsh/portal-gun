package api

import (
	"github.com/axsh/portal-gun/driver"
	"golang.org/x/net/context"
)

type VpnService struct {
	api *PortalAPIServer
}

func (a *VpnService) Create(ctx context.Context, req *CreateVpnRequest) (*CreateVpnReply, error) {
	vpn := req.GetVpnDriver()
	d, err := driver.NewVpnDriver(ctx, vpn.GetDriverType())
	if err != nil {
		return nil, err
	}

	if err := d.GenerateConfig(vpn.GetVpnParams()); err != nil {
		return nil, err
	}
	if err := d.StartVpn(); err != nil {
		return nil, err
	}

	return &CreateVpnReply{}, nil
}

func (a *VpnService) Destroy(ctx context.Context, req *DestroyVpnRequest) (*DestroyVpnReply, error) {
	vpn := req.GetVpnDriver()
	d, err := driver.NewVpnDriver(ctx, vpn.GetDriverType())
	if err != nil {
		return nil, err
	}

	d.StopVpn()
	d.RemoveVpn()

	return &DestroyVpnReply{}, nil
}
