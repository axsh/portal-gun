package api

import (
	"golang.org/x/net/context"
)

type VpnService struct {
	api          *VpnHubAPIServer
}

func(a *VpnService) Create(ctx context.Context, in *CreateVpnRequest) (*CreateVpnReply, error) {
	return &CreateVpnReply{}, nil
}

func(a *VpnService) Destroy(ctx context.Context, in *DestroyVpnRequest) (*DestroyVpnReply, error) {

	return &DestroyVpnReply{}, nil
}
