package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type divisionUserRepository struct {
	BaseRepository
}

func NewDivisionUserRepository(db *gorm.DB) repository.DivisionUserRepository {
	return &divisionUserRepository{
		BaseRepository: NewBaseRepository(db, "division_user"),
	}
}

func (r *divisionUserRepository) Create(divisionUser *entity.DivisionUser) error {
	err := r.db.Create(divisionUser).Error
	if err != nil {
		return err
	}

	// Cache the newly created division user
	_ = r.cacheDetail("byid", divisionUser, divisionUser.ID)
	_ = r.cacheDetail("byuser", divisionUser, divisionUser.UserID)
	_ = r.cacheDetail("bydivision", divisionUser, divisionUser.DivisionID)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *divisionUserRepository) GetByID(id uint) (*entity.DivisionUser, error) {
	var divisionUser entity.DivisionUser
	err := r.getCachedDetail("byid", &divisionUser, func() error {
		return r.db.First(&divisionUser, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &divisionUser, nil
}

func (r *divisionUserRepository) GetByUserID(userID uint) ([]*entity.DivisionUser, error) {
	var divisionUsers []*entity.DivisionUser
	err := r.getCachedList("byuser", &divisionUsers, func() error {
		return r.db.Where("user_id = ?", userID).Find(&divisionUsers).Error
	}, userID)
	
	return divisionUsers, err
}

func (r *divisionUserRepository) GetByDivisionID(divisionID uint) ([]*entity.DivisionUser, error) {
	var divisionUsers []*entity.DivisionUser
	err := r.getCachedList("bydivision", &divisionUsers, func() error {
		return r.db.Where("division_id = ?", divisionID).Find(&divisionUsers).Error
	}, divisionID)
	
	return divisionUsers, err
}

func (r *divisionUserRepository) GetByContestID(contestID uint) ([]*entity.DivisionUser, error) {
	var divisionUsers []*entity.DivisionUser
	if err := r.db.Where("contest_id = ?", contestID).Find(&divisionUsers).Error; err != nil {
		return nil, err
	}
	return divisionUsers, nil
}

func (r *divisionUserRepository) GetAll() ([]*entity.DivisionUser, error) {
	var divisionUsers []*entity.DivisionUser
	err := r.getCachedList("list", &divisionUsers, func() error {
		return r.db.Find(&divisionUsers).Error
	})
	
	return divisionUsers, err
}

func (r *divisionUserRepository) Update(divisionUser *entity.DivisionUser) error {
	err := r.db.Save(divisionUser).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", divisionUser, divisionUser.ID)
	_ = r.cacheDetail("byuser", divisionUser, divisionUser.UserID)
	_ = r.cacheDetail("bydivision", divisionUser, divisionUser.DivisionID)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *divisionUserRepository) Delete(id uint) error {
	var divisionUser entity.DivisionUser
	if err := r.db.First(&divisionUser, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&divisionUser).Error; err != nil {
		return err
	}

	// Invalidate all caches for this division user
	r.invalidateAllCache()
	
	return nil
}
