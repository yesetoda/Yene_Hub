package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}
func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &CommentRepository{db: db}
}
func (r *CommentRepository) Create(comment *entity.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetByID(id uint) (*entity.Comment, error) {
	var comment entity.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) GetByUserID(userID uint) ([]*entity.Comment, error) {
	var comments []*entity.Comment
	if err := r.db.Where("user_id = ?", userID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}	

func (r *CommentRepository) List() ([]*entity.Comment, error) {
	var comments []*entity.Comment
	if err := r.db.Find(&comments).Error; err != nil {
		return nil, err
	}	
	return comments, nil
}

func (r *CommentRepository) Update(comment *entity.Comment) error {
	return r.db.Save(comment).Error
}

func (r *CommentRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Comment{}, id).Error
}
