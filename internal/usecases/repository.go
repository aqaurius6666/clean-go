package usecases

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
)

type Repository interface {
	UserRepository
}
type GenericRepository[T gentity.E] interface {
	GetEntityById(ctx context.Context, id uuid.UUID) (T, error)
	// GetEntity(ctx context.Context, ext gentity.Extend[T]) (T, error)
}
type UserRepository interface {
	GetUserById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error)
	InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
}

type Migrator interface {
	Migrate(context.Context) error
}
