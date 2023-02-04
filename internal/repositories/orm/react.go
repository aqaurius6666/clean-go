package orm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

func (s *ORMRepository) InsertReact(ctx context.Context, ex gentity.Extend[*entities.React]) (*entities.React, error) {
	return gentity.InsertEntity(ctx, s.DB, ex)
}

func (s *ORMRepository) SelectReact(ctx context.Context, ex gentity.Extend[*entities.React]) (*entities.React, error) {
	return gentity.GetEntity(ctx, s.DB, ex)
}

func (s *ORMRepository) ListReacts(ctx context.Context, ex gentity.Extend[*entities.React]) ([]*entities.React, error) {
	return gentity.ListEntities(ctx, s.DB, ex)
}
