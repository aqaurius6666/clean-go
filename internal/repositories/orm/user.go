package orm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

func (s *ORMRepository) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	return gentity.GetEntityById[*entities.User](ctx, s.DB, id)
}

func (s *ORMRepository) ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error) {
	return gentity.ListEntities(ctx, s.DB, ex)
}

func (s *ORMRepository) InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	return gentity.InsertEntity(ctx, s.DB, ex)
}

func (s *ORMRepository) SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	return gentity.GetEntity(ctx, s.DB, ex)
}

func (s *ORMRepository) UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error) {
	return gentity.UpdateEntity(ctx, s.DB, ex, v)
}
