package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// StipendRepository defines methods for Stipend data operations
type StipendUseCaseInterface interface {
	CreateStipend(Stipend *entity.Stipend) error

	ListStipend() ([]*entity.Stipend, error)
	GetStipendByID(id uint) (*entity.Stipend, error)

	UpdateStipend(Stipend *entity.Stipend) error

	DeleteStipend(id uint) error
}


type StipendUsecase struct{
	StipendRepository repository.StipendRepository
}

func NewStipendUsecase(stipendRepository repository.StipendRepository) StipendUseCaseInterface{
	return &StipendUsecase{
		StipendRepository: stipendRepository,
	}
}

func (s *StipendUsecase) CreateStipend(stipend *entity.Stipend) error {
	err := s.StipendRepository.CreateStipend(stipend)
	if err != nil {
		return err
	}
	return nil
}

func (s *StipendUsecase) ListStipend() ([]*entity.Stipend, error) {
	stipends, err := s.StipendRepository.ListStipend()
	if err != nil {
		return nil, err
	}
	return stipends, nil
}

func (s *StipendUsecase) GetStipendByID(id uint) (*entity.Stipend, error) {
	stipend, err := s.StipendRepository.GetStipendByID(id)
	if err != nil {
		return nil, err
	}
	return stipend, nil
}

func (s *StipendUsecase) UpdateStipend(stipend *entity.Stipend) error {
	err := s.StipendRepository.UpdateStipend(stipend)
	if err != nil {
		return err
	}
	return nil
}

func (s *StipendUsecase) DeleteStipend(id uint) error {
	err := s.StipendRepository.DeleteStipend(id)
	if err != nil {
		return err
	}
	return nil
}
