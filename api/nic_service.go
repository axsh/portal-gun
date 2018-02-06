package api

import (
	"github.com/axsh/portal-gun/driver"
	"github.com/pkg/errors"
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

	if p := nd.GetNicParams(); p != nil {
		d.RegisterNic(p)
	} else {
		return nil, errors.Errorf("Missing nic parameters")
	}

	return &RegisterNicReply{}, nil
}

func (a *NicService) Deregister(ctx context.Context, req *DeregisterNicRequest) (*DeregisterNicReply, error) {
	nd := req.GetNetworkDriver()
	d, err := driver.NewNetworkDriver(ctx, nd.GetDriverType())
	if err != nil {
		return nil, err
	}

	if p := nd.GetNicParams(); p != nil {
		d.DeregisterNic(p)
	} else {
		return nil, errors.Errorf("Missing nic parameters")
	}

	return &DeregisterNicReply{}, nil
}
