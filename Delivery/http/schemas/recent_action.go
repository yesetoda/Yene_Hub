package schemas

import "time"

// CreateRecentActionRequest represents the request body for creating a new recent action
// swagger:model
type CreateRecentActionRequest struct {
	// Required: true
	UserID      uint   `json:"user_id" binding:"required" example:"1"`
	ActionType  string `json:"action_type" binding:"required" example:"problem_solved"`
	Description string `json:"description" binding:"required" example:"Solved Two Sum problem"`
	EntityID    uint   `json:"entity_id" binding:"required" example:"123"`
	EntityType  string `json:"entity_type" binding:"required" example:"problem"`
}

// UpdateRecentActionRequest represents the request body for updating a recent action
// swagger:model
type UpdateRecentActionRequest struct {
	ID          uint      `json:"id" example:"1"`
	ActionType  *string `json:"action_type,omitempty" example:"problem_attempted"`
	Description *string `json:"description,omitempty" example:"Attempted Two Sum problem"`
	EntityID    *uint   `json:"entity_id,omitempty" example:"123"`
	EntityType  *string `json:"entity_type,omitempty" example:"problem"`
}

// RecentActionResponse represents a recent action in responses
// swagger:model
type RecentActionResponse struct {
	ID          uint      `json:"id" example:"1"`
	UserID      uint      `json:"user_id" example:"1"`
	ActionType  string    `json:"action_type" example:"problem_solved"`
	Description string    `json:"description" example:"Solved Two Sum problem"`
	EntityID    uint      `json:"entity_id" example:"123"`
	EntityType  string    `json:"entity_type" example:"problem"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RecentActionListQuery represents query parameters for listing recent actions
// swagger:model
type RecentActionListQuery struct {
	Page       int       `form:"page,default=1" example:"1"`
	PageSize   int       `form:"page_size,default=10" example:"10"`
	UserID     *uint     `form:"user_id" example:"1"`
	ActionType *string   `form:"action_type" example:"problem_solved"`
	EntityType *string   `form:"entity_type" example:"problem"`
	StartDate  *time.Time `form:"start_date"`
	EndDate    *time.Time `form:"end_date"`
}

// RecentActionListResponse represents paginated recent action results
// swagger:model
type RecentActionListResponse struct {
	Data []*RecentActionResponse `json:"data"`
	Meta PaginationMeta         `json:"meta"`
}
