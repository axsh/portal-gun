package util

import (
	"github.com/axsh/portal-gun/api"
	"github.com/axsh/portal-gun/cmd/portal-client/cmd"
	"golang.org/x/net/context"
)

func PrepareApiClient() (*api.PortalClient, context.Context) {
	return api.NewPortalClient(
		cmd.Settings.ServerIp,
		cmd.Settings.ServerPort,
		cmd.Settings.Insecure,
		cmd.Settings.CertKey,
		cmd.Settings.AuthToken,
	)
}
