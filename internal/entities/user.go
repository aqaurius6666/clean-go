package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name string    `json:"name"`
}

func (s *User) SetId(id interface{}) {
	switch t := id.(type) {
	case string:
		s.ID = uuid.MustParse(t)
	case uuid.UUID:
		s.ID = t
	}
}

func (*User) IsEntity() {}
