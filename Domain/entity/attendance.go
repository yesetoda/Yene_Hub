package entity

import (
	"time"
)

// Attendance represents a user's attendance record for a session
type Attendance struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	HeadID    uint      `json:"head_id"` // The head who recorded this attendance
	Head      *User     `json:"head,omitempty" gorm:"foreignKey:HeadID"`
	SessionID *uint     `json:"session_id,omitempty"`
	Session   *Session  `json:"session,omitempty" gorm:"foreignKey:SessionID"`
	Status    int       `json:"status"`         // e.g., 0=absent, 1=present, 2=excused
	Type      int       `json:"type,omitempty"` // Type of attendance record
	At        time.Time `json:"at"`             // Time when attendance was taken

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
