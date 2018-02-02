package main

import (
	"fmt"
	"net"
	"time"

	"github.com/axsh/vpnhub/api"
	"github.com/axsh/vpnhub/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


type VpnHub struct {
	serverIp   string
	serverPort string
	conn      *grpc.ClientConn
}

func (h *VpnHub) connect(ctx context.Context) error {
	var err error
	copts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	endpoint := net.JoinHostPort(h.serverIp, h.serverPort)
	h.conn, err = grpc.DialContext(ctx, endpoint, copts...)
	if err != nil {
		fmt.Println("failed to connect to server at:", endpoint, err)
		return err
	}
	return nil
}

func (h *VpnHub) VpnServiceCall(ctx context.Context, req func(c api.VpnServiceClient) error) int {
	rc := -1

	if err := h.connect(ctx); err != nil {
		fmt.Println("failed connect")
		return rc
	}
	vpnClient := api.NewVpnServiceClient(h.conn)
	if err := req(vpnClient); err == nil {
		rc = 0
	}
	defer h.conn.Close()
	return rc
}

func (h *VpnHub) NicServiceCall(ctx context.Context, req func(c api.NicServiceClient) error) int {
	rc := -1
	if err := h.connect(ctx); err != nil {
		fmt.Println("failed connect")
		return rc
	}

	nicClient := api.NewNicServiceClient(h.conn)
	if err := req(nicClient); err == nil {
		rc = 0
	}
	defer h.conn.Close()
	return rc
}

func NewVpnHub(serverIp string, serverPort string) *VpnHub {
	return &VpnHub{
		serverIp: serverIp,
		serverPort: serverPort,
	}
}

func main() {
	hub := NewVpnHub("localhost", "8002")
	fmt.Println("Starting vpn client")
}
