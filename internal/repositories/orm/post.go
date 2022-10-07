package orm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

func (s *ORMRepository) InsertPost(ctx context.Context, ex gentity.Extend[*entities.Post]) (*entities.Post, error) {
	return gentity.InsertEntity(ctx, s.DB, ex)
}

func (s *ORMRepository) ListPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) ([]*entities.Post, error) {
	return gentity.ListEntities(ctx, s.DB, ex)
}

func (s *ORMRepository) TotalPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) (int64, error) {
	return gentity.TotalEntity(ctx, s.DB, ex)
}
