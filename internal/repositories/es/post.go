package es

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
)

func (s *ESClient) ListPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) ([]*entities.Post, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ESClient) TotalPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) (int64, error) {
	panic("not implemented") // TODO: Implement
}

// SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
// UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error)
func (s *ESClient) InsertPost(ctx context.Context, ex gentity.Extend[*entities.Post]) (*entities.Post, error) {
	panic("not implemented") // TODO: Implement
}
