package main

import (
	"github.com/aqaurius6666/clean-go/internal/entities"
	"gorm.io/gen"
)

type CommonQuerier interface {
}

type UserQuerier interface {
	CommonQuerier
}

type PostQuerier interface {
	CommonQuerier
}

type ReactQuerier interface {
	CommonQuerier
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/repositories/orm/gormgen",
		Mode:    gen.WithDefaultQuery | gen.WithoutContext | gen.WithQueryInterface,
	})

	g.ApplyBasic(entities.User{}, entities.Post{})

	g.ApplyInterface(func(UserQuerier) {}, entities.User{})
	g.ApplyInterface(func(PostQuerier) {}, entities.Post{})
	g.ApplyInterface(func(ReactQuerier) {}, entities.React{})

	g.Execute()
}
