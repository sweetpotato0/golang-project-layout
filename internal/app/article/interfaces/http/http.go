package http

import (
	"time"

	"layout/configs"
	"layout/internal/app/article/application"
	fhttp "layout/pkg/transport/http"
)

// Run start http server
func NewHttpServer(usecase *application.Article, c *configs.Configuration) *fhttp.Server {

	var opts = []fhttp.Option{}

	if c.Server.Http.Addr != "" {
		opts = append(opts, fhttp.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Port != 0 {
		opts = append(opts, fhttp.Port(c.Server.Http.Port))
	}
	if c.Server.Http.Timeout != 0 {
		opts = append(opts, fhttp.Timeout(time.Duration(c.Server.Http.Timeout)*time.Second))
	}

	server := fhttp.New(c.Server.Mode, opts...)
	h := Handler{
		UseCase: usecase,
	}
	ArticleRoutes(server, h)

	return server
}
