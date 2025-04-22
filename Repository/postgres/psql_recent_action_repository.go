package postgres

import (
	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type RecentActionRepository struct {
	db *gorm.DB
}

func NewRecentActionRepository(db *gorm.DB) repository.RecentActionRepository {
	return &RecentActionRepository{
		db: db,
	}
}

func (r *RecentActionRepository) CreateRecentAction(RecentAction *schemas.CreateRecentActionRequest) error {
	return r.db.Create(RecentAction).Error
}
func (r *RecentActionRepository) GetRecentActionByID(id uint) (*entity.RecentAction, error) {
	var RecentAction entity.RecentAction
	result := r.db.First(&RecentAction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &RecentAction, nil
}

func (r *RecentActionRepository) GetRecentActionByUserID(userID uint) ([]*entity.RecentAction, error) {
	var RecentActions []*entity.RecentAction
	result := r.db.Where("user_id = ?", userID).Find(&RecentActions)
	if result.Error != nil {
		return nil, result.Error
	}
	return RecentActions, nil
}

func (r *RecentActionRepository) GetRecentActionByType(actionType string) ([]*entity.RecentAction, error) {
	var RecentActions []*entity.RecentAction
	result := r.db.Where("action_type = ?", actionType).Find(&RecentActions)
	if result.Error != nil {
		return nil, result.Error
	}
	return RecentActions, nil
}
func (r *RecentActionRepository) ListRecentAction() ([]*entity.RecentAction, error) {
	var RecentActions []*entity.RecentAction
	result := r.db.Find(&RecentActions)
	if result.Error != nil {
		return nil, result.Error
	}
	return RecentActions, nil
}
func (r *RecentActionRepository) UpdateRecentAction(recentAction *schemas.UpdateRecentActionRequest) error {
	return r.db.Save(recentAction).Error
}
func (r *RecentActionRepository) DeleteRecentAction(id uint) error {
	return r.db.Delete(&entity.RecentAction{}, id).Error
}
