package repositories

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/repositories/cql"
	"github.com/aqaurius6666/clean-go/internal/repositories/es"
	"github.com/aqaurius6666/clean-go/internal/repositories/odm"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm"
	"github.com/sirupsen/logrus"
)

type RepositoryImpl interface{}

func BuildRepository(logger *logrus.Logger, cfg config.DBConfig) (RepositoryImpl, error) {
	switch cfg.Scheme {
	case "postgres":
		return orm.BuildRepository(logger, cfg)
	case "mongodb":
		return odm.BuildRepository(logger, cfg)
	case "cassandra":
		return cql.BuildRepository(logger, cfg)
	case "esdb":
		return es.BuildRepository(logger, cfg)
	default:
		panic("unknown db scheme")
	}
}
