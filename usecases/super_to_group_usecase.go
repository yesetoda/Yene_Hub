package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// SuperToGroupRepository defines methods for SuperToGroup data operations
type SuperToGroupUseCaseInterface interface {
	CreateSuperToGroup(SuperToGroup *entity.SuperToGroup) error

	ListSuperToGroup() ([]*entity.SuperToGroup, error)

	GetSuperToGroupByID(id uint) (*entity.SuperToGroup, error)

	UpdateSuperToGroup(SuperToGroup *entity.SuperToGroup) error

	DeleteSuperToGroup(id uint) error
}

type SuperToGroupUsecase struct {
	SuperToGroupRepository repository.SuperToGroupRepository
}

func NewSuperToGroupUsecase(superToGroupRepository repository.SuperToGroupRepository) *SuperToGroupUsecase {
	return &SuperToGroupUsecase{
		SuperToGroupRepository: superToGroupRepository,
	}
}

func (s *SuperToGroupUsecase) CreateSuperToGroup(superToGroup *entity.SuperToGroup) error {
	err := s.SuperToGroupRepository.CreateSuperToGroup(superToGroup)
	if err != nil {
		return err
	}
	return nil
}

func (s *SuperToGroupUsecase) ListSuperToGroup() ([]*entity.SuperToGroup, error) {
	superToGroups, err := s.SuperToGroupRepository.ListSuperToGroup()
	if err != nil {
		return nil, err
	}
	return superToGroups, nil
}

func (s *SuperToGroupUsecase) GetSuperToGroupByID(id uint) (*entity.SuperToGroup, error) {
	superToGroup, err := s.SuperToGroupRepository.GetSuperToGroupByID(id)
	if err != nil {
		return nil, err
	}
	return superToGroup, nil
}
func (s *SuperToGroupUsecase) UpdateSuperToGroup(superToGroup *entity.SuperToGroup) error {
	err := s.SuperToGroupRepository.UpdateSuperToGroup(superToGroup)
	if err != nil {
		return err
	}
	return nil
}
func (s *SuperToGroupUsecase) DeleteSuperToGroup(id uint) error {
	err := s.SuperToGroupRepository.DeleteSuperToGroup(id)
	if err != nil {
		return err
	}
	return nil
}