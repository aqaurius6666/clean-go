package main

import "github.com/aqaurius6666/clean-go/internal/repositories"

func CastRepositoryTo[T interface{}](r repositories.RepositoryImpl) T {
	repo, ok := r.(T)
	if !ok {
		panic("failed to cast repository")
	}
	return repo
}
