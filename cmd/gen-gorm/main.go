package main

import (
	"github.com/aqaurius6666/clean-go/internal/entities"
	"gorm.io/gen"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/repositories/orm/gen",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(entities.Post{}, entities.User{})

	g.Execute()
}
