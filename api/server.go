package api

import (
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc"
)

//go:generate protoc -I../proto --go_out=plugins=grpc:${GOPATH}/src ../proto/server.proto

type PortalAPIServer struct {
	listener net.Listener
	server   *grpc.Server
}

type ServerSettings struct {
	Insecure bool
	CertFile string
	KeyFile  string
	Token    string
}

var serverCtx = "portalGun.server.ctx"

func authClient(ctx context.Context, validToken string) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if token := strings.Join(md["auth_token"], ""); token != validToken {
			return "", errors.Errorf("Invalid token, wants %s, got %s", validToken, token)
		}
		return strings.Join(md["client_id"], ""), nil
	}

	return "", errors.Errorf("missing metadata")
}

func NewPortalAPIServer(settings ServerSettings) (*PortalAPIServer, error) {
	sopts := []grpc.ServerOption{}

	if len(settings.Token) > 0 {
		interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			clientId, err := authClient(ctx, settings.Token)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to authenticate client")
			}

			ctx = context.WithValue(ctx, serverCtx, clientId)
			return handler(ctx, req)
		}
		sopts = append(sopts, grpc.UnaryInterceptor(interceptor))
	}

	if !settings.Insecure {
		fmt.Println("adding cert")
		if len(settings.CertFile) == 0 {
			return nil, errors.Errorf("cert file not specified")
		}
		if len(settings.KeyFile) == 0 {
			return nil, errors.Errorf("key file not specified")
		}
		creds, err := credentials.NewServerTLSFromFile(settings.CertFile, settings.KeyFile)
		if err != nil {
			return nil, errors.Wrapf(err, "failed NewServerTLSFromFile()")
		}
		sopts = append(sopts, grpc.Creds(creds))
	}

	s := &PortalAPIServer{server: grpc.NewServer(sopts...)}
	RegisterVpnServiceServer(s.server, &VpnService{api: s})
	RegisterNicServiceServer(s.server, &NicService{api: s})
	return s, nil
}

func (s *PortalAPIServer) Serve(listen net.Listener) error {
	s.listener = listen
	return s.server.Serve(listen)
}

func (s *PortalAPIServer) Stop() {
	s.server.Stop()
	s.listener = nil
}

func (s *PortalAPIServer) GracefulStop() {
	s.server.GracefulStop()
	s.listener = nil
}

func (s *PortalAPIServer) Listener() net.Listener {
	return s.listener
}
