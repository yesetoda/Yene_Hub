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

// parseUintParam parses a uint parameter from the path or query and returns an error response if invalid.
func parseUintParam(c *gin.Context, param string, source string) (uint, bool) {
	var valueStr string
	if source == "path" {
		valueStr = c.Param(param)
	} else {
		valueStr = c.Query(param)
	}
	value, err := strconv.ParseUint(valueStr, 10, 32)
	if err != nil || value < 1 {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid " + param,
			Details: param + " must be a positive integer",
		})
		return 0, false
	}
	return uint(value), true
}

// respondError standardizes error responses
func respondError(c *gin.Context, code int, message, details string) {
	c.JSON(code, schemas.ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	})
}

// respondSuccess standardizes success responses
func respondSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, schemas.SuccessResponse{
		Message: message,
		Data:    data,
	})
}

// CreateUser handles creating a new user
// @Summary Create a new user
// @Description Create a new user account with the provided information
// @Tags users
// @Accept json
// @Produce json
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
		respondError(c, 400, "Invalid request format", err.Error())
		return
	}

	user, err := h.userUseCase.Create(&input)
	if err != nil {
		if err.Error() == "email already exists" {
			respondError(c, 409, "User already exists", "A user with this email already exists")
			return
		}
		respondError(c, 500, "Failed to create user", "Internal server error")
		return
	}

	respondSuccess(c, 201, "User created successfully", user)
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
	id, ok := parseUintParam(c, "id", "path")
	if !ok {
		return
	}

	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		respondError(c, 404, "User not found", "No user with the given ID")
		return
	}

	respondSuccess(c, 200, "User details retrieved successfully", user)
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
	id, ok := parseUintParam(c, "id", "path")
	if !ok {
		return
	}

	var input schemas.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		respondError(c, 400, "Invalid request format", err.Error())
		return
	}

	err := h.userUseCase.Update(id, &input)
	if err != nil {
		if err.Error() == "user not found" {
			respondError(c, 404, "User not found", "No user with the given ID")
			return
		}
		respondError(c, 500, "Failed to update user", "Internal server error")
		return
	}

	// Fetch updated user for response
	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		respondError(c, 500, "Failed to fetch updated user", "Internal server error")
		return
	}
	respondSuccess(c, 200, "User updated successfully", user)
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
	id, ok := parseUintParam(c, "id", "path")
	if !ok {
		return
	}

	if err := h.userUseCase.Delete(id); err != nil {
		respondError(c, 500, "Failed to delete user", "Internal server error")
		return
	}
	respondSuccess(c, 200, "User deleted successfully", nil)
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

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		respondError(c, 400, "Invalid page number", "Page number must be a positive integer")
		return
	}
	query.Page = page

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		respondError(c, 400, "Invalid page size", "Page size must be between 1 and 100")
		return
	}
	query.PageSize = pageSize

	query.Search = c.Query("search")

	if roleIDStr := c.Query("role_id"); roleIDStr != "" {
		roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
		if err != nil || roleID < 1 {
			respondError(c, 400, "Invalid role ID", "Role ID must be a positive integer")
			return
		}
		roleIDUint := uint(roleID)
		query.RoleID = &roleIDUint
	}

	if groupIDStr := c.Query("group_id"); groupIDStr != "" {
		groupID, err := strconv.ParseUint(groupIDStr, 10, 32)
		if err != nil || groupID < 1 {
			respondError(c, 400, "Invalid group ID", "Group ID must be a positive integer")
			return
		}
		groupIDUint := uint(groupID)
		query.GroupID = &groupIDUint
	}

	result, err := h.userUseCase.List(&query)
	if err != nil {
		respondError(c, 500, "Failed to list users", "Internal server error")
		return
	}
	respondSuccess(c, 200, "List of users retrieved successfully", result)
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
		respondError(c, 400, "Invalid request format", "Check your email and password format.")
		return
	}
	// Basic validation
	if input.Email == "" || input.Password == "" {
		respondError(c, 400, "Missing credentials", "Email and password are required.")
		return
	}

	result, err := h.userUseCase.Login(input.Email, input.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			// Optionally log invalid attempts here
			respondError(c, 401, "Invalid credentials", "Email or password is incorrect.")
			return
		}
		// Internal error
		respondError(c, 500, "Internal server error", "An unexpected error occurred.")
		return
	}
	// Optionally log successful login here
	respondSuccess(c, 200, "Login successful", result)
}
