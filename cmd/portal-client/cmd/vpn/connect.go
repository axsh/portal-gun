package vpn

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConnectVpn = &cobra.Command{
	Use: "connect [options]",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
