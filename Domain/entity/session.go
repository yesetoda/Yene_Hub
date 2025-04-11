package entity

import (
	"time"
)

// Session represents a learning/teaching session in the system
type Session struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"size:255"`
	Description     string    `json:"description,omitempty" gorm:"type:text"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	MeetLink        string    `json:"meet_link,omitempty" gorm:"size:255"`
	Location        string    `json:"location,omitempty" gorm:"size:255"`
	ResourceLink    string    `json:"resource_link,omitempty" gorm:"size:255"`
	RecordingLink   string    `json:"recording_link,omitempty" gorm:"size:255"`
	CalendarEventID string    `json:"calendar_event_id,omitempty" gorm:"size:255"`
	LecturerID      *uint     `json:"lecturer_id,omitempty"`
	Lecturer        *User     `json:"lecturer,omitempty" gorm:"foreignKey:LecturerID"`
	FundID          *uint     `json:"fund_id,omitempty"`
	Fund            *Fund     `json:"fund,omitempty" gorm:"foreignKey:FundID"`

	// Relations
	Attendances   []Attendance   `json:"attendances,omitempty" gorm:"foreignKey:SessionID"`
	GroupSessions []GroupSession `json:"group_sessions,omitempty" gorm:"foreignKey:SessionID"`
	Stipends      []Stipend      `json:"stipends,omitempty" gorm:"foreignKey:SessionID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
