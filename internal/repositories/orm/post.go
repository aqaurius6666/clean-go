package orm

import (
	"context"
	"fmt"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm/gormgen"
)

func (s *ORMRepository) GetPostById(ctx context.Context, id string) (*entities.Post, error) {
	postQ := gormgen.Post
	post, err := postQ.WithContext(ctx).Where(postQ.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *ORMRepository) ListPostsByCreatorId(ctx context.Context, id string, offset int, limit int) ([]*entities.Post, error) {
	postQ := gormgen.Post
	userQ := gormgen.User
	fmt.Printf("postQ.Creator.GetSelects(): %v\n", postQ.Creator.GetSelects())
	posts, err := postQ.
		WithContext(ctx).
		// Select(postQ.ALL).
		Select(postQ.Creator.GetSelects()...).
		Where(postQ.CreatorID.Eq(id)).
		LeftJoin(userQ, postQ.CreatorID.EqCol(userQ.ID)).
		Offset(offset).
		Debug().
		Limit(limit).Find()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *ORMRepository) CountPostsByCreatorId(ctx context.Context, id string) (int64, error) {
	postQ := gormgen.Post
	count, err := postQ.
		WithContext(ctx).
		Where(postQ.CreatorID.Eq(id)).
		Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *ORMRepository) InsertPost(ctx context.Context, post *entities.Post) (*entities.Post, error) {
	postQ := gormgen.Post
	err := postQ.WithContext(ctx).Create(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
