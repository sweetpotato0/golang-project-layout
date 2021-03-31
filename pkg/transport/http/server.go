package http

import (
	"context"
	"fmt"
	"layout/configs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestFunc func(context.Context, *http.Request) context.Context

type ResponseFunc func(context.Context, http.ResponseWriter) context.Context

type Server struct {
	*gin.Engine
	server  *http.Server
	conf    *configs.HTTPConf
	addr    string
	port    int
	timeout time.Duration
	before  []RequestFunc
	after   []ResponseFunc
}

type Option func(*Server)

func ServerBefore(before ...RequestFunc) Option {
	return func(s *Server) {
		s.before = append(s.before, before...)
	}
}

func ServerAfter(after ...ResponseFunc) Option {
	return func(s *Server) {
		s.after = append(s.after, after...)
	}
}

func Address(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

func Port(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

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

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
