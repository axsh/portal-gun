package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/pkg/errors"
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
		errors.Wrapf(err, "failed to connect to server at: %s", endpoint)
		return err
	}

	return nil
}

func (p *Portal) VpnServiceRequest(ctx context.Context, req func(c api.VpnServiceClient) error) int {
	rc := -1

	if err := p.connect(ctx); err != nil {
		fmt.Println("failed connection to vpn service")
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
		fmt.Println("failed connection to network service")
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
	portal := NewPortal("localhost", "8002");
	if portal == nil {
		os.Exit(-1)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)

	// debug code
	portal.NicServiceRequest(ctx, func(c api.NicServiceClient) error {
		c.Register(ctx, &api.RegisterNicRequest{
			&model.NetworkDriver{
				DriverType: model.NetworkDriver_OPENVNET,
				InterfaceParams: &model.NetworkDriver_OpenvnetParams{
					OpenvnetParams: &model.OpenVnetParam{
						IpLease: &model.OpenVnetParam_IpLease{
							Ipv4Address: "10.0.100.10",
						},
					},
				},
			},
		})
		return nil
	})

	portal.VpnServiceRequest(ctx, func(c api.VpnServiceClient) error {
		c.Create(ctx, &api.CreateVpnRequest{
			&model.VpnDriver{
				DriverType: model.VpnDriver_SOFTETHER_VPN,
				ServerParams: &model.VpnDriver_SoftetherParams{
					SoftetherParams: &model.SoftEtherParam{},
				},
			},
		})
		return nil
	})


	fmt.Println("Starting vpn client")
}
