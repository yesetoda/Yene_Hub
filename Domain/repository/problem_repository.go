package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// ProblemRepository defines methods for Problem data operations
type ProblemRepository interface {
	CreateProblem(Problem *entity.Problem) error

	ListProblem() ([]*entity.Problem, error)

	GetProblemByName(name string) (*entity.Problem, error)
	GetProblemByDifficulty(difficulty string) ([]*entity.Problem, error)
	GetProblemByTag(tag string) ([]*entity.Problem, error)
	GetProblemByPlatform(platform string) ([]*entity.Problem, error)
	GetProblemByID(id uint) (*entity.Problem, error)

	UpdateProblem(Problem *entity.Problem) error

	DeleteProblem(id uint) error
}
