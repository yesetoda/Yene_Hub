package usecases

import (
	"errors"
	"fmt"
	"os"
	"time"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"a2sv.org/hub/infrastructure/email_services"
	"a2sv.org/hub/infrastructure/password_services"
	"a2sv.org/hub/infrastructure/token_services"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase defines methods for user business logic
type UserUseCaseInterface interface {
	Create(input *schemas.CreateUserRequest) (*schemas.UserResponse, error)
	GetByID(id uint) (*schemas.UserResponse, error)
	GetByEmail(email string) (*schemas.UserResponse, error)
	Update(uid uint, input *schemas.UpdateUserRequest) error
	Delete(id uint) error
	List(query *schemas.UserListQuery) (*schemas.UserListResponse, error)
	Login(email, password string) (*schemas.LoginResponse, error)
}

// UserUseCase implements UserUseCase
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase creates a new UserUseCase instance
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// Create creates a new user
func (u *UserUseCase) Create(input *schemas.CreateUserRequest) (*schemas.UserResponse, error) {
	// Check if email already exists
	existingUser, err := u.userRepo.GetUserByEmail(input.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Create user entity from input
	user := &entity.User{
		Email: input.Email,
		Name:  input.Name,
	}

	// Set optional fields if provided
	if input.RoleID != nil {
		user.RoleID = *input.RoleID
	}
	if input.GroupID != nil {
		user.GroupID = input.GroupID
	}
	if input.CountryID != nil {
		user.CountryID = input.CountryID
	}
	if input.University != nil {
		user.University = *input.University
	}
	if input.StudentID != nil {
		user.StudentID = *input.StudentID
	}
	if input.Department != nil {
		user.Department = *input.Department
	}
	if input.ExpectedGraduationDate != nil {
		user.ExpectedGraduationDate = input.ExpectedGraduationDate
	}
	if input.Phone != nil {
		user.Phone = *input.Phone
	}
	if input.TelegramUsername != nil {
		user.TelegramUsername = input.TelegramUsername
	}
	if input.TelegramUID != nil {
		user.TelegramUID = *input.TelegramUID
	}
	if input.Leetcode != nil {
		user.Leetcode = input.Leetcode
	}
	if input.Codeforces != nil {
		user.Codeforces = input.Codeforces
	}
	if input.Github != nil {
		user.Github = input.Github
	}
	if input.Hackerrank != nil {
		user.Hackerrank = input.Hackerrank
	}
	if input.Linkedin != nil {
		user.Linkedin = *input.Linkedin
	}
	if input.Instagram != nil {
		user.Instagram = *input.Instagram
	}
	if input.Birthday != nil {
		user.Birthday = input.Birthday
	}
	if input.Gender != nil {
		user.Gender = *input.Gender
	}
	if input.ShortBio != nil {
		user.ShortBio = *input.ShortBio
	}
	if input.PreferredLanguage != nil {
		user.PreferredLanguage = *input.PreferredLanguage
	}

	// Generate random password
	userPassword, err := password_services.GenerateRandomPassword(12)
	if err != nil {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Save user
	if err := u.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	// Send welcome email with password
	emailContent := fmt.Sprintf("Welcome to A2SV Hub!\n\nYour account has been created successfully.\nYour temporary password is: %s\n\nPlease change your password after logging in.", userPassword)
	if err := email_services.SendEmail(user.Email, "Welcome to A2SV Hub", emailContent, "yene-hub-ls0y.onrender.com/api/auth/login"); err != nil {
		// Log the error but don't fail the request
		fmt.Fprintf(os.Stderr, "Failed to send welcome email: %v\n", err)
	}

	// Convert to response
	return u.entityToResponse(user), nil
}

// GetByID retrieves a user by ID
func (u *UserUseCase) GetByID(id uint) (*schemas.UserResponse, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return u.entityToResponse(user), nil
}

// GetByEmail retrieves a user by email
func (u *UserUseCase) GetByEmail(email string) (*schemas.UserResponse, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return u.entityToResponse(user), nil
}

// Update updates a user
func (u *UserUseCase) Update(uid uint, input *schemas.UpdateUserRequest) error {
	user, err := u.userRepo.GetUserByID(uid)
	if err != nil {
		return err
	}

	// Apply only non-nil fields
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.RoleID != nil {
		user.RoleID = *input.RoleID
	}
	if input.GroupID != nil {
		user.GroupID = input.GroupID
	}
	if input.CountryID != nil {
		user.CountryID = input.CountryID
	}
	if input.University != nil {
		user.University = *input.University
	}
	if input.StudentID != nil {
		user.StudentID = *input.StudentID
	}
	if input.Department != nil {
		user.Department = *input.Department
	}
	if input.ExpectedGraduationDate != nil {
		user.ExpectedGraduationDate = input.ExpectedGraduationDate
	}
	if input.Phone != nil {
		user.Phone = *input.Phone
	}
	if input.TelegramUsername != nil {
		user.TelegramUsername = input.TelegramUsername
	}
	if input.Leetcode != nil {
		user.Leetcode = input.Leetcode
	}
	if input.Codeforces != nil {
		user.Codeforces = input.Codeforces
	}
	if input.Github != nil {
		user.Github = input.Github
	}
	if input.Hackerrank != nil {
		user.Hackerrank = input.Hackerrank
	}
	if input.Linkedin != nil {
		user.Linkedin = *input.Linkedin
	}
	if input.Instagram != nil {
		user.Instagram = *input.Instagram
	}
	if input.Birthday != nil {
		user.Birthday = input.Birthday
	}
	if input.Gender != nil {
		user.Gender = *input.Gender
	}
	if input.ShortBio != nil {
		user.ShortBio = *input.ShortBio
	}
	if input.PreferredLanguage != nil {
		user.PreferredLanguage = *input.PreferredLanguage
	}

	// Save updated user
	return u.userRepo.UpdateUser(user)
}

// Delete deletes a user
func (u *UserUseCase) Delete(id uint) error {
	return u.userRepo.DeleteUser(id)
}

// List retrieves all users with pagination and filtering
func (u *UserUseCase) List(query *schemas.UserListQuery) (*schemas.UserListResponse, error) {
	// Set default values if not provided
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}

	// Get users with filters
	users, total, err := u.userRepo.ListUsers(query)
	if err != nil {
		return nil, err
	}

	// Convert entities to responses
	var responses []*schemas.UserResponse
	for _, user := range users {
		responses = append(responses, u.entityToResponse(user))
	}

	return &schemas.UserListResponse{
		Data: responses,
		Meta: schemas.PaginationMeta{
			Page:     query.Page,
			PageSize: query.PageSize,
			Total:    total,
		},
	}, nil
}

// Login handles user authentication
func (u *UserUseCase) Login(email, password string) (*schemas.LoginResponse, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := token_services.CreateJWTToken(user, os.Getenv("JWT_SECRET"), time.Hour*24)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	return &schemas.LoginResponse{
		Token: token,
		User:  u.entityToResponse(user),
	}, nil
}

// Helper function to convert *string to string, returning empty string if nil
func stringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Helper function to convert entity to response
func (u *UserUseCase) entityToResponse(user *entity.User) *schemas.UserResponse {
	response := &schemas.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		GroupID:   user.GroupID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,

		// Academic Information
		University:             &user.University,
		StudentID:              &user.StudentID,
		Department:             &user.Department,
		ExpectedGraduationDate: user.ExpectedGraduationDate,

		// Contact Information
		Phone:            &user.Phone,
		TelegramUsername: user.TelegramUsername,

		// Coding Profiles
		Leetcode:   user.Leetcode,
		Codeforces: user.Codeforces,
		Github:     user.Github,
		Hackerrank: user.Hackerrank,

		// Social Media
		Linkedin:  &user.Linkedin,
		Instagram: &user.Instagram,

		// Professional Details
		Gender:            &user.Gender,
		ShortBio:          &user.ShortBio,
		PreferredLanguage: &user.PreferredLanguage,
	}

	return response
}
