// The basic deserialization/serialization in GO
package main

import (
	"context"
	"json-service/internal/application"
	"json-service/internal/config"
	"json-service/pkg/infra/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	// get environment variable values ​​(HTTP_PORT, IsProd)
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("getting config failed: %s", err.Error())
	}

	// initialize logger
	optsLogger := logger.LoggerOptions{IsProd: cfg.IsProd}
	l, err := logger.New(optsLogger)
	if err != nil {
		log.Fatalf("logger initialization failed: %s", err.Error())
	}

	// create and start application
	optsApp := application.AppOptions{HTTP_port: cfg.HTTP_port, IsProd: cfg.IsProd}
	app := application.New(optsApp)
	err = app.Start()
	if err != nil {
		l.Sugar().Fatalf("app not started: %s", err.Error())
	}

	// when the context completes, gracefully exit
	<-ctx.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()

	err = app.Stop(stopCtx)
	if err != nil {
		l.Sugar().Error(err)
	}
	l.Info("app stopped")
}
