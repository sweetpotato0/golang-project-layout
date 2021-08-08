package grpc

import (
	"context"
	"fmt"
	"log"

	"layout/configs"
	"layout/register"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

// NewGrpcConn create a grpc connection.
func NewGrpcConn(ctx context.Context, name string, c *configs.Configuration) (*grpc.ClientConn, error) {

	builder := register.NewBuilder(c)
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s://%s/%s", "consul", c.Consul.Addr, name),
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithResolvers(builder),
	)
	if err != nil {
		log.Fatalln(err)
	}

	return conn, nil
}
