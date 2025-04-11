package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"gorm.io/gorm"
)

type StipendRepository struct {
	db *gorm.DB
}

func NewStipendRepository(db *gorm.DB) *StipendRepository {
	return &StipendRepository{
		db: db,
	}
}

func (r *StipendRepository) CreateStipend(stipend *entity.Stipend) error {
	return r.db.Create(stipend).Error
}

func (r *StipendRepository) GetStipendByID(id uint) (*entity.Stipend, error) {
	var stipend entity.Stipend
	result := r.db.First(&stipend, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stipend, nil
}
func (r *StipendRepository) ListStipends() ([]*entity.Stipend, error) {
	var stipends []*entity.Stipend
	result := r.db.Find(&stipends)
	if result.Error != nil {
		return nil, result.Error
	}
	return stipends, nil
}

func (r *StipendRepository) UpdateStipend(stipend *entity.Stipend) error {
	return r.db.Save(stipend).Error
}
func (r *StipendRepository) DeleteStipend(id uint) error {
	return r.db.Delete(&entity.Stipend{}, id).Error
}