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

func New(options ...Option) *Server {

	server := &Server{
		Engine: gin.New(),
	}
	for _, option := range options {
		option(server)
	}

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	return server
}

func (s *Server) Start() error {
	return s.Run(fmt.Sprintf("%s:%d", s.addr, s.port))
}

func (s *Server) Stop() error {
	return nil
}
