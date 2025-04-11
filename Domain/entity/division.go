package entity

import "time"

type Division struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

