package orm

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ORMRepository struct {
	DB *gorm.DB
}

func ConnectGorm(cfg config.DBConfig) (*gorm.DB, error) {
	gormOpts := []gorm.Option{}
	var (
		db  *gorm.DB
		err error
	)
	if cfg.DSN == "" {
		cfg.DSN = (&url.URL{
			Scheme:   cfg.Scheme,
			User:     url.UserPassword(cfg.User, cfg.Pass),
			Host:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Path:     cfg.Name,
			RawQuery: cfg.Query,
		}).String()
	}
	db, err = gorm.Open(postgres.Open(cfg.DSN), gormOpts...)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *ORMRepository) Migrate(ctx context.Context) error {
	return s.DB.WithContext(ctx).AutoMigrate(&entities.User{})
}
