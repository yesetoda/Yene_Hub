package entity

import (
	"time"
)

type AssistantMessage struct {
	id uint `json:"id" gorm:"primaryKey"`
	user_id uint `json:"user_id"`
	user User `json:"user" gorm:"foreignKey:user_id"`
	question string `json:"question" gorm:"type:text"`
	answer string `json:"answer" gorm:"type:text"`
	feedback string `json:"feedback" gorm:"type:text"`
	session string `json:"session" gorm:"size:255"`
	raw_response string `json:"raw_response" gorm:"type:text"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}
