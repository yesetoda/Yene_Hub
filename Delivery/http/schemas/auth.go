package schemas

// LoginRequest represents the login credentials
// swagger:model
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"mypassword123"`
}

// LoginResponse represents the successful login response
// swagger:model
type LoginResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  *UserResponse `json:"user"`
}

// AuthTokenResponse represents a JWT token response
// swagger:model
type AuthTokenResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
