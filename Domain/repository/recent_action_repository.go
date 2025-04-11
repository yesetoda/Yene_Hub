package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// RecentActionRepository defines methods for RecentAction data opeRecentActiontions
type RecentActionRepository interface {
	CreateRecentAction(RecentAction *entity.RecentAction) error

	ListRecentAction() ([]*entity.RecentAction, error)

	GetRecentActionByUserID(userID uint) ([]*entity.RecentAction, error)
	GetRecentActionByID(id uint) (*entity.RecentAction, error)
	GetRecentActionByType(actionType string) ([]*entity.RecentAction, error)
	

	UpdateRecentAction(RecentAction *entity.RecentAction) error

	DeleteRecentAction(id uint) error
}
