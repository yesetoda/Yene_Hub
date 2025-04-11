package repository

import "a2sv.org/hub/Domain/entity"

// AttendanceRepository defines methods for attendance database operations
type AttendanceRepository interface {
	Create(attendance *entity.Attendance) error
	GetByID(id uint) (*entity.Attendance, error)
	GetByUserID(userID uint) ([]*entity.Attendance, error)
	GetBySessionID(sessionID uint) ([]*entity.Attendance, error)
	GetByHeadID(headID uint) ([]*entity.Attendance, error)
	Update(attendance *entity.Attendance) error
	Delete(id uint) error
	List() ([]*entity.Attendance, error)
}
