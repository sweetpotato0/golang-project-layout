package main

import (
	"layout/app"
	"layout/configs"
	"layout/pkg/transport/grpc"
)

func newApp(gs *grpc.Server) *app.App {
	return app.New(gs)
}

func main() {
	a := initApp(configs.Get())
	if err := a.Run(); err != nil {
		panic(err)
	}

}
