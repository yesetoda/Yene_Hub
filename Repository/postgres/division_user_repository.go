package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type DivisionUserRepository struct {
	db *gorm.DB
}

func NewDivisionUserRepository(db *gorm.DB) repository.DivisionUserRepository {
	return &DivisionUserRepository{db: db}
}

func (r *DivisionUserRepository) Create(divisionUser *entity.DivisionUser) error {
	return r.db.Create(divisionUser).Error
}

func (r *DivisionUserRepository) GetByID(id uint) (*entity.DivisionUser, error) {
	var divisionUser entity.DivisionUser
	if err := r.db.First(&divisionUser, id).Error; err != nil {
		return nil, err
	}
	return &divisionUser, nil
}

func (r *DivisionUserRepository) GetByUserID(userID uint) ([]*entity.DivisionUser, error) {
	var divisionUsers []*entity.DivisionUser
	if err := r.db.Where("user_id = ?", userID).Find(&divisionUsers).Error; err != nil {
		return nil, err
	}
	return divisionUsers, nil
}

func (r *DivisionUserRepository) GetByDivisionID(divisionID uint) ([]*entity.DivisionUser, error) {
	var divisionUsers []*entity.DivisionUser
	if err := r.db.Where("division_id = ?", divisionID).Find(&divisionUsers).Error; err != nil {
		return nil, err
	}
	return divisionUsers, nil
}

func (r *DivisionUserRepository) GetByContestID(contestID uint) ([]*entity.DivisionUser, error) {
	var divisionUsers []*entity.DivisionUser
	if err := r.db.Where("contest_id = ?", contestID).Find(&divisionUsers).Error; err != nil {
		return nil, err
	}
	return divisionUsers, nil
}

func (r *DivisionUserRepository) GetAll() ([]*entity.DivisionUser, error) {
	var divisionUsers []*entity.DivisionUser
	if err := r.db.Find(&divisionUsers).Error; err != nil {
		return nil, err
	}
	return divisionUsers, nil
}

func (r *DivisionUserRepository) Update(divisionUser *entity.DivisionUser) error {
	return r.db.Save(divisionUser).Error
}

func (r *DivisionUserRepository) Delete(id uint) error {
	return r.db.Delete(&entity.DivisionUser{}, id).Error
}
