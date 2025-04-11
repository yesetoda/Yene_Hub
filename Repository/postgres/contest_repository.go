package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type ContestRepository struct {
	db *gorm.DB
}

func NewContestRepository(db *gorm.DB) repository.ContestRepository {
	return &ContestRepository{db: db}
}

func (r *ContestRepository) CreateContest(contest *entity.Contest) error {
	return r.db.Create(contest).Error
}

func (r *ContestRepository) GetContestByID(id uint) (*entity.Contest, error) {
	var contest entity.Contest
	if err := r.db.First(&contest, id).Error; err != nil {
		return nil, err
	}
	return &contest, nil
}

func (r *ContestRepository) GetContestByName(name string) (*entity.Contest, error) {
	var contest entity.Contest
	if err := r.db.Where("name = ?", name).First(&contest).Error; err != nil {
		return nil, err
	}
	return &contest, nil
}

func (r *ContestRepository) GetContests() ([]*entity.Contest, error) {
	var contests []*entity.Contest
	if err := r.db.Find(&contests).Error; err != nil {
		return nil, err
	}
	return contests, nil
}

func (r *ContestRepository) UpdateContest(contest *entity.Contest) error {
	return r.db.Save(contest).Error
}

func (r *ContestRepository) DeleteContest(id uint) error {
	return r.db.Delete(&entity.Contest{}, id).Error
}

