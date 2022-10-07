package es

import (
	"context"
	"encoding/json"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/google/uuid"
)

type userEventType string

var (
	userCreated userEventType = "user-created"
	userUpdated userEventType = "user-updated"
	userDeleted userEventType = "user-deleted"
)

type userEvent struct {
	EventName userEventType `json:"eventName"`
	Email     string        `json:"email"`
	Name      string        `json:"name"`
	Password  string        `json:"password"`
	ID        string        `json:"id"`
}

func reduce(user *entities.User, event *userEvent) *entities.User {
	switch event.EventName {
	case userCreated:
		user.ID = event.ID
		user.Email = event.Email
		user.Name = event.Name
		user.Password = event.Password
	case userUpdated:
		user.Email = event.Email
		user.Name = event.Name
		user.Password = event.Password
	case userDeleted:
		user = nil
	}
	return user
}

func (s *ESClient) aggregate(ctx context.Context, id string) (*entities.User, error) {
	// consumer, err := s.JS.AddConsumer("users-"+id, nats.Durable("users-"+id), nats.DeliverAllAvailable())
	// sub.Unsubscribe()
	// s.NC.Flush()
	// user := &entities.User{}
	// events := make([]*userEvent, 0)
	// for {
	// 	msg, err := reader.FetchMessage(ctx)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		return nil, err
	// 	}

	// 	var e userEvent
	// 	if err := json.Unmarshal(msg.Value, &e); err != nil {
	// 		return nil, err
	// 	}
	// 	events = append(events, &e)
	// }
	// for _, event := range events {
	// 	user = reduce(user, event)
	// }
	return nil, nil
}

func (s *ESClient) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ESClient) ListUsers(ctx context.Context, ex gentity.Extend[*entities.User]) ([]*entities.User, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ESClient) SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	if ex.Entity.ID != "" {
		return s.aggregate(ctx, ex.Entity.ID)
	}
	return nil, nil
}

func (s *ESClient) InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error) {
	ex.Entity.ID = uuid.New().String()
	var event userEvent
	event.Email = ex.Entity.Email
	event.Name = ex.Entity.Name
	event.Password = ex.Entity.Password
	event.ID = ex.Entity.ID
	event.EventName = userCreated
	bz, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	_, err = s.JS.Publish("users-"+ex.Entity.ID, bz)
	if err != nil {
		return nil, err
	}
	return ex.Entity, nil
}

func (s *ESClient) UpdateUser(ctx context.Context, ex gentity.Extend[*entities.User], v *entities.User) (*entities.User, error) {
	panic("not implemented") // TODO: Implement
}

// GetUserById(ctx context.Context, id string) (*entities.User, error)
