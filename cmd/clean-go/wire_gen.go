// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/generics"
	"github.com/aqaurius6666/clean-go/internal/repositories"
	"github.com/aqaurius6666/clean-go/internal/repositories/odm"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm"
	"github.com/aqaurius6666/clean-go/internal/restapi"
	"github.com/aqaurius6666/clean-go/internal/restapi/v1"
	"github.com/aqaurius6666/clean-go/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Injectors from wire.go:

func BuildApp(ctx context.Context, cfg config.AppConfig) (*App, error) {
	engine := gin.New()
	logConfig := cfg.Log
	logger := config.NewLogger(logConfig)
	authConfig := cfg.Auth
	dbConfig := cfg.Db
	repositoryImpl, err := repositories.BuildRepository(logger, dbConfig)
	if err != nil {
		return nil, err
	}
	repository := CastRepository(repositoryImpl)
	usecasesService := &usecases.UsecasesService{
		Logger:     logger,
		AuthConfig: authConfig,
		Repo:       repository,
	}
	handler := &v1.Handler{
		Usecase: usecasesService,
	}
	middleware := &v1.Middleware{
		L:        logger,
		Usecases: usecasesService,
	}
	restAPIServer := &restapi.RestAPIServer{
		G:          engine,
		Logger:     logger,
		Handler:    handler,
		Middleware: middleware,
	}
	migrator := CastMigrator(repositoryImpl)
	app := &App{
		RestApiServer: restAPIServer,
		Migrator:      migrator,
	}
	return app, nil
}

// wire.go:

type App struct {
	RestApiServer restapi.Server
	Migrator      usecases.Migrator
}

// wire provider set
var (
	UsecaseSet = wire.NewSet(wire.Struct(new(usecases.UsecasesService), "*"), wire.Bind(new(usecases.Usecases), new(*usecases.UsecasesService)))
	// ORMSet = wire.NewSet(
	// 	wire.Bind(new(usecases.Repository), new(*orm.ORMRepository)),
	// 	wire.Bind(new(usecases.Migrator), new(*orm.ORMRepository)),
	// 	wire.Struct(new(orm.ORMRepository), "*"),
	// 	orm.ConnectGorm,
	// )
	// ODMSet = wire.NewSet(
	// 	wire.Bind(new(usecases.Repository), new(*odm.ODMRepository)),
	// 	wire.Bind(new(usecases.Migrator), new(*odm.ODMRepository)),
	// 	wire.Struct(new(odm.ODMRepository), "*"),
	// 	odm.ConnectMongoDB,
	// )
	RepositorySet = wire.NewSet(repositories.BuildRepository, CastRepository,
		CastMigrator,
	)
	RestSet      = wire.NewSet(wire.Bind(new(restapi.Server), new(*restapi.RestAPIServer)), wire.Struct(new(restapi.RestAPIServer), "*"), wire.Bind(new(restapi.Handler), new(*v1.Handler)), wire.Bind(new(restapi.Middleware), new(*v1.Middleware)), RestApiV1Set, gin.New)
	RestApiV1Set = wire.NewSet(wire.Struct(new(v1.Handler), "*"), wire.Struct(new(v1.Middleware), "*"))
)

// interface constraints
var (
	_ usecases.Usecases                          = (*usecases.UsecasesService)(nil)
	_ usecases.Repository                        = (*orm.ORMRepository)(nil)
	_ usecases.Repository                        = (*odm.ODMRepository)(nil)
	_ generics.GenericRepository[*entities.User] = (*generics.ORMGenericRepository[*entities.User])(nil)
	_ usecases.Migrator                          = (*orm.ORMRepository)(nil)
	_ restapi.Server                             = (*restapi.RestAPIServer)(nil)
	_ restapi.Handler                            = (*v1.Handler)(nil)
	_ restapi.Middleware                         = (*v1.Middleware)(nil)
)

func CastRepository(r repositories.RepositoryImpl) usecases.Repository {
	repo, ok := r.(usecases.Repository)
	if !ok {
		panic("failed to cast repository")
	}
	return repo
}

func CastMigrator(r repositories.RepositoryImpl) usecases.Migrator {
	repo, ok := r.(usecases.Migrator)
	if !ok {
		panic("failed to cast migrator")
	}
	return repo
}
