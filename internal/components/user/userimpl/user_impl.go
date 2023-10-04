package userimpl

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/components/user"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

type UseCaseImpl struct {
	UserRepo user.Repository
}

func (s *UseCaseImpl) GetUser(ctx context.Context, id string) (*entities.User, error) {
	return s.UserRepo.GetUserById(ctx, id)
}

func (s *UseCaseImpl) UpdateUser(ctx context.Context, id string, user *entities.User) (*entities.User, error) {
	return s.UserRepo.UpdateUser(ctx, gentity.Extend[*entities.User]{
		Entity: &entities.User{
			ID: id,
		},
	}, user)
}
