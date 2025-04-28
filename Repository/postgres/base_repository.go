package postgres

import (
	"encoding/json"
	"fmt"
	"strings"

	"a2sv.org/hub/infrastructure/caching"
	"gorm.io/gorm"
)

// BaseRepository provides common caching functionality for all repositories
type BaseRepository struct {
	db         *gorm.DB
	entityName string
}

// NewBaseRepository creates a new BaseRepository instance
func NewBaseRepository(db *gorm.DB, entityName string) BaseRepository {
	return BaseRepository{
		db:         db,
		entityName: strings.ToLower(entityName),
	}
}

// getCacheKey generates a cache key for a specific entity and operation
func (r *BaseRepository) getCacheKey(operation string, params ...interface{}) string {
	key := fmt.Sprintf("%s:%s", r.entityName, operation)
	for _, param := range params {
		key = fmt.Sprintf("%s:%v", key, param)
	}
	return key
}

// cacheDetail caches a single entity
func (r *BaseRepository) cacheDetail(operation string, data interface{}, params ...interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	key := r.getCacheKey(operation, params...)
	return caching.SetRedisValue(key, string(jsonData))
}

// getCachedDetail retrieves a cached entity
func (r *BaseRepository) getCachedDetail(operation string, result interface{}, loader func() error, params ...interface{}) error {
	key := r.getCacheKey(operation, params...)

	// Try to get from cache first
	data, err := caching.GetOrSetRedisValue(key, func() (string, error) {
		if err := loader(); err != nil {
			return "", err
		}

		jsonData, err := json.Marshal(result)
		if err != nil {
			return "", err
		}

		return string(jsonData), nil
	})

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), result)
	
}

// cacheList caches a list of entities
func (r *BaseRepository) cacheList(operation string, data interface{}, params ...interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	key := r.getCacheKey(operation, params...)
	return caching.SetRedisValue(key, string(jsonData))
	// return fmt.Errorf("intensionaly disabled caching")
}

// getCachedList retrieves a cached list
func (r *BaseRepository) getCachedList(operation string, result interface{}, loader func() error, params ...interface{}) error {
	key := r.getCacheKey(operation, params...)

	data, err := caching.GetOrSetRedisValue(key, func() (string, error) {
		if err := loader(); err != nil {
			return "", err
		}

		jsonData, err := json.Marshal(result)
		if err != nil {
			return "", err
		}

		return string(jsonData), nil
	})

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), result)
}

// invalidateCache removes cache entries for the entity
func (r *BaseRepository) invalidateCache(operations ...string) {
	for _, operation := range operations {
		key := fmt.Sprintf("%s:%s:*", r.entityName, operation)
		_ = caching.DeleteRedisValue(key)
	}
}

// invalidateAllCache removes all cache entries for the entity
func (r *BaseRepository) invalidateAllCache() {
	key := fmt.Sprintf("%s:*", r.entityName)
	_ = caching.DeleteRedisValue(key)
}
