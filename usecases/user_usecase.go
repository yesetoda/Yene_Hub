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
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Update(uid uint,user *schemas.UpdateUserInput) error
	Delete(id uint) error
	List(page, page_size int) ([]*entity.User, error)
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
func (u *UserUseCase) Create(user *entity.User) error {
	// Check if email already exists
	existingUser, err := u.userRepo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	// Hash password
	userPassword, err := password_services.GenerateRandomPassword(12)
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update user fields
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Create user
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	email_services.SendEmail(user.Email, "Welcome to A2SV", "Your password is: "+userPassword, "localhost:8080/api/auth/login")
	return nil

}

// GetByID retrieves a user by ID
func (u *UserUseCase) GetByID(id uint) (*entity.User, error) {
	return u.userRepo.GetUserByID(id)
}

// GetByEmail retrieves a user by email
func (u *UserUseCase) GetByEmail(email string) (*entity.User, error) {
	return u.userRepo.GetUserByEmail(email)
}

// Update updates a user
func (u *UserUseCase) Update(uid uint, input *schemas.UpdateUserInput) error {
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

	user.UpdatedAt = time.Now()
	return u.userRepo.UpdateUser(user)
}

// Delete deletes a user
func (u *UserUseCase) Delete(id uint) error {
	return u.userRepo.DeleteUser(id)
}

// List retrieves all users
func (u *UserUseCase) List(page, page_size int) ([]*entity.User, error) {
	return u.userRepo.ListUser(page, page_size)
}

func (u *UserUseCase) Login(email, password string) (*entity.User, string, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", errors.New("user not found")
	}
	// Removed logging of sensitive password details to prevent security vulnerabilities.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid password")
	}
	token, err := token_services.CreateJWTToken(user, os.Getenv("JWT_SECRET"), time.Hour*24)
	fmt.Println("this is the token and the error", token, err)
	if err != nil {
		return nil, "", err
	}
	return user, token, nil
}
