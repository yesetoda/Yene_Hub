package schemas

import "time"

// CreateGroupRequest represents the request body for creating a new group
// swagger:model
// (Refined: matches entity, includes ShortName, HOAID, pointers for optional fields)
type CreateGroupRequest struct {
	Name        string  `json:"name" binding:"required" example:"Team Alpha"`
	ShortName   *string `json:"short_name,omitempty" example:"Alpha"`
	Description *string `json:"description,omitempty" example:"A competitive programming team"`
	HOAID       *uint   `json:"hoa_id,omitempty" example:"5"`
	CountryID   uint    `json:"country_id" binding:"required" example:"1"`
}

// UpdateGroupRequest represents the request body for updating a group
// swagger:model
// (Refined: all fields optional, pointers for partial updates)
type UpdateGroupRequest struct {
	Name        *string `json:"name,omitempty" example:"Team Beta"`
	ShortName   *string `json:"short_name,omitempty" example:"Beta"`
	Description *string `json:"description,omitempty" example:"An advanced competitive programming team"`
	HOAID       *uint   `json:"hoa_id,omitempty" example:"6"`
	CountryID   *uint   `json:"country_id,omitempty" example:"2"`
}

// GroupResponse represents a group in responses
// swagger:model
// (Refined: matches entity, pointers for optional fields)
type GroupResponse struct {
	ID          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"Team Alpha"`
	ShortName   *string   `json:"short_name,omitempty" example:"Alpha"`
	Description *string   `json:"description,omitempty" example:"A competitive programming team"`
	HOAID       *uint     `json:"hoa_id,omitempty" example:"5"`
	CountryID   *uint     `json:"country_id,omitempty" example:"1"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GroupListQuery represents query parameters for listing groups
// swagger:model
type GroupListQuery struct {
	Page      int    `form:"page,default=1" example:"1"`
	PageSize  int    `form:"page_size,default=10" example:"10"`
	Search    string `form:"search" example:"alpha"`
	CountryID *uint  `form:"country_id" example:"1"`
}

// GroupListResponse represents paginated group results
// swagger:model
type GroupListResponse struct {
	Data []*GroupResponse `json:"data"`
	Meta PaginationMeta   `json:"meta"`
}
