package cql

import (
	"context"
	"strings"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
)

func buildPostsArgs(ex gentity.Extend[*entities.Post]) ([]string, []interface{}) {
	args := make([]interface{}, 0)
	stmt := make([]string, 0)
	if ex.Entity.ID != "" {
		stmt = append(stmt, "id = ?")
		args = append(args, ex.Entity.ID)
	}
	if ex.Entity.CreatorID != "" {
		stmt = append(stmt, "creator_id = ?")
		args = append(args, ex.Entity.CreatorID)
	}
	if ex.Entity.Title != "" {
		stmt = append(stmt, "title = ?")
		args = append(args, ex.Entity.Title)
	}
	if ex.Entity.Content != "" {
		stmt = append(stmt, "content = ?")
		args = append(args, ex.Entity.Content)
	}
	return stmt, args
}

func buildPaginationArgs(ex gentity.Extend[*entities.Post]) ([]string, []interface{}) {
	args := make([]interface{}, 0)
	stmt := make([]string, 0)
	if ex.ExFields.Limit != nil {
		stmt = append(stmt, "limit ?")
		args = append(args, *ex.ExFields.Limit)
	}
	if ex.ExFields.Offset != nil {
		stmt = append(stmt, "offset ?")
		args = append(args, *ex.ExFields.Offset)
	}
	return stmt, args
}

// GetUserById(ctx context.Context, id string) (*entities.User, error)
func (s *CQLRepository) ListPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) ([]*entities.Post, error) {
	wheres, args := buildPostsArgs(ex)
	paging, args2 := buildPaginationArgs(ex)
	args = append(args, args2...)
	q := s.Db.Query(`
	select * from posts
	where `+strings.Join(wheres, " and ")+
		strings.Join(paging, " ")+
		` allow filtering`, args...)
	s.Logger.Debug(q.String())
	iter := q.Iter()
	ret := make([]*entities.Post, iter.NumRows())
	scanner := iter.Scanner()
	for i := 0; scanner.Next(); i++ {
		ret[i] = &entities.Post{}
		if err := scanner.Scan(&ret[i].ID, &ret[i].Title, &ret[i].Content, &ret[i].CreatorID); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (s *CQLRepository) TotalPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) (int64, error) {
	var ret int64 = 0
	wheres, args := buildPostsArgs(ex)
	q := s.Db.Query(`
	select count(id) from posts
	where `+strings.Join(wheres, " and ")+
		` allow filtering`, args...)
	s.Logger.Debug(q.String())
	if err := q.Scan(&ret); err != nil {
		return 0, err
	}
	return ret, nil
}

// SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
// UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error)
func (s *CQLRepository) InsertPost(ctx context.Context, ex gentity.Extend[*entities.Post]) (*entities.Post, error) {
	ex.Entity.ID = uuid.New().String()
	q := s.Db.Query(`
	insert into posts (id, title, content, creator_id)
		VALUES (?, ?, ?, ?)`, ex.Entity.ID, ex.Entity.Title, ex.Entity.Content, ex.Entity.CreatorID)
	s.Logger.Debug(q.String())
	if err := q.Exec(); err != nil {
		return nil, err
	}
	return ex.Entity, nil
}
