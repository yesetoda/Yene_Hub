package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type ContestRepository struct {
	db *gorm.DB
}

type contestRepository struct {
	db *gorm.DB
}

func NewContestRepository(db *gorm.DB) repository.ContestRepository {
	return &contestRepository{
		db: db,
	}
}

func (r *contestRepository) CreateContest(contest *entity.Contest) error {
	err := r.db.Create(contest).Error
	if err != nil {
		return err
	}

	// Cache the newly created contest
	// _ = r.cacheDetail("byid", contest, contest.ID)
	// _ = r.cacheDetail("byname", contest, contest.Name)
	// _ = r.cacheDetail("bygroup", contest, contest.GroupID)

	// Invalidate list caches
	// r.invalidateCache("list")
	
	return nil
}

func (r *contestRepository) GetContestByID(id uint) (*entity.Contest, error) {
	var contest entity.Contest
	// err := r.getCachedDetail("byid", &contest, func() error {
	// 	return r.db.First(&contest, id).Error
	// }, id)
	err := r.db.First(&contest, id).Error
	
	if err != nil {
		return nil, err
	}
	return &contest, nil
}

func (r *contestRepository) GetContestByName(name string) (*entity.Contest, error) {
	var contest entity.Contest
	// err := r.getCachedDetail("byname", &contest, func() error {
	// 	return r.db.Where("name = ?", name).First(&contest).Error
	// }, name)
	err := r.db.Where("name = ?", name).First(&contest).Error
	
	if err != nil {
		return nil, err
	}
	return &contest, nil
}

func (r *contestRepository) GetContests() ([]*entity.Contest, error) {
	var contests []*entity.Contest
	// err := r.getCachedList("list", &contests, func() error {
	// 	return r.db.Find(&contests).Error
	// })
	err := r.db.Find(&contests).Error
	
	return contests, err
}

func (r *contestRepository) UpdateContest(contest *entity.Contest) error {
	err := r.db.Save(contest).Error
	if err != nil {
		return err
	}

	// Update detail caches
	// _ = r.cacheDetail("byid", contest, contest.ID)
	// _ = r.cacheDetail("byname", contest, contest.Name)
	// _ = r.cacheDetail("bygroup", contest, contest.GroupID)

	// Invalidate list caches
	// r.invalidateCache("list")
	
	return nil
}

func (r *contestRepository) DeleteContest(id uint) error {
	var contest entity.Contest
	if err := r.db.First(&contest, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&contest).Error; err != nil {
		return err
	}

	// Invalidate all caches for this contest
	// r.invalidateAllCache()
	
	return nil
}
