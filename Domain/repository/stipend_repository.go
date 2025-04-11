package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// StipendRepository defines methods for Stipend data operations
type StipendRepository interface {
	CreateStipend(Stipend *entity.Stipend) error

	ListStipend() ([]*entity.Stipend, error)
	GetStipendByID(id uint) (*entity.Stipend, error)

	UpdateStipend(Stipend *entity.Stipend) error

	DeleteStipend(id uint) error
}
