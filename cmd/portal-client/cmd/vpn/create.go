package vpn

import (
	"github.com/axsh/portal-gun/api"
	"github.com/axsh/portal-gun/model"
	"github.com/axsh/portal-gun/cmd/portal-client/cmd"
	"github.com/axsh/portal-gun/cmd/portal-client/cmd/util"
	"github.com/spf13/cobra"
)

var createVpn = &cobra.Command{
	Use:   "create [options]",
	Short: "Create a new vpn server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		portal, ctx := util.PrepareApiClient()
		return portal.VpnServiceRequest(ctx, func(c api.VpnServiceClient) error {
			_, e := c.Create(ctx, &api.CreateVpnRequest{
				&model.VpnDriver{
					DriverType: model.VpnDriver_SOFTETHER_VPN,
					ServerParams: &model.VpnDriver_SoftetherParams{
						SoftetherParams: &model.SoftEtherParam{},
					},
				},
			})
			return e
		})
		return nil
	},
}

func init() {
	cmd.AddSubCommand(cmd.VpnCmd, createVpn)
}
