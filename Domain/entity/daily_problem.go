package entity

import (
	"time"
)

// DailyProblem represents a problem assigned for a specific date
type DailyProblem struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	ProblemID    uint        `json:"problem_id"`
	Problem      *Problem    `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
	SuperGroupID uint        `json:"super_group_id"`
	SuperGroup   *SuperGroup `json:"super_group,omitempty" gorm:"foreignKey:SuperGroupID"`
	ForDate      time.Time   `json:"for_date"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
