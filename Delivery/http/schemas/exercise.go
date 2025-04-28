package schemas

import "time"

// CreateExerciseRequest represents the request body for creating a new exercise
// swagger:model
type CreateExerciseRequest struct {
	TrackID   uint `json:"track_id" binding:"required"`
	ProblemID uint `json:"problem_id" binding:"required"`
	GroupID   uint `json:"group_id" binding:"required"`
}

// ExerciseResponse represents an exercise in responses
// swagger:model
type ExerciseResponse struct {
	ID        uint      `json:"id"`
	TrackID   uint      `json:"track_id"`
	ProblemID uint      `json:"problem_id"`
	GroupID   uint      `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateExerciseRequest represents the request body for updating an exercise
// swagger:model
type UpdateExerciseRequest struct {
	ID        uint      `json:"id" example:"1"`
	TrackID   uint      `json:"track_id"`
	ProblemID uint      `json:"problem_id"`
	GroupID   uint      `json:"group_id"`
}

// DeleteExerciseRequest represents the request body for deleting an exercise
// swagger:model
type DeleteExerciseRequest struct {
	ID uint `json:"id" binding:"required"`
}
