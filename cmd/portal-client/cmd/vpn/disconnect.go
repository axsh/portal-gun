package vpn

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DisconnectVpn = &cobra.Command{
	Use:   "disconnect [subnet] [options]",
	Short: "Completely disconnect a subnet and all its nics from the virtual network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
