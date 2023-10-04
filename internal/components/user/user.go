package user

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

type UseCase interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
	UpdateUser(ctx context.Context, id string, user *entities.User) (*entities.User, error)
}

type Repository interface {
	GetUserById(ctx context.Context, id string) (*entities.User, error)
	ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error)
	SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
	UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error)
}
