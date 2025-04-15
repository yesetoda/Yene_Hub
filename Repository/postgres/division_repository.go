package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type DivisionRepository struct {
	db *gorm.DB
}

func NewDivisionRepository(db *gorm.DB) repository.DivisionRepository {
	return &DivisionRepository{db: db}
}

func (r *DivisionRepository) Create(division *entity.Division) error {
	return r.db.Create(division).Error
}

func (r *DivisionRepository) GetByID(id uint) (*entity.Division, error) {
	var division entity.Division
	if err := r.db.First(&division, id).Error; err != nil {
		return nil, err
	}
	return &division, nil
}

func (r *DivisionRepository) GetByName(name string) (*entity.Division, error) {
	var division entity.Division
	if err := r.db.Where("name = ?", name).First(&division).Error; err != nil {
		return nil, err
	}
	return &division, nil
}	

func (r *DivisionRepository) GetAll() ([]*entity.Division, error) {
	var divisions []*entity.Division
	if err := r.db.Find(&divisions).Error; err != nil {
		return nil, err
	}	
	return divisions, nil
}

func (r *DivisionRepository) Update(division *entity.Division) error {
	return r.db.Save(division).Error
}

func (r *DivisionRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Division{}, id).Error
}
