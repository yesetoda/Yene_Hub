package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"gorm.io/gorm"
)

type ProblemRepository struct {
	db *gorm.DB
}

func NewProblemRepository(db *gorm.DB) *ProblemRepository {
	return &ProblemRepository{
		db: db,
	}
}

func (r *ProblemRepository) CreateProblem(problem *entity.Problem) error {
	return r.db.Create(problem).Error
}
func (r *ProblemRepository) GetProblemByID(id uint) (*entity.Problem, error) {
	var problem entity.Problem
	result := r.db.First(&problem, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &problem, nil
}

func (r *ProblemRepository) GetProblemByName(name string) ([]*entity.Problem, error) {
	var problems []*entity.Problem
	result := r.db.Where("name = ?", name).Find(&problems)
	if result.Error != nil {
		return nil, result.Error
	}
	return problems, nil
}
func (r *ProblemRepository) GetProblemByTag(tag string) ([]*entity.Problem, error) {
	var problems []*entity.Problem
	result := r.db.Where("tag = ?", tag).Find(&problems)
	if result.Error != nil {
		return nil, result.Error
	}
	return problems, nil
}
func (r *ProblemRepository) GetProblemByDifficulty(difficulty string) ([]*entity.Problem, error) {
	var problems []*entity.Problem
	result := r.db.Where("difficulty = ?", difficulty).Find(&problems)
	if result.Error != nil {
		return nil, result.Error
	}
	return problems, nil
}

func (r *ProblemRepository) GetProblemByPlatform(platform string) ([]*entity.Problem, error) {
	var problems []*entity.Problem
	result := r.db.Where("platform = ?", platform).Find(&problems)
	if result.Error != nil {
		return nil, result.Error
	}
	return problems, nil
}

func (r *ProblemRepository) UpdateProblem(problem *entity.Problem) error {
	return r.db.Save(problem).Error
}

func (r *ProblemRepository) DeleteProblem(id uint) error {
	var problem entity.Problem
	result := r.db.First(&problem, id)
	if result.Error != nil {
		return result.Error
	}
	return r.db.Delete(&problem).Error
}