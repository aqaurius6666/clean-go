//go:build wireinject

package main

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/components/auth"
	"github.com/aqaurius6666/clean-go/internal/components/auth/authimpl"
	"github.com/aqaurius6666/clean-go/internal/components/post"
	"github.com/aqaurius6666/clean-go/internal/components/post/postimpl"
	"github.com/aqaurius6666/clean-go/internal/components/user"
	"github.com/aqaurius6666/clean-go/internal/components/user/userimpl"
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/generics"
	"github.com/aqaurius6666/clean-go/internal/repositories"
	"github.com/aqaurius6666/clean-go/internal/restapi"
	v1 "github.com/aqaurius6666/clean-go/internal/restapi/v1"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

type App struct {
	RestApiServer restapi.Server
	Migrator      Migrator
	Logger        *logrus.Logger
}
type Migrator interface {
	Migrate(context.Context) error
}

// wire provider set
var (
	UserSet = wire.NewSet(
		wire.Struct(new(userimpl.UsecasesImpl), "*"),
		wire.Bind(new(user.Usecases), new(*userimpl.UsecasesImpl)),
		CastUserRepository,
	)
	PostSet = wire.NewSet(
		wire.Struct(new(postimpl.UsecasesImpl), "*"),
		wire.Bind(new(post.Usecases), new(*postimpl.UsecasesImpl)),
		CastPostRepository,
	)
	AuthSet = wire.NewSet(
		wire.Struct(new(authimpl.UsecasesImpl), "*"),
		wire.Bind(new(auth.Usecases), new(*authimpl.UsecasesImpl)),
		CastAuthRepository,
	)

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
	RepositorySet = wire.NewSet(
		repositories.BuildRepository,
		// CastRepository,
		CastMigrator,
	)
	RestSet = wire.NewSet(
		wire.Bind(new(restapi.Server), new(*restapi.RestAPIServer)),
		wire.Struct(new(restapi.RestAPIServer), "*"),
		wire.Bind(new(restapi.Handler), new(*v1.Handler)),
		wire.Bind(new(restapi.Middleware), new(*v1.Middleware)),
		RestApiV1Set,
		gin.New,
	)
	RestApiV1Set = wire.NewSet(wire.Struct(new(v1.Handler), "*"), wire.Struct(new(v1.Middleware), "*"))
)

// interface constraints
var (
	_ generics.GenericRepository[*entities.User] = (*generics.ORMGenericRepository[*entities.User])(nil)
	_ restapi.Server                             = (*restapi.RestAPIServer)(nil)
	_ restapi.Handler                            = (*v1.Handler)(nil)
	_ restapi.Middleware                         = (*v1.Middleware)(nil)
)

func BuildApp(ctx context.Context, cfg config.AppConfig) (*App, error) {

	wire.Build(
		wire.FieldsOf(&cfg, "Db", "Auth", "Log"),
		wire.Struct(new(App), "*"),
		config.NewLogger,
		RestSet,
		UserSet,
		PostSet,
		AuthSet,
		RepositorySet,
		// ODMSet,
	)
	return nil, nil
}

// func CastRepository(r repositories.RepositoryImpl) usecases.Repository {
// 	repo, ok := r.(usecases.Repository)
// 	if !ok {
// 		panic("failed to cast repository")
// 	}
// 	return repo
// }

func CastMigrator(r repositories.RepositoryImpl) Migrator {
	return CastRepositoryTo[Migrator](r)
}

func CastUserRepository(r repositories.RepositoryImpl) user.Repository {
	return CastRepositoryTo[user.Repository](r)
}
func CastPostRepository(r repositories.RepositoryImpl) post.Repository {
	return CastRepositoryTo[post.Repository](r)
}
func CastAuthRepository(r repositories.RepositoryImpl) auth.Repository {
	return CastRepositoryTo[auth.Repository](r)
}
