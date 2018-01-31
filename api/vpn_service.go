package api

import (
	"golang.org/x/net/context"
	"github.com/axsh/vpnhub/driver/vpn"
)

type VpnService struct {
	api    *VpnHubAPIServer
}

func(a *VpnService) Create(ctx context.Context, req *CreateVpnRequest) (*CreateVpnReply, error) {
	d, err := vpn.NewVpnDriver(ctx)
	if err != nil {
		return nil, err
	}

	if err := d.GenerateConfig(req.GetVpnServer()); err != nil {
		return nil, err
	}
	if err := d.StartVpn(); err != nil {
		return nil, err
	}
	return &CreateVpnReply{}, nil
}

func(a *VpnService) Destroy(ctx context.Context, req *DestroyVpnRequest) (*DestroyVpnReply, error) {
	d, err := vpn.NewVpnDriver(ctx)
	if err != nil {
		return nil, err
	}

	d.StopVpn()
	d.RemoveVpn()
	return &DestroyVpnReply{}, nil
}
