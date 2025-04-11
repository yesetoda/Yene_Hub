package entity

import (
	"time"
)

// GoogleOAuth represents a Google OAuth connection for a user
type GoogleOAuth struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	UserID              uint      `json:"user_id"`
	User                *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	GroupID             uint      `json:"group_id"`
	Group               *Group    `json:"group,omitempty" gorm:"foreignKey:GroupID"`
	EncryptedTokenString string    `json:"encrypted_token_string" gorm:"type:text"`
	CalendarID          string    `json:"calendar_id" gorm:"size:255"`
	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
