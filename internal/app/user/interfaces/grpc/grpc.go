package grpc

import (
	"time"

	"layout/api/user/v1"
	"layout/configs"
	"layout/internal/app/user/application"
	"layout/internal/app/user/service"
	"layout/pkg/log"
	"layout/pkg/transport/grpc"
	"layout/register"

	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// NewGrpcServer create grpc server.
func NewGrpcServer(usecase *application.User, c *configs.Configuration) *grpc.Server {

	reg := register.NewRegistry(c)

	ops := []register.RegistryOption{
		register.Port(c.Server.Grpc.Port),
		register.Tag("user_service"),
		register.Name("user_service"),
	}
	for _, opt := range ops {
		opt(reg)
	}

	err := reg.Register()
	if err != nil {
		log.Errorf("register failed: %#v", err)
	}

	opts := []grpc.Option{}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.Server.Grpc.Timeout)))
	}
	srv := grpc.New(opts...)

	userService := service.User{
		UseCase: usecase,
	}
	v1.RegisterUserServer(srv, userService)

	grpc_health_v1.RegisterHealthServer(srv, &register.Healthy{Status: grpc_health_v1.HealthCheckResponse_SERVING})
	reflection.Register(srv)
	return srv
}
