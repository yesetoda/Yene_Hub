package password_services

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateRandomPassword returns a strong random password of the given length.
func GenerateRandomPassword(length int) (string, error) {
	const passwordChars = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%^&*()-_=+[]{}<>?,."

	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(passwordChars)))
	for i := range password {
		num, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = passwordChars[num.Int64()]
	}
	return string(password), nil
}
