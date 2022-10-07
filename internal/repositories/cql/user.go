package cql

import (
	"context"
	"strings"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
)

func buildUserArgs(ex gentity.Extend[*entities.User]) ([]string, []interface{}) {
	args := make([]interface{}, 0)
	stmt := make([]string, 0)
	if ex.Entity.ID != "" {
		stmt = append(stmt, "id = ?")
		args = append(args, ex.Entity.ID)
	}
	if ex.Entity.Email != "" {
		stmt = append(stmt, "email = ?")
		args = append(args, ex.Entity.Email)
	}
	if ex.Entity.Name != "" {
		stmt = append(stmt, "name = ?")
		args = append(args, ex.Entity.Name)
	}
	if ex.Entity.Password != "" {
		stmt = append(stmt, "password = ?")
		args = append(args, ex.Entity.Password)
	}
	return stmt, args
}
func (s *CQLRepository) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	q := s.Db.Query(`
	select * from users
	where id = ?`, id)
	s.Logger.Debug(q.String())
	user := &entities.User{}
	err := q.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *CQLRepository) ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *CQLRepository) SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	wheres, args := buildUserArgs(ex)
	user := &entities.User{}
	q := s.Db.Query(`
	select * from users
	where `+strings.Join(wheres, " and ")+
		` allow filtering`, args...)
	s.Logger.Debug(q.String())
	err := q.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *CQLRepository) InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	ex.Entity.ID = uuid.New().String()
	q := s.Db.Query(`
	insert into users (id, email, name, password)
		VALUES (?, ?, ?, ?)`, ex.Entity.ID, ex.Entity.Name, ex.Entity.Email, ex.Entity.Password)
	s.Logger.Debug(q.String())
	if err := q.Exec(); err != nil {
		return nil, err
	}
	return ex.Entity, nil
}

func (s *CQLRepository) UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error) {
	updates, args := buildUserArgs(gentity.Extend[*entities.User]{Entity: v})
	args = append(args, ex.Entity.ID)
	q := s.Db.Query(`
	UPDATE users
	SET `+strings.Join(updates, " , ")+`
	WHERE id = ?`, args...)
	s.Logger.Debug(q.String())
	err := q.Exec()
	if err != nil {
		return nil, err
	}
	return v, nil
}
