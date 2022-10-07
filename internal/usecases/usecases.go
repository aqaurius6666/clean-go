package usecases

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/sirupsen/logrus"
)

type Usecases interface {
	UserUsecases
	AuthUsecases
	PostUsecases
}

type UsecasesService struct {
	Logger     *logrus.Logger
	AuthConfig config.AuthConfig
	Repo       Repository
}
