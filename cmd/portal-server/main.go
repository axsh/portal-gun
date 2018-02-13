package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"os"

	"github.com/axsh/portal-gun/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portal-server",
	Short: "Run the server process",
	Long:  ``,
	Run:   startServer,
}

var settings struct {
	insecure   bool
	listenIp   string
	listenPort string
	authToken  string
	certKey    string
	certFile   string
}

func randomizeToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func startServer(cmd *cobra.Command, args []string) {
	endpoint := net.JoinHostPort(settings.listenIp, settings.listenPort)
	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		fmt.Println("failed listen", err)
		os.Exit(-1)
	}

	if settings.authToken == "" {
		token, err := randomizeToken(32)
		if err != nil {
			fmt.Println(err, "failed to generate random token, exiting...")
			os.Exit(-1)
		} else {
			settings.authToken = token
		}

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
			errChan <- err
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
