package entity

import (
	"time"
)

// Exercise represents a problem assigned to a track and group
type Exercise struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	TrackID   uint     `json:"track_id"`
	Track     *Track   `json:"track,omitempty" gorm:"foreignKey:TrackID"`
	ProblemID uint     `json:"problem_id"`
	Problem   *Problem `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
	GroupID   uint     `json:"group_id"`
	Group     *Group   `json:"group,omitempty" gorm:"foreignKey:GroupID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
