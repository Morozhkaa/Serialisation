package application

import (
	"context"
	"fmt"
	"yaml-service/internal/adapters/http"
	"yaml-service/internal/domain/usecases"
)

type App struct {
	opts          AppOptions
	shutdownFuncs []func(ctx context.Context) error
}

type AppOptions struct {
	HTTP_port int
	IsProd    bool
}

func New(opts AppOptions) *App {
	return &App{
		opts: opts,
	}
}

func (app *App) Start() error {
	serializationApp := usecases.New()
	optsAdapter := http.AdapterOptions{HTTP_port: app.opts.HTTP_port, IsProd: app.opts.IsProd}
	s, err := http.New(serializationApp, optsAdapter)
	if err != nil {
		return fmt.Errorf("server not started %w", err)
	}
	app.shutdownFuncs = append(app.shutdownFuncs, s.Stop)
	err = s.Start()
	if err != nil {
		return fmt.Errorf("server not started: %w", err)
	}
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	var err error
	for i := len(a.shutdownFuncs) - 1; i >= 0; i-- {
		err = a.shutdownFuncs[i](ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
