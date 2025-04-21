package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// ProblemRepository defines methods for Problem data operations
type ProblemUseCaseInterface interface {
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

type ProblemUsecase struct {
	ProblemRepository repository.ProblemRepository
}

func NewProblemUsecase(problemRepository repository.ProblemRepository) *ProblemUsecase {
	return &ProblemUsecase{
		ProblemRepository: problemRepository,
	}
}

func (p *ProblemUsecase) CreateProblem(problem *entity.Problem) error {
	err := p.ProblemRepository.CreateProblem(problem)
	if err != nil {
		return err
	}
	return nil
}
func (p *ProblemUsecase) ListProblem() ([]*entity.Problem, error) {
	problems, err := p.ProblemRepository.ListProblem()
	if err != nil {
		return nil, err
	}
	return problems, nil
}
func (p *ProblemUsecase) GetProblemByName(name string) (*entity.Problem, error) {
	problem, err := p.ProblemRepository.GetProblemByName(name)
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func (p *ProblemUsecase) GetProblemByDifficulty(difficulty string) ([]*entity.Problem, error) {
	problems, err := p.ProblemRepository.GetProblemByDifficulty(difficulty)
	if err != nil {
		return nil, err
	}
	return problems, nil
}
func (p *ProblemUsecase) GetProblemByTag(tag string) ([]*entity.Problem, error) {
	problems, err := p.ProblemRepository.GetProblemByTag(tag)
	if err != nil {
		return nil, err
	}
	return problems, nil
}
func GetProblemByDifficulty(p *ProblemUsecase, difficulty string) ([]*entity.Problem, error) {
	problems, err := p.ProblemRepository.GetProblemByDifficulty(difficulty)
	if err != nil {
		return nil, err
	}
	return problems, nil
}
func (p *ProblemUsecase) GetProblemByPlatform(platform string) ([]*entity.Problem, error) {
	problems, err := p.ProblemRepository.GetProblemByPlatform(platform)
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (p *ProblemUsecase) GetProblemByID(id uint) (*entity.Problem, error) {
	problem, err := p.ProblemRepository.GetProblemByID(id)
	if err != nil {
		return nil, err
	}
	return problem, nil
}

func (p *ProblemUsecase) UpdateProblem(problem *entity.Problem) error {
	err := p.ProblemRepository.UpdateProblem(problem)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemUsecase) DeleteProblem(id uint) error {
	err := p.ProblemRepository.DeleteProblem(id)
	if err != nil {
		return err
	}
	return nil
}
