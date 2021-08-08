package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"layout/pkg/transport"

	"github.com/gin-gonic/gin"
)

var _ transport.Server = (*Server)(nil)

// RequestFunc request function to deal with request.
type RequestFunc func(context.Context, *http.Request) context.Context

// ResponseFunc to deal with response.
type ResponseFunc func(context.Context, http.ResponseWriter) context.Context

// Server is http server.
type Server struct {
	*gin.Engine
	server  *http.Server
	addr    string
	port    int
	timeout time.Duration
	before  []RequestFunc
	after   []ResponseFunc
}

// Option is a function to add option to the server.
type Option func(*Server)

// ServerBefore add RequestFunc to do something befre the http handler.
func ServerBefore(before ...RequestFunc) Option {
	return func(s *Server) {
		s.before = append(s.before, before...)
	}
}

// ServerAfter add RequestFunc to do something after the http handler.
func ServerAfter(after ...ResponseFunc) Option {
	return func(s *Server) {
		s.after = append(s.after, after...)
	}
}

// Address set the address to the server.
func Address(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// Port set the port the server.
func Port(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// Timeout set the address to the server.
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// New create http server.
func New(mode string, options ...Option) *Server {

	gin.SetMode(mode)

	srv := &Server{
		Engine: gin.New(),
	}
	for _, option := range options {
		option(srv)
	}

	srv.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", srv.addr, srv.port),
		Handler: srv,
	}

	srv.Use(gin.Logger())
	srv.Use(gin.Recovery())

	return srv
}

// Start start http server.
func (s *Server) Start(ctx context.Context) error {
	return s.server.ListenAndServe()
}

// Stop stop http server.
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
