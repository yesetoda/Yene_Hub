package middleware

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/infrastructure/token_services"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserUsecases usecases.UserUseCaseInterface
	RoleUsecases usecases.RoleUseCaseInterface
}
func NewRoleMiddleware(userUsecase usecases.UserUseCaseInterface, roleUsecase usecases.RoleUseCaseInterface) AuthController {
	return AuthController{
		UserUsecases: userUsecase,
		RoleUsecases: roleUsecase,
	}
}

// func (ac *AuthController) RoleAuthenticationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				c.JSON(http.StatusForbidden, gin.H{"error": "unexpected error"})
// 				c.Abort()
// 			}
// 		}()
// 		claims, err := token_services.GetClaims(c)
// 		if err != nil {

// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			c.Abort()
// 			return
// 		}
// 		_, err = ac.UserUsecases.GetByID(claims.ID)
// 		if err != nil {

// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			c.Abort()
// 			return
// 		}

//			c.Set("claims", claims)
//			c.Next()
//		}
//	}
func (ac *AuthController) RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
					Code:    500,
					Message: "unexpected error",
					Details: "unexpected error",
				})
				c.Abort()
			}
		}()
		// Extract claims from token
		claims, err := token_services.GetClaims(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, schemas.ErrorResponse{
				Code:    404,
				Message: "invalid token",
				Details: err.Error(),
			})
			return
		}
		// Fetch User details
		User, err := ac.UserUsecases.GetByID(claims.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, schemas.ErrorResponse{
				Code:    404,
				Message: "error trying to find  the User record",
				Details: err.Error(),
			})
			return
		}
		if User == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, schemas.ErrorResponse{
				Code:    404,
				Message: "User record is missing",
				Details: "no such user",
			})
			return
		}
		if role != "" {
			userRole, err := ac.RoleUsecases.GetByID(User.RoleID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, schemas.ErrorResponse{
					Code:    404,
					Message: "error find role record",
					Details: err.Error(),
				})
				return
			}
			if userRole.Type != role {
				c.AbortWithStatusJSON(http.StatusForbidden, schemas.ErrorResponse{
					Code:    404,
					Message: "unauthorized User",
					Details: "unauthorized",
				})
				return
			}
		}
		c.Next()
	}
}

func (ac *AuthController) SelfMiddleware() gin.HandlerFunc {
	var id uint64
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
					Code:    500,
					Message: "unexpected error",
					Details: "unexpected error",
				})
				c.Abort()
			}
		}()
		uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, schemas.ErrorResponse{
				Code:    400,
				Message: "Invalid user ID",
				Details: "User ID must be a positive integer",
			})
			return
		}
		id = uid
		// Extract claims from token
		claims, err := token_services.GetClaims(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, schemas.ErrorResponse{
				Code:    404,
				Message: "invalid token",
				Details: err.Error(),
			})
			return
		}
		// Fetch User details
		User, err := ac.UserUsecases.GetByID(claims.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, schemas.ErrorResponse{
				Code:    404,
				Message: "error trying to find  the User record",
				Details: err.Error(),
			})
			return
		}
		if User == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, schemas.ErrorResponse{
				Code:    404,
				Message: "User record is missing",
				Details: "no such user",
			})
			return
		}
		if id != uint64(claims.ID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, schemas.ErrorResponse{
				Code:    404,
				Message: "you're not the user",
				Details: "you're not the user ",
			})
			return
		}
		c.Next()
	}
}
