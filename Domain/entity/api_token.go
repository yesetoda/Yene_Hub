package entity

import (
	"time"
)

// APIToken represents an API token for user authentication
type APIToken struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"user_id"`
	User      *User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Name      string     `json:"name" gorm:"size:255"`
	Type      string     `json:"type" gorm:"size:255"`
	Token     string     `json:"token" gorm:"size:64"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
