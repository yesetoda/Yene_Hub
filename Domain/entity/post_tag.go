package entity

import (
	"time"
)

// PostTag represents a tag that can be applied to posts
type PostTag struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:255"`

	// Relations
	PostToTags []PostToTag `json:"post_to_tags,omitempty" gorm:"foreignKey:PostTagID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
