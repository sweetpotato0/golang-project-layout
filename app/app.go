package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"layout/pkg/transport"

	"golang.org/x/sync/errgroup"
)

// App is an application components lifecycle manager
type App struct {
	ctx        context.Context
	sigs       []os.Signal
	cancel     func()
	transports []transport.Server
}

// New create an app.
func New(servers ...transport.Server) *App {

	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx:        ctx,
		cancel:     cancel,
		sigs:       []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL},
		transports: servers,
	}
}

// Run .
func (a *App) Run() error {

	fmt.Println("server starting...")

	g, ctx := errgroup.WithContext(a.ctx)

	for _, transport := range a.transports {
		server := transport
		g.Go(func() error {
			<-ctx.Done()
			return server.Stop(ctx)
		})
		g.Go(func() error {
			return server.Start(ctx)
		})
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, a.sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return a.Stop()
			}
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	fmt.Println("server stoped.")
	return nil
}

// Stop stop the app.
func (a *App) Stop() error {
	a.cancel()
	return nil
}
