package handlers

import (
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

// RoleHandler handles HTTP requests for role operations
type RoleHandler struct {
	roleUseCase usecases.RoleUseCase
}

// NewRoleHandler creates a new RoleHandler instance
func NewRoleHandler(roleUseCase usecases.RoleUseCase) *RoleHandler {
	return &RoleHandler{
		roleUseCase: roleUseCase,
	}
}

// CreateRole handles creating a new role
// @Summary Create role
// @Description Create a new role with the provided information
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param request body schemas.CreateRoleRequest true "Role creation data"
// @Success 201 {object} schemas.SuccessResponse "Role created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request format"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var input schemas.CreateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid request format",
			Details: err.Error(),
		})
		return
	}

	role := &schemas.CreateRoleRequest{
		Type: input.Type,
	}

	createdRole, err := h.roleUseCase.Create(role)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to create role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(201, schemas.SuccessResponse{
		Message: "Role created successfully",
		Data:    createdRole,
	})
}

// GetRoleByID handles retrieving a role by ID
// @Summary Get role
// @Description Get detailed information about a specific role
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Role ID" minimum(1)
// @Success 200 {object} schemas.SuccessResponse "Role details retrieved successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid role ID format"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 404 {object} schemas.ErrorResponse "Role not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/roles/{id} [get]
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid role ID",
			Details: "Role ID must be a positive integer",
		})
		return
	}

	role, err := h.roleUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{
			Code:    404,
			Message: "Role not found",
			Details: err.Error(),
		})
		return
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "Role details retrieved successfully",
		Data:    role,
	})
}

// UpdateRole handles updating a role's information
// @Summary Update role
// @Description Update an existing role's information
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Role ID" minimum(1)
// @Param request body schemas.UpdateRoleRequest true "Role update data"
// @Success 200 {object} schemas.SuccessResponse "Role updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request format or role ID"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 404 {object} schemas.ErrorResponse "Role not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/roles/{id} [patch]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid role ID",
			Details: "Role ID must be a positive integer",
		})
		return
	}

	var input schemas.UpdateRoleRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid request format",
			Details: err.Error(),
		})
		return
	}

	role := &schemas.UpdateRoleRequest{Type: input.Type}
	if input.Type != nil {
		role.Type = input.Type
	}

	updatedRole, err := h.roleUseCase.Update(uint(id), role)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(404, schemas.ErrorResponse{
				Code:    404,
				Message: "Role not found",
				Details: err.Error(),
			})
			return
		}
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to update role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "Role updated successfully",
		Data:    updatedRole,
	})
}

// DeleteRole handles deleting a role
// @Summary Delete role
// @Description Delete an existing role
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Role ID" minimum(1)
// @Success 200 {object} schemas.SuccessResponse "Role deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid role ID format"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 404 {object} schemas.ErrorResponse "Role not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/roles/{id} [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid role ID",
			Details: "Role ID must be a positive integer",
		})
		return
	}

	if err := h.roleUseCase.Delete(uint(id)); err != nil {
		if err.Error() == "record not found" {
			c.JSON(404, schemas.ErrorResponse{
				Code:    404,
				Message: "Role not found",
				Details: err.Error(),
			})
			return
		}
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to delete role",
			Details: err.Error(),
		})
		return
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "Role deleted successfully",
	})
}

// ListRoles handles listing roles with pagination
// @Summary List roles
// @Description Get a paginated list of roles
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param page query int false "Page number" minimum(0) default(0)
// @Param page_size query int false "Number of items per page" minimum(1) maximum(100) default(10)
// @Success 200 {object} schemas.SuccessResponse "List of roles retrieved successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid query parameters"
// @Failure 401 {object} schemas.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} schemas.ErrorResponse "Forbidden - Insufficient permissions"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/roles [get]
func (h *RoleHandler) ListRoles(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid page number",
			Details: "Page number must be a positive integer",
		})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 || pageSize > 100 {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid page size",
			Details: "Page size must be between 1 and 100",
		})
		return
	}

	roles, _, err := h.roleUseCase.List()
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to list roles",
			Details: err.Error(),
		})
		return
	}

	// Convert to response format
	roleResponses := make([]*schemas.RoleResponse, len(roles))
	for i, role := range roles {
		roleResponses[i] = &schemas.RoleResponse{
			ID:          role.ID,
			Type:        role.Type,
			Description: "",
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		}
	}

	c.JSON(200, schemas.SuccessResponse{
		Message: "List of roles retrieved successfully",
	})
}
