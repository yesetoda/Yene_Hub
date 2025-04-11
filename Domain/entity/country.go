package entity

import (
	"time"
)

// Country represents a country in the system
type Country struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"size:255"`
	ShortCode string `json:"short_code" gorm:"size:255"`

	// Relations
	Users  []User  `json:"users,omitempty" gorm:"foreignKey:CountryID"`
	Groups []Group `json:"groups,omitempty" gorm:"foreignKey:CountryID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
