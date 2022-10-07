package entities

type User struct {
	ID       string  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id" bson:"_id,omitempty"`
	Name     string  `json:"name" gorm:"not null" bson:"name,omitempty"`
	Email    string  `json:"email" gorm:"unique;not null" bson:"email,omitempty"`
	Password string  `json:"password" gorm:"not null" bson:"password,omitempty"`
	Posts    []*Post `gorm:"-:all" json:"posts,omitempty" bson:"posts,omitempty"`
}

func (s *User) SetId(id interface{}) {
	switch t := id.(type) {
	case string:
		s.ID = t
		// case uuid.UUID:
		// 	s.ID = t
	}
}

func (*User) IsEntity() {}
