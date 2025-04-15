package schemas

import (
	"time"

	"a2sv.org/hub/Domain/entity"
)

// ErrorResponse represents a standardized error response
// swagger:model
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// PaginationMeta contains pagination information
// swagger:model
type PaginationMeta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
}

// PaginatedUsers represents paginated user results
// swagger:model
type PaginatedUsers struct {
	Data []*entity.User `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

type ResponseUser struct {
	// Core Identity Fields
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"size:255;not null"`
	Email     string `json:"email" gorm:"size:255;unique;not null"`
	RoleID    uint   `json:"role_id" gorm:"default:3;not null"` // Default role ID (adjust as necessary)
	GroupID   *uint  `json:"group_id,omitempty"`
	CountryID *uint  `json:"country_id,omitempty"`

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
}

// AuthTokenResponse represents authentication token response
// swagger:model
type AuthTokenResponse struct {
	Token string      `json:"token"`
	User  entity.User `json:"user"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type LoginResponse struct {
	Message string        `json:"message"`
	Token   string        `json:"token"`
	User    *ResponseUser `json:"user"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	
	Name                   *string    `json:"name,omitempty"`
	RoleID                 *uint      `json:"role_id,omitempty"`
	University             *string    `json:"university,omitempty"`
	StudentID              *string    `json:"student_id,omitempty"`
	Department             *string    `json:"department,omitempty"`
	ExpectedGraduationDate *time.Time `json:"expected_graduation_date,omitempty"`
	Phone                  *string    `json:"phone,omitempty"`
	TelegramUsername       *string    `json:"telegram_username,omitempty"`
	TelegramUID            *string    `json:"telegram_uid,omitempty"`
	Leetcode               *string    `json:"leetcode,omitempty"`
	Codeforces             *string    `json:"codeforces,omitempty"`
	Github                 *string    `json:"github,omitempty"`
	Hackerrank             *string    `json:"hackerrank,omitempty"`
	Linkedin               *string    `json:"linkedin,omitempty"`
	Instagram              *string    `json:"instagram,omitempty"`
	Birthday               *time.Time `json:"birthday,omitempty"`
	
}
type UpdateUserInput struct {
	
	Name                   *string    `json:"name,omitempty"`
	Email                  *string    `json:"email,omitempty"`
	RoleID                 *uint      `json:"role_id,omitempty"`
	University             *string    `json:"university,omitempty"`
	StudentID              *string    `json:"student_id,omitempty"`
	Department             *string    `json:"department,omitempty"`
	ExpectedGraduationDate *time.Time `json:"expected_graduation_date,omitempty"`
	Phone                  *string    `json:"phone,omitempty"`
	TelegramUsername       *string    `json:"telegram_username,omitempty"`
	TelegramUID            *string    `json:"telegram_uid,omitempty"`
	Leetcode               *string    `json:"leetcode,omitempty"`
	Codeforces             *string    `json:"codeforces,omitempty"`
	Github                 *string    `json:"github,omitempty"`
	Hackerrank             *string    `json:"hackerrank,omitempty"`
	Linkedin               *string    `json:"linkedin,omitempty"`
	Instagram              *string    `json:"instagram,omitempty"`
	Birthday               *time.Time `json:"birthday,omitempty"`
}