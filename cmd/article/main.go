package main

import (
	"layout/app"
	"layout/configs"
	"layout/pkg/transport/http"
)

func newApp(hs *http.Server) *app.App {
	return app.New(hs)
}

func main() {
	a := initApp(configs.Get())
	if err := a.Run(); err != nil {
		panic(err)
	}

}
