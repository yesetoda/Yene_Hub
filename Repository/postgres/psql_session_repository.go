package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type sessionRepository struct {
	BaseRepository
}

func NewSessionRepository(db *gorm.DB) repository.SessionRepository {
	return &sessionRepository{
		BaseRepository: NewBaseRepository(db, "session"),
	}
}

func (r *sessionRepository) CreateSession(session *entity.Session) error {
	err := r.db.Create(session).Error
	if err != nil {
		return err
	}

	// Cache the newly created session
	_ = r.cacheDetail("byid", session, session.ID)
	_ = r.cacheDetail("byname", session, session.Name)
	_ = r.cacheDetail("bystarttime", session, session.StartTime)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *sessionRepository) ListSession() ([]*entity.Session, error) {
	var sessions []*entity.Session
	err := r.getCachedList("list", &sessions, func() error {
		return r.db.Find(&sessions).Error
	})
	
	return sessions, err
}

func (r *sessionRepository) GetSessionByID(id uint) (*entity.Session, error) {
	var session entity.Session
	err := r.getCachedDetail("byid", &session, func() error {
		return r.db.First(&session, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *sessionRepository) GetSessionByStartTime(startTime string) ([]*entity.Session, error) {
	var sessions []*entity.Session
	err := r.getCachedList("bystarttime", &sessions, func() error {
		return r.db.Where("start_time = ?", startTime).Find(&sessions).Error
	}, startTime)
	
	return sessions, err
}

func (r *sessionRepository) GetSessionByName(name string) ([]*entity.Session, error) {
	var sessions []*entity.Session
	err := r.getCachedList("byname", &sessions, func() error {
		return r.db.Where("name = ?", name).Find(&sessions).Error
	}, name)
	
	return sessions, err
}

func (r *sessionRepository) UpdateSession(session *entity.Session) error {
	err := r.db.Save(session).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", session, session.ID)
	_ = r.cacheDetail("byname", session, session.Name)
	_ = r.cacheDetail("bystarttime", session, session.StartTime)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *sessionRepository) DeleteSession(id uint) error {
	var session entity.Session
	if err := r.db.First(&session, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&session).Error; err != nil {
		return err
	}

	// Invalidate all caches for this session
	r.invalidateAllCache()
	
	return nil
}
