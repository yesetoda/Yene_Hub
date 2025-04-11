package entity

import "time"

type DivisionUser struct {
	ID uint `json:"id" gorm:"primary_key"`
	DivisionID uint `json:"division_id"`
	UserID uint `json:"user_id"`
	ContestID uint `json:"contest_id"`
	Active bool `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Division Division `json:"division" gorm:"foreignKey:DivisionID"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	Contest Contest `json:"contest" gorm:"foreignKey:ContestID"`
	
}