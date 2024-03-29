package usecases

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

type UserUsecases interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
	UpdateUser(ctx context.Context, id string, user *entities.User) (*entities.User, error)
}

func (s *UsecasesService) GetUser(ctx context.Context, id string) (*entities.User, error) {
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.GetUser")
	defer span.End()
	return s.Repo.GetUserById(ctx, id)
}

func (s *UsecasesService) UpdateUser(ctx context.Context, id string, user *entities.User) (*entities.User, error) {
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.UpdateUser")
	defer span.End()
	return s.Repo.UpdateUser(ctx, gentity.Extend[*entities.User]{
		Entity: &entities.User{
			ID: id,
		},
	}, user)
}
