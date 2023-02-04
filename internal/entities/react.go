package entities

type React struct {
	UserID string    `gorm:"type:uuid;index:idx_react_user_id_post_id" json:"user_id" bson:"-"`
	User   *User     `gorm:"foreignKey:UserID" json:"user" bson:"-"`
	PostID string    `gorm:"type:uuid;index:idx_react_user_id_post_id" json:"post_id" bson:"-"`
	Post   *Post     `gorm:"foreignKey:PostID" json:"post" bson:"-"`
	Type   ReactType `gorm:"type:varchar(12)" json:"type" bson:"type,omitempty"`
}

func (*React) IsEntity() {}

func (r *React) SetId(id interface{}) {
	// switch t := id.(type) {
	// case string:
	// 	r.ID = t
	// 	// case uuid.UUID:
	// 	// 	s.ID = t
	// }
}

type ReactType string

const (
	ReactLike    ReactType = "LIKE"
	ReactDislike ReactType = "DISLIKE"
)
