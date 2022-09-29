package main

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/pkg/parsecli"
	"github.com/urfave/cli/v2"
)

func buildCmd() *cli.App {
	cmd := &cli.App{
		Name: "clean-go",
		Commands: []*cli.Command{
			buildServeCmd(),
		},
		Flags: parsecli.Convert(config.AppConfig{}, ""),
	}

	return cmd
}
