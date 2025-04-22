package handlers

import (
	"fmt"
	"net/http"
	"os"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/infrastructure/oauth"
	"a2sv.org/hub/infrastructure/token_services"
	"github.com/gin-gonic/gin"
)

// UserRepo is a global variable for accessing user data.
// Ensure it is set in main.go.
var UserRepo interface {
	GetUserByEmail(email string) (*entity.User, error)
}

// InitGoogleOAuth initiates the Google OAuth flow
// @Summary Start Google OAuth
// @Description Initiates the OAuth2 flow by redirecting to Google's authentication page
// @Tags auth
// @Accept json
// @Produce json
// @Success 302 {string} string "Redirect to Google OAuth"
// @Failure 500 {object} schemas.ErrorResponse "Failed to initiate OAuth flow"
// @Router /auth/google [get]
func InitGoogleOAuth(c *gin.Context) {
	url := fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=email profile",
		os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
	)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleGoogleCallback handles the OAuth callback and only allows login for registered users
// @Summary Handle Google OAuth callback
// @Description Process Google OAuth callback, verify user, and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param code query string true "OAuth2 authorization code from Google"
// @Param state query string false "OAuth state for CSRF protection"
// @Param error query string false "Error message from OAuth provider"
// @Success 200 {object} schemas.AuthTokenResponse "JWT token for authenticated user"
// @Failure 400 {object} schemas.ErrorResponse "Missing or invalid authorization code"
// @Failure 401 {object} schemas.ErrorResponse "User not registered or email not verified"
// @Failure 403 {object} schemas.ErrorResponse "Invalid OAuth state"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /auth/google/callback [get]
func HandleGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Missing authorization code",
			Details: "The authorization code from Google is required",
		})
		return
	}

	// Exchange the authorization code for an access token
	token, err := oauth.ExchangeCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Token exchange failed",
			Details: err.Error(),
		})
		return
	}

	// Retrieve user info from Google
	userInfo, err := oauth.GetUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve user info",
			Details: err.Error(),
		})
		return
	}

	email, ok := userInfo["email"].(string)
	if !ok || email == "" {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Invalid email in Google response",
			Details: "Email field is missing or invalid in Google user info",
		})
		return
	}

	user, err := UserRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "User not registered or email not verified",
			Details: "User not found or not verified in the system",
		})
		return
	}

	// Generate JWT token for the user
	jwtToken, err := token_services.GenerateToken(user, user.Email, fmt.Sprintf("%d", user.RoleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate JWT token",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, schemas.AuthTokenResponse{
		Token: jwtToken,
	})
}
