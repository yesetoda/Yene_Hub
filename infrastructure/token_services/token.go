package token_services

import (
	"crypto/rand"
	"errors"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"a2sv.org/hub/Domain/entity"
)

func GenerateToken(user *entity.User, password, jwtSecret string) (string, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid user name or password")
	}

	accessToken, err := CreateJWTToken(user, jwtSecret, 24*30*time.Hour)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func CreateJWTToken(user *entity.User, jwtSecret string, duration time.Duration) (string, error) {

	expirationTime := time.Now().Add(duration)
	claims := &entity.Claims{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.Phone,
		Role:        user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetClaims(c *gin.Context) (*entity.Claims, error) {

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return &entity.Claims{}, errors.New("missing authorization header")
	}

	TokenString := strings.Split(authHeader, " ")
	if len(TokenString) != 2 || TokenString[0] != "Bearer" {
		return &entity.Claims{}, errors.New("invalid token format")
	}
	tokenString := TokenString[1]

	token, err := jwt.ParseWithClaims(tokenString, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return &entity.Claims{}, err
	}
	if claims, ok := token.Claims.(*entity.Claims); ok && token.Valid {
		return claims, err
	}
	return &entity.Claims{}, errors.New("invalid token")
}

func GenerateConfirmationToken(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range confirmationToken {
		num, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		confirmationToken[i] = charset[num.Int64()]
	}

	return string(confirmationToken), nil
}