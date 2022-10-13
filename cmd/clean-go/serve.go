package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/pkg/otel"
	"github.com/aqaurius6666/clean-go/pkg/parsecli"
	"github.com/urfave/cli/v2"
)

func buildServeCmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "serve",
		Usage: "Run server",
		Flags: parsecli.Convert(config.AppConfig{}, ""),
		Action: func(c *cli.Context) error {
			cfg := config.AppConfig{}
			err := parsecli.Parse(c, &cfg)
			if err != nil {
				return err
			}
			return serve(cfg)
		},
	}
	return cmd
}

func serve(cfg config.AppConfig) error {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app, err := BuildApp(ctx, cfg)
	if err != nil {
		return err
	}
	app.Logger.Info("Initialize application successfully")
	if err = app.Migrator.Migrate(ctx); err != nil {
		return err
	}
	app.Logger.Info("Migrate successfully")
	if cfg.Otel.Enabled {
		clearFunc, err := otel.InitOtel(ctx, otel.OtelOptions{
			CollectorAddr:  cfg.Otel.CollectorAddr,
			ID:             cfg.Otel.ID,
			ServiceName:    cfg.Otel.ServiceName,
			MetricPeriodic: 5 * time.Second,
		})
		app.Logger.Info("Initialize OTEL successfully")
		defer func() {
			clearCtx := context.Background()
			clearCtx, cancel := context.WithTimeout(clearCtx, 5*time.Second)
			defer cancel()
			_ = clearFunc(clearCtx)
		}()
		if err != nil {
			return err
		}
	} else {
		app.Logger.Info("OTEL is disabled")
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Http.Port),
		Handler: app.RestApiServer,
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		app.Logger.Info("Start server")
		app.RestApiServer.RegisterEndpoint()
		app.Logger.Info("Register endpoint successfully")

		if err := httpServer.ListenAndServe(); err != nil {
			app.Logger.Fatal(err)
		}
	}()
	shutdownFunc := func(ctx context.Context) {
		if err := httpServer.Shutdown(ctx); err != nil {
			app.Logger.Fatal(err)
		}
	}
	defer func() {
		if err := recover(); err != nil {
			app.Logger.Info("recover")
			app.Logger.Fatal(err)
		}
	}()
	defer func() {
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 5*time.Second)
		defer cancel()
		shutdownFunc(shutdownCtx)
	}()
	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case <-ctx.Done():
			app.Logger.Info("Shutdown server by context done")
			return ctx.Err()
		case <-killChan:
			app.Logger.Info("Shutdown server by kill signal")
			return nil
		}
	}
}
