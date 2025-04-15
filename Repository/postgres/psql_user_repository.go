package postgres

import (
	"encoding/json"
	"fmt"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"a2sv.org/hub/infrastructure/caching"
	"gorm.io/gorm"
)

// userRepository implements the repository.UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// ====================
// CREATE
// ====================

// CreateUser creates a new user and ensures that related cache entries are updated or invalidated.
func (r *userRepository) CreateUser(user *entity.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	// Marshal the created user for caching.
	UserCache := entity.UserCache{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleID:    user.RoleID,
		Inactive:  user.Inactive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	data, err := json.Marshal(UserCache)
	if err == nil {
		// Eagerly update detail caches.
		cacheKeyByID := fmt.Sprintf("getuserbyid:%d", user.ID)
		_ = caching.SetRedisValue(cacheKeyByID, string(data))

		cacheKeyByEmail := fmt.Sprintf("getuserbyemail:%s", user.Email)
		_ = caching.SetRedisValue(cacheKeyByEmail, string(data))
	}

	// Invalidate list caches since a new record might affect queries.
	_ = caching.DeleteRedisValue("listuser:*")
	_ = caching.DeleteRedisValue("getuserbyname:*")
	_ = caching.DeleteRedisValue("getuserbycountryid:*")
	_ = caching.DeleteRedisValue("getuserbygroupid:*")
	
	return nil
}

// ====================
// READ
// ====================

// GetUserByID retrieves a user by ID using a singleflight-enabled cache helper.
func (r *userRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	cacheKey := fmt.Sprintf("getuserbyid:%d", id)

	// Use our caching helper to either retrieve or load data.
	data, err := caching.GetOrSetRedisValue(cacheKey, func() (string, error) {
		result := r.db.First(&user, id)
		if result.Error != nil {
			return "", result.Error
		}
		UserCache := entity.UserCache{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			RoleID:    user.RoleID,
			Inactive:  user.Inactive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			}
		value, err := json.Marshal(UserCache)
		if err != nil {
			return "", err
		}
		return string(value), nil
	})
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByName retrieves users by name using the caching helper.
func (r *userRepository) GetUserByName(name string) ([]*entity.User, error) {
	var users []*entity.User
	cacheKey := fmt.Sprintf("getuserbyname:%s", name)

	data, err := caching.GetOrSetRedisValue(cacheKey, func() (string, error) {
		result := r.db.Where("name = ?", name).Find(&users)
		if result.Error != nil {
			return "", result.Error
		}
		UsersCache := make([]entity.UserCache, len(users))
		for i, user := range users {
			UsersCache[i] = entity.UserCache{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				RoleID:    user.RoleID,
				Inactive:  user.Inactive,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}
		}
		value, err := json.Marshal(UsersCache)
		if err != nil {
			return "", err
		}
		return string(value), nil
	})
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByUniversity retrieves users by university (no caching applied; add similar logic if needed).
func (r *userRepository) GetUserByUniversity(university string) ([]*entity.User, error) {
	var users []*entity.User
	result := r.db.Where("university = ?", university).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUserByCountryID retrieves users by country ID using the caching helper.
func (r *userRepository) GetUserByCountryID(countryID uint) ([]*entity.User, error) {
	var users []*entity.User
	cacheKey := fmt.Sprintf("getuserbycountryid:%d", countryID)

	data, err := caching.GetOrSetRedisValue(cacheKey, func() (string, error) {
		result := r.db.Where("country_id = ?", countryID).Find(&users)
		if result.Error != nil {
			return "", result.Error
		}
		UsersCache := make([]entity.UserCache, len(users))
		for i, user := range users {
			UsersCache[i] = entity.UserCache{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				RoleID:    user.RoleID,
				Inactive:  user.Inactive,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}
		}
		value, err := json.Marshal(UsersCache)
		if err != nil {
			return "", err
		}
		return string(value), nil
	})
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByGroupID retrieves users by group ID using the caching helper.
func (r *userRepository) GetUserByGroupID(groupID uint) ([]*entity.User, error) {
	var users []*entity.User
	cacheKey := fmt.Sprintf("getuserbygroupid:%d", groupID)

	data, err := caching.GetOrSetRedisValue(cacheKey, func() (string, error) {
		result := r.db.Where("group_id = ?", groupID).Find(&users)
		if result.Error != nil {
			return "", result.Error
		}
		UsersCache := make([]entity.UserCache, len(users))
		for i, user := range users {
			UsersCache[i] = entity.UserCache{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				RoleID:    user.RoleID,
				Inactive:  user.Inactive,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}
		}
		value, err := json.Marshal(UsersCache)
		if err != nil {
			return "", err
		}
		return string(value), nil
	})
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByEmail retrieves a user by email using the caching helper.
func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	cacheKey := fmt.Sprintf("getuserbyemail:%s", email)

	data, err := caching.GetOrSetRedisValue(cacheKey, func() (string, error) {
		result := r.db.Where("email = ?", email).First(&user)
		if result.Error != nil {
			return "", result.Error
		}
		UserCache := entity.UserCache{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				RoleID:    user.RoleID,
				Inactive:  user.Inactive,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}
		
		value, err := json.Marshal(UserCache)
		if err != nil {
			return "", err
		}
		return string(value), nil
	})
	if err != nil {
		return nil, err
	}
	var uc entity.UserCache
	if err := json.Unmarshal([]byte(data), &uc); err != nil {
		return nil, err
	}
	user = entity.User{
		ID:        uc.ID,
		Name:      uc.Name,
		Email:     uc.Email,
		Password:  uc.Password,
		RoleID:    uc.RoleID,
		Inactive:  uc.Inactive,
		CreatedAt: uc.CreatedAt,
		UpdatedAt: uc.UpdatedAt,
		}

	return &user, nil
}

// ListUser retrieves users with pagination using the caching helper.
func (r *userRepository) ListUser(page, page_size int) ([]*entity.User, error) {
	var users []*entity.User
	cacheKey := fmt.Sprintf("listuser:%d:%d", page, page_size)

	data, err := caching.GetOrSetRedisValue(cacheKey, func() (string, error) {
		var result *gorm.DB
		if page_size > 0 && page >= 0 {
			result = r.db.Limit(page_size).Offset(page * page_size).Find(&users)
		} else {
			result = r.db.Find(&users)
		}
		if result.Error != nil {
			return "", result.Error
		}
		UsersCache := make([]entity.UserCache, len(users))
		for i, user := range users {
			UsersCache[i] = entity.UserCache{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Password:  user.Password,
				RoleID:    user.RoleID,
				Inactive:  user.Inactive,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}
		}
		value, err := json.Marshal(UsersCache)
		if err != nil {
			return "", err
		}
		return string(value), nil
	})
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(data), &users); err != nil {
		return nil, err
	}
	return users, nil
}

// ====================
// UPDATE
// ====================

// UpdateUser updates a user and updates corresponding cache entries.
// repository/user_repository.go

func (r *userRepository) UpdateUser(user *entity.User) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}

	// Optional: update Redis cache
	userCache := entity.UserCache{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		RoleID:    user.RoleID,
		Inactive:  user.Inactive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	data, err := json.Marshal(userCache)
	if err == nil {
		cacheKeyByID := fmt.Sprintf("getuserbyid:%d", user.ID)
		_ = caching.SetRedisValue(cacheKeyByID, string(data))

		cacheKeyByEmail := fmt.Sprintf("getuserbyemail:%s", user.Email)
		_ = caching.SetRedisValue(cacheKeyByEmail, string(data))
	}

	_ = caching.DeleteRedisValue("listuser:*")
	_ = caching.DeleteRedisValue("getuserbyname:*")
	_ = caching.DeleteRedisValue("getuserbycountryid:*")
	_ = caching.DeleteRedisValue("getuserbygroupid:*")

	return nil
}


// ====================
// DELETE
// ====================

// DeleteUser deletes a user and invalidates affected cache entries.
func (r *userRepository) DeleteUser(id uint) error {
	err := r.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}

	// Invalidate detail cache for the given id.
	cacheKeyByID := fmt.Sprintf("getuserbyid:%d", id)
	_ = caching.DeleteRedisValue(cacheKeyByID)

	// Invalidate list caches—if you maintain other keys that reference this user,
	// consider invalidating those as well.
	_ = caching.DeleteRedisValue("listuser:*")
	_ = caching.DeleteRedisValue("getuserbyname:*")
	_ = caching.DeleteRedisValue("getuserbycountryid:*")
	_ = caching.DeleteRedisValue("getuserbygroupid:*")

	// Optionally, if you store a cache keyed by email, you would also invalidate it.
	// This requires you have the user's email, which might be looked up prior to deletion.

	return nil
}
