package entity

import (
	"time"
)

// Comment represents a user comment on various entities
type Comment struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	UserID       uint        `json:"user_id"`
	User         *User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	PostID       *uint       `json:"post_id,omitempty"`
	Post         *Post       `json:"post,omitempty" gorm:"foreignKey:PostID"`
	SubmissionID *uint       `json:"submission_id,omitempty"`
	Submission   *Submission `json:"submission,omitempty" gorm:"foreignKey:SubmissionID"`
	ProblemID    *uint       `json:"problem_id,omitempty"`
	Problem      *Problem    `json:"problem,omitempty" gorm:"foreignKey:ProblemID"`
	TrackID      *uint       `json:"track_id,omitempty"`
	Track        *Track      `json:"track,omitempty" gorm:"foreignKey:TrackID"`
	ReplyID      *uint       `json:"reply_id,omitempty"`
	Reply        *Comment    `json:"reply,omitempty" gorm:"foreignKey:ReplyID"`
	Text         string      `json:"text" gorm:"type:text"`

	// Relations
	Replies []Comment `json:"replies,omitempty" gorm:"foreignKey:ReplyID"`
	Votes   []Vote    `json:"votes,omitempty" gorm:"foreignKey:CommentID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}