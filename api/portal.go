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
	certFile  string
	authToken string
}

// dummy token
var ctxId = "portalGun.Id.ctx"

func (p *PortalClient) addSecureOpts(opts *[]grpc.DialOption) error {
	if p.insecure {
		*opts = append(*opts, grpc.WithInsecure())
		return nil
	}

	if len(p.certFile) == 0 {
		return errors.Errorf("missing certificate")
	}
	creds, err := credentials.NewClientTLSFromFile(p.certFile, "")
	if err != nil {
		return errors.Wrapf(err, "unable to load cert key: %s", p.certFile)
	}
	*opts = append(*opts, grpc.WithTransportCredentials(creds))
	return nil
}

func (p *PortalClient) connect(ctx context.Context) error {
	var err error

	copts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second * 3),
	}

	if err = p.addSecureOpts(&copts); err != nil {
		return err
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

func NewPortalClient(ip string, port string, insecure bool, cert string, token string) (*PortalClient, context.Context) {
	ctx := context.Background()
	if len(token) > 0 {
		md := metadata.Pairs(
			"auth_token", token,
			"id", ctxId,
		)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	return &PortalClient{
		hostIp:    ip,
		hostPort:  port,
		insecure:  insecure,
		certFile:  cert,
		authToken: token,
	}, ctx
}
