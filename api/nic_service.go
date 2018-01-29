package api

import (
	"golang.org/x/net/context"
)

type NicService struct {
	api          *VpnHubAPIServer
}

func(a *NicService) Register(ctx context.Context, in *RegisterNicRequest) (*RegisterNicReply, error) {

	return &RegisterNicReply{}, nil
}

func(a *NicService) Deregister(ctx context.Context, in *DeregisterNicRequest) (*DeregisterNicReply, error) {

	return &DeregisterNicReply{}, nil
}
