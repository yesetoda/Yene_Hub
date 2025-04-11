package repository

import (
	"a2sv.org/hub/Domain/entity"
)

// UserRepository defines methods for user data operations
type UserRepository interface {
	CreateUser(user *entity.User) error

	ListUser(page, page_size int) ([]*entity.User, error)

	GetUserByID(id uint) (*entity.User, error)
	GetUserByName(name string) ([]*entity.User, error)
	GetUserByUniversity(university string) ([]*entity.User, error)
	GetUserByCountryID(countryID uint) ([]*entity.User, error)
	GetUserByGroupID(groupID uint) ([]*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)

	UpdateUser(user *entity.User) error

	DeleteUser(id uint) error
}
