package generics

import (
	"context"

	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

type GenericRepository[T gentity.E] interface {
	GetEntityById(ctx context.Context, id string) (T, error)
	ListEntities(ctx context.Context, ext gentity.Extend[T]) ([]T, error)
	InsertEntity(ctx context.Context, ext gentity.Extend[T]) (T, error)
	TotalEntity(ctx context.Context, ext gentity.Extend[T]) (int64, error)
	UpdateEntity(ctx context.Context, ext gentity.Extend[T], v T) (T, error)
	DeleteEntity(ctx context.Context, ext gentity.Extend[T]) error
}
