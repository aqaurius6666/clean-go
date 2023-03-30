package usecases

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/var/e"
	"github.com/pkg/errors"
)

type ReactUsecases interface {
	CreateReact(ctx context.Context, userId string, postId string, reactType entities.ReactType) (*entities.React, error)
	GetReactsByPostIds(ctx context.Context, postIds []string) ([]*entities.React, error)
	// ListReacts(ctx context.Context, id string, limit *int, offset *int) ([]*entities.React, error)
	// TotalReacts(ctx context.Context, id string) (int64, error)
}

func (s *UsecasesService) CreateReact(ctx context.Context, userId string, postId string, reactType entities.ReactType) (*entities.React, error) {
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.CreateReact")
	defer span.End()
	_, err := s.Repo.GetPostById(ctx, postId)
	if err != nil {
		return nil, errors.New(e.ErrPostNotFound)
	}
	react, err := s.Repo.UpsertReact(ctx, &entities.React{
		UserID: userId,
		PostID: postId,
		Type:   reactType,
	})
	if err != nil {
		return nil, err
	}
	return react, nil
}

func (s *UsecasesService) GetReactsByPostIds(ctx context.Context, postIds []string) ([]*entities.React, error) {
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.GetReactsByPostIds")
	defer span.End()
	if len(postIds) == 0 {
		return nil, nil
	}
	reacts, err := s.Repo.GetReactsByPostIds(ctx, postIds)
	if err != nil {
		return nil, err
	}
	return reacts, nil
}
