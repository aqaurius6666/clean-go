package usecases

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
)

type PostUsecases interface {
	CreatePost(context.Context, string, *entities.Post) (*entities.Post, error)
	GetPostById(ctx context.Context, id string) (*entities.Post, error)
	ListPosts(ctx context.Context, id string, limit int, offset int) ([]*entities.Post, error)
	TotalPosts(ctx context.Context, id string) (int64, error)
}

func (s *UsecasesService) CreatePost(ctx context.Context, id string, post *entities.Post) (*entities.Post, error) {
	post.CreatorID = id
	return s.Repo.InsertPost(ctx, post)
}

func (s *UsecasesService) ListPosts(ctx context.Context, id string, limit int, offset int) ([]*entities.Post, error) {
	return s.Repo.ListPostsByCreatorId(ctx, id, offset, limit)
}

func (s *UsecasesService) TotalPosts(ctx context.Context, id string) (int64, error) {
	return s.Repo.CountPostsByCreatorId(ctx, id)
}

func (s *UsecasesService) GetPostById(ctx context.Context, id string) (*entities.Post, error) {
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.GetPostById")
	defer span.End()
	post, err := s.Repo.GetPostById(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
