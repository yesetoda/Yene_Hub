package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// SubmissionRepository defines methods for Submission data operations
type SubmissionUseCaseInterface interface {
	CreateSubmission(Submission *entity.Submission) error

	ListSubmission() ([]*entity.Submission, error)

	GetSubmissionByUserID(userID uint) ([]*entity.Submission, error)
	GetSubmissionByProblemID(problemID uint) ([]*entity.Submission, error)
	GetSubmissionByID(id uint) (*entity.Submission, error)

	UpdateSubmission(Submission *entity.Submission) error

	DeleteSubmission(id uint) error
}

type SubmissionUsecase struct {
	SubmissionRepository repository.SubmissionRepository
}

func NewSubmissionUsecase(submissionRepository repository.SubmissionRepository) SubmissionUseCaseInterface {
	return &SubmissionUsecase{
		SubmissionRepository: submissionRepository,
	}
}

func (s *SubmissionUsecase) CreateSubmission(submission *entity.Submission) error {
	err := s.SubmissionRepository.CreateSubmission(submission)
	if err != nil {
		return err
	}
	return nil
}
func (s *SubmissionUsecase) ListSubmission() ([]*entity.Submission, error) {
	submissions, err := s.SubmissionRepository.ListSubmission()
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func (s *SubmissionUsecase) GetSubmissionByUserID(userID uint) ([]*entity.Submission, error) {
	submissions, err := s.SubmissionRepository.GetSubmissionByUserID(userID)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}
func (s *SubmissionUsecase) GetSubmissionByProblemID(problemID uint) ([]*entity.Submission, error) {
	submissions, err := s.SubmissionRepository.GetSubmissionByProblemID(problemID)
	if err != nil {
		return nil, err
	}
	return submissions, nil
}
func (s *SubmissionUsecase) GetSubmissionByID(id uint) (*entity.Submission, error) {
	submission, err := s.SubmissionRepository.GetSubmissionByID(id)
	if err != nil {
		return nil, err
	}
	return submission, nil
}
func (s *SubmissionUsecase) UpdateSubmission(submission *entity.Submission) error {
	err := s.SubmissionRepository.UpdateSubmission(submission)
	if err != nil {
		return err
	}
	return nil
}
func (s *SubmissionUsecase) DeleteSubmission(id uint) error {
	err := s.SubmissionRepository.DeleteSubmission(id)
	if err != nil {
		return err
	}
	return nil
}
