package cql

import (
	"context"
	"fmt"

	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

type CQLRepository struct {
	Db     *gocql.Session
	Logger *logrus.Logger
}

func ConnectCassandra(cfg config.DBConfig) (*gocql.Session, error) {

	// if cfg.DSN == "" {
	// 	cfg.DSN = (&url.URL{
	// 		Scheme:   cfg.Scheme,
	// 		User:     url.UserPassword(cfg.User, cfg.Pass),
	// 		Host:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	// 		Path:     cfg.Name,
	// 		RawQuery: cfg.Query,
	// 	}).String()
	// }
	cluster := gocql.NewCluster(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.User,
		Password: cfg.Pass,
	}
	cluster.Keyspace = cfg.Name
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *CQLRepository) Migrate(ctx context.Context) error {
	return nil
}
