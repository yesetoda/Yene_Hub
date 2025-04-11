package entity

import (
	"time"
)

// Invite represents an invitation for a user to join a group with a specific role
type Invite struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Key     string `json:"key" gorm:"size:255"` // Unique invite key
	RoleID  uint   `json:"role_id"`
	Role    *Role  `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	UserID  *uint  `json:"user_id,omitempty"` // User who created the invite
	User    *User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	GroupID *uint  `json:"group_id,omitempty"`
	Group   *Group `json:"group,omitempty" gorm:"foreignKey:GroupID"`
	Used    bool   `json:"used" gorm:"default:false"` // Whether the invite has been used

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
