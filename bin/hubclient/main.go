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

	vpnClient *api.VpnServiceClient
	nicClient *api.NicServiceClient
}

func (c *VpnHub) Connect(ctx context.Context) error {
	var err error
	copts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	endpoint := net.JoinHostPort(hubServerIp, hubServerPort)
	c.conn, err = grpc.DialContext(ctx, endpoint, copts...)
	if err != nil {
		fmt.Println("failed to connect to server at:", endpoint, err)
		return err
	}

	return nil
}

func newVpnHub() *VpnHub {
	hubServerIp = "localhost"
	hubServerPort = "9000"

	return &VpnHub{}
}

func main() {
	fmt.Println("Starting vpn client")
	hub := newVpnHub()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	if err := hub.Connect(ctx); err != nil {
		fmt.Println("failed connect")
		os.Exit(-1)
	}
	defer hub.conn.Close()

	// test code
	context.WithValue(ctx, "vpn", "softether")

	c := api.NewVpnServiceClient(hub.conn)
	c.Create(ctx, &api.CreateVpnRequest{})
}
