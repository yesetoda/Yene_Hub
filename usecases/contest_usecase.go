package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

type ContestUsecase struct {
	contestRepo repository.ContestRepository
}

func NewContestUsecase(contestRepo repository.ContestRepository) *ContestUsecase {
	return &ContestUsecase{contestRepo: contestRepo}
}

func (u *ContestUsecase) CreateContest(contest *entity.Contest) error {
	return u.contestRepo.CreateContest(contest)
}

func (u *ContestUsecase) GetContestByID(id uint) (*entity.Contest, error) {
	return u.contestRepo.GetContestByID(id)
}	

func (u *ContestUsecase) GetContestByName(name string) (*entity.Contest, error) {
	return u.contestRepo.GetContestByName(name)
}

func (u *ContestUsecase) GetContests() ([]*entity.Contest, error) {
	return u.contestRepo.GetContests()
}

func (u *ContestUsecase) UpdateContest(contest *entity.Contest) error {
	return u.contestRepo.UpdateContest(contest)
}

func (u *ContestUsecase) DeleteContest(id uint) error {
	return u.contestRepo.DeleteContest(id)
}

