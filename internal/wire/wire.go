//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/restapi"
	"github.com/aqaurius6666/clean-go/internal/usecases"
	"github.com/google/wire"
)

type App struct {
	RestApiServer restapi.Server
	Migrator      usecases.Migrator
}

func BuildApp(ctx context.Context, cfg config.AppConfig) (*App, error) {

	wire.Build(
		wire.FieldsOf(&cfg, "Db", "Auth", "Log"),
		wire.Struct(new(App), "*"),
		config.NewLogger,
		RestSet,
		UsecaseSet,
		RepositorySet,
		// ODMSet,
	)
	return nil, nil
}
