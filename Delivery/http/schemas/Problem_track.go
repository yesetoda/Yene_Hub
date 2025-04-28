package schemas

import "time"

// AddProblemToTrackRequest represents the request body for adding a problem to a track
// swagger:model

type AddProblemToTrackRequest struct {
	TrackID   uint `json:"track_id" binding:"required"`
	ProblemID uint `json:"problem_id" binding:"required"`
}

// ProblemTrackResponse represents a problem track in responses
// swagger:model
type ProblemTrackResponse struct {
	ID        uint      `json:"id"`
	TrackID   uint      `json:"track_id"`
	ProblemID uint      `json:"problem_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateProblemTrackRequest represents the request body for updating a problem track
// swagger:model
type UpdateProblemTrackRequest struct {
	ID        uint      `json:"id" example:"1"`
	TrackID   uint      `json:"track_id"`
	ProblemID uint      `json:"problem_id"`
}

// DeleteProblemTrackRequest represents the request body for deleting a problem track
// swagger:model
type DeleteProblemTrackRequest struct {
	ID uint `json:"id" binding:"required"`
}
	