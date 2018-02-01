package api

import (
	// "fmt"

	"golang.org/x/net/context"
	"github.com/axsh/vpnhub/driver"
)

type NicService struct {
	api    *VpnHubAPIServer
}


func(a *NicService) Register(ctx context.Context, req *RegisterNicRequest) (*RegisterNicReply, error) {
	nic := req.GetNic()
	d, err := driver.NewNetworkDriver(ctx, nic.GetDriverType())
	if err != nil {
		return nil, err
	}

	d.RegisterNic(nic)
	return &RegisterNicReply{}, nil
}

func(a *NicService) Deregister(ctx context.Context, req *DeregisterNicRequest) (*DeregisterNicReply, error) {
	nic := req.GetNic()
	d, err := driver.NewNetworkDriver(ctx, nic.GetDriverType())
	if err != nil {
		return nil, err
	}

	d.DeregisterNic(nic)
	return &DeregisterNicReply{}, nil
}
