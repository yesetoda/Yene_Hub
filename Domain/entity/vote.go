package entity

import (
	"time"
)

// Vote represents a user vote on various entities (like/upvote/downvote)
type Vote struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	UserID       uint        `json:"user_id"`
	User         *User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CommentID    *uint       `json:"comment_id,omitempty"`
	Comment      *Comment    `json:"comment,omitempty" gorm:"foreignKey:CommentID"`
	PostID       *uint       `json:"post_id,omitempty"`
	Post         *Post       `json:"post,omitempty" gorm:"foreignKey:PostID"`
	ProblemID    *uint       `json:"problem_id,omitempty"`
	Problem      *Problem    `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
	TrackID      *uint       `json:"track_id,omitempty"`
	Track        *Track      `json:"track,omitempty" gorm:"foreignKey:TrackID"`
	ContestID    *uint       `json:"contest_id,omitempty"`
	Contest      *Contest    `json:"contest,omitempty" gorm:"foreignKey:ContestID"`
	SubmissionID *uint       `json:"submission_id,omitempty"`
	Submission   *Submission `json:"submission,omitempty" gorm:"foreignKey:SubmissionID"`
	Type         int         `json:"type"` // e.g., 1=upvote, -1=downvote

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

