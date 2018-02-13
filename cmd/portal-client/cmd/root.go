package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"


var RootCmd = &cobra.Command{
	Use:   "portal-client [sub]",
	Short: "",
}

var NetworkCmd = &cobra.Command{
	Use:   "network [command]",
	Short: "Manage nics on virtual network",
}

var VpnCmd = &cobra.Command{
	Use:   "vpn [command]",
	Short: "Manage vpn service",
}

var Settings struct {
	Insecure   bool
	ServerIp   string
	CertKey    string
	AuthToken  string
	ServerPort string
}

func AddSubCommand(parent *cobra.Command, child *cobra.Command) {
	parent.AddCommand(child)
}

func Execute() {
	RootCmd.AddCommand(NetworkCmd)
	RootCmd.AddCommand(VpnCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&Settings.ServerIp, "server", "s", "0.0.0.0", "host ip of portal gun server")
	RootCmd.PersistentFlags().StringVarP(&Settings.ServerPort, "port", "p", "8002", "listening port of portal gun server")
	RootCmd.PersistentFlags().BoolVarP(&Settings.Insecure, "insecure", "", true, "disable certification")
	RootCmd.PersistentFlags().StringVarP(&Settings.CertKey, "cert", "c", "", "certificate file")
	RootCmd.PersistentFlags().StringVarP(&Settings.AuthToken, "auth-token", "t", "", "authorization token")
}
