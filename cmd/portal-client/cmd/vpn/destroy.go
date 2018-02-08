package vpn

import (
	"fmt"

	"github.com/axsh/portal-gun/cmd/portal-client/cmd"
	"github.com/spf13/cobra"
)

var destroyVpn = &cobra.Command{
	Use:   "destroy [vpnid] [options]",
	Short: "Destroy the vpn service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	cmd.AddSubCommand(cmd.VpnCmd, destroyVpn)
}
