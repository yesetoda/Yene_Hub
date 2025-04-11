package entity

import (
	"time"
)

// PostToTag represents a many-to-many relationship between posts and tags
type PostToTag struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	PostID    uint     `json:"post_id"`
	Post      *Post    `json:"post,omitempty" gorm:"foreignKey:PostID"`
	PostTagID uint     `json:"post_tag_id"`
	PostTag   *PostTag `json:"post_tag,omitempty" gorm:"foreignKey:PostTagID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
