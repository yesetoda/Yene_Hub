package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type ExerciseRepository struct {
	db *gorm.DB
}


func NewExerciseRepository(db *gorm.DB) repository.ExerciseRepository {
	return &ExerciseRepository{db: db}
}

func (r *ExerciseRepository) Create(exercise *entity.Exercise) error {
	return r.db.Create(exercise).Error
}

func (r *ExerciseRepository) GetByID(id uint) (*entity.Exercise, error) {
	var exercise entity.Exercise
	if err := r.db.First(&exercise, id).Error; err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (r *ExerciseRepository) GetByGroupID(groupID uint) ([]*entity.Exercise, error) {
	var exercises []*entity.Exercise
	if err := r.db.Where("group_id = ?", groupID).Find(&exercises).Error; err != nil {
		return nil, err
	}
	return exercises, nil
}

func (r *ExerciseRepository) GetByTrackID(trackID uint) ([]*entity.Exercise, error) {
	var exercises []*entity.Exercise
	if err := r.db.Where("track_id = ?", trackID).Find(&exercises).Error; err != nil {
		return nil, err
	}
	return exercises, nil
}

func (r *ExerciseRepository) GetByProblemID(problemID uint) ([]*entity.Exercise, error) {
	var exercises []*entity.Exercise
	if err := r.db.Where("problem_id = ?", problemID).Find(&exercises).Error; err != nil {
		return nil, err
	}
	return exercises, nil
}

func (r *ExerciseRepository) GetAll() ([]*entity.Exercise, error) {
	var exercises []*entity.Exercise
	if err := r.db.Find(&exercises).Error; err != nil {
		return nil, err
	}
	return exercises, nil
}

func (r *ExerciseRepository) Update(exercise *entity.Exercise) error {
	return r.db.Save(exercise).Error
}

func (r *ExerciseRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Exercise{}, id).Error
}

