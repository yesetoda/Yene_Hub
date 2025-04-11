package repository

import "a2sv.org/hub/Domain/entity"

// CommentRepository defines methods for comment database operations
type CommentRepository interface {
	Create(comment *entity.Comment) error
	GetByID(id uint) (*entity.Comment, error)
	GetByUserID(userID uint) ([]*entity.Comment, error)
	Update(comment *entity.Comment) error
	Delete(id uint) error
	List() ([]*entity.Comment, error)
}
