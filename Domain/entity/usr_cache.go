package entity

import "time"

type UserCache struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Email     string    `json:"email" gorm:"size:255;unique;not null"`
	Password  string    `json:"-" gorm:"size:255;not null"`
	RoleID    uint      `json:"role_id" gorm:"default:3;not null"`
	Inactive  bool      `json:"inactive" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
