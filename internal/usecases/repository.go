package usecases

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

type Repository interface {
	UserRepository
	PostRepository
}

type UserRepository interface {
	GetUserById(ctx context.Context, id string) (*entities.User, error)
	ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error)
	SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
	InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
	UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error)
}

type PostRepository interface {
	// GetUserById(ctx context.Context, id string) (*entities.User, error)
	ListPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) ([]*entities.Post, error)
	TotalPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) (int64, error)
	// SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
	InsertPost(ctx context.Context, ex gentity.Extend[*entities.Post]) (*entities.Post, error)
	// UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error)
}

type Migrator interface {
	Migrate(context.Context) error
}
