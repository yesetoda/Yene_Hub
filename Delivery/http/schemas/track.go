package schemas

import "time"

// TrackType represents the type of track
type TrackType string

const (
	DSATrack     TrackType = "dsa"
	ProjectTrack TrackType = "project"
	SystemTrack  TrackType = "system_design"
)

// CreateTrackRequest represents the request body for creating a new track
// swagger:model
type CreateTrackRequest struct {
	// Required: true
	Name        string    `json:"name" binding:"required" example:"DSA Advanced Track"`
	Description string    `json:"description" binding:"required" example:"Advanced Data Structures and Algorithms"`
	Type        TrackType `json:"type" binding:"required" example:"dsa"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	GroupID     uint      `json:"group_id" binding:"required" example:"1"`
}

// UpdateTrackRequest represents the request body for updating a track
// swagger:model
type UpdateTrackRequest struct {
	Name        *string    `json:"name,omitempty" example:"DSA Expert Track"`
	Description *string    `json:"description,omitempty" example:"Expert level algorithms"`
	Type        *TrackType `json:"type,omitempty" example:"dsa"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	GroupID     *uint      `json:"group_id,omitempty" example:"1"`
	Status      *string    `json:"status,omitempty" example:"active"`
}

// TrackResponse represents a track in responses
// swagger:model
type TrackResponse struct {
	ID          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"DSA Advanced Track"`
	Description string    `json:"description" example:"Advanced Data Structures and Algorithms"`
	Type        TrackType `json:"type" example:"dsa"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	GroupID     uint      `json:"group_id" example:"1"`
	Status      string    `json:"status" example:"active"`
	Progress    float64   `json:"progress" example:"75.5"` // Percentage of completion
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TrackListQuery represents query parameters for listing tracks
// swagger:model
type TrackListQuery struct {
	Page     int        `form:"page,default=1" example:"1"`
	PageSize int        `form:"page_size,default=10" example:"10"`
	Search   string     `form:"search" example:"DSA"`
	Type     *TrackType `form:"type" example:"dsa"`
	GroupID  *uint      `form:"group_id" example:"1"`
	Status   *string    `form:"status" example:"active"`
}

// TrackListResponse represents paginated track results
// swagger:model
type TrackListResponse struct {
	Data []*TrackResponse `json:"data"`
	Meta PaginationMeta   `json:"meta"`
}

// TrackProgressResponse represents track progress statistics
// swagger:model
type TrackProgressResponse struct {
	TrackID           uint    `json:"track_id" example:"1"`
	CompletedItems    int     `json:"completed_items" example:"15"`
	TotalItems        int     `json:"total_items" example:"20"`
	CompletionPercent float64 `json:"completion_percent" example:"75.5"`
	TimeSpentHours    float64 `json:"time_spent_hours" example:"24.5"`
}
