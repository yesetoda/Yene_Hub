package schemas

import "time"

// VoteType represents the type of vote (upvote/downvote)
type VoteType string

const (
	Upvote   VoteType = "upvote"
	Downvote VoteType = "downvote"
)

// CreateVoteRequest represents the request body for creating a new vote
// swagger:model
type CreateVoteRequest struct {
	// Required: true
	UserID     uint     `json:"user_id" binding:"required" example:"1"`
	VoteType   VoteType `json:"vote_type" binding:"required" example:"upvote"`
	EntityID   uint     `json:"entity_id" binding:"required" example:"123"`
	EntityType string   `json:"entity_type" binding:"required" example:"comment"` // comment, post, track, submission, problem
}

// UpdateVoteRequest represents the request body for updating a vote
// swagger:model
type UpdateVoteRequest struct {
	VoteType *VoteType `json:"vote_type,omitempty" example:"downvote"`
}

// VoteResponse represents a vote in responses
// swagger:model
type VoteResponse struct {
	ID         uint      `json:"id" example:"1"`
	UserID     uint      `json:"user_id" example:"1"`
	VoteType   VoteType  `json:"vote_type" example:"upvote"`
	EntityID   uint      `json:"entity_id" example:"123"`
	EntityType string    `json:"entity_type" example:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// VoteListQuery represents query parameters for listing votes
// swagger:model
type VoteListQuery struct {
	Page       int       `form:"page,default=1" example:"1"`
	PageSize   int       `form:"page_size,default=10" example:"10"`
	UserID     *uint     `form:"user_id" example:"1"`
	EntityID   *uint     `form:"entity_id" example:"123"`
	EntityType *string   `form:"entity_type" example:"comment"`
	VoteType   *VoteType `form:"vote_type" example:"upvote"`
}

// VoteListResponse represents paginated vote results
// swagger:model
type VoteListResponse struct {
	Data []*VoteResponse `json:"data"`
	Meta PaginationMeta  `json:"meta"`
}

// VoteStats represents vote statistics for an entity
// swagger:model
type VoteStats struct {
	EntityID    uint `json:"entity_id" example:"123"`
	UpvoteCount int  `json:"upvote_count" example:"10"`
	DownvoteCount int `json:"downvote_count" example:"2"`
	TotalScore   int  `json:"total_score" example:"8"` // UpvoteCount - DownvoteCount
}
