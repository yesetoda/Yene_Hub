package entity

import (
	"time"
)

// Submission represents a user's submission for a problem
type Submission struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	ProblemID uint     `json:"problem_id"`
	Problem   *Problem `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
	UserID    uint     `json:"user_id"`
	User      *User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	TimeSpent int      `json:"time_spent,omitempty"` // Time spent in seconds
	Tries     int      `json:"tries,omitempty"`      // Number of attempts
	InContest int      `json:"in_contest,omitempty"` // Whether it was solved in a contest
	Code      string   `json:"code,omitempty" gorm:"type:text"`
	Language  string   `json:"language,omitempty" gorm:"size:255"`
	Verified  bool     `json:"verified"`

	// Relations
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:SubmissionID"`
	Votes    []Vote    `json:"votes,omitempty" gorm:"foreignKey:SubmissionID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

