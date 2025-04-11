package entity

import (
	"time"
)

// Track represents a learning track in the system
type Track struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	Name         string      `json:"name" gorm:"size:255"`
	Active       bool        `json:"active" gorm:"default:true"`
	SuperGroupID *uint       `json:"super_group_id,omitempty"`
	SuperGroup   *SuperGroup `json:"super_group,omitempty" gorm:"foreignKey:SuperGroupID"`

	// Relations
	Problems      []Problem      `json:"problems,omitempty" gorm:"foreignKey:TrackID"`
	Exercises     []Exercise     `json:"exercises,omitempty" gorm:"foreignKey:TrackID"`
	ProblemTracks []ProblemTrack `json:"problem_tracks,omitempty" gorm:"foreignKey:TrackID"`
	Comments      []Comment      `json:"comments,omitempty" gorm:"foreignKey:TrackID"`
	Votes         []Vote         `json:"votes,omitempty" gorm:"foreignKey:TrackID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
