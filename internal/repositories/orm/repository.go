package orm

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm/gormgen"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ORMRepository struct {
	DB *gorm.DB
}

// var _ usecases.Repository = (*ORMRepository)(nil)

func ConnectGorm(cfg config.DBConfig) (*gorm.DB, error) {
	gormOpts := []gorm.Option{
		&gorm.Config{
			PrepareStmt:                              true,
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   logger.Default.LogMode(logger.Info),
		},
	}
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

	if cfg.OTELEnabled {
		if err := db.Use(otelgorm.NewPlugin()); err != nil {
			return nil, err
		}
	}
	gormgen.SetDefault(db)

	return db, nil
}

func (s *ORMRepository) Migrate(ctx context.Context) error {
	return s.DB.WithContext(ctx).AutoMigrate(&entities.User{}, &entities.Post{}, &entities.React{})
}
