package generics

import (
	"context"

	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
)

type GenericRepository[T gentity.E] interface {
	GetEntityById(ctx context.Context, id uuid.UUID) (T, error)
	ListEntities(ctx context.Context, ext gentity.Extend[T]) ([]T, error)
	// GetEntity(ctx context.Context, ext gentity.Extend[T]) (T, error)
}
