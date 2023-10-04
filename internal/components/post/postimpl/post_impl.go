package postimpl

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/components/post"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

type UseCaseImpl struct {
	PostRepo post.Repository
}

func (s *UseCaseImpl) CreatePost(ctx context.Context, id string, post *entities.Post) (*entities.Post, error) {
	post.CreatorID = id
	return s.PostRepo.InsertPost(ctx, gentity.Extend[*entities.Post]{
		Entity: post,
	})
}

func (s *UseCaseImpl) ListPosts(ctx context.Context, id string, limit *int, offset *int) ([]*entities.Post, error) {
	return s.PostRepo.ListPosts(ctx, gentity.Extend[*entities.Post]{
		Entity: &entities.Post{
			CreatorID: id,
		},
		ExFields: &gentity.ExtendFields{
			Offset: offset,
			Limit:  limit,
		},
	})
}

func (s *UseCaseImpl) TotalPosts(ctx context.Context, id string) (int64, error) {
	return s.PostRepo.TotalPosts(ctx, gentity.Extend[*entities.Post]{
		Entity: &entities.Post{
			CreatorID: id,
		},
	})
}
