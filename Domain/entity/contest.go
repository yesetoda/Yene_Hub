package entity

import (
	"time"
)

// Contest represents a programming contest in the system
type Contest struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	Name         string      `json:"name" gorm:"size:255"`
	Link         string      `json:"link" gorm:"size:255"`
	Link2        string      `json:"link2,omitempty" gorm:"size:255"`
	Link3        string      `json:"link3,omitempty" gorm:"size:255"`
	ProblemCount int         `json:"problem_count"`
	Unrated      bool        `json:"unrated" gorm:"default:false"`
	Type         string      `json:"type,omitempty" gorm:"size:255"`
	SuperGroupID *uint       `json:"super_group_id,omitempty"`
	SuperGroup   *SuperGroup `json:"super_group,omitempty" gorm:"foreignKey:SuperGroupID"`

	// Relations
	Problems []Problem `json:"problems,omitempty" gorm:"foreignKey:ContestID"`
	// Ratings       []Rating       `json:"ratings,omitempty" gorm:"foreignKey:ContestID"`
	// DivisionUsers []DivisionUser `json:"division_users,omitempty" gorm:"foreignKey:ContestID"`
	Votes []Vote `json:"votes,omitempty" gorm:"foreignKey:ContestID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
