package main

import (
	"github.com/aqaurius6666/clean-go/internal/entities"
	"gorm.io/gen"
)

type CommonQuerier interface {
	// select * from @@table where id = @id limit 1
	FindById(id string) (gen.T, error)
}

type UserQuerier interface {
	CommonQuerier

	// select * from @@table where username = @username limit 1
	FindByUsername(username string) (gen.T, error)
}

type PostQuerier interface {
	CommonQuerier
	// select * from @@table where user_id = @userId limit @limit offset @offset
	FindByUserId(userId string, limit, offset int64) ([]*gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/repositories/orm/gormgen",
		Mode:    gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
	})

	g.ApplyBasic(entities.User{}, entities.Post{})

	g.ApplyInterface(func(UserQuerier) {}, entities.User{})

	g.ApplyInterface(func(PostQuerier) {}, entities.Post{})

	g.Execute()
}
