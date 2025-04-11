package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// SessionRepository defines methods for Session data operations
type SessionRepository interface {
	CreateSession(Session *entity.Session) error

	ListSession() ([]*entity.Session, error)

	GetSessionByName(name string) ([]*entity.Session, error)
	GetSessionByID(id uint) (*entity.Session, error)
	GetSessionByStartTime(startTime string) ([]*entity.Session, error)
	

	UpdateSession(Session *entity.Session) error

	DeleteSession(id uint) error
}
