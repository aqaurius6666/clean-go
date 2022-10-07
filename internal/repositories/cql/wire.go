//go:build wireinject

package cql

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func BuildRepository(logger *logrus.Logger, cfg config.DBConfig) (*CQLRepository, error) {
	wire.Build(
		ConnectCassandra,
		wire.Struct(new(CQLRepository), "*"),
	)
	return nil, nil
}
