package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"proto-service/internal/domain/usecases"
	"proto-service/pkg/infra/logger"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Adapter struct {
	s   *http.Server
	l   net.Listener
	app *usecases.SerializationApp
}

type AdapterOptions struct {
	HTTP_port int
	IsProd    bool
}

func New(app *usecases.SerializationApp, opts AdapterOptions) (*Adapter, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.HTTP_port))
	if err != nil {
		return nil, fmt.Errorf("server start failed: %w", err)
	}
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	server := http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	a := Adapter{
		s:   &server,
		l:   l,
		app: app,
	}
	optsLogger := logger.LoggerOptions{IsProd: opts.IsProd}
	err = initRouter(&a, router, optsLogger)
	return &a, err
}

func (a *Adapter) Start() error {
	eg := &errgroup.Group{}
	eg.Go(func() error {
		return a.s.Serve(a.l)
	})
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (a *Adapter) Stop(ctx context.Context) error {
	var (
		err  error
		once sync.Once
	)
	once.Do(func() {
		err = a.s.Shutdown(ctx)
	})
	return err
}
