package vpn

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DisconnectVpn = &cobra.Command{
	Use:   "disconnect [options]",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
