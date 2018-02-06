package vpn

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DestroyVpn = &cobra.Command{
	Use:   "destroy [vpnid] [options]",
	Short: "Destroy the vpn service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
