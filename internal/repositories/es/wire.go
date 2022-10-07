//go:build wireinject

package es

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func BuildRepository(logger *logrus.Logger, cfg config.DBConfig) (*ESClient, error) {
	wire.Build(
		NewESClient,
		// wire.Struct(new(ESClient), "*"),
	)
	return nil, nil
}
