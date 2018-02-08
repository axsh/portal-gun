package network

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/axsh/portal-gun/cmd/portal-client/cmd"
)

var deregisterNic = &cobra.Command{
	Use:   "deregister [options]",
	Short: "Remove ip/mac lease from virtual network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	cmd.AddSubCommand(cmd.NetworkCmd, deregisterNic)
}
