package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type divisionRepository struct {
	BaseRepository
}

func NewDivisionRepository(db *gorm.DB) repository.DivisionRepository {
	return &divisionRepository{
		BaseRepository: NewBaseRepository(db, "division"),
	}
}

func (r *divisionRepository) Create(division *entity.Division) error {
	err := r.db.Create(division).Error
	if err != nil {
		return err
	}

	// Cache the newly created division
	_ = r.cacheDetail("byid", division, division.ID)
	_ = r.cacheDetail("byname", division, division.Name)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *divisionRepository) GetByID(id uint) (*entity.Division, error) {
	var division entity.Division
	err := r.getCachedDetail("byid", &division, func() error {
		return r.db.First(&division, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &division, nil
}

func (r *divisionRepository) GetByName(name string) (*entity.Division, error) {
	var division entity.Division
	err := r.getCachedDetail("byname", &division, func() error {
		return r.db.Where("name = ?", name).First(&division).Error
	}, name)
	
	if err != nil {
		return nil, err
	}
	return &division, nil
}

func (r *divisionRepository) GetAll() ([]*entity.Division, error) {
	var divisions []*entity.Division
	err := r.getCachedList("list", &divisions, func() error {
		return r.db.Find(&divisions).Error
	})
	
	return divisions, err
}

func (r *divisionRepository) Update(division *entity.Division) error {
	err := r.db.Save(division).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", division, division.ID)
	_ = r.cacheDetail("byname", division, division.Name)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *divisionRepository) Delete(id uint) error {
	var division entity.Division
	if err := r.db.First(&division, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&division).Error; err != nil {
		return err
	}

	// Invalidate all caches for this division
	r.invalidateAllCache()
	
	return nil
}
