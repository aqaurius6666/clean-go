package generics

import (
	"context"

	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"gorm.io/gorm"
)

type ORMGenericRepository[T gentity.E] struct {
	db *gorm.DB
}

func (s *ORMGenericRepository[T]) GetEntity(ctx context.Context, ext gentity.Extend[T]) (T, error) {
	return gentity.GetEntity(ctx, s.db, ext)
}

func (s *ORMGenericRepository[T]) GetEntityById(ctx context.Context, id string) (T, error) {
	return gentity.GetEntityById[T](ctx, s.db, id)
}

func (s *ORMGenericRepository[T]) ListEntities(ctx context.Context, ext gentity.Extend[T]) ([]T, error) {
	return gentity.ListEntities(ctx, s.db, ext)
}
func (s *ORMGenericRepository[T]) TotalEntity(ctx context.Context, ext gentity.Extend[T]) (int64, error) {
	return gentity.TotalEntity(ctx, s.db, ext)
}

func (s *ORMGenericRepository[T]) InsertEntity(ctx context.Context, ext gentity.Extend[T]) (T, error) {
	return gentity.InsertEntity(ctx, s.db, ext)
}
func (s *ORMGenericRepository[T]) UpdateEntity(ctx context.Context, ext gentity.Extend[T], v T) (T, error) {
	return gentity.UpdateEntity(ctx, s.db, ext, v)
}
func (s *ORMGenericRepository[T]) DeleteEntity(ctx context.Context, ext gentity.Extend[T]) error {
	return gentity.DeleteEntity(ctx, s.db, ext)
}
