package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type FundRepository struct {
	db *gorm.DB
}

func NewFundRepository(db *gorm.DB) repository.FundRepository {
	return &FundRepository{db: db}
}

func (r *FundRepository) Create(fund *entity.Fund) error {
	return r.db.Create(fund).Error
}	

func (r *FundRepository) GetByID(id uint) (*entity.Fund, error) {
	var fund entity.Fund
	if err := r.db.First(&fund, id).Error; err != nil {
		return nil, err
	}	
	return &fund, nil
}

func (r *FundRepository) GetByUserID(userID uint) (*entity.Fund, error) {
	var fund entity.Fund
	if err := r.db.Where("user_id = ?", userID).First(&fund).Error; err != nil {
		return nil, err
	}
	return &fund, nil
}

func (r *FundRepository) GetAll() ([]*entity.Fund, error) {	
	var funds []*entity.Fund
	if err := r.db.Find(&funds).Error; err != nil {
		return nil, err
	}
	return funds, nil
}	

func (r *FundRepository) Update(fund *entity.Fund) error {
	return r.db.Save(fund).Error
}

func (r *FundRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Fund{}, id).Error
}






