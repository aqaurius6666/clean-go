package orm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm/gormgen"
	"gorm.io/gorm"
)

func (s *ORMRepository) UpsertReact(ctx context.Context, react *entities.React) (*entities.React, error) {
	reactQ := gormgen.React
	reactInDB, err := reactQ.WithContext(ctx).Where(reactQ.PostID.Eq(react.PostID), reactQ.UserID.Eq(react.UserID)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = reactQ.WithContext(ctx).Create(react)
			if err != nil {
				return nil, err
			}
			return react, nil
		}
	}

	if reactInDB.Type == react.Type {
		return reactInDB, nil
	}
	_, err = reactQ.WithContext(ctx).Where(reactQ.PostID.Eq(react.PostID), reactQ.UserID.Eq(react.UserID)).Update(reactQ.Type, react.Type)
	if err != nil {
		return nil, err
	}
	return react, nil
}

func (s *ORMRepository) GetReactsByPostId(ctx context.Context, postId string) ([]*entities.React, error) {
	reactQ := gormgen.React
	reacts, err := reactQ.WithContext(ctx).
		Select(reactQ.PostID, reactQ.Type, reactQ.PostID.Count().As("count")).
		Where(reactQ.PostID.Eq(postId)).
		Group(reactQ.Type).
		Debug().
		Find()
	if err != nil {
		return nil, err
	}
	return reacts, nil
}

func (s *ORMRepository) GetReactsByPostIds(ctx context.Context, ids []string) ([]*entities.React, error) {
	reactQ := gormgen.React
	reacts, err := reactQ.WithContext(ctx).
		Select(reactQ.PostID, reactQ.Type, reactQ.PostID.Count().As("count")).
		Where(reactQ.PostID.In(ids...)).
		Group(reactQ.PostID, reactQ.Type).
		Debug().
		Find()
	if err != nil {
		return nil, err
	}
	return reacts, nil
}
