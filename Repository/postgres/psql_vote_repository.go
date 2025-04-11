package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"gorm.io/gorm"
)

type VoteRepository struct {
	db *gorm.DB
}

func NewVoteRepository(db *gorm.DB) *VoteRepository {
	return &VoteRepository{
		db: db,
	}
}

func (r *VoteRepository) CreateVote(vote *entity.Vote) error {
	return r.db.Create(vote).Error
}

func (r *VoteRepository) GetVoteByID(id uint) (*entity.Vote, error) {
	var vote entity.Vote
	result := r.db.First(&vote, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &vote, nil
}

func (r *VoteRepository) GetVoteByName(name string) ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Where("name = ?", name).Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}

func (r *VoteRepository) ListVote() ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}

func (r *VoteRepository) UpdateVote(vote *entity.Vote) error {
	return r.db.Save(vote).Error
}

func (r *VoteRepository) DeleteVote(id uint) error {
	return r.db.Delete(&entity.Vote{}, id).Error
}
