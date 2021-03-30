package app

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"layout/pkg/transport/http"

	"golang.org/x/sync/errgroup"
)

// App is an application components lifecycle manager
type App struct {
	ctx    context.Context
	sigs   []os.Signal
	cancel func()
	hs     *http.Server
}

// New create an app.
func New(hs *http.Server) *App {

	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx:    ctx,
		cancel: cancel,
		sigs:   []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		hs:     hs,
	}
}

// Run .
func (a *App) Run() error {

	g, ctx := errgroup.WithContext(a.ctx)
	g.Go(func() error {
		<-ctx.Done() // wait for stop signal
		return a.hs.Stop()
	})
	g.Go(func() error {
		return a.hs.Start()
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, a.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) Stop() {}
