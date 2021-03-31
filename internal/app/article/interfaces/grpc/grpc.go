package grpc

import (
	"layout/configs"
	"layout/internal/app/article/application"
	"layout/pkg/transport/grpc"
)

// Run start grpc server
func NewGrpcServer(usecase *application.Article, c *configs.Configuration) *grpc.Server {
	server := grpc.New()
	return server
}
