package repository

import "a2sv.org/hub/Domain/entity"

type ExerciseRepository interface {
	Create(exercise *entity.Exercise) error
	GetByID(id uint) (*entity.Exercise, error)
	GetByGroupID(groupID uint) ([]*entity.Exercise, error)
	GetByTrackID(trackID uint) ([]*entity.Exercise, error)
	GetByProblemID(problemID uint) ([]*entity.Exercise, error)
	GetAll() ([]*entity.Exercise, error)
	Update(exercise *entity.Exercise) error
	Delete(id uint) error
}
