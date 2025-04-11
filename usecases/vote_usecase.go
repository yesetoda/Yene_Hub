package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// VoteRepository defines methods for Vote data operations
type VoteUseCaseInterface interface {
	CreateVote(Vote *entity.Vote) error

	ListVote() ([]*entity.Vote, error)

	GetVoteByCommentID(commentID uint) ([]*entity.Vote, error)
	GetVoteByPostID(postID uint) ([]*entity.Vote, error)
	GetVoteByUserID(userID uint) ([]*entity.Vote, error)
	GetVoteByTrackID(trackID uint) ([]*entity.Vote, error)
	GetVoteBySubmissionID(submissionID uint) ([]*entity.Vote, error)
	GetVoteByProblemID(problemID uint) ([]*entity.Vote, error)
	GetVoteByID(id uint) (*entity.Vote, error)

	UpdateVote(Vote *entity.Vote) error

	DeleteVote(id uint) error
}

type VoteUsecase struct {
	VoteRepository repository.VoteRepository
}
func NewVoteUsecase(voteRepository repository.VoteRepository) VoteUseCaseInterface {
	return &VoteUsecase{
		VoteRepository: voteRepository,
	}
}

func (v *VoteUsecase) CreateVote(vote *entity.Vote) error {
	err := v.VoteRepository.CreateVote(vote)
	if err != nil {
		return err
	}
	return nil
}

func (v *VoteUsecase) ListVote() ([]*entity.Vote, error) {
	votes, err := v.VoteRepository.ListVote()
	if err != nil {
		return nil, err
	}
	return votes, nil
}


func (v *VoteUsecase) GetVoteByCommentID(commentID uint) ([]*entity.Vote, error) {
	votes, err := v.VoteRepository.GetVoteByCommentID(commentID)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (v *VoteUsecase) GetVoteByPostID(postID uint) ([]*entity.Vote, error) {

	votes, err := v.VoteRepository.GetVoteByPostID(postID)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (v *VoteUsecase) GetVoteByUserID(userID uint) ([]*entity.Vote, error) {
	votes, err := v.VoteRepository.GetVoteByUserID(userID)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (v *VoteUsecase) GetVoteByTrackID(trackID uint) ([]*entity.Vote, error) {
	votes, err := v.VoteRepository.GetVoteByTrackID(trackID)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (v *VoteUsecase) GetVoteBySubmissionID(submissionID uint) ([]*entity.Vote, error) {
	votes, err := v.VoteRepository.GetVoteBySubmissionID(submissionID)
	if err != nil {
		return nil, err
	}
	return votes, nil
}

func (v *VoteUsecase) GetVoteByProblemID(problemID uint) ([]*entity.Vote, error) {
	votes, err := v.VoteRepository.GetVoteByProblemID(problemID)
	if err != nil {
		return nil, err
	}
	return votes, nil
}
func (v *VoteUsecase) GetVoteByID(id uint) (*entity.Vote, error) {
	vote, err := v.VoteRepository.GetVoteByID(id)
	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (v *VoteUsecase) UpdateVote(vote *entity.Vote) error {
	return v.VoteRepository.UpdateVote(vote)
}
func (v *VoteUsecase) DeleteVote(id uint) error {
	err := v.VoteRepository.DeleteVote(id)
	if err != nil {
		return err
	}
	return nil
}