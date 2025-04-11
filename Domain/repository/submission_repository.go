package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// SubmissionRepository defines methods for Submission data operations
type SubmissionRepository interface {
	CreateSubmission(Submission *entity.Submission) error

	ListSubmission() ([]*entity.Submission, error)

	GetSubmissionByUserID(userID uint) ([]*entity.Submission, error)
	GetSubmissionByProblemID(problemID uint) ([]*entity.Submission, error)
	GetSubmissionByID(id uint) (*entity.Submission, error)

	UpdateSubmission(Submission *entity.Submission) error

	DeleteSubmission(id uint) error
}
