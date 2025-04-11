package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// TrackRepository defines methods for Track data operations
type TrackUseCaseInterface interface {
	CreateTrack(Track *entity.Track) error

	ListTrack() ([]*entity.Track, error)

	GetTrackByName(name string) (*entity.Track, error)
	GetTrackByID(id uint) (*entity.Track, error)

	UpdateTrack(Track *entity.Track) error

	DeleteTrack(id uint) error
}

type TrackUsecase struct {
	TrackRepository repository.TrackRepository
}

func NewTrackUsecase(trackRepository repository.TrackRepository) TrackUseCaseInterface {
	return &TrackUsecase{
		TrackRepository: trackRepository,
	}
}
func (s *TrackUsecase) CreateTrack(track *entity.Track) error {
	err := s.TrackRepository.CreateTrack(track)
	if err != nil {
		return err
	}
	return nil
}

func (s *TrackUsecase) ListTrack() ([]*entity.Track, error) {
	tracks, err := s.TrackRepository.ListTrack()
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func (s *TrackUsecase) GetTrackByName(name string) (*entity.Track, error) {
	track, err := s.TrackRepository.GetTrackByName(name)
	if err != nil {
		return nil, err
	}
	return track, nil
}
func (s *TrackUsecase) GetTrackByID(id uint) (*entity.Track, error) {
	track, err := s.TrackRepository.GetTrackByID(id)
	if err != nil {
		return nil, err
	}
	return track, nil
}
func (s *TrackUsecase) UpdateTrack(track *entity.Track) error {
	err := s.TrackRepository.UpdateTrack(track)
	if err != nil {
		return err
	}
	return nil
}
func (s *TrackUsecase) DeleteTrack(id uint) error {
	err := s.TrackRepository.DeleteTrack(id)
	if err != nil {
		return err
	}
	return nil
}