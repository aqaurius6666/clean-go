package usecases

import (
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type Usecases interface {
	UserUsecases
	AuthUsecases
	PostUsecases
	ReactUsecases
}

type UsecasesService struct {
	Logger        *logrus.Logger
	AuthConfig    config.AuthConfig
	Repo          Repository
	TraceProvider trace.TracerProvider
}

var pkgName = "internal/usecases"
