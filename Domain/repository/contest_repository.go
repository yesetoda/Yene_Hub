package repository

import "a2sv.org/hub/Domain/entity"

type ContestRepository interface {
	CreateContest(contest *entity.Contest) error
	GetContestByID(id uint) (*entity.Contest, error)
	GetContestByName(name string) (*entity.Contest, error)
	GetContests() ([]*entity.Contest, error)
	UpdateContest(contest *entity.Contest) error
	DeleteContest(id uint) error
	
}
