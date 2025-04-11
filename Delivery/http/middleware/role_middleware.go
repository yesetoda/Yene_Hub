package middleware

import (
	"fmt"
	"net/http"

	"a2sv.org/hub/infrastructure/token_services"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserUsecases usecases.UserUseCaseInterface
	RoleUsecases usecases.RoleUseCaseInterface
}

func NewRoleMiddleware(userUsecase usecases.UserUseCaseInterface,roleUsecase usecases.RoleUseCaseInterface) AuthController {
	return AuthController{
		UserUsecases: userUsecase,
		RoleUsecases: roleUsecase,
	}
}

func (ac *AuthController) RoleAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": "unexpected error"})
				c.Abort()
			}
		}()
		claims, err := token_services.GetClaims(c)
		if err != nil {
			fmt.Println("error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		_, err = ac.UserUsecases.GetByID(claims.ID)
		if err != nil {
			fmt.Println("error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
func (ac *AuthController) RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error occurred"})
				c.Abort()
			}
		}()

		fmt.Println("middleware", role)

		// Extract claims from token
		claims, err := token_services.GetClaims(c)
		fmt.Println("claims", claims, claims.ID, err)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// Fetch User details
		fmt.Println("trying to find the User", claims.ID)
		User, err := ac.UserUsecases.GetByID(claims.ID)
		fmt.Println("User", User, err)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Check if User is nil
		if User == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User record is missing"})
			return
		}
		userRole,err := ac.RoleUsecases.GetByID(User.RoleID)
		fmt.Println("userRole", userRole, err)
		fmt.Println("userRole", userRole.Type, role)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}
		
		if userRole.Type != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		c.Next()
	}
}
