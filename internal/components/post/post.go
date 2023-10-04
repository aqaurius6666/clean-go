package post

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

type UseCase interface {
	CreatePost(context.Context, string, *entities.Post) (*entities.Post, error)
	ListPosts(ctx context.Context, id string, limit *int, offset *int) ([]*entities.Post, error)
	TotalPosts(ctx context.Context, id string) (int64, error)
}

type Repository interface {
	// GetUserById(ctx context.Context, id string) (*entities.User, error)
	ListPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) ([]*entities.Post, error)
	TotalPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) (int64, error)
	// SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
	InsertPost(ctx context.Context, ex gentity.Extend[*entities.Post]) (*entities.Post, error)
	// UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error)
}
