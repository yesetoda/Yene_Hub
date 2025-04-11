package entity

import (
	"time"
)

// GroupSession represents a many-to-many relationship between groups and sessions
type GroupSession struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	GroupID   uint     `json:"group_id"`
	Group     *Group   `json:"group,omitempty" gorm:"foreignKey:GroupID"`
	SessionID uint     `json:"session_id"`
	Session   *Session `json:"session,omitempty" gorm:"foreignKey:SessionID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
