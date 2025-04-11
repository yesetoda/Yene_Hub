package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"gorm.io/gorm"
)

type TrackRepository struct {
	db *gorm.DB
}

func NewTrackRepository(db *gorm.DB) *TrackRepository {
	return &TrackRepository{
		db: db,
	}
}

func (r *TrackRepository) CreateTrack(track *entity.Track) error {
	return r.db.Create(track).Error
}

func (r *TrackRepository) GetTrackByID(id uint) (*entity.Track, error) {
	var track entity.Track
	result := r.db.First(&track, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &track, nil
}

func (r *TrackRepository) GetTrackByName(name string) ([]*entity.Track, error) {
	var tracks []*entity.Track
	result := r.db.Where("name = ?", name).Find(&tracks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tracks, nil
}

func (r *TrackRepository) ListTrack() ([]*entity.Track, error) {
	var tracks []*entity.Track
	result := r.db.Find(&tracks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tracks, nil
}

func (r *TrackRepository) UpdateTrack(track *entity.Track) error {
	return r.db.Save(track).Error
}

func (r *TrackRepository) DeleteTrack(id uint) error {
	return r.db.Delete(&entity.Track{}, id).Error
}