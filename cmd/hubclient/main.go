package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/axsh/vpnhub/api"
	"github.com/axsh/vpnhub/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var hubServerIp string
var hubServerPort string

type VpnHub struct {
	conn *grpc.ClientConn
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

	// context.WithValue(ctx, "vpn", "softether")

	// debug code vpn
	vpnClient := api.NewVpnServiceClient(hub.conn)
	vpnClient.Create(ctx, &api.CreateVpnRequest{
		&model.VpnServer{
			DriverType: model.VpnServer_SOFTETHER_VPN,
		},
	})

	// debug code network
	nicClient := api.NewNicServiceClient(hub.conn)
	nicClient.Register(ctx, &api.RegisterNicRequest{
		&model.Nic{
			DriverType: model.Nic_OPENVNET,
		},
	})

}
