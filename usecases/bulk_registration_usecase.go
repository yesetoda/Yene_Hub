package usecases

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/Domain/repository"
	"golang.org/x/crypto/bcrypt"
)

// BulkRegistrationUseCase defines methods for bulk user registration
type BulkRegistrationUseCase interface {
	RegisterUsers(emails string, roleID uint, groupID uint, countryID *uint) ([]RegistrationResult, error)
}

// RegistrationResult represents the result of a single user registration
type RegistrationResult struct {
	Email   string `json:"email"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	// Remove password from response struct to avoid accidentally sending it back
	UserID uint `json:"user_id,omitempty"` // Only included for successful registrations
}

// bulkRegistrationUseCase implements BulkRegistrationUseCase
type bulkRegistrationUseCase struct {
	userRepo    repository.UserRepository
	roleRepo    repository.RoleRepository
	groupRepo   repository.GroupRepository
	countryRepo repository.CountryRepository
}

// NewBulkRegistrationUseCase creates a new BulkRegistrationUseCase instance
func NewBulkRegistrationUseCase(
	userRepo repository.UserRepository,
	roleRepo repository.RoleRepository,
	groupRepo repository.GroupRepository,
	countryRepo repository.CountryRepository,
) BulkRegistrationUseCase {
	return &bulkRegistrationUseCase{
		userRepo:    userRepo,
		roleRepo:    roleRepo,
		groupRepo:   groupRepo,
		countryRepo: countryRepo,
	}
}

// RegisterUsers registers multiple users from a comma-separated list of emails
func (u *bulkRegistrationUseCase) RegisterUsers(emails string, roleID uint, groupID uint, countryID *uint) ([]RegistrationResult, error) {
	// Split emails by comma and trim whitespace
	emailList := strings.Split(emails, ",")
	var results []RegistrationResult

	// Validate role exists
	role, err := u.roleRepo.GetRoleByID(roleID)
	if err != nil {
		return nil, fmt.Errorf("invalid role ID: %v", err)
	}

	// Validate group exists
	_, err = u.groupRepo.GetByID(groupID)
	if err != nil {
		return nil, fmt.Errorf("invalid group ID: %v", err)
	}

	// Validate country exists if countryID is provided
	if countryID != nil {
		_, err = u.countryRepo.GetByID(*countryID)
		if err != nil {
			return nil, fmt.Errorf("invalid country ID: %v", err)
		}
	}

	// Process each email
	for _, email := range emailList {
		email = strings.TrimSpace(email)
		if email == "" {
			continue // Skip empty emails
		}

		result := RegistrationResult{
			Email: email,
		}

		// Check if email already exists
		_, err := u.userRepo.GetUserByEmail(email)
		if err == nil {
			result.Success = false
			result.Message = "Email already registered"
			results = append(results, result)
			continue
		}

		// Generate random password
		password, err := generateRandomPassword(12)
		if err != nil {
			result.Success = false
			result.Message = "Failed to generate password"
			results = append(results, result)
			continue
		}

		// Create new user with only required fields
		user := &entity.User{
			Email:     email,	
			Password:  password, // Will be hashed in Create method
			RoleID:    roleID,
			GroupID:   &groupID, // Use pointer to groupID
			CountryID: countryID, // Use pointer to countryID	
			Name:      extractNameFromEmail(email),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			result.Success = false
			result.Message = "Failed to hash password"
			results = append(results, result)
			continue
		}

		// Set hashed password
		user.Password = string(hashedPassword)

		// Save user to database
		err = u.userRepo.CreateUser(user)
		if err != nil {
			// Check if it's a unique constraint violation (e.g., email already exists)
			if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
				result.Success = false
				result.Message = "Email already registered"
				results = append(results, result)
				continue
			}
			result.Success = false
			result.Message = fmt.Sprintf("Failed to create user: %v", err)
			results = append(results, result)
			continue
		}

		// Verify the user was actually created
		verifiedUser, err := u.userRepo.GetUserByID(user.ID)
		if err != nil || verifiedUser == nil {
			// If we can't find the user, it wasn't actually created
			result.Success = false
			result.Message = "User creation failed: user not found after create operation"
			results = append(results, result)
			continue
		}

		// Send welcome email with password
		plainPassword := password // Store before we clear it from the result
		err = u.sendWelcomeEmail(email, plainPassword, role.Type)
		emailStatus := "Email sent successfully"
		if err != nil {
			emailStatus = fmt.Sprintf("User created but email failed: %v", err)
		}

		// Record successful registration
		result.Success = true
		result.Message = emailStatus
		result.UserID = user.ID
		results = append(results, result)

		// Log successful creation without exposing the password
		log.Printf("Successfully created user with ID %d and email %s", user.ID, email)
	}

	return results, nil
}

// generateRandomPassword generates a secure random password of the specified length
func generateRandomPassword(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes)[:length], nil
}

// extractNameFromEmail extracts and formats a name from an email address
func extractNameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) < 2 {
		return email
	}

	// Extract the part before the @ symbol
	namePart := parts[0]

	// Split by dots or underscores and capitalize each part
	var nameParts []string
	for _, part := range strings.FieldsFunc(namePart, func(r rune) bool {
		return r == '.' || r == '_' || r == '-'
	}) {
		if len(part) > 0 {
			// Capitalize first letter, keep rest lowercase
			nameParts = append(nameParts, strings.ToUpper(part[:1])+strings.ToLower(part[1:]))
		}
	}

	// Join with spaces
	return strings.Join(nameParts, " ")
}

// sendWelcomeEmail sends a welcome email to the user with their password
func (u *bulkRegistrationUseCase) sendWelcomeEmail(email, password, roleType string) error {
	// Simple implementation for now - just log the email details
	log.Printf("[EMAIL SERVICE] Sending welcome email to %s with role %s", email, roleType)
	log.Printf("[EMAIL CONTENT] Welcome to A2SV Hub! Your password is: %s", password)
	return nil
}


