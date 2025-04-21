package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type commentRepository struct {
	BaseRepository
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &commentRepository{
		BaseRepository: NewBaseRepository(db, "comment"),
	}
}

func (r *commentRepository) Create(comment *entity.Comment) error {
	err := r.db.Create(comment).Error
	if err != nil {
		return err
	}

	// Cache the newly created comment
	_ = r.cacheDetail("byid", comment, comment.ID)
	_ = r.cacheDetail("byuser", comment, comment.UserID)
	_ = r.cacheDetail("byproblem", comment, comment.ProblemID)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *commentRepository) GetByID(id uint) (*entity.Comment, error) {
	var comment entity.Comment
	err := r.getCachedDetail("byid", &comment, func() error {
		return r.db.First(&comment, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) GetByUserID(userID uint) ([]*entity.Comment, error) {
	var comments []*entity.Comment
	err := r.getCachedList("byuser", &comments, func() error {
		return r.db.Where("user_id = ?", userID).Find(&comments).Error
	}, userID)
	
	return comments, err
}

func (r *commentRepository) GetByProblemID(problemID uint) ([]*entity.Comment, error) {
	var comments []*entity.Comment
	err := r.getCachedList("byproblem", &comments, func() error {
		return r.db.Where("problem_id = ?", problemID).Find(&comments).Error
	}, problemID)
	
	return comments, err
}

func (r *commentRepository) List() ([]*entity.Comment, error) {
	var comments []*entity.Comment
	err := r.getCachedList("list", &comments, func() error {
		return r.db.Find(&comments).Error
	})
	
	return comments, err
}

func (r *commentRepository) Update(comment *entity.Comment) error {
	err := r.db.Save(comment).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", comment, comment.ID)
	_ = r.cacheDetail("byuser", comment, comment.UserID)
	_ = r.cacheDetail("byproblem", comment, comment.ProblemID)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *commentRepository) Delete(id uint) error {
	var comment entity.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&comment).Error; err != nil {
		return err
	}

	// Invalidate all caches for this comment
	r.invalidateAllCache()
	
	return nil
}
