package api

import (
	"net"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Portal struct {
	hostIp   string
	hostPort string
	conn     *grpc.ClientConn
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

func (p *Portal) VpnServiceRequest(ctx context.Context, req func(c VpnServiceClient) error) error {
	if err := p.connect(ctx); err != nil {
		return err
	}

	vpnClient := NewVpnServiceClient(p.conn)
	if err := req(vpnClient); err == nil {
		return err
	}

	defer p.conn.Close()
	return nil
}

func (p *Portal) NicServiceRequest(ctx context.Context, req func(c NicServiceClient) error) error {
	if err := p.connect(ctx); err != nil {
		return err
	}

	nicClient := NewNicServiceClient(p.conn)
	if err := req(nicClient); err == nil {
		return err
	}

	defer p.conn.Close()
	return nil
}

func NewPortal(hostIp string, hostPort string) *Portal {
	return &Portal{
		hostIp:   hostIp,
		hostPort: hostPort,
	}
}
