package main

import (
	"net"
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"
	"github.com/axsh/vpnhub/api"
	"google.golang.org/grpc"
)

var	hubServerIp string
var	hubServerPort string


type VpnHub struct {
	conn      *grpc.ClientConn
	ctx       context.Context

	vpnClient *api.VpnServiceClient
	nicClient *api.NicServiceClient
}

func (c *VpnHub) Connect() error {
	var err error
	copts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	endpoint := net.JoinHostPort(hubServerIp, hubServerPort)
	c.ctx, _ = context.WithTimeout(context.Background(), time.Second*1)
	c.conn, err = grpc.DialContext(c.ctx, endpoint, copts...)
	if err != nil {
		fmt.Println("failed to connect to server at:", endpoint, err)
		return err
	}

	return nil
}

func newVpnHubClient() *VpnHub {
	hubServerIp = "localhost"
	hubServerPort = "9000"

	return &VpnHub{}
}

func main() {
	fmt.Println("Starting vpn client")
	hub := newVpnHub()
	if err := hub.Connect(); err != nil {
		fmt.Println("failed connect")
		os.Exit(-1)
	}
	defer hub.conn.Close()

	// test code
	c := api.NewVpnServiceClient(client.conn)
	c.Create(client.ctx, &api.CreateVpnRequest{})
}
