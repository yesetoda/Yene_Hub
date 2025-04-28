package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

// RegistrationHandler handles HTTP requests for user registration operations
type RegistrationHandler struct {
	bulkRegistrationUseCase usecases.BulkRegistrationUseCase
}

// NewRegistrationHandler creates a new RegistrationHandler instance
func NewRegistrationHandler(bulkRegistrationUseCase usecases.BulkRegistrationUseCase) *RegistrationHandler {
	return &RegistrationHandler{
		bulkRegistrationUseCase: bulkRegistrationUseCase,
	}
}

// BulkRegistrationRequest defines the structure for bulk registration requests
type BulkRegistrationRequest struct {
	Emails    string `json:"emails" binding:"required"`
	RoleID    uint   `json:"role_id" binding:"required"`
	GroupID   uint   `json:"group_id" binding:"required"`
	CountryID *uint  `json:"country_id"`
}

// RegistrationParam defines the request body for user registration with a role
// swagger:model
// @Description Registration data for bulk user registration with a role
// @Param emails body string true "Emails (comma-separated)"
// @Param group_id body uint true "Group ID"
// @Param country_id body uint false "Country ID"
type RegistrationParam struct {
	Emails    string `json:"emails" binding:"required"`
	GroupID   uint   `json:"group_id" binding:"required"`
	CountryID *uint  `json:"country_id"`
}

// RegisterBulkUsers handles registering multiple users at once
// @Summary Register multiple users in bulk
// @Description Register multiple users with the provided information
// @Tags Registration
// @Accept json
// @Produce json
// @Param bulk-registration body BulkRegistrationRequest true "Bulk registration data"
// @Success 200 {object} schemas.SuccessResponse "Bulk registration processed"
// @Failure 400 {object} schemas.ErrorResponse "Invalid input"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/registration/bulk [post]
func (h *RegistrationHandler) RegisterBulkUsers(c *gin.Context) {
	var request BulkRegistrationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid request body"})
		return
	}

	// Validate role ID (must be positive)
	if request.RoleID <= 0 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid role ID"})
		return
	}

	// Validate group ID (must be positive)
	if request.GroupID <= 0 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid group ID"})
		return
	}

	// Validate country ID if provided (must be positive)
	if request.CountryID != nil && *request.CountryID <= 0 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid country ID"})
		return
	}

	// Register users
	results, err := h.bulkRegistrationUseCase.RegisterUsers(request.Emails, request.RoleID, request.GroupID, request.CountryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		return
	}

	// Count successful registrations
	successCount := 0
	for _, result := range results {
		if result.Success {
			successCount++
		}
	}

	c.JSON(http.StatusOK, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Bulk registration processed",
		Data: map[string]interface{}{
			"total":      len(results),
			"successful": successCount,
			"failed":     len(results) - successCount,
			"results":    results,
		},
	})
}

// RegisterUsersWithRole handles registering multiple users with a specific role ID from the URL
// @Summary Register multiple users with a specific role
// @Description Register multiple users with the provided information and a specific role ID
// @Tags Registration
// @Accept json
// @Produce json
// @Param role_id path uint true "Role ID"
// @Param registration body RegistrationParam true "Registration data"
// @Success 200 {object} schemas.SuccessResponse "Bulk registration processed"
// @Failure 400 {object} schemas.ErrorResponse "Invalid input"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/registration/role/{role_id} [post]
func (h *RegistrationHandler) RegisterUsersWithRole(c *gin.Context) {
	// Get role ID from URL parameter
	roleIDStr := c.Param("role_id")
	roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid role ID format"})
		return
	}

	// Get emails and required IDs from request body
	var request RegistrationParam
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid request body"})
		return
	}

	// Validate group ID (must be positive)
	if request.GroupID <= 0 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid group ID"})
		return
	}

	// Validate country ID if provided (must be positive)
	if request.CountryID != nil && *request.CountryID <= 0 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid country ID"})
		return
	}

	// Register users
	results, err := h.bulkRegistrationUseCase.RegisterUsers(request.Emails, uint(roleID), request.GroupID, request.CountryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		return
	}

	// Count successful registrations
	successCount := 0
	for _, result := range results {
		if result.Success {
			successCount++
		}
	}

	c.JSON(http.StatusOK, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Bulk registration processed",
		Data: map[string]interface{}{
			"total":      len(results),
			"successful": successCount,
			"failed":     len(results) - successCount,
			"results":    results,
		},
	})
}

// ForceSwaggoParse is a dummy function to ensure Swaggo parses this file.
func ForceSwaggoParseRegistrationHandler() {}
