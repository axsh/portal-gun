package vpn

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConnectVpn = &cobra.Command{
	Use:   "connect [subnet] [options]",
	Short: "Create a vpn tunnel and connects all existing nics found in the subnet",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
