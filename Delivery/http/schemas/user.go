package schemas

import "time"

// CreateUserRequest represents the request body for creating a new user
// swagger:model
// All fields required for creation, except those that are optional in the entity.
// (Already refined in previous steps)
type CreateUserRequest struct {
	Email    string  `json:"email" binding:"required,email" example:"user@example.com"`
	Password string  `json:"password" binding:"required" example:"MySecret123"`
	Name     string  `json:"name" binding:"required" example:"John Doe"`
	RoleID   *uint   `json:"role_id,omitempty" example:"2"`
	GroupID  *uint   `json:"group_id,omitempty" example:"1"`
	CountryID *uint  `json:"country_id,omitempty" example:"1"`

	University             *string    `json:"university,omitempty" example:"Example University"`
	StudentID              *string    `json:"student_id,omitempty" example:"STU123"`
	Department             *string    `json:"department,omitempty" example:"Computer Science"`
	ExpectedGraduationDate *time.Time `json:"expected_graduation_date,omitempty"`

	Phone            *string `json:"phone,omitempty" example:"+1234567890"`
	TelegramUsername *string `json:"telegram_username,omitempty" example:"@username"`
	TelegramUID      *string `json:"telegram_uid,omitempty" example:"123456789"`

	Leetcode   *string `json:"leetcode,omitempty" example:"leetcode_user"`
	Codeforces *string `json:"codeforces,omitempty" example:"cf_user"`
	Github     *string `json:"github,omitempty" example:"github_user"`
	Hackerrank *string `json:"hackerrank,omitempty" example:"hr_user"`
	Linkedin   *string `json:"linkedin,omitempty" example:"linkedin_profile"`
	Instagram  *string `json:"instagram,omitempty" example:"insta_user"`

	Birthday          *time.Time `json:"birthday,omitempty"`
	Gender            *string    `json:"gender,omitempty" example:"male"`
	ShortBio          *string    `json:"short_bio,omitempty" example:"Software developer passionate about algorithms"`
	PreferredLanguage *string    `json:"preferred_language,omitempty" example:"en"`
}

// UpdateUserRequest represents the request body for updating a user
// swagger:model
// (Already refined in previous steps)
type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty" example:"John Doe"`
	Email    *string `json:"email,omitempty" example:"user@example.com"`
	Password *string `json:"password,omitempty" example:"MySecret123"`
	RoleID   *uint   `json:"role_id,omitempty" example:"2"`
	GroupID  *uint   `json:"group_id,omitempty" example:"1"`
	CountryID *uint  `json:"country_id,omitempty" example:"1"`

	University             *string    `json:"university,omitempty" example:"Example University"`
	StudentID              *string    `json:"student_id,omitempty" example:"STU123"`
	Department             *string    `json:"department,omitempty" example:"Computer Science"`
	ExpectedGraduationDate *time.Time `json:"expected_graduation_date,omitempty"`

	Phone            *string `json:"phone,omitempty" example:"+1234567890"`
	TelegramUsername *string `json:"telegram_username,omitempty" example:"@username"`
	TelegramUID      *string `json:"telegram_uid,omitempty" example:"123456789"`

	Leetcode   *string `json:"leetcode,omitempty" example:"leetcode_user"`
	Codeforces *string `json:"codeforces,omitempty" example:"cf_user"`
	Github     *string `json:"github,omitempty" example:"github_user"`
	Hackerrank *string `json:"hackerrank,omitempty" example:"hr_user"`
	Linkedin   *string `json:"linkedin,omitempty" example:"linkedin_profile"`
	Instagram  *string `json:"instagram,omitempty" example:"insta_user"`

	Birthday          *time.Time `json:"birthday,omitempty"`
	Gender            *string    `json:"gender,omitempty" example:"male"`
	ShortBio          *string    `json:"short_bio,omitempty" example:"Software developer passionate about algorithms"`
	PreferredLanguage *string    `json:"preferred_language,omitempty" example:"en"`
}

// UserListQuery represents query parameters for listing users
// swagger:model
// (Already refined in previous steps)
type UserListQuery struct {
	Page     int    `form:"page,default=1" example:"1"`
	PageSize int    `form:"page_size,default=10" example:"10"`
	Search   string `form:"search" example:"john"`
	RoleID   *uint  `form:"role_id" example:"2"`
	GroupID  *uint  `form:"group_id" example:"1"`
}

// UserResponse represents a user in responses
// swagger:model
// (Already refined in previous steps)
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"user@example.com"`
	RoleID    uint      `json:"role_id" example:"2"`
	GroupID   *uint     `json:"group_id,omitempty" example:"1"`
	CountryID *uint     `json:"country_id,omitempty" example:"1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	University             *string    `json:"university,omitempty" example:"Example University"`
	StudentID              *string    `json:"student_id,omitempty" example:"STU123"`
	Department             *string    `json:"department,omitempty" example:"Computer Science"`
	ExpectedGraduationDate *time.Time `json:"expected_graduation_date,omitempty"`

	Phone            *string `json:"phone,omitempty" example:"+1234567890"`
	TelegramUsername *string `json:"telegram_username,omitempty" example:"@username"`
	TelegramUID      *string `json:"telegram_uid,omitempty" example:"123456789"`

	Leetcode   *string `json:"leetcode,omitempty" example:"leetcode_user"`
	Codeforces *string `json:"codeforces,omitempty" example:"cf_user"`
	Github     *string `json:"github,omitempty" example:"github_user"`
	Hackerrank *string `json:"hackerrank,omitempty" example:"hr_user"`
	Linkedin   *string `json:"linkedin,omitempty" example:"linkedin_profile"`
	Instagram  *string `json:"instagram,omitempty" example:"insta_user"`

	Birthday          *time.Time `json:"birthday,omitempty"`
	Gender            *string    `json:"gender,omitempty" example:"male"`
	ShortBio          *string    `json:"short_bio,omitempty" example:"Software developer passionate about algorithms"`
	PreferredLanguage *string    `json:"preferred_language,omitempty" example:"en"`
}

// UserListResponse represents paginated user results
// swagger:model
// (Already refined in previous steps)
type UserListResponse struct {
	Data []*UserResponse `json:"data"`
	Meta PaginationMeta  `json:"meta"`
}
