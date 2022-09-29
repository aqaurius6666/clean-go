package wire

import (
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/generics"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm"
	"github.com/aqaurius6666/clean-go/internal/restapi"
	v1 "github.com/aqaurius6666/clean-go/internal/restapi/v1"
	"github.com/aqaurius6666/clean-go/internal/usecases"
	apipb "github.com/aqaurius6666/clean-go/pkg/proto/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// wire provider set
var (
	UsecaseSet = wire.NewSet(
		wire.Struct(new(usecases.UsecasesService), "*"),
		wire.Bind(new(usecases.Usecases), new(*usecases.UsecasesService)),
	)
	ORMSet = wire.NewSet(
		wire.Bind(new(usecases.Repository), new(*orm.ORMRepository)),
		wire.Bind(new(usecases.Migrator), new(*orm.ORMRepository)),
		wire.Struct(new(orm.ORMRepository), "*"),
		orm.ConnectGorm,
	)
	RestSet = wire.NewSet(
		wire.Bind(new(restapi.Server), new(*restapi.RestAPIServer)),
		wire.Struct(new(restapi.RestAPIServer), "*"),
		wire.Bind(new(restapi.Handler), new(*v1.Handler)),
		wire.Bind(new(restapi.Middleware), new(*v1.Middleware)),
		RestApiV1Set,
		gin.New,
		UserGenericSet,
	)
	RestApiV1Set   = wire.NewSet(wire.Struct(new(v1.Handler), "*"), wire.Struct(new(v1.Middleware), "*"))
	UserGenericSet = wire.NewSet(
		NewUserHandler,
		generics.NewUserGenericRepository,
		wire.Bind(new(generics.GenericRepository[*entities.User]), new(*generics.ORMGenericRepository[*entities.User])),
		wire.Bind(new(restapi.UserHandler), new(*generics.GenericHandler[*entities.User, *apipb.UserEntity])))
)

// interface constraints
var (
	_ usecases.Usecases                          = (*usecases.UsecasesService)(nil)
	_ usecases.Repository                        = (*orm.ORMRepository)(nil)
	_ generics.GenericRepository[*entities.User] = (*generics.ORMGenericRepository[*entities.User])(nil)
	_ usecases.Migrator                          = (*orm.ORMRepository)(nil)
	_ restapi.Server                             = (*restapi.RestAPIServer)(nil)
	_ restapi.Handler                            = (*v1.Handler)(nil)
	_ restapi.UserHandler                        = (*generics.GenericHandler[*entities.User, *apipb.UserEntity])(nil)
	_ restapi.Middleware                         = (*v1.Middleware)(nil)
)

// generic initializer
func NewUserHandler(repo generics.GenericRepository[*entities.User]) *generics.GenericHandler[*entities.User, *apipb.UserEntity] {
	return &generics.GenericHandler[*entities.User, *apipb.UserEntity]{
		Usecase: repo,
	}
}
