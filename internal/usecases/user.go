package usecases

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/google/uuid"
)

// var _ UserUsecases = (*UsecasesService)(nil)

type UserUsecases interface {
	GetUser(ctx context.Context, id uuid.UUID) (*entities.User, error)
}

func (s *UsecasesService) GetUser(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return s.Repo.GetUserById(ctx, id)
}
