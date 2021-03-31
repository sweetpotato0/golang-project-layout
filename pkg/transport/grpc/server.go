package grpc

import (
	"context"

	"layout/pkg/transport"
)

var _ transport.Server = (*Server)(nil)

type Server struct {
}

type Option func(*Server)

func New() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
