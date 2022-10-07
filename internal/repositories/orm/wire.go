//go:build wireinject

package orm

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func BuildRepository(logger *logrus.Logger, cfg config.DBConfig) (*ORMRepository, error) {
	wire.Build(
		ConnectGorm,
		wire.Struct(new(ORMRepository), "*"),
	)
	return nil, nil
}
