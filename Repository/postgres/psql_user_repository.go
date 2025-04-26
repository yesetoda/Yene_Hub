package postgres

import (
	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"gorm.io/gorm"
)

// userRepository implements the repository.UserRepository interface
type userRepository struct {
	BaseRepository
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository(db, "user"),
	}
}

// CreateUser creates a new user and ensures that related cache entries are updated or invalidated.
func (r *userRepository) CreateUser(user *entity.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	// Cache the newly created user, including the password
	_ = r.cacheDetail("byid", user, user.ID)
	_ = r.cacheDetail("byemail", user, user.Email)
	_ = r.cacheDetail("password", user.Password, user.Email)

	// Invalidate list caches
	r.invalidateCache("list", "byname", "bycountryid", "bygroupid")

	return nil
}

func (r *userRepository) ListUsers(query *schemas.UserListQuery) ([]*entity.User, int, error) {
	var users []*entity.User
	err := r.getCachedList("list", &users, func() error {
		return r.db.Order("created_at desc").Find(&users).Error
	}, query)

	return users, 0, err
}

// GetUserByID retrieves a user by ID using the cache
func (r *userRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.getCachedDetail("byid", &user, func() error {
		return r.db.First(&user, id).Error
	}, id)
	if err != nil {
		return nil, err
	}
	// Ensure password is present (fetch from password cache if missing)
	if user.Password == "" {
		var pw string
		_ = r.getCachedDetail("password", &pw, func() error {
			return r.db.Model(&user).Select("password").Where("id = ?", id).Scan(&pw).Error
		}, user.Email)
		user.Password = pw
	}
	return &user, nil
}

// GetUserByName retrieves users by name using the cache
func (r *userRepository) GetUserByName(name string) ([]*entity.User, error) {
	var users []*entity.User
	err := r.getCachedList("byname", &users, func() error {
		return r.db.Where("name LIKE ?", "%"+name+"%").Find(&users).Error
	}, name)

	return users, err
}

// GetUserByUniversity retrieves users by university
func (r *userRepository) GetUserByUniversity(university string) ([]*entity.User, error) {
	var users []*entity.User
	err := r.getCachedList("byuniversity", &users, func() error {
		return r.db.Where("university = ?", university).Find(&users).Error
	}, university)

	return users, err
}

// GetUserByCountryID retrieves users by country ID using the cache
func (r *userRepository) GetUserByCountryID(countryID uint) ([]*entity.User, error) {
	var users []*entity.User
	err := r.getCachedList("bycountryid", &users, func() error {
		return r.db.Where("country_id = ?", countryID).Find(&users).Error
	}, countryID)

	return users, err
}

// GetUserByGroupID retrieves users by group ID using the cache
func (r *userRepository) GetUserByGroupID(groupID uint) ([]*entity.User, error) {
	var users []*entity.User
	err := r.getCachedList("bygroupid", &users, func() error {
		return r.db.Where("group_id = ?", groupID).Find(&users).Error
	}, groupID)

	return users, err
}

// GetUserByEmail retrieves a user by email using the cache
func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.getCachedDetail("byemail", &user, func() error {
		return r.db.Where("email = ?", email).First(&user).Error
	}, email)
	if err != nil {
		return nil, err
	}
	// Ensure password is present (fetch from password cache if missing)
	if user.Password == "" {
		var pw string
		_ = r.getCachedDetail("password", &pw, func() error {
			return r.db.Model(&user).Select("password").Where("email = ?", email).Scan(&pw).Error
		}, email)
		user.Password = pw
	}
	return &user, nil
}

// ListUser retrieves users with pagination using the cache
func (r *userRepository) ListUser(page, pageSize int) ([]*entity.User, error) {
	var users []*entity.User
	offset := (page - 1) * pageSize

	err := r.getCachedList("list", &users, func() error {
		return r.db.Offset(offset).Limit(pageSize).Find(&users).Error
	}, page, pageSize)

	return users, err
}

// UpdateUser updates a user and invalidates affected cache entries
func (r *userRepository) UpdateUser(user *entity.User) error {
	err := r.db.Save(user).Error
	if err != nil {
		return err
	}

	// Update detail caches
	_ = r.cacheDetail("byid", user, user.ID)
	_ = r.cacheDetail("byemail", user, user.Email)

	// Invalidate list caches
	r.invalidateCache("list", "byname", "bycountryid", "bygroupid")

	return nil
}

// DeleteUser deletes a user and invalidates affected cache entries
func (r *userRepository) DeleteUser(id uint) error {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}

	// Invalidate all caches for this user
	r.invalidateAllCache()

	return nil
}
