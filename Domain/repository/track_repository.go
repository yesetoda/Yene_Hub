package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// TrackRepository defines methods for Track data operations
type TrackRepository interface {
	CreateTrack(Track *entity.Track) error

	ListTrack() ([]*entity.Track, error)

	GetTrackByName(name string) (*entity.Track, error)
	GetTrackByID(id uint) (*entity.Track, error)

	UpdateTrack(Track *entity.Track) error

	DeleteTrack(id uint) error
}
