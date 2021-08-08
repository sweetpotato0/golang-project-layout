package http

import (
	"time"

	"layout/configs"
	"layout/internal/app/article/application"
	fhttp "layout/pkg/transport/http"
)

// NewHTTPServer create http server.
func NewHTTPServer(usecase *application.Article, c *configs.Configuration) *fhttp.Server {

	var opts = []fhttp.Option{}

	if c.Server.HTTP.Addr != "" {
		opts = append(opts, fhttp.Address(c.Server.HTTP.Addr))
	}
	if c.Server.HTTP.Port != 0 {
		opts = append(opts, fhttp.Port(c.Server.HTTP.Port))
	}
	if c.Server.HTTP.Timeout != 0 {
		opts = append(opts, fhttp.Timeout(time.Duration(c.Server.HTTP.Timeout)*time.Second))
	}

	server := fhttp.New(c.Server.Mode, opts...)
	h := Handler{
		UseCase: usecase,
	}
	ArticleRoutes(server, h)

	return server
}
