// +build wireinject

package main

import (
	"layout/app"
	"layout/configs"
	"layout/internal/app/article/application"
	"layout/internal/app/article/interfaces/http"
	"layout/internal/app/article/repository"

	"github.com/google/wire"
)

// initApp init application.
func initApp(conf *configs.Configuration) *app.App {

	wire.Build(
		repository.NewArticle,         // 数据库
		application.NewArticleUseCase, // 业务逻辑
		http.NewHTTPServer,            // 相当于 controller，处理参数等
		newApp,                        // App
	)
	return &app.App{}
}
