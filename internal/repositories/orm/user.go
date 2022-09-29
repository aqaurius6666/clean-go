package orm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
)

func (s *ORMRepository) GetUserById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return GetEntityById[*entities.User](ctx, s.db, id.String())
}

func (s *ORMRepository) ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error) {
	return ListEntities(ctx, s.db, ex)
}

func (s *ORMRepository) InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	return InsertEntity(ctx, s.db, ex)
}
