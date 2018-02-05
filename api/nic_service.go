package api

import (
	"github.com/axsh/portal-gun/driver"
	"golang.org/x/net/context"
)

type NicService struct {
	api *PortalAPIServer
}

func (a *NicService) Register(ctx context.Context, req *RegisterNicRequest) (*RegisterNicReply, error) {
	nic := req.GetNic()
	d, err := driver.GetNetworkDriver(ctx, nic.GetDriverType())
	if err != nil {
		return nil, err
	}

	d.RegisterNic(nic)
	return &RegisterNicReply{}, nil
}

func (a *NicService) Deregister(ctx context.Context, req *DeregisterNicRequest) (*DeregisterNicReply, error) {
	nic := req.GetNic()
	d, err := driver.GetNetworkDriver(ctx, nic.GetDriverType())
	if err != nil {
		return nil, err
	}

	d.DeregisterNic(nic)
	return &DeregisterNicReply{}, nil
}
