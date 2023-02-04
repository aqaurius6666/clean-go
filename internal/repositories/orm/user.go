package orm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/repositories/orm/gormgen"
)

func (s *ORMRepository) InsertUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	userQ := gormgen.User
	err := userQ.WithContext(ctx).Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *ORMRepository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (*entities.User, error) {
	userQ := gormgen.User
	user, err := userQ.WithContext(ctx).
		Where(userQ.Email.Eq(email), userQ.Password.Eq(password)).First()
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *ORMRepository) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	userQ := gormgen.User
	user, err := userQ.WithContext(ctx).Where(userQ.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *ORMRepository) UpdateUser(ctx context.Context, id string, user *entities.User) (*entities.User, error) {
	userQ := gormgen.User
	_, err := userQ.WithContext(ctx).Where(userQ.ID.Eq(id)).Updates(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
