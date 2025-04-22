package schemas

import "time"

// CreateSessionRequest represents the request body for creating a new session
// swagger:model
type CreateSessionRequest struct {
	// Required: true
	Name       string    `json:"name" binding:"required" example:"DSA Session #1"`
	Description string    `json:"description" binding:"required" example:"Introduction to Data Structures"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	GroupID     uint      `json:"group_id" binding:"required" example:"1"`
	HostID      uint      `json:"host_id" binding:"required" example:"1"`
	Location    string    `json:"location,omitempty" example:"Room 101"`
	MeetingLink string    `json:"meeting_link,omitempty" example:"https://meet.google.com/abc-xyz"`
}

// UpdateSessionRequest represents the request body for updating a session
// swagger:model
type UpdateSessionRequest struct {
	Name       *string    `json:"name,omitempty" example:"DSA Session #1 - Updated"`
	Description *string    `json:"description,omitempty" example:"Updated Introduction to Data Structures"`
	StartTime   *time.Time `json:"start_time,omitempty"`
	EndTime     *time.Time `json:"end_time,omitempty"`
	GroupID     *uint      `json:"group_id,omitempty" example:"1"`
	HostID      *uint      `json:"host_id,omitempty" example:"1"`
	Location    *string    `json:"location,omitempty" example:"Room 102"`
	MeetingLink *string    `json:"meeting_link,omitempty" example:"https://meet.google.com/new-link"`
	Status      *string    `json:"status,omitempty" example:"completed"`
}

// SessionResponse represents a session in responses
// swagger:model
type SessionResponse struct {
	ID          uint      `json:"id" example:"1"`
	Name       string    `json:"name" example:"DSA Session #1"`
	Description string    `json:"description" example:"Introduction to Data Structures"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	GroupID     uint      `json:"group_id" example:"1"`
	HostID      uint      `json:"host_id" example:"1"`
	Location    string    `json:"location,omitempty" example:"Room 101"`
	MeetingLink string    `json:"meeting_link,omitempty" example:"https://meet.google.com/abc-xyz"`
	Status      string    `json:"status" example:"scheduled"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SessionListQuery represents query parameters for listing sessions
// swagger:model
type SessionListQuery struct {
	Page      int       `form:"page,default=1" example:"1"`
	PageSize  int       `form:"page_size,default=10" example:"10"`
	GroupID   *uint     `form:"group_id" example:"1"`
	HostID    *uint     `form:"host_id" example:"1"`
	Status    *string   `form:"status" example:"scheduled"`
	StartDate *time.Time `form:"start_date"`
	EndDate   *time.Time `form:"end_date"`
}

// SessionListResponse represents paginated session results
// swagger:model
type SessionListResponse struct {
	Data []*SessionResponse `json:"data"`
	Meta PaginationMeta    `json:"meta"`
}
