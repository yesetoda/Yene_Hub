package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// VoteRepository defines methods for Vote data operations
type VoteRepository interface {
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
