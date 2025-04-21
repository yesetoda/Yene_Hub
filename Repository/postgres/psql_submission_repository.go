package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type SubmissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) repository.SubmissionRepository {
	return &SubmissionRepository{
		db: db,
	}
}
func (r *SubmissionRepository) CreateSubmission(submission *entity.Submission) error {
	return r.db.Create(submission).Error
}

func (r *SubmissionRepository) ListSubmission() ([]*entity.Submission, error) {
	var submissions []*entity.Submission
	result := r.db.Find(&submissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return submissions, nil
}

func (r *SubmissionRepository) GetSubmissionByID(id uint) (*entity.Submission, error) {
	var submission entity.Submission
	result := r.db.First(&submission, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &submission, nil
}

func (r *SubmissionRepository) GetSubmissionByUserID(userID uint) ([]*entity.Submission, error) {
	var submissions []*entity.Submission
	result := r.db.Where("user_id = ?", userID).Find(&submissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return submissions, nil
}

func (r *SubmissionRepository) GetSubmissionByProblemID(problemID uint) ([]*entity.Submission, error) {
	var submissions []*entity.Submission
	result := r.db.Where("problem_id = ?", problemID).Find(&submissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return submissions, nil
}

func (r *SubmissionRepository) UpdateSubmission(submission *entity.Submission) error {
	return r.db.Save(submission).Error
}

func (r *SubmissionRepository) DeleteSubmission(id uint) error {
	return r.db.Delete(&entity.Submission{}, id).Error
}
