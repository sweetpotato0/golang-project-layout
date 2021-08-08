package grpc

import (
	"context"
	"net"
	"time"

	"layout/pkg/transport"

	"google.golang.org/grpc"
)

var _ transport.Server = (*Server)(nil)

// Server is a grpc server.
type Server struct {
	*grpc.Server
	ctx     context.Context
	lis     net.Listener
	network string
	address string
	port    int
	timeout time.Duration
}

// Option is to add option to the grpc server.
type Option func(*Server)

// Address set the address to the grpc server, address is ip:port.
func Address(add string) Option {
	return func(s *Server) {
		s.address = add
	}
}

// Timeout set the address to the grpc server.
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// Network set the network to the grpc server.
func Network(network string) Option {
	return func(s *Server) {
		s.network = network
	}
}

// New create t grpc server.
func New(opts ...Option) *Server {

	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
	}

	for _, o := range opts {
		o(srv)
	}

	srv.Server = grpc.NewServer()

	return srv
}

func (s *Server) endpoint() error {

	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}

	s.lis = lis

	return nil
}

// Start start grpc server.
func (s *Server) Start(ctx context.Context) error {
	if err := s.endpoint(); err != nil {
		return err
	}
	s.ctx = ctx
	return s.Serve(s.lis)
}

// Stop stop grpc server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}
