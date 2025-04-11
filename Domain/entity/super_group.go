package entity

import (
	"time"
)

// SuperGroup represents a higher-level grouping of multiple groups
type SuperGroup struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:255"`

	// Relations
	Tracks        []Track        `json:"tracks,omitempty" gorm:"foreignKey:SuperGroupID"`
	Contests      []Contest      `json:"contests,omitempty" gorm:"foreignKey:SuperGroupID"`
	DailyProblems []DailyProblem `json:"daily_problems,omitempty" gorm:"foreignKey:SuperGroupID"`
	SuperToGroups []SuperToGroup `json:"super_to_groups,omitempty" gorm:"foreignKey:SuperGroupID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
