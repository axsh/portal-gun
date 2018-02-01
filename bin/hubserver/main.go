package main

import (
	"net"
	"fmt"
	"os"

	// "golang.org/x/net/context"
	"github.com/axsh/vpnhub/api"
	// "google.golang.org/grpc"

	_ "github.com/axsh/vpnhub/driver/network/openvnet"
	_ "github.com/axsh/vpnhub/driver/vpn/softether"
)

var (
	hubServerIp string = "localhost"
	hubServerPort string = "9000"
)

func main() {
	fmt.Println("Starting vpn api server")
	endpoint := net.JoinHostPort(hubServerIp, hubServerPort)
	l, err := net.Listen("tcp", endpoint)
	if err != nil {
		fmt.Println("failed listen", err)
		os.Exit(-1)
	}

	hubServer := api.NewVpnHubAPIServer()

	errChan := make(chan error)
	go func() {
		if err := hubServer.Serve(l); err != nil {
			fmt.Println("server failed")
			errChan <-err
			return
		}
	}()
	defer hubServer.GracefulStop()

	<-errChan
}
