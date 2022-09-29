package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/wire"
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
	app, err := wire.BuildApp(ctx, cfg)
	if err != nil {
		return err
	}
	if err = app.Migrator.Migrate(ctx); err != nil {
		return err
	}
	app.RestApiServer.RegisterEndpoint()
	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTP.Port), app.RestApiServer)
}
