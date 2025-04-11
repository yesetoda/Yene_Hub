package entity

import (
	"time"
)

// Stipend represents a payment to a user from a fund
type Stipend struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	FundID    uint     `json:"fund_id"`
	Fund      *Fund    `json:"fund,omitempty" gorm:"foreignKey:FundID"`
	UserID    uint     `json:"user_id"`
	User      *User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	SessionID *uint    `json:"session_id,omitempty"`
	Session   *Session `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	Paid      bool     `json:"paid" gorm:"default:false"`
	Share     float64  `json:"share,omitempty"` // Share of the fund amount

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
