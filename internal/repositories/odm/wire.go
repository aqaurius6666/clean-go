//go:build wireinject

package odm

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func BuildRepository(logger *logrus.Logger, cfg config.DBConfig) (*ODMRepository, error) {
	wire.Build(
		ConnectMongoDB,
		wire.Struct(new(ODMRepository), "*"),
	)
	return nil, nil
}
