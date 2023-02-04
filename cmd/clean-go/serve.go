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
	"github.com/pkg/errors"
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
			cfg = cfg.Modify()
			return serve(cfg)
		},
	}
	return cmd
}

func serve(cfg config.AppConfig) error {
	var err error
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	manualCfg, err := manualConfig(ctx, cfg)
	if err != nil {
		return err
	}
	app, err := BuildApp(ctx, cfg, manualCfg)
	if err != nil || app == nil {
		return errors.New("buildApp failed")
	}
	app.Logger.Info("Initialize application successfully")
	if err = app.Migrator.Migrate(ctx); err != nil {
		return err
	}
	app.Logger.Info("Migrate successfully")

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

func manualConfig(ctx context.Context, cfg config.AppConfig) (manualCfg ManualConfig, err error) {
	manualCfg.Logger = config.NewLogger(cfg.Log)
	manualCfg.Tracer = otel.NoOpTracer

	if cfg.Otel.Enabled {
		_, err = otel.InitOtel(ctx, otel.OtelOptions{
			CollectorAddr:  cfg.Otel.CollectorAddr,
			ID:             cfg.Otel.ID,
			ServiceName:    cfg.Otel.ServiceName,
			MetricPeriodic: 5 * time.Second,
		})
		if err != nil {
			return
		}
		manualCfg.Tracer = otel.TracerProvider(cfg.Otel.ServiceName)
		manualCfg.Logger.Info("Initialize OTEL successfully")

	} else {
		manualCfg.Logger.Info("OTEL is disabled")
	}
	return
}
