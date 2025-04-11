package repository

import "a2sv.org/hub/Domain/entity"

// CacheRepository defines methods for cache database operations
type CacheRepository interface {
	Create(cache *entity.UserCache) error
	GetByID(id uint) (*entity.UserCache, error)
	GetByIdentifier(identifier string) (*entity.UserCache, error)
	Update(cache *entity.UserCache) error
	Delete(id uint) error
	List() ([]*entity.UserCache, error)
}
