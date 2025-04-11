package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) CreateSession(session *entity.Session) error {
	return r.db.Create(session).Error
}
func (r *SessionRepository) GetSessionByID(id uint) (*entity.Session, error) {
	var session entity.Session
	result := r.db.First(&session, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}
func (r *SessionRepository) GetSessionByStartTime(startTime string) ([]*entity.Session, error) {
	var sessions []*entity.Session
	result := r.db.Where("start_time = ?", startTime).Find(&sessions)
	if result.Error != nil {
		return nil, result.Error
	}
	return sessions, nil
}
func (r *SessionRepository) GetSessionByName(name string) ([]*entity.Session, error) {
	var sessions []*entity.Session
	result := r.db.Where("name = ?", name).Find(&sessions)
	if result.Error != nil {
		return nil, result.Error
	}
	return sessions, nil
}
func (r *SessionRepository) UpdateSession(session *entity.Session) error {
	return r.db.Save(session).Error
}

func (r *SessionRepository) DeleteSession(id uint) error {
	return r.db.Delete(&entity.Session{}, id).Error
}