package api

import (
	// "fmt"

	"golang.org/x/net/context"
	"github.com/axsh/vpnhub/driver/network"
)

type NicService struct {
	api    *VpnHubAPIServer
}


func(a *NicService) Register(ctx context.Context, req *RegisterNicRequest) (*RegisterNicReply, error) {
	d, err := network.NewNetworkDriver(ctx)
	if err != nil {
		return nil, err
	}

	d.RegisterNic(req.GetNic())
	return &RegisterNicReply{}, nil
}

func(a *NicService) Deregister(ctx context.Context, req *DeregisterNicRequest) (*DeregisterNicReply, error) {
	d, err := network.NewNetworkDriver(ctx)
	if err != nil {
		return nil, err
	}

	d.RegisterNic(req.GetNic())
	return &DeregisterNicReply{}, nil
}
