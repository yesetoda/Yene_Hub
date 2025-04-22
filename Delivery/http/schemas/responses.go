package schemas

import (
	"time"
)

// PaginationMeta contains pagination information
// swagger:model

// PaginatedUsers represents paginated user results
// swagger:model
type PaginatedUsers struct {
	Data []*ResponseUser `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

// ResponseUser represents a detailed user response
// swagger:model
type ResponseUser struct {
	// Core Identity Fields
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"user@example.com"`
	RoleID    uint      `json:"role_id" example:"2"`
	GroupID   *uint     `json:"group_id,omitempty" example:"1"`
	CountryID *uint     `json:"country_id,omitempty" example:"1"`

	// Academic Information
	University             *string    `json:"university,omitempty" example:"Example University"`
	StudentID             *string    `json:"student_id,omitempty" example:"STU123"`
	Department            *string    `json:"department,omitempty" example:"Computer Science"`
	ExpectedGraduationDate *time.Time `json:"expected_graduation_date,omitempty"`

	// Contact Information
	Phone            *string `json:"phone,omitempty" example:"+1234567890"`
	TelegramUsername *string `json:"telegram_username,omitempty" example:"@username"`
	TelegramUID      *string `json:"telegram_uid,omitempty"`

	// Coding Profiles
	Leetcode   *string `json:"leetcode,omitempty" example:"leetcode_user"`
	Codeforces *string `json:"codeforces,omitempty" example:"cf_user"`
	Github     *string `json:"github,omitempty" example:"github_user"`
	Hackerrank *string `json:"hackerrank,omitempty" example:"hr_user"`

	// Social Media
	Linkedin  *string `json:"linkedin,omitempty" example:"linkedin_profile"`
	Instagram *string `json:"instagram,omitempty" example:"insta_user"`

	// Personal Information
	Birthday          *time.Time `json:"birthday,omitempty"`
	Gender            *string    `json:"gender,omitempty" example:"male"`
	ShortBio         *string    `json:"short_bio,omitempty"`
	PreferredLanguage *string    `json:"preferred_language,omitempty" example:"en"`

	// Professional Details
	CV         *string    `json:"cv,omitempty"`
	MentorName *string    `json:"mentor_name,omitempty"`
	JoinedDate *time.Time `json:"joined_date,omitempty"`

	// Physical Attributes
	TshirtColor *string `json:"tshirt_color,omitempty"`
	TshirtSize  *string `json:"tshirt_size,omitempty"`

	// System Fields
	Photo         *string `json:"photo,omitempty"`
	CodeOfConduct *string `json:"code_of_conduct,omitempty"`
	Inactive      *bool   `json:"inactive" example:"false"`
	Config        *string `json:"config,omitempty"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PaginationMeta represents metadata for paginated results
// swagger:model