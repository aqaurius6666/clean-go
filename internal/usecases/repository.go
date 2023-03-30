package usecases

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
)

type Repository interface {
	UserRepository
	PostRepository
	ReactRepository
}

type UserRepository interface {
	InsertUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByEmailAndPassword(ctx context.Context, username, password string) (*entities.User, error)
	GetUserById(ctx context.Context, id string) (*entities.User, error)
	UpdateUser(ctx context.Context, id string, user *entities.User) (*entities.User, error)
}

type PostRepository interface {
	GetPostById(ctx context.Context, id string) (*entities.Post, error)
	ListPostsByCreatorId(ctx context.Context, id string, offset, limit int) ([]*entities.Post, error)
	CountPostsByCreatorId(ctx context.Context, id string) (int64, error)
	InsertPost(ctx context.Context, post *entities.Post) (*entities.Post, error)
}

type ReactRepository interface {
	UpsertReact(ctx context.Context, react *entities.React) (*entities.React, error)
	GetReactsByPostId(ctx context.Context, id string) ([]*entities.React, error)
	GetReactsByPostIds(ctx context.Context, ids []string) ([]*entities.React, error)
}
type Migrator interface {
	Migrate(context.Context) error
}
