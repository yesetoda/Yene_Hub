package entity

import (
	"time"
)

type AssistantMessage struct {
	ID uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	Question string `json:"question" gorm:"type:text"`
	Answer string `json:"answer" gorm:"type:text"`
	Feedback string `json:"feedback" gorm:"type:text"`
	Session string `json:"session" gorm:"size:255"`
	RawResponse string `json:"raw_response" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
