package repository

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Delivery/http/schemas"
)

// UserRepository defines methods for user data operations
type UserRepository interface {
	// Create method
	CreateUser(user *entity.User) error

	// List methods
	ListUser(page, page_size int) ([]*entity.User, error)
	ListUsers(query *schemas.UserListQuery) ([]*entity.User, int, error)

	// Get methods
	GetUserByID(id uint) (*entity.User, error)
	GetUserByName(name string) ([]*entity.User, error)
	GetUserByUniversity(university string) ([]*entity.User, error)
	GetUserByCountryID(countryID uint) ([]*entity.User, error)
	GetUserByGroupID(groupID uint) ([]*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)

	// Update and Delete methods
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
}
