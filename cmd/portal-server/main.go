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
	Use:   "portal-server",
	Short: "Run the server process",
	Long:  ``,
	Run: startServer,
}

var settings struct {
	insecure   bool
	listenIp   string
	listenPort string
	authToken  string
	certKey    string
	certFile   string
}


func startServer(cmd *cobra.Command, args []string) {
	endpoint := net.JoinHostPort(settings.listenIp, settings.listenPort)
	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		fmt.Println("failed listen", err)
		os.Exit(-1)
	}

	portalServer, err := api.NewPortalAPIServer(api.ServerSettings{
		Insecure: settings.insecure,
		CertFile: settings.certFile,
		KeyFile:  settings.certKey,
		Token:    settings.authToken,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	errChan := make(chan error)
	go func() {
		if err := portalServer.Serve(l); err != nil {
			fmt.Println("server failed")
			errChan <-err
			return
		}
	}()
	defer portalServer.GracefulStop()
	<-errChan
}

func init() {
	rootCmd.Flags().StringVarP(&settings.listenIp, "server", "s", "0.0.0.0", "listening ip for server")
	rootCmd.Flags().StringVarP(&settings.listenPort, "port", "p", "8002", "listening port for server")
	rootCmd.Flags().StringVarP(&settings.authToken, "auth-token", "t", "", "authorization token, leave empty for random")
	// not implemented yet
	rootCmd.Flags().BoolVarP(&settings.insecure, "insecure", "", true, "disable certification")
}

func main() {
	fmt.Println("Starting vpn api server")
	fmt.Printf("Listening on %s:%s\n", settings.listenIp, settings.listenPort)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
