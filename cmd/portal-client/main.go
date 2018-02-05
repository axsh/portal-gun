package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/axsh/portal-gun/api"
	"github.com/axsh/portal-gun/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


type Portal struct {
	hostIp   string
	hostPort string
	conn      *grpc.ClientConn
}

func (p *Portal) connect(ctx context.Context) error {
	var err error
	copts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	endpoint := net.JoinHostPort(p.hostIp, p.hostPort)
	p.conn, err = grpc.DialContext(ctx, endpoint, copts...)
	if err != nil {
		fmt.Println("failed to connect to server at:", endpoint, err)
		return err
	}

	return nil
}

func (p *Portal) VpnServiceRequest(ctx context.Context, req func(c api.VpnServiceClient) error) int {
	rc := -1

	if err := p.connect(ctx); err != nil {
		fmt.Println("failed connect")
		return rc
	}
	vpnClient := api.NewVpnServiceClient(p.conn)
	if err := req(vpnClient); err == nil {
		rc = 0
	}

	defer p.conn.Close()
	return rc
}

func (p *Portal) NicServiceRequest(ctx context.Context, req func(c api.NicServiceClient) error) int {
	rc := -1
	if err := p.connect(ctx); err != nil {
		fmt.Println("failed connect")
		return rc
	}

	nicClient := api.NewNicServiceClient(p.conn)
	if err := req(nicClient); err == nil {
		rc = 0
	}

	defer p.conn.Close()
	return rc
}

func NewPortal(hostIp string, hostPort string) *Portal {
	return &Portal{
		hostIp: hostIp,
		hostPort: hostPort,
	}
}

func main() {
	if portal := NewPortal("localhost", "8002"); portal == nil {
		os.Exit(-1)
	}

	fmt.Println("Starting vpn client")
}
