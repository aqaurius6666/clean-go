package es

import (
	"context"
)

func (s *ESClient) RegisterViewConsumer(ctx context.Context) error {
	s.registerUserConsumer(ctx)
	return nil
}

func (s *ESClient) registerUserConsumer(ctx context.Context) error {
	// reader := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers:     s.Brokers,
	// 	Topic:       "users",
	// 	StartOffset: 0,
	// 	MinBytes:    1e3, // 1KB
	// 	MaxWait:     1 * time.Second,
	// 	GroupTopics: ,
	// })

	return nil
}
