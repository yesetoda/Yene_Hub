package schemas

import "time"

// CreateRoleRequest represents the request body for creating a new role
// swagger:model
// Note: The entity uses 'Type', not 'Name'.
type CreateRoleRequest struct {
	// Required: true
	Type        string `json:"type" binding:"required" example:"Moderator"`
	Description string `json:"description,omitempty" example:"Can moderate content"`
}

// UpdateRoleRequest represents the request body for updating a role
// swagger:model
type UpdateRoleRequest struct {
	Type        *string `json:"type,omitempty" example:"Senior Moderator"`
	Description *string `json:"description,omitempty" example:"Can moderate all content"`
}

// RoleResponse represents a role in responses
// swagger:model
type RoleResponse struct {
	ID          uint      `json:"id" example:"1"`
	Type        string    `json:"type" example:"Moderator"`
	Description string    `json:"description,omitempty" example:"Can moderate content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RoleListResponse represents paginated role results
// swagger:model
type RoleListResponse struct {
	Data []*RoleResponse `json:"data"`
	Meta PaginationMeta  `json:"meta"`
}
