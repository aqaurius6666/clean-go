package es

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/nats-io/nats.go"
)

type ESClient struct {
	NC *nats.Conn
	JS nats.JetStreamContext
}

func NewESClient(cfg config.DBConfig) (*ESClient, error) {
	if cfg.DSN == "" {
		cfg.DSN = (&url.URL{
			Scheme: cfg.Scheme,
			User:   url.UserPassword(cfg.User, cfg.Pass),
			Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		}).String()
	}

	nc, err := nats.Connect(cfg.DSN)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	return &ESClient{
		NC: nc,
		JS: js,
	}, nil
}

func (s *ESClient) Close() error {
	return nil
}

func (s *ESClient) Migrate(_ context.Context) error {
	s.JS.AddStream(&nats.StreamConfig{
		Name:     "users",
		Subjects: []string{"users-*"},
	})

	return nil
}
