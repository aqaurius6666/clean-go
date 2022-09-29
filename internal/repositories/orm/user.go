package orm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/generics"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
)

func (s *ORMRepository) GetUserById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return generics.GetEntityById[*entities.User](ctx, s.DB, id.String())
}

func (s *ORMRepository) ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error) {
	return generics.ListEntities(ctx, s.DB, ex)
}

func (s *ORMRepository) InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	return generics.InsertEntity(ctx, s.DB, ex)
}
