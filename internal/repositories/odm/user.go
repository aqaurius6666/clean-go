package odm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/var/e"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getUserFilter(ex gentity.Extend[*entities.User]) (interface{}, error) {
	filter := bson.M{}
	if ex.Entity.ID != "" {
		oid, err := primitive.ObjectIDFromHex(ex.Entity.ID)
		if err != nil {
			return nil, err
		}
		filter["_id"] = oid
	}
	if ex.Entity.Name != "" {
		filter["name"] = ex.Entity.Name
	}
	if ex.Entity.Email != "" {
		filter["email"] = ex.Entity.Email
	}
	if ex.Entity.Password != "" {
		filter["password"] = ex.Entity.Password
	}
	return filter, nil
}
func (s *ODMRepository) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res := s.DB.Collection("users").FindOne(ctx, bson.M{"_id": objID})
	if err := res.Err(); err != nil {
		return nil, err
	}
	user := new(entities.User)
	if err := res.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *ODMRepository) ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error) {
	filter, err := getUserFilter(ex)
	if err != nil {
		return nil, err
	}
	cur, err := s.DB.Collection("users").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var users []*entities.User
	for cur.Next(ctx) {
		user := new(entities.User)
		err := cur.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *ODMRepository) SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	filter, err := getUserFilter(ex)
	if err != nil {
		return nil, err
	}
	res := s.DB.Collection("users").FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		return nil, err
	}
	if err := res.Decode(ex.Entity); err != nil {
		return nil, err
	}
	return ex.Entity, nil
}

func (s *ODMRepository) InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	res, err := s.DB.Collection("users").InsertOne(ctx, ex.Entity) // TODO: Implement
	if err != nil {
		return nil, err
	}
	ex.Entity.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return ex.Entity, nil
}

func (s *ODMRepository) UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error) {
	filter, err := getUserFilter(ex)
	if err != nil {
		return nil, err
	}
	res, err := s.DB.Collection("users").UpdateOne(ctx, filter, bson.M{"$set": v})
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		return nil, e.ErrInvalidOperation
	}
	return v, nil
}
