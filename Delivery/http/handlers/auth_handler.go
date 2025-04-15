package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

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
// @Description Redirects to Google's OAuth2 authentication page
// @Tags Authentication
// @Produce json
// @Success 302 "Redirect to Google OAuth"
// @Failure 500 {object} schemas.ErrorResponse "OAuth configuration error"
// @Router /api/auth/google [get]
func InitGoogleOAuth(c *gin.Context) {
	url := fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=email profile",
		os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
	)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleGoogleCallback handles the OAuth callback and only allows login for registered users.
// HandleGoogleCallback handles the OAuth callback
// @Summary Google OAuth callback
// @Description Handle Google OAuth callback and return JWT token
// @Tags Authentication
// @Produce json
// @Description [Login with Google](http://localhost:8080/api/auth/google)
// @Param code query string true "Authorization code from Google"
// @Param state query string false "OAuth state parameter"
// @Param error query string false "OAuth error description"
// @Success 200 {object} schemas.AuthTokenResponse "Authentication token"
// @Failure 400 {object} schemas.ErrorResponse "Missing code parameter"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized/Unverified email"
// @Failure 500 {object} schemas.ErrorResponse "Token exchange error"
// @Router /api/auth/google/callback [get]
func HandleGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No code provided"})
		return
	}

	// Exchange the authorization code for an access token.
	token, err := oauth.ExchangeCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange error"})
		return
	}

	// Retrieve user info from Google.
	userInfo, err := oauth.GetUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch user info"})
		return
	}

	// Extract email from userInfo.
	email, ok := userInfo["email"].(string)
	if !ok || email == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid email in Google response"})
		return
	}

	// (Optional) You may want to check if the email is verified:
	verified, ok := userInfo["verified_email"].(bool)
	if !ok || !verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not verified by Google"})
		return
	}

	// Look up the user by email in the database.
	existingUser, err := UserRepo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database lookup error"})
		return
	}

	// If no user is found, login is not allowed.
	if existingUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not registered with this service"})
		return
	}

	// Generate a JWT token for the user.
	jwtToken, err := token_services.CreateJWTToken(existingUser, os.Getenv("JWT_SECRET"), 72*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
