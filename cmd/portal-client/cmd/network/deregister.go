package network

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeregisterNic = &cobra.Command{
	Use: "deregister [options]",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
