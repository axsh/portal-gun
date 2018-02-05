package main

import (
	"net"
	"fmt"
	"os"

	"github.com/axsh/portal-gun/api"
	// "github.com/spf13/pflag"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hubserver",
	Short: "Run the server process",
	Long:  ``,
	Example: `
    options

    '--cert' can be used to specify the location a key for credentials

      $ portal-server --key=/home/.ssh/server-key.pem

    '--insecure' can be set to enable/disable certification requirement.

      $ portal-server --insecure

    '--port' and '--host' can be used to explicitly set where to listen for connections.

      $ portal-server --host=192.168.1.10 --port=11000

	`,
	Run: startServer,
}

var (
	portalHost string = "0.0.0.0"
	portalPort string = "8002"
)

func startServer(cmd *cobra.Command, args []string) {
	endpoint := net.JoinHostPort(portalHost, portalPort)
	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		fmt.Println("failed listen", err)
		os.Exit(-1)
	}

	portal := api.NewPortalAPIServer()
	errChan := make(chan error)
	go func() {
		if err := portal.Serve(l); err != nil {
			fmt.Println("server failed")
			errChan <-err
			return
		}
	}()
	defer portal.GracefulStop()
	<-errChan
}

func main() {
	fmt.Println("Starting vpn api server")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
