package repository

import "a2sv.org/hub/Domain/entity"

type DivisionUserRepository interface {
	Create(divisionUser *entity.DivisionUser) error
	GetByID(id uint) (*entity.DivisionUser, error)
	GetByUserID(userID uint) ([]*entity.DivisionUser, error)
	GetByContestID(contestID uint) ([]*entity.DivisionUser, error)
	GetByDivisionID(divisionID uint) ([]*entity.DivisionUser, error)
	GetAll() ([]*entity.DivisionUser, error)
	Update(divisionUser *entity.DivisionUser) error
	Delete(id uint) error
}

