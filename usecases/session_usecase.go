package usecases

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
)

// SessionRepository defines methods for Session data operations
type SessionUseCaseInterface interface {
	CreateSession(Session *entity.Session) error

	ListSession() ([]*entity.Session, error)

	GetSessionByName(name string) ([]*entity.Session, error)
	GetSessionByID(id uint) (*entity.Session, error)
	GetSessionByStartTime(startTime string) ([]*entity.Session, error)

	UpdateSession(Session *entity.Session) error

	DeleteSession(id uint) error
}

type SessionUsecase struct {
	SessionRepository repository.SessionRepository
}

func NewSessionUsecase(sessionRepository repository.SessionRepository) *SessionUsecase {
	return &SessionUsecase{
		SessionRepository: sessionRepository,
	}
}

func (s *SessionUsecase) CreateSession(session *entity.Session) error {
	err := s.SessionRepository.CreateSession(session)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionUsecase) ListSession() ([]*entity.Session, error) {
	sessions, err := s.SessionRepository.ListSession()
	if err != nil {
		return nil, err
	}
	return sessions, nil
}
func (s *SessionUsecase) GetSessionByName(name string) ([]*entity.Session, error) {
	sessions, err := s.SessionRepository.GetSessionByName(name)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *SessionUsecase) GetSessionByID(id uint) (*entity.Session, error) {
	session, err := s.SessionRepository.GetSessionByID(id)
	if err != nil {
		return nil, err
	}
	return session, nil
}
func (s *SessionUsecase) GetSessionByStartTime(startTime string) ([]*entity.Session, error) {
	sessions, err := s.SessionRepository.GetSessionByStartTime(startTime)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *SessionUsecase) UpdateSession(session *entity.Session) error {
	err := s.SessionRepository.UpdateSession(session)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionUsecase) DeleteSession(id uint) error {
	return s.SessionRepository.DeleteSession(id)
}
