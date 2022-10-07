package entities

type Post struct {
	ID        string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title,omitempty"`
	Content   string `json:"content" bson:"content,omitempty"`
	CreatorID string `gorm:"type:uuid" json:"creator_id" bson:"-"`
	Creator   *User  `gorm:"foreignKey:CreatorID" json:"creator" bson:"-"`
}

func (s *Post) SetId(id interface{}) {
	switch t := id.(type) {
	case string:
		s.ID = t

	}
}

func (*Post) IsEntity() {}
