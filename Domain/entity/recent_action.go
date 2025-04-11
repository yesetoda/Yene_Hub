package entity

import (
	"time"
)

// RecentAction represents a recent action performed by a user
type RecentAction struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	User        *User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Type        string `json:"type" gorm:"size:255"`
	Description string `json:"description" gorm:"size:255"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}