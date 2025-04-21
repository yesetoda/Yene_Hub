package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type apiTokenRepository struct {
	BaseRepository
}

func NewAPITokenRepository(db *gorm.DB) repository.APITokenRepository {
	return &apiTokenRepository{
		BaseRepository: NewBaseRepository(db, "api_token"),
	}
}

func (r *apiTokenRepository) Create(token *entity.APIToken) error {
	err := r.db.Create(token).Error
	if err != nil {
		return err
	}

	// Cache the newly created token
	_ = r.cacheDetail("byid", token, token.ID)
	_ = r.cacheDetail("bytoken", token, token.Token)
	_ = r.cacheDetail("byuser", token, token.UserID)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *apiTokenRepository) GetByID(id uint) (*entity.APIToken, error) {
	var token entity.APIToken
	err := r.getCachedDetail("byid", &token, func() error {
		return r.db.First(&token, id).Error
	}, id)
	
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *apiTokenRepository) GetByToken(tokenStr string) (*entity.APIToken, error) {
	var token entity.APIToken
	err := r.getCachedDetail("bytoken", &token, func() error {
		return r.db.Where("token = ?", tokenStr).First(&token).Error
	}, tokenStr)
	
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *apiTokenRepository) GetByUserID(userID uint) ([]*entity.APIToken, error) {
	var tokens []*entity.APIToken
	err := r.getCachedList("byuser", &tokens, func() error {
		return r.db.Where("user_id = ?", userID).Find(&tokens).Error
	}, userID)
	
	return tokens, err
}

func (r *apiTokenRepository) List() ([]*entity.APIToken, error) {
	var tokens []*entity.APIToken
	err := r.getCachedList("list", &tokens, func() error {
		return r.db.Find(&tokens).Error
	})
	
	return tokens, err
}

func (r *apiTokenRepository) Update(token *entity.APIToken) error {
	err := r.db.Save(token).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", token, token.ID)
	_ = r.cacheDetail("bytoken", token, token.Token)
	_ = r.cacheDetail("byuser", token, token.UserID)

	// Invalidate list caches
	r.invalidateCache("list")
	
	return nil
}

func (r *apiTokenRepository) Delete(id uint) error {
	var token entity.APIToken
	if err := r.db.First(&token, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&token).Error; err != nil {
		return err
	}

	// Invalidate all caches for this token
	r.invalidateAllCache()
	
	return nil
}
