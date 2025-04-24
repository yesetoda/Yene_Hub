// Package handlers provides HTTP request handlers for the application
// @title A2SV Hub API
// @version 1.0
// @description API server for A2SV Hub
// @BasePath /api
package handlers

import (
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userUseCase usecases.UserUseCase
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// CreateUser handles creating a new user
// @Summary Create a new user
// @Description Create a new user account with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body schemas.CreateUserRequest true "User creation data"
// @Success 201 {object} schemas.SuccessResponse "User created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request format or validation error"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 409 {object} schemas.ErrorResponse "Conflict - User already exists"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input schemas.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid request format",
			Details: err.Error(),
		})
		return
	}

	user, err := h.userUseCase.Create(&input)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(409, schemas.ErrorResponse{
				Code:    409,
				Message: "User already exists",
				Details: "A user with this email already exists",
			})
			return
		}
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to create user",
			Details: err.Error(),
		})
		return
	}

	c.JSON(201, schemas.SuccessResponse{
		Message: "User created successfully",
		Data:    user,
	})
}

// GetUserByID handles retrieving a user by ID
// @Summary Get user details
// @Description Get detailed information about a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "User ID" minimum(1)
// @Success 200 {object} schemas.SuccessResponse "User details retrieved successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid user ID format"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 404 {object} schemas.ErrorResponse "User not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid user ID",
			Details: "User ID must be a positive integer",
		})
		return
	}

	user, err := h.userUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{
			Code:    404,
			Message: "User not found",
			Details: err.Error(),
		})
		return
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "User details retrieved successfully",
		Data:    user,
	})
}

// UpdateUser handles updating a user's information
// @Summary Update user details
// @Description Update an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "User ID" minimum(1)
// @Param request body schemas.UpdateUserRequest true "User update data"
// @Success 200 {object} schemas.SuccessResponse "User updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request format or validation error"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 404 {object} schemas.ErrorResponse "User not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users/{id} [patch]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid user ID",
			Details: "User ID must be a positive integer",
		})
		return
	}

	var input schemas.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid request format",
			Details: err.Error(),
		})
		return
	}

	err = h.userUseCase.Update(uint(id), &input)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(404, schemas.ErrorResponse{
				Code:    404,
				Message: "User not found",
				Details: err.Error(),
			})
			return
		}
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to update user",
			Details: err.Error(),
		})
		return
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "User updated successfully",
		Data:    input,
	})
}

// DeleteUser handles deleting a user
// @Summary Delete user
// @Description Delete an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "User ID" minimum(1)
// @Success 200 {object} schemas.SuccessResponse "User deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid user ID format"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 404 {object} schemas.ErrorResponse "User not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid user ID",
			Details: "User ID must be a positive integer",
		})
		return
	}

	if err := h.userUseCase.Delete(uint(id)); err != nil {
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to delete user",
			Details: err.Error(),
		})
		return
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "User deleted successfully",
		Data:    nil,
	})
}

// ListUsers handles listing users with pagination and filters
// @Summary List users
// @Description Get a paginated list of users with optional filters
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param page query int false "Page number" minimum(1) default(1)
// @Param page_size query int false "Number of items per page" minimum(1) maximum(100) default(10)
// @Param search query string false "Search term for filtering users"
// @Param role_id query int false "Filter by role ID" minimum(1)
// @Param group_id query int false "Filter by group ID" minimum(1)
// @Success 200 {object} schemas.SuccessResponse "List of users retrieved successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid query parameters"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	var query schemas.UserListQuery

	// Parse page parameter
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid page number",
			Details: "Page number must be a positive integer",
		})
		return
	}
	query.Page = page

	// Parse page_size parameter
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid page size",
			Details: "Page size must be between 1 and 100",
		})
		return
	}
	query.PageSize = pageSize

	// Parse search parameter
	query.Search = c.Query("search")

	// Parse role_id parameter
	if roleIDStr := c.Query("role_id"); roleIDStr != "" {
		roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
		if err != nil {
			c.JSON(400, schemas.ErrorResponse{
				Code:    400,
				Message: "Invalid role ID",
				Details: "Role ID must be a positive integer",
			})
			return
		}
		roleIDUint := uint(roleID)
		query.RoleID = &roleIDUint
	}

	// Parse group_id parameter
	if groupIDStr := c.Query("group_id"); groupIDStr != "" {
		groupID, err := strconv.ParseUint(groupIDStr, 10, 32)
		if err != nil {
			c.JSON(400, schemas.ErrorResponse{
				Code:    400,
				Message: "Invalid group ID",
				Details: "Group ID must be a positive integer",
			})
			return
		}
		groupIDUint := uint(groupID)
		query.GroupID = &groupIDUint
	}

	result, err := h.userUseCase.List(&query)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to list users",
			Details: err.Error(),
		})
		return
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "List of users retrieved successfully",
		Data:    result,
	})
}

// Login handles user authentication
// @Summary Login user
// @Description Authenticate a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schemas.LoginRequest true "Login credentials"
// @Success 200 {object} schemas.SuccessResponse "Login successful"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request format"
// @Failure 401 {object} schemas.ErrorResponse "Invalid credentials"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var input schemas.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid request format",
			Details: err.Error(),
		})
		return
	}

	result, err := h.userUseCase.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(401, schemas.ErrorResponse{
			Code:    401,
			Message: "Invalid credentials",
			Details: err.Error(),
		})
		return
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "Login successful",
		Data:    result,
	})
}
