package repository

import "a2sv.org/hub/Domain/entity"

// APITokenRepository defines methods for API token database operations
type APITokenRepository interface {
	Create(token *entity.APIToken) error
	GetByID(id uint) (*entity.APIToken, error)
	GetByToken(token string) (*entity.APIToken, error) 
	GetByUserID(userID uint) ([]*entity.APIToken, error)
	Update(token *entity.APIToken) error
	Delete(id uint) error
	List() ([]*entity.APIToken, error)
}
