package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

type ProblemTracksUsecaseInterface interface {
	AddProblemToTrack(trackID , problemID uint) error
	ListProblemsInTrack(trackID uint) ([]*entity.Problem, error)
	GetProblemInTracksByName(trackID uint, name string) (*entity.Problem, error)
	GetProblemInTracksByDifficulty(trackID uint, difficulty string) ([]*entity.Problem, error)
	GetProblemInTracksByTag(trackID uint, tag string) ([]*entity.Problem, error)
	GetProblemInTracksByPlatform(trackID uint, platform string) ([]*entity.Problem, error)
	RemoveProblemFromTrack(id uint) error
}

type ProblemTracksUsecase struct {
	repo repository.ProblemInTracksRepository
}

func NewProblemTracksUsecase(repo repository.ProblemInTracksRepository) *ProblemTracksUsecase {
	return &ProblemTracksUsecase{repo: repo}
}

func (u *ProblemTracksUsecase) AddProblemToTrack(trackID uint, problemID uint) error {
	return u.repo.AddProblemToTrack(trackID, problemID)
}

func (u *ProblemTracksUsecase) ListProblemsInTrack(trackID uint) ([]*entity.Problem, error) {
	return u.repo.ListProblemsInTrack(trackID)
}

func (u *ProblemTracksUsecase) GetProblemInTracksByName(trackID uint, name string) (*entity.Problem, error) {
	return u.repo.GetProblemInTracksByName(trackID, name)
}

func (u *ProblemTracksUsecase) GetProblemInTracksByDifficulty(trackID uint, difficulty string) ([]*entity.Problem, error) {
	return u.repo.GetProblemInTracksByDifficulty(trackID, difficulty)
}

func (u *ProblemTracksUsecase) GetProblemInTracksByTag(trackID uint, tag string) ([]*entity.Problem, error) {
	return u.repo.GetProblemInTracksByTag(trackID, tag)
}

func (u *ProblemTracksUsecase) GetProblemInTracksByPlatform(trackID uint, platform string) ([]*entity.Problem, error) {
	return u.repo.GetProblemInTracksByPlatform(trackID, platform)
}

func (u *ProblemTracksUsecase) RemoveProblemFromTrack(id uint) error {
	return u.repo.RemoveProblemFromTrack(id)
}
