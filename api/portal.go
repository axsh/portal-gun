package api

import (
	"net"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type PortalClient struct {
	hostIp    string
	hostPort  string
	conn      *grpc.ClientConn
	insecure  bool
	certKey   string
	authToken string
}

func (p *PortalClient) addClientOpts() ([]grpc.DialOption, error) {
	opts := []grpc.DialOption{}

	if !p.insecure {
		opts = append(opts, grpc.WithInsecure())
		return opts, nil
	} else {
		if len(p.certKey) == 0 {
			return nil, errors.Errorf("missing certificate")
		}
		creds, err := credentials.NewClientTLSFromFile(p.certKey, "")
		if err != nil {
			return nil, errors.Wrapf(err, "unable to load cert key: %s", p.certKey)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	return opts, nil
}

func (p *PortalClient) connect(ctx context.Context) error {
	var err error
	copts, err := p.addClientOpts()
	if err != nil {
		return err
	}

	if len(p.authToken) > 0 {
		md := metadata.Pairs("authToken", p.authToken)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	endpoint := net.JoinHostPort(p.hostIp, p.hostPort)
	p.conn, err = grpc.DialContext(ctx, endpoint, copts...)
	if err != nil {
		errors.Wrapf(err, "failed to connect to server at: %s", endpoint)
		return err
	}

	return nil
}

func (p *PortalClient) VpnServiceRequest(ctx context.Context, req func(c VpnServiceClient) error) error {
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

func (p *PortalClient) NicServiceRequest(ctx context.Context, req func(c NicServiceClient) error) error {
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

func NewPortalClient(host string, port string) (*PortalClient, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	return &PortalClient{
		hostIp:    host,
		hostPort:  port,
		// insecure:  insecure,
		// authToken: token,
		// certKey:   key,
	}, ctx
}
