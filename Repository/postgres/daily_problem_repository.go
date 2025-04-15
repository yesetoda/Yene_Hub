package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type DailyProblemRepository struct {
	db *gorm.DB
}

func NewDailyProblemRepository(db *gorm.DB) repository.DailyProblemsRepository {
	return &DailyProblemRepository{db: db}
}

func (r *DailyProblemRepository) Create(dailyProblem *entity.DailyProblem) error {
	return r.db.Create(dailyProblem).Error
}

func (r *DailyProblemRepository) GetByID(id uint) (*entity.DailyProblem, error) {
	var dailyProblem entity.DailyProblem
	if err := r.db.First(&dailyProblem, id).Error; err != nil {
		return nil, err
	}
	return &dailyProblem, nil
}

func (r *DailyProblemRepository) GetByProblemID(problemID uint) (*entity.DailyProblem, error) {
	var dailyProblem entity.DailyProblem
	if err := r.db.Where("problem_id = ?", problemID).First(&dailyProblem).Error; err != nil {
		return nil, err
	}
	return &dailyProblem, nil
}	

func (r *DailyProblemRepository) GetBySuperGroupID(superGroupID uint) ([]*entity.DailyProblem, error) {
	var dailyProblem []*entity.DailyProblem
	if err := r.db.Where("super_group_id = ?", superGroupID).Find(&dailyProblem).Error; err != nil {
		return nil, err
	}
	return dailyProblem, nil
}
func (r *DailyProblemRepository) List() ([]*entity.DailyProblem, error) {
	var dailyProblems []*entity.DailyProblem
	if err := r.db.Find(&dailyProblems).Error; err != nil {
		return nil, err
	}
	return dailyProblems, nil
}

func (r *DailyProblemRepository) Update(dailyProblem *entity.DailyProblem) error {
	return r.db.Save(dailyProblem).Error
}

func (r *DailyProblemRepository) Delete(id uint) error {
	return r.db.Delete(&entity.DailyProblem{}, id).Error
}
