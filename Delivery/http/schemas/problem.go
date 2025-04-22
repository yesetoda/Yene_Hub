package schemas

import "time"

// Difficulty represents the difficulty level of a problem
type Difficulty string

const (
	Easy     Difficulty = "easy"
	Medium   Difficulty = "medium"
	Hard     Difficulty = "hard"
	VeryHard Difficulty = "very_hard"
)

// CreateProblemRequest represents the request body for creating a new problem
// swagger:model
type CreateProblemRequest struct {
	// Required: true
	Name       string     `json:"name" binding:"required" example:"Two Sum"`
	Description string     `json:"description" binding:"required" example:"Find two numbers that add up to target"`
	Difficulty  Difficulty `json:"difficulty" binding:"required" example:"medium"`
	Link        string     `json:"link" binding:"required" example:"https://leetcode.com/problems/two-sum"`
	Platform    string     `json:"platform" binding:"required" example:"leetcode"`
	Tags        []string   `json:"tags" example:"['array','hash-table']"`
}

// UpdateProblemRequest represents the request body for updating a problem
// swagger:model
type UpdateProblemRequest struct {
	ID          uint      `json:"id" example:"1"`
	Name        *string     `json:"name,omitempty" example:"Two Sum Problem"`
	Description  *string     `json:"description,omitempty"`
	Difficulty   *Difficulty `json:"difficulty,omitempty" example:"hard"`
	Link         *string     `json:"link,omitempty"`
	Platform     *string     `json:"platform,omitempty" example:"leetcode"`
	Tags         []string    `json:"tags,omitempty" example:"['array','hash-table','two-pointer']"`
}

// ProblemResponse represents a problem in responses
// swagger:model
type ProblemResponse struct {
	ID           uint       `json:"id" example:"1"`
	Name        string     `json:"name" example:"Two Sum"`
	Description  string     `json:"description" example:"Find two numbers that add up to target"`
	Difficulty   Difficulty `json:"difficulty" example:"medium"`
	Link         string     `json:"link" example:"https://leetcode.com/problems/two-sum"`
	Platform     string     `json:"platform" example:"leetcode"`
	Tags         []string   `json:"tags" example:"['array','hash-table']"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ProblemListQuery represents query parameters for listing problems
// swagger:model
type ProblemListQuery struct {
	Page       int         `form:"page,default=1" example:"1"`
	PageSize   int         `form:"page_size,default=10" example:"10"`
	Search     string      `form:"search" example:"two sum"`
	Difficulty *Difficulty `form:"difficulty" example:"medium"`
	Platform   *string     `form:"platform" example:"leetcode"`
	Tags       []string    `form:"tags" example:"array,hash-table"`
}

// ProblemListResponse represents paginated problem results
// swagger:model
type ProblemListResponse struct {
	Data []*ProblemResponse `json:"data"`
	Meta PaginationMeta     `json:"meta"`
}
