package api

import (
	"net"

	// "golang.org/x/net/context"
	"google.golang.org/grpc"
)

//go:generate protoc -I../proto --go_out=plugins=grpc:${GOPATH}/src ../proto/server.proto

type PortalAPIServer struct {
	listener net.Listener
	server   *grpc.Server
}

func NewPortalAPIServer() *PortalAPIServer {
	s := &PortalAPIServer{
		server: grpc.NewServer(),
	}

	RegisterVpnServiceServer(s.server, &VpnService{api: s})
	RegisterNicServiceServer(s.server, &NicService{api: s})
	return s
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
