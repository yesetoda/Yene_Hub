package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type SuperGroupRepository struct {
	db *gorm.DB
}

func NewSuperGroupRepository(db *gorm.DB) repository.SuperGroupRepository {
	return &SuperGroupRepository{
		db: db,
	}
}

func (r *SuperGroupRepository) CreateSuperGroup(superGroup *entity.SuperGroup) error {
	return r.db.Create(superGroup).Error
}

func (r *SuperGroupRepository) GetSuperGroupByID(id uint) (*entity.SuperGroup, error) {
	var superGroup entity.SuperGroup
	result := r.db.First(&superGroup, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &superGroup, nil
}

func (r *SuperGroupRepository) GetSuperGroupByName(name string) (*entity.SuperGroup, error) {
	var superGroups *entity.SuperGroup
	result := r.db.Where("name = ?", name).Find(&superGroups)
	if result.Error != nil {
		return nil, result.Error
	}
	return superGroups, nil
}

func (r *SuperGroupRepository) ListSuperGroup() ([]*entity.SuperGroup, error) {
	var superGroups []*entity.SuperGroup
	result := r.db.Find(&superGroups)
	if result.Error != nil {
		return nil, result.Error
	}
	return superGroups, nil
}

func (r *SuperGroupRepository) UpdateSuperGroup(superGroup *entity.SuperGroup) error {
	return r.db.Save(superGroup).Error
}

func (r *SuperGroupRepository) DeleteSuperGroup(id uint) error {
	return r.db.Delete(&entity.SuperGroup{}, id).Error
}
