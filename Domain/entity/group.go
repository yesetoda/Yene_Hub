package entity

import (
	"time"
)

// Group represents a learning or working group in the system
type Group struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name" gorm:"size:255"`
	ShortName   string   `json:"short_name" gorm:"size:255"`
	Description string   `json:"description" gorm:"size:255"`
	HOAID       *uint    `json:"hoa_id,omitempty"` // Head of Academy ID
	CountryID   *uint    `json:"country_id,omitempty"`
	Country     *Country `json:"country,omitempty" gorm:"foreignKey:CountryID"`

	// Relations
	Users         []User         `json:"users,omitempty" gorm:"foreignKey:GroupID"`
	HOAs          []HOA          `json:"hoas,omitempty" gorm:"foreignKey:GroupID"`
	Exercises     []Exercise     `json:"exercises,omitempty" gorm:"foreignKey:GroupID"`
	GoogleOAuths  []GoogleOAuth  `json:"google_oauths,omitempty" gorm:"foreignKey:GroupID"`
	Invites       []Invite       `json:"invites,omitempty" gorm:"foreignKey:GroupID"`
	SuperToGroups []SuperToGroup `json:"super_to_groups,omitempty" gorm:"foreignKey:GroupID"`
	GroupSessions []GroupSession `json:"group_sessions,omitempty" gorm:"foreignKey:GroupID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
