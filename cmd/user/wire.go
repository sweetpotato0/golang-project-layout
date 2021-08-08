// +build wireinject

package main

import (
	"layout/app"
	"layout/configs"
	"layout/internal/app/user/application"
	"layout/internal/app/user/interfaces/grpc"
	"layout/internal/app/user/repository"

	"github.com/google/wire"
)

// initApp init application.
func initApp(conf *configs.Configuration) *app.App {

	wire.Build(
		repository.NewUser,         // 数据库
		application.NewUserUseCase, // 业务逻辑
		grpc.NewGrpcServer,
		newApp, // App
	)
	return &app.App{}
}
