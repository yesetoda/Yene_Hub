package entity

import (
	"time"
)

// HOA represents a Head of Academy for a group
type HOA struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	UserID  uint   `json:"user_id"`
	User    *User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	GroupID uint   `json:"group_id"`
	Group   *Group `json:"group,omitempty" gorm:"foreignKey:GroupID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

