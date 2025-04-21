package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// RecentActionRepository defines methods for RecentAction data opeRecentActiontions
type RecentActionUseCaseInterface interface {
	CreateRecentAction(RecentAction *entity.RecentAction) error

	ListRecentAction() ([]*entity.RecentAction, error)

	GetRecentActionByUserID(userID uint) ([]*entity.RecentAction, error)
	GetRecentActionByID(id uint) (*entity.RecentAction, error)
	GetRecentActionByType(actionType string) ([]*entity.RecentAction, error)

	UpdateRecentAction(RecentAction *entity.RecentAction) error

	DeleteRecentAction(id uint) error
}

type RecentActionUsecase struct {
	RecentActionRepository repository.RecentActionRepository
}
func NewRecentActionUsecase(recentActionRepository repository.RecentActionRepository) *RecentActionUsecase{
	return &RecentActionUsecase{
		RecentActionRepository: recentActionRepository,
	}
}

func (r *RecentActionUsecase) CreateRecentAction(recentAction *entity.RecentAction) error {
	err := r.RecentActionRepository.CreateRecentAction(recentAction)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecentActionUsecase) ListRecentAction() ([]*entity.RecentAction, error) {
	recentActions, err := r.RecentActionRepository.ListRecentAction()
	if err != nil {
		return nil, err
	}
	return recentActions, nil
}

func (r *RecentActionUsecase) GetRecentActionByUserID(userID uint) ([]*entity.RecentAction, error) {
	recentActions, err := r.RecentActionRepository.GetRecentActionByUserID(userID)
	if err != nil {
		return nil, err
	}
	return recentActions, nil
}

func (r *RecentActionUsecase) GetRecentActionByID(id uint) (*entity.RecentAction, error) {
	recentAction, err := r.RecentActionRepository.GetRecentActionByID(id)
	if err != nil {
		return nil, err
	}
	return recentAction, nil
}

func (r *RecentActionUsecase) GetRecentActionByType(actionType string) ([]*entity.RecentAction, error) {
	recentActions, err := r.RecentActionRepository.GetRecentActionByType(actionType)
	if err != nil {
		return nil, err
	}
	return recentActions, nil
}

func (r *RecentActionUsecase) UpdateRecentAction(recentAction *entity.RecentAction) error {
	err := r.RecentActionRepository.UpdateRecentAction(recentAction)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecentActionUsecase) DeleteRecentAction(id uint) error {
	err := r.RecentActionRepository.DeleteRecentAction(id)
	if err != nil {
		return err
	}
	return nil
}