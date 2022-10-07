package odm

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ODMRepository) InsertPost(ctx context.Context, ex gentity.Extend[*entities.Post]) (*entities.Post, error) {
	filter := make(map[string]interface{})
	oid, err := primitive.ObjectIDFromHex(ex.Entity.CreatorID)
	if err != nil {
		return nil, err
	}
	filter["_id"] = oid
	var updates struct {
		Post *entities.Post     `bson:",inline"`
		ID   primitive.ObjectID `bson:"_id"`
	}
	updates.Post = ex.Entity
	updates.ID = primitive.NewObjectID()
	res := s.DB.Collection("users").FindOneAndUpdate(ctx, filter, bson.M{"$push": bson.M{"posts": updates}})
	if err := res.Err(); err != nil {
		return nil, err
	}
	ex.Entity.ID = updates.ID.Hex()
	return ex.Entity, nil
}

func (s *ODMRepository) ListPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) ([]*entities.Post, error) {
	filter := make(map[string]interface{})
	projection := make(map[string]interface{})
	oid, err := primitive.ObjectIDFromHex(ex.Entity.CreatorID)
	if err != nil {
		return nil, err
	}
	filter["_id"] = oid

	posts := make([]*entities.Post, 0)
	projection["posts"] = 1
	if ex.ExFields.Limit != nil {
		projection["posts"] = bson.M{"$slice": bson.A{"$posts", *ex.ExFields.Limit}}
	}
	if ex.ExFields.Offset != nil {
		projection["posts"] = bson.M{"$slice": bson.A{"$posts", *ex.ExFields.Offset, 1e5}}
	}
	if ex.ExFields.Offset != nil && ex.ExFields.Limit != nil {
		projection["posts"] = bson.M{"$slice": bson.A{"$posts", *ex.ExFields.Offset, *ex.ExFields.Limit}}
	}
	res, err := s.DB.Collection("users").Aggregate(ctx, []bson.M{
		{"$match": filter},
		{"$project": projection},
		{"$unwind": bson.M{"path": "$posts"}},
		{"$replaceRoot": bson.M{"newRoot": "$posts"}},
	})
	if err != nil {
		return nil, err
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	if err := res.All(ctx, &posts); err != nil {
		return nil, err
	}
	for i, p := range posts {
		p.CreatorID = ex.Entity.CreatorID
		posts[i] = p
	}
	return posts, nil
}

func (s *ODMRepository) TotalPosts(ctx context.Context, ex gentity.Extend[*entities.Post]) (int64, error) {
	filter := make(map[string]interface{})
	oid, err := primitive.ObjectIDFromHex(ex.Entity.CreatorID)
	if err != nil {
		return 0, err
	}
	filter["_id"] = oid
	var ret []struct {
		Total int64 `bson:"total"`
	}
	res, err := s.DB.Collection("users").Aggregate(ctx, []bson.M{
		{"$match": filter},
		{"$unwind": bson.M{"path": "$posts"}},
		{"$replaceRoot": bson.M{"newRoot": "$posts"}},
		{"$count": "total"},
	})
	if err != nil {
		return 0, err
	}
	if err := res.Err(); err != nil {
		return 0, err
	}
	defer res.Close(ctx)
	if err := res.All(ctx, &ret); err != nil {
		return 0, err
	}
	return ret[0].Total, nil
}
