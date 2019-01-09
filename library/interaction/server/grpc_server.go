package server

import (
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGrpcServer(server *grpc.Server, listener net.Listener) InteractionServerBehavior {
	return &GrpcServer{server: server, listener: listener}
}

func (s *GrpcServer) Serve() error {
	return s.server.Serve(s.listener)
}

func (s *GrpcServer) Stop() {
	s.server.GracefulStop()
}
