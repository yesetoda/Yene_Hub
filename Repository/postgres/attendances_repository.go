package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)	
type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) repository.AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) Create(attendance *entity.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *AttendanceRepository) GetByID(id uint) (*entity.Attendance, error) {
	var attendance entity.Attendance
	if err := r.db.First(&attendance, id).Error; err != nil {
		return nil, err
	}
	return &attendance, nil	
}

func (r *AttendanceRepository) GetByUserID(userID uint) ([]*entity.Attendance, error) {
	var attendances []*entity.Attendance
	if err := r.db.Where("user_id = ?", userID).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}	

func (r *AttendanceRepository) GetBySessionID(sessionID uint) ([]*entity.Attendance, error) {
	var attendances []*entity.Attendance
	if err := r.db.Where("session_id = ?", sessionID).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}

func (r *AttendanceRepository) GetByHeadID(headID uint) ([]*entity.Attendance, error) {
	var attendances []*entity.Attendance
	if err := r.db.Where("head_id = ?", headID).Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}
func (r *AttendanceRepository) List() ([]*entity.Attendance, error) {
	var attendances []*entity.Attendance
	if err := r.db.Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}	

func (r *AttendanceRepository) Update(attendance *entity.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *AttendanceRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Attendance{}, id).Error
}	








