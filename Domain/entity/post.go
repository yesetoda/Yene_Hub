package entity

import (
	"time"
)

// Post represents a forum post in the system
type Post struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"user_id"`
	User   *User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Title  string `json:"title" gorm:"size:255"`
	Body   string `json:"body" gorm:"type:text"`

	// Relations
	Comments   []Comment   `json:"comments,omitempty" gorm:"foreignKey:PostID"`
	PostToTags []PostToTag `json:"post_to_tags,omitempty" gorm:"foreignKey:PostID"`
	Votes      []Vote      `json:"votes,omitempty" gorm:"foreignKey:PostID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
