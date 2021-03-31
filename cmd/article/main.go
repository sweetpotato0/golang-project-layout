package main

import (
	"layout/app"
	"layout/configs"
	"layout/pkg/transport/grpc"
	"layout/pkg/transport/http"
)

func newApp(hs *http.Server, gs *grpc.Server) *app.App {
	return app.New(hs, gs)
}

func main() {
	a := initApp(configs.Get())
	if err := a.Run(); err != nil {
		panic(err)
	}

}
