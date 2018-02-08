package network

import (
	"github.com/axsh/portal-gun/cmd/portal-client/cmd"
	"github.com/axsh/portal-gun/cmd/portal-client/cmd/util"
	"github.com/axsh/portal-gun/api"
	"github.com/axsh/portal-gun/model"
	"github.com/spf13/cobra"
)

var registerNic = &cobra.Command{
	Use: "register [options]",
	Short: "Register ip/mac leases to be part of a virtual network",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		portal, ctx := util.PrepareApiClient()

		return portal.NicServiceRequest(ctx, func(c api.NicServiceClient) error {
			_, e := c.Register(ctx, &api.RegisterNicRequest{
				&model.NetworkDriver{
					DriverType: model.NetworkDriver_OPENVNET,
					InterfaceParams: &model.NetworkDriver_OpenvnetParams{
						OpenvnetParams: &model.OpenVnetParam{
							IpLease: &model.OpenVnetParam_IpLease{
								Ipv4Address: "10.0.100.10",
							},
						},
					},
				},
			})
			return e
		})
	},
}

func init() {
	cmd.AddSubCommand(cmd.NetworkCmd, registerNic)
}
