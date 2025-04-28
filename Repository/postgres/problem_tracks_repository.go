package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type ProblemInTracksRepository struct {
	db *gorm.DB
}

func NewProblemInTracksRepository(db *gorm.DB) repository.ProblemInTracksRepository {
	return &ProblemInTracksRepository{
		db: db,
	}
}

// AddProblemToTrack adds a problem to a track (creates a ProblemTrack row)
func (r *ProblemInTracksRepository) AddProblemToTrack(trackID uint, problemID uint) error {
	pt := &entity.ProblemTrack{
		ProblemID: problemID,
		TrackID:   trackID,
	}
	return r.db.Create(pt).Error
}

// ListProblemsInTrack lists all problems in a track
func (r *ProblemInTracksRepository) ListProblemsInTrack(trackID uint) ([]*entity.Problem, error) {
	var pts []entity.ProblemTrack
	err := r.db.Preload("Problem").Where("track_id = ?", trackID).Find(&pts).Error
	if err != nil {
		return nil, err
	}
	problems := make([]*entity.Problem, 0, len(pts))
	for _, pt := range pts {
		if pt.Problem != nil {
			problems = append(problems, pt.Problem)
		}
	}
	return problems, nil
}

// GetProblemInTracksByName gets a ProblemTrack by problem name in a track
func (r *ProblemInTracksRepository) GetProblemInTracksByName(trackID uint, name string) (*entity.Problem, error) {
	var pt entity.Problem
	err := r.db.Preload("Problem").Where("track_id = ?", trackID).Joins("JOIN problems ON problems.id = problem_tracks.problem_id").Where("problems.name = ?", name).First(&pt).Error
	if err != nil {
		return nil, err
	}
	return &pt, nil
}

// GetProblemInTracksByDifficulty gets ProblemTracks by problem difficulty in a track
func (r *ProblemInTracksRepository) GetProblemInTracksByDifficulty(trackID uint, difficulty string) ([]*entity.Problem, error) {
	var pts []entity.Problem
	err := r.db.Preload("Problem").Where("track_id = ?", trackID).Joins("JOIN problems ON problems.id = problem_tracks.problem_id").Where("problems.difficulty = ?", difficulty).Find(&pts).Error
	if err != nil {
		return nil, err
	}
	result := make([]*entity.Problem, 0, len(pts))
	for i := range pts {
		result = append(result, &pts[i])
	}
	return result, nil
}

// GetProblemInTracksByTag gets ProblemTracks by problem tag in a track
func (r *ProblemInTracksRepository) GetProblemInTracksByTag(trackID uint, tag string) ([]*entity.Problem, error) {
	var pts []entity.Problem
	err := r.db.Preload("Problem").Where("track_id = ?", trackID).Joins("JOIN problems ON problems.id = problem_tracks.problem_id").Where("problems.tag = ?", tag).Find(&pts).Error
	if err != nil {
		return nil, err
	}
	result := make([]*entity.Problem, 0, len(pts))
	for i := range pts {
		result = append(result, &pts[i])
	}
	return result, nil
}

// GetProblemInTracksByPlatform gets ProblemTracks by problem platform in a track
func (r *ProblemInTracksRepository) GetProblemInTracksByPlatform(trackID uint, platform string) ([]*entity.Problem, error) {
	var pts []entity.Problem
	err := r.db.Preload("Problem").Where("track_id = ?", trackID).Joins("JOIN problems ON problems.id = problem_tracks.problem_id").Where("problems.platform = ?", platform).Find(&pts).Error
	if err != nil {
		return nil, err
	}
	result := make([]*entity.Problem, 0, len(pts))
	for i := range pts {
		result = append(result, &pts[i])
	}
	return result, nil
}

// RemoveProblemFromTrack removes a problem from a track (deletes ProblemTrack row)
func (r *ProblemInTracksRepository) RemoveProblemFromTrack(id uint) error {
	return r.db.Delete(&entity.ProblemTrack{}, id).Error
}
