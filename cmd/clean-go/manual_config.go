package main

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type ManualConfig struct {
	Tracer trace.TracerProvider
	Logger *logrus.Logger
}
