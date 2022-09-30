package entities

import (
	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title     string    `json:"title"`
	CreatorID uuid.UUID `gorm:"type:uuid" json:"creator_id"`
	Creator   *User     `gorm:"foreignKey:CreatorID" json:"creator"`
}

func (s *Post) SetId(id interface{}) {
	switch t := id.(type) {
	case string:
		s.ID = uuid.MustParse(t)
	case uuid.UUID:
		s.ID = t
	}
}

func (*Post) IsEntity() {}
