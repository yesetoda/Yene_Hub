package entity

import (
	"time"
)

// SuperToGroup represents a many-to-many relationship between super groups and groups
type SuperToGroup struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	GroupID      uint        `json:"group_id"`
	Group        *Group      `json:"group,omitempty" gorm:"foreignKey:GroupID"`
	SuperGroupID uint        `json:"super_group_id"`
	SuperGroup   *SuperGroup `json:"super_group,omitempty" gorm:"foreignKey:SuperGroupID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
