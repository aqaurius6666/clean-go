// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package odm

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

func BuildRepository(logger *logrus.Logger, cfg config.DBConfig) (*ODMRepository, error) {
	database, err := ConnectMongoDB(cfg)
	if err != nil {
		return nil, err
	}
	odmRepository := &ODMRepository{
		DB: database,
	}
	return odmRepository, nil
}