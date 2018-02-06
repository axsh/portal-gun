package vpn

import (
	"time"

	"github.com/axsh/portal-gun/api"
	"github.com/axsh/portal-gun/model"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var CreateVpn = &cobra.Command{
	Use:   "create [options]",
	Short: "Create a new vpn server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
		portal := api.NewPortal("0.0.0.0", "8002")

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
	},
}
