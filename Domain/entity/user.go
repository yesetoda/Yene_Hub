package entity

import (
	"time"
)

// User represents a system user with their profile and account information.
type User struct {
	// Core Identity Fields
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"size:255;not null"`
	Email    string `json:"email" gorm:"size:255;unique;not null"`
	Password string `json:"-" gorm:"size:255;not null"` // Password not included in JSON responses

	// Role and Relationships
	RoleID    uint     `json:"role_id" gorm:"default:3;not null"` // Default role ID (adjust as necessary)
	Role      *Role    `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	GroupID   *uint    `json:"group_id,omitempty"`
	Group     *Group   `json:"group,omitempty" gorm:"foreignKey:GroupID"`
	CountryID *uint    `json:"country_id,omitempty"`
	Country   *Country `json:"country,omitempty" gorm:"foreignKey:CountryID"`

	// Academic Information
	University             string     `json:"university,omitempty" gorm:"size:255"`
	StudentID              string     `json:"student_id,omitempty" gorm:"size:255"`
	Department             string     `json:"department,omitempty" gorm:"size:255"`
	ExpectedGraduationDate *time.Time `json:"expected_graduation_date,omitempty"`

	// Contact Information
	Phone            string  `json:"phone,omitempty" gorm:"size:20"`
	TelegramUsername *string `json:"telegram_username,omitempty" gorm:"size:255;unique"`
	TelegramUID      string  `json:"telegram_uid,omitempty" gorm:"size:255"`

	// Coding Profiles (optional)
	// Use pointer types so that if no value is provided, they remain nil.
	Leetcode   *string `json:"leetcode,omitempty" gorm:"type:varchar(255)"`
	Codeforces *string `json:"codeforces,omitempty" gorm:"type:varchar(255)"`
	Github     *string `json:"github,omitempty" gorm:"type:varchar(255)"`
	Hackerrank *string `json:"hackerrank,omitempty" gorm:"type:varchar(255)"`

	// Social Media (all optional)
	Linkedin  string `json:"linkedin,omitempty" gorm:"size:255"`
	Instagram string `json:"instagram,omitempty" gorm:"size:255"`

	// Personal Information
	Birthday          *time.Time `json:"birthday,omitempty"`
	Gender            string     `json:"gender,omitempty" gorm:"size:50"`
	ShortBio          string     `json:"short_bio,omitempty" gorm:"type:text"`
	PreferredLanguage string     `json:"preferred_language,omitempty" gorm:"size:50"`

	// Professional Details
	CV         string     `json:"cv,omitempty" gorm:"size:255"`
	MentorName string     `json:"mentor_name,omitempty" gorm:"size:255"`
	JoinedDate *time.Time `json:"joined_date,omitempty"`

	// Physical Attributes
	TshirtColor string `json:"tshirt_color,omitempty" gorm:"size:50"`
	TshirtSize  string `json:"tshirt_size,omitempty" gorm:"size:10"`

	// System Fields
	Photo         string `json:"photo,omitempty" gorm:"size:255"`
	CodeOfConduct string `json:"code_of_conduct,omitempty" gorm:"size:255"`
	Inactive      bool   `json:"inactive" gorm:"default:false"`
	Config        string `json:"config,omitempty" gorm:"type:text"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations (using GORM associations)
	Submissions     []Submission   `json:"submissions,omitempty" gorm:"foreignKey:UserID"`
	Attendances     []Attendance   `json:"attendances,omitempty" gorm:"foreignKey:UserID"`
	HeadAttendances []Attendance   `json:"head_attendances,omitempty" gorm:"foreignKey:HeadID"`
	Comments        []Comment      `json:"comments,omitempty" gorm:"foreignKey:UserID"`
	Posts           []Post         `json:"posts,omitempty" gorm:"foreignKey:UserID"`
	RecentActions   []RecentAction `json:"recent_actions,omitempty" gorm:"foreignKey:UserID"`
	APITokens       []APIToken     `json:"api_tokens,omitempty" gorm:"foreignKey:UserID"`
}

