package cmd

import (
	"fmt"
	"os"

	"github.com/axsh/portal-gun/cmd/portal-client/cmd/network"
	"github.com/axsh/portal-gun/cmd/portal-client/cmd/vpn"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portal-client [sub]",
	Short: "",
}

var networkCmd = &cobra.Command{
	Use:   "network [command]",
	Short: "Manage nics on virtual network",
}

var vpnCmd = &cobra.Command{
	Use:   "vpn [command]",
	Short: "Manage vpn service",
}

func Execute() {
	rootCmd.AddCommand(networkCmd)
	rootCmd.AddCommand(vpnCmd)

	networkCmd.AddCommand(network.DeregisterNic)
	networkCmd.AddCommand(network.RegisterNic)

	vpnCmd.AddCommand(vpn.CreateVpn)
	vpnCmd.AddCommand(vpn.DestroyVpn)
	vpnCmd.AddCommand(vpn.ConnectVpn)
	vpnCmd.AddCommand(vpn.DisconnectVpn)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
