package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type attendanceRepository struct {
	BaseRepository
}

func NewAttendanceRepository(db *gorm.DB) repository.AttendanceRepository {
	return &attendanceRepository{
		BaseRepository: NewBaseRepository(db, "attendance"),
	}
}

func (r *attendanceRepository) Create(attendance *entity.Attendance) error {
	err := r.db.Create(attendance).Error
	if err != nil {
		return err
	}

	// Cache the newly created attendance
	_ = r.cacheDetail("byid", attendance, attendance.ID)
	_ = r.cacheDetail("byuser", attendance, attendance.UserID)
	_ = r.cacheDetail("byhead", attendance, attendance.HeadID)
	if attendance.SessionID != nil {
		_ = r.cacheDetail("bysession", attendance, *attendance.SessionID)
	}

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *attendanceRepository) GetByID(id uint) (*entity.Attendance, error) {
	var attendance entity.Attendance
	err := r.getCachedDetail("byid", &attendance, func() error {
		return r.db.First(&attendance, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *attendanceRepository) GetByUserID(userID uint) ([]*entity.Attendance, error) {
	var attendances []*entity.Attendance
	err := r.getCachedList("byuser", &attendances, func() error {
		return r.db.Where("user_id = ?", userID).Find(&attendances).Error
	}, userID)
	
	return attendances, err
}

func (r *attendanceRepository) GetBySessionID(sessionID uint) ([]*entity.Attendance, error) {
	var attendances []*entity.Attendance
	err := r.getCachedList("bysession", &attendances, func() error {
		return r.db.Where("session_id = ?", sessionID).Find(&attendances).Error
	}, sessionID)
	
	return attendances, err
}

func (r *attendanceRepository) GetByHeadID(headID uint) ([]*entity.Attendance, error) {
	var attendances []*entity.Attendance
	err := r.getCachedList("byhead", &attendances, func() error {
		return r.db.Where("head_id = ?", headID).Find(&attendances).Error
	}, headID)
	
	return attendances, err
}

func (r *attendanceRepository) List() ([]*entity.Attendance, error) {
	var attendances []*entity.Attendance
	err := r.getCachedList("list", &attendances, func() error {
		return r.db.Find(&attendances).Error
	})
	
	return attendances, err
}

func (r *attendanceRepository) Update(attendance *entity.Attendance) error {
	err := r.db.Save(attendance).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", attendance, attendance.ID)
	_ = r.cacheDetail("byuser", attendance, attendance.UserID)
	_ = r.cacheDetail("byhead", attendance, attendance.HeadID)
	if attendance.SessionID != nil {
		_ = r.cacheDetail("bysession", attendance, *attendance.SessionID)
	}

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *attendanceRepository) Delete(id uint) error {
	var attendance entity.Attendance
	if err := r.db.First(&attendance, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&attendance).Error; err != nil {
		return err
	}

	// Invalidate all caches for this attendance
	r.invalidateAllCache()
	
	return nil
}
