package entity

import (
	"time"
)

// Role represents a user role in the system (e.g., student, admin, etc.)
type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Type string `json:"type" gorm:"size:255"`

	// Relations
	Users   []User   `json:"users,omitempty" gorm:"foreignKey:RoleID"`
	Invites []Invite `json:"invites,omitempty" gorm:"foreignKey:RoleID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

