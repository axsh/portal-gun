package network

import (
	"fmt"
	"time"

	"github.com/axsh/portal-gun/api"
	"github.com/axsh/portal-gun/model"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var RegisterNic = &cobra.Command{
	Use: "register [options]",

	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)

		ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
		portal := api.NewPortal("0.0.0.0", "8002")

		portal.NicServiceRequest(ctx, func(c api.NicServiceClient) error {
			c.Register(ctx, &api.RegisterNicRequest{
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
			return nil
		})

	},
}
