package entity

import (
	"time"
)

// Fund represents a monetary fund in the system
type Fund struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"size:255"`
	Description string  `json:"description,omitempty" gorm:"type:text"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency" gorm:"size:255"`

	// Relations
	Sessions []Session `json:"sessions,omitempty" gorm:"foreignKey:FundID"`
	Stipends []Stipend `json:"stipends,omitempty" gorm:"foreignKey:FundID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
