package postgres

import (
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

type ApiTokenRepository struct {
	db *gorm.DB
}

func NewApiTokensRepository(db *gorm.DB) repository.APITokenRepository {
	return &ApiTokenRepository{db: db}
}

func (r *ApiTokenRepository) Create(apiToken *entity.APIToken) error {
	return r.db.Create(apiToken).Error
}

func (r *ApiTokenRepository) GetByID(id uint) (*entity.APIToken, error) {
	return &entity.APIToken{}, nil
}

func (r *ApiTokenRepository) List() ([]*entity.APIToken, error) {
	return []*entity.APIToken{}, nil
}

func (r *ApiTokenRepository) GetByToken(token string) (*entity.APIToken, error) {
	return &entity.APIToken{}, nil
}

func (r *ApiTokenRepository) GetByUserID(userID uint) ([]*entity.APIToken, error) {
	return []*entity.APIToken{}, nil
}

func (r *ApiTokenRepository) Update(apiToken *entity.APIToken) error {
	return r.db.Save(apiToken).Error	
}

func (r *ApiTokenRepository) Delete(id uint) error {
	return r.db.Delete(&entity.APIToken{}, id).Error
}
