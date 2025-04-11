package entity

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID          uint   `json:"id" gorm:"primaryKey" `
	Name        string `json:"name" gorm:"size:255" `
	Email       string `json:"email" gorm:"size:255" `
	PhoneNumber string `json:"phone_number" gorm:"size:255" `
	Role        *Role `json:"role" gorm:"size:255" `
	jwt.StandardClaims
}
