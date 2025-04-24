package schemas

import "time"

// CreateSuperGroupRequest represents the request body for creating a new super group
// swagger:model
type CreateSuperGroupRequest struct {
	// Required: true
	Name        string `json:"name" binding:"required" example:"A2SV Generation 2"`
	Description string `json:"description" binding:"required" example:"Second generation of A2SV students"`
	CountryID   uint   `json:"country_id" binding:"required" example:"1"`
}

// UpdateSuperGroupRequest represents the request body for updating a super group
// swagger:model
type UpdateSuperGroupRequest struct {
	Name        *string `json:"name,omitempty" example:"A2SV Generation 2 - Advanced"`
	Description *string `json:"description,omitempty" example:"Advanced track of second generation"`
	CountryID   *uint   `json:"country_id,omitempty" example:"1"`
	Status      *string `json:"status,omitempty" example:"active"`
}

// SuperGroupResponse represents a super group in responses
// swagger:model
type SuperGroupResponse struct {
	ID          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"A2SV Generation 2"`
	Description string    `json:"description" example:"Second generation of A2SV students"`
	CountryID   uint      `json:"country_id" example:"1"`
	Status      string    `json:"status" example:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SuperGroupListQuery represents query parameters for listing super groups
// swagger:model
type SuperGroupListQuery struct {
	Page      int     `form:"page,default=1" example:"1"`
	PageSize  int     `form:"page_size,default=10" example:"10"`
	Search    string  `form:"search" example:"Generation 2"`
	CountryID *uint   `form:"country_id" example:"1"`
	Status    *string `form:"status" example:"active"`
}

// SuperGroupListResponse represents paginated super group results
// swagger:model
type SuperGroupListResponse struct {
	Data []*SuperGroupResponse `json:"data"`
	Meta PaginationMeta       `json:"meta"`
}

// SuperToGroupRequest represents the request body for adding groups to a super group
// swagger:model
// Note: The example tag should be a comma-separated list for array fields, not a JSON array string.
type SuperToGroupRequest struct {
	GroupIDs []uint `json:"group_ids" binding:"required" example:"1,2,3"`
}

// SuperToGroupResponse represents the response for groups in a super group
// swagger:model
type SuperToGroupResponse struct {
	SuperGroupID uint           `json:"super_group_id" example:"1"`
	Groups       []GroupSummary `json:"groups"`
}

// GroupSummary represents a minimal group response for super group relationships
// swagger:model
type GroupSummary struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"Team Alpha"`
}
