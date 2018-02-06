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
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
		portal := api.NewPortal("0.0.0.0", "8002")

		portal.VpnServiceRequest(ctx, func(c api.VpnServiceClient) error {
			c.Create(ctx, &api.CreateVpnRequest{
				&model.VpnDriver{
					DriverType: model.VpnDriver_SOFTETHER_VPN,
					ServerParams: &model.VpnDriver_SoftetherParams{
						SoftetherParams: &model.SoftEtherParam{},
					},
				},
			})
			return nil
		})
	},
}
