package api

import (
	"net"

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

func NewPortalAPIServer(settings ServerSettings) (*PortalAPIServer, error) {
	sopts := []grpc.ServerOption{}

	if len(settings.Token) > 0 {
		// todo: fix this, its just for trying metadata/unary interceptor
		interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, errors.Errorf("missing context metadata")
			}
			if len(md["authToken"]) != 1 || md["authToken"][0] != settings.Token {
				return nil, errors.Errorf("invalid token")
			}

			return handler(ctx, req)
		}
		sopts = append(sopts, grpc.UnaryInterceptor(interceptor))
	}

	if !settings.Insecure {
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
