package entity

import (
	"time"
)

// Problem represents a programming problem/challenge in the system
type Problem struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	ContestID  *uint    `json:"contest_id,omitempty"`
	Contest    *Contest `json:"contest,omitempty" gorm:"foreignKey:ContestID"`
	TrackID    *uint    `json:"track_id,omitempty"`
	Track      *Track   `json:"track,omitempty" gorm:"foreignKey:TrackID"`
	Name       string   `json:"name" gorm:"size:255"`
	Difficulty string   `json:"difficulty" gorm:"size:255"`
	Tag        string   `json:"tag" gorm:"size:255"`
	Platform   string   `json:"platform" gorm:"size:255"`
	Link       string   `json:"link" gorm:"size:255"`

	// Relations
	Submissions   []Submission   `json:"submissions,omitempty" gorm:"foreignKey:ProblemID"`
	DailyProblems []DailyProblem `json:"daily_problems,omitempty" gorm:"foreignKey:ProblemID"`
	Exercises     []Exercise     `json:"exercises,omitempty" gorm:"foreignKey:ProblemID"`
	ProblemTracks []ProblemTrack `json:"problem_tracks,omitempty" gorm:"foreignKey:ProblemID"`
	Comments      []Comment      `json:"comments,omitempty" gorm:"foreignKey:ProblemID"`
	Votes         []Vote         `json:"votes,omitempty" gorm:"foreignKey:ProblemID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
