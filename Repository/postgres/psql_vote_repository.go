package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type VoteRepository struct {
	db *gorm.DB
}

// GetVoteByUserID implements repository.VoteRepository.

func NewVoteRepository(db *gorm.DB) repository.VoteRepository {
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
func (r *VoteRepository) GetVoteBySuperGroupID(superGroupID uint) ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Where("super_group_id = ?", superGroupID).Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}
func (r *VoteRepository) GetVoteByCommentID(commentId uint) ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Where("comment_id = ?", commentId).Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}
func (r *VoteRepository) GetVoteByPostID(postId uint) ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Where("post_id = ?", postId).Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}
func (r *VoteRepository) GetVoteBySubmissionID(submissionId uint) ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Where("submission_id = ?", submissionId).Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}
func (r *VoteRepository) GetVoteByProblemID(problemId uint) ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Where("problem_id = ?", problemId).Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}
func (r *VoteRepository) GetVoteByTrackID(trackId uint) ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Where("track_id = ?", trackId).Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}
	return votes, nil
}
func (r *VoteRepository) GetVoteByUserID(userID uint) ([]*entity.Vote, error) {
	var votes []*entity.Vote
	result := r.db.Where("user_id = ?", userID).Find(&votes)
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
