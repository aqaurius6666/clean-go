package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name string    `json:"name"`
}

func (s *User) GetId() string {
	return s.ID.String()
}

func (*User) IsEntity() {}
