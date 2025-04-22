package schemas

import "time"

// SubmissionStatus represents the status of a submission
type SubmissionStatus string

const (
	Pending    SubmissionStatus = "pending"
	Running    SubmissionStatus = "running"
	Accepted   SubmissionStatus = "accepted"
	WrongAnswer SubmissionStatus = "wrong_answer"
	TimeLimitExceeded SubmissionStatus = "time_limit_exceeded"
	MemoryLimitExceeded SubmissionStatus = "memory_limit_exceeded"
	RuntimeError SubmissionStatus = "runtime_error"
	CompilationError SubmissionStatus = "compilation_error"
)

// Language represents the programming language of a submission
type Language string

const (
	Python  Language = "python"
	Java    Language = "java"
	CPP     Language = "cpp"
	JavaScript Language = "javascript"
	Go      Language = "go"
)

// CreateSubmissionRequest represents the request body for creating a new submission
// swagger:model
type CreateSubmissionRequest struct {
	// Required: true
	UserID    uint     `json:"user_id" binding:"required" example:"1"`
	ProblemID uint     `json:"problem_id" binding:"required" example:"1"`
	Code      string   `json:"code" binding:"required" example:"def two_sum(nums, target):..."`
	Language  Language `json:"language" binding:"required" example:"python"`
}

// UpdateSubmissionRequest represents the request body for updating a submission
// swagger:model
type UpdateSubmissionRequest struct {
	Status *SubmissionStatus `json:"status,omitempty" example:"accepted"`
	Score  *int             `json:"score,omitempty" example:"100"`
}

// SubmissionResponse represents a submission in responses
// swagger:model
type SubmissionResponse struct {
	ID        uint            `json:"id" example:"1"`
	UserID    uint            `json:"user_id" example:"1"`
	ProblemID uint            `json:"problem_id" example:"1"`
	Code      string          `json:"code" example:"def two_sum(nums, target):..."`
	Language  Language        `json:"language" example:"python"`
	Status    SubmissionStatus `json:"status" example:"accepted"`
	
	// Performance metrics
	ExecutionTime  int `json:"execution_time,omitempty" example:"45"` // in milliseconds
	MemoryUsed     int `json:"memory_used,omitempty" example:"24"`    // in MB
	Score          int `json:"score" example:"100"`
	TestCasesPassed int `json:"test_cases_passed" example:"10"`
	TotalTestCases  int `json:"total_test_cases" example:"10"`
	
	// Error information
	ErrorMessage string `json:"error_message,omitempty" example:"Index out of range"`
	StackTrace   string `json:"stack_trace,omitempty"`
	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SubmissionListQuery represents query parameters for listing submissions
// swagger:model
type SubmissionListQuery struct {
	Page      int              `form:"page,default=1" example:"1"`
	PageSize  int              `form:"page_size,default=10" example:"10"`
	UserID    *uint            `form:"user_id" example:"1"`
	ProblemID *uint            `form:"problem_id" example:"1"`
	Language  *Language        `form:"language" example:"python"`
	Status    *SubmissionStatus `form:"status" example:"accepted"`
	StartDate *time.Time       `form:"start_date"`
	EndDate   *time.Time       `form:"end_date"`
}

// SubmissionListResponse represents paginated submission results
// swagger:model
type SubmissionListResponse struct {
	Data []*SubmissionResponse `json:"data"`
	Meta PaginationMeta        `json:"meta"`
}

// SubmissionStats represents submission statistics
// swagger:model
type SubmissionStats struct {
	TotalSubmissions     int     `json:"total_submissions" example:"100"`
	AcceptedSubmissions  int     `json:"accepted_submissions" example:"75"`
	AcceptanceRate      float64 `json:"acceptance_rate" example:"75.0"`
	AverageScore        float64 `json:"average_score" example:"85.5"`
	AverageExecutionTime float64 `json:"average_execution_time" example:"50.2"`
	AverageMemoryUsed    float64 `json:"average_memory_used" example:"25.5"`
}
