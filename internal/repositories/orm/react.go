package orm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm/gormgen"
)

func (s *ORMRepository) InsertReact(ctx context.Context, react *entities.React) (*entities.React, error) {
	reactQ := gormgen.React
	err := reactQ.WithContext(ctx).Create(react)
	if err != nil {
		return nil, err
	}
	return react, nil
}
