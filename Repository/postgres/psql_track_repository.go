package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type trackRepository struct {
	BaseRepository
}

func NewTrackRepository(db *gorm.DB) repository.TrackRepository {
	return &trackRepository{
		BaseRepository: NewBaseRepository(db, "track"),
	}
}

func (r *trackRepository) CreateTrack(track *entity.Track) error {
	err := r.db.Create(track).Error
	if err != nil {
		return err
	}

	// Cache the newly created track
	_ = r.cacheDetail("byid", track, track.ID)
	_ = r.cacheDetail("byname", track, track.Name)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *trackRepository) GetTrackByID(id uint) (*entity.Track, error) {
	var track entity.Track
	err := r.getCachedDetail("byid", &track, func() error {
		return r.db.First(&track, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &track, nil
}

func (r *trackRepository) GetTrackByName(name string) (*entity.Track, error) {
	var track entity.Track
	err := r.getCachedDetail("byname", &track, func() error {
		return r.db.Where("name = ?", name).First(&track).Error
	}, name)
	
	if err != nil {
		return nil, err
	}
	return &track, nil
}

func (r *trackRepository) UpdateTrack(track *entity.Track) error {
	err := r.db.Save(track).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", track, track.ID)
	_ = r.cacheDetail("byname", track, track.Name)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *trackRepository) DeleteTrack(id uint) error {
	var track entity.Track
	if err := r.db.First(&track, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&track).Error; err != nil {
		return err
	}

	// Invalidate all caches for this track
	r.invalidateAllCache()
	
	return nil
}

func (r *trackRepository) ListTrack() ([]*entity.Track, error) {
	var tracks []*entity.Track
	err := r.getCachedList("list", &tracks, func() error {
		return r.db.Find(&tracks).Error
	})
	
	return tracks, err
}
