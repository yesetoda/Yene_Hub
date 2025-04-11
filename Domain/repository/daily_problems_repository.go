package repository

import "a2sv.org/hub/Domain/entity"

// DailyProblemsRepository defines methods for daily problems database operations
type DailyProblemsRepository interface {
	Create(problem *entity.DailyProblem) error
	GetByID(id uint) (*entity.DailyProblem, error)
	GetByProblemID(problemID uint) (*entity.DailyProblem, error)
	GetBySuperGroupID(superGroupID uint) ([]*entity.DailyProblem, error)
	Update(problem *entity.DailyProblem) error
	Delete(id uint) error
	List() ([]*entity.DailyProblem, error)
}
