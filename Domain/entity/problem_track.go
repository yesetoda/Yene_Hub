package entity

import (
	"time"
)

// ProblemTrack represents a many-to-many relationship between problems and tracks
type ProblemTrack struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	ProblemID uint     `json:"problem_id"`
	Problem   *Problem `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
	TrackID   uint     `json:"track_id"`
	Track     *Track   `json:"track,omitempty" gorm:"foreignKey:TrackID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}