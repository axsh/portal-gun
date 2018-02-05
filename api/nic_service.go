package api

import (
	"github.com/axsh/portal-gun/driver"
	"golang.org/x/net/context"
)

type NicService struct {
	api *PortalAPIServer
}

func (a *NicService) Register(ctx context.Context, req *RegisterNicRequest) (*RegisterNicReply, error) {
	nd := req.GetNetworkDriver()
	d, err := driver.NewNetworkDriver(ctx, nd.GetDriverType())
	if err != nil {
		return nil, err
	}

	d.RegisterNic(nd.GetNicParams())

	return &RegisterNicReply{}, nil
}

func (a *NicService) Deregister(ctx context.Context, req *DeregisterNicRequest) (*DeregisterNicReply, error) {
	nd := req.GetNetworkDriver()
	d, err := driver.NewNetworkDriver(ctx, nd.GetDriverType())
	if err != nil {
		return nil, err
	}

	d.DeregisterNic(nd.GetNicParams())

	return &DeregisterNicReply{}, nil
}
