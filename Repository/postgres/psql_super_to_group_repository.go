package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"gorm.io/gorm"
)

type SuperToGroupRepository struct {
	db *gorm.DB
}

func NewSuperToGroupRepository(db *gorm.DB) *SuperToGroupRepository {
	return &SuperToGroupRepository{
		db: db,
	}
}

func (r *SuperToGroupRepository) CreateSuperToGroup(superToGroup *entity.SuperToGroup) error {
	return r.db.Create(superToGroup).Error
}

func (r *SuperToGroupRepository) GetSuperToGroupByID(id uint) (*entity.SuperToGroup, error) {
	var superToGroup entity.SuperToGroup
	result := r.db.First(&superToGroup, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &superToGroup, nil
}

func (r *SuperToGroupRepository) ListSuperToGroup() ([]*entity.SuperToGroup, error) {
	var superToGroups []*entity.SuperToGroup
	result := r.db.Find(&superToGroups)
	if result.Error != nil {
		return nil, result.Error
	}
	return superToGroups, nil
}

func (r *SuperToGroupRepository) UpdateSuperToGroup(superToGroup *entity.SuperToGroup) error {
	return r.db.Save(superToGroup).Error
}

func (r *SuperToGroupRepository) DeleteSuperToGroup(id uint) error {
	return r.db.Delete(&entity.SuperToGroup{}, id).Error
}
