package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"a2sv.org/hub/Domain/entity"
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
// @Description Create a new user account with required information
// @Tags Users
// @Accept json
// @Produce json
// @Param user body entity.User true "User creation data"
// @Security BearerAuth
// @Success 201 {object} entity.User "Successfully created user"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request format"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized"
// @Failure 409 {object} schemas.ErrorResponse "User already exists"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.userUseCase.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// GetUserByID handles getting a user by ID
// @Summary Get user details
// @Description Get detailed information about a specific user
// @Tags Users
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Security BearerAuth
// @Success 200 {object} entity.User "User details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid user ID"
// @Failure 404 {object} schemas.ErrorResponse "User not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser handles updating a user
// @Summary Update user details
// @Description Update existing user information. Only provided fields will be updated.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(uint32)
// @Param user body entity.User true "Partial user data for update"
// @Security BearerAuth
// @Success 200 {object} schemas.SuccessResponse "User updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid ID format or request body"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden"
// @Failure 404 {object} schemas.ErrorResponse "User not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user.ID = uint(id)
	if err := h.userUseCase.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}
// DeleteUser handles deleting a user
// @Summary Delete user account
// @Description Permanently delete a user account and associated data
// @Tags Users
// @Produce json
// @Param id path int true "User ID" Format(uint32)
// @Security BearerAuth
// @Success 200 {object} schemas.SuccessResponse "User deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid ID format"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden"
// @Failure 404 {object} schemas.ErrorResponse "User not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userUseCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

// ListUsers handles listing all users
// @Summary List users
// @Description Get paginated list of users with optional filters
// @Tags Users
// @Produce json
// @Param page query int false "Page number" minimum(1) default(1)
// @Param page_size query int false "Items per page" minimum(1) maximum(100) default(20)
// @Security BearerAuth
// @Success 200 {object} schemas.PaginatedUsers "List of users"
// @Failure 400 {object} schemas.ErrorResponse "Invalid pagination parameters"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = -1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = -1
	}
	users, err := h.userUseCase.List(page,pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}


// Login handles user login
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param email body string true "User email"
// @Param password body string true "User password"
// @Success 200 {object} schemas.LoginResponse "Login successful"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request format"
// @Failure 401 {object} schemas.ErrorResponse "Invalid email or password"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/auth/login [post]
// @Security BearerAuth
func (h *UserHandler) Login(c *gin.Context){
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}
	fmt.Println("this is the email and the password",email,password)
	user,token, err := h.userUseCase.Login(email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}