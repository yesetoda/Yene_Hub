package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Domain/entity"
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
// @Summary Create a new role
// @Description Create a new role entry
// @Tags Roles
// @Accept json
// @Produce json
// @Param role body entity.Role true "Role data"
// @Success 201 {object} entity.Role "Role created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role entity.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdRole, err := h.roleUseCase.Create(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Role created successfully",
		"role":    createdRole,
	})
}

// GetRoleByID handles getting a role by ID
// @Summary Get role by ID
// @Description Get a role by its ID
// @Tags Roles
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} entity.Role "Role details"
// @Failure 400 {object} map[string]string "Invalid role ID"
// @Failure 404 {object} map[string]string "Role not found"
// @Router /api/roles/{id} [get]
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.roleUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// UpdateRole handles updating a role
// @Summary Update a role
// @Description Update a role by its ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Param role body entity.Role true "Role data"
// @Success 200 {object} entity.Role "Role updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/roles/{id} [patch]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var role entity.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	role.ID = uint(id)
	updatedRole, err := h.roleUseCase.Update(&role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role updated successfully",
		"role":    updatedRole,
	})
}

// DeleteRole handles deleting a role
// @Summary Delete a role
// @Description Delete a role by its ID
// @Tags Roles
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} map[string]string "Role deleted successfully"
// @Failure 400 {object} map[string]string "Invalid role ID"
// @Failure 404 {object} map[string]string "Role not found"
// @Router /api/roles/{id} [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := h.roleUseCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role deleted successfully",
	})
}

// ListRoles handles listing all roles
// @Summary List roles
// @Description Get a list of all roles
// @Tags Roles
// @Produce json
// @Success 200 {array} entity.Role "List of roles"
// @Router /api/roles [get]
func (h *RoleHandler) ListRoles(c *gin.Context) {
	roles, err := h.roleUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}
