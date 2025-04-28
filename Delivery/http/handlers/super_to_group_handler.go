package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type SuperToGroupHandler struct {
	SuperToGroupUseCase usecases.SuperToGroupUsecase
}

func NewSuperToGroupHandler(superToGroupUseCase usecases.SuperToGroupUsecase) *SuperToGroupHandler {
	return &SuperToGroupHandler{
		SuperToGroupUseCase: superToGroupUseCase,
	}

}

// CreateSuperToGroup handles associating groups with a super group
// @Summary Add groups to a super group
// @Description Associate one or more groups with a super group
// @Tags SuperGroups
// @Accept json
// @Produce json
// @Param super_group_id path int true "Super Group ID"
// @Param body body schemas.SuperToGroupRequest true "Group IDs to associate"
// @Success 201 {object} schemas.SuccessResponse "Groups added to super group successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body or super group ID"
// @Failure 404 {object} schemas.ErrorResponse "Super group not found"
// @Router /api/super-groups/{super_group_id}/groups [post]
func (h *SuperToGroupHandler) CreateSuperToGroup(c *gin.Context) {
	id := c.Param("super_group_id")
	if id == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Super group ID is required"})
		return
	}
	superGroupID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid super group ID", Details: err.Error()})
		return
	}
	var req schemas.SuperToGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid request body", Details: err.Error()})
		return
	}
	// Call the usecase to perform the association
	if err := h.SuperToGroupUseCase.CreateSuperToGroup(&entity.SuperToGroup{SuperGroupID: uint(superGroupID)}); err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Failed to add groups to super group", Details: err.Error()})
		return
	}
	c.JSON(201, schemas.SuccessResponse{Success: true, Code: 201, Message: "Groups added to super group successfully"})
}

// ListSuperToGroup handles listing all super to groups
// @Summary List super to groups
// @Description Get a list of all super to groups
// @Tags SuperToGroups
// @Produce json
// @Success 200 {object} schemas.SuccessResponse "List of super to groups"
// @Router /api/super_to_groups [get]
func (h *SuperToGroupHandler) ListSuperToGroup(c *gin.Context) {
	superToGroups, err := h.SuperToGroupUseCase.ListSuperToGroup()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Code: 500, Message: "Internal server error", Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, schemas.SuccessResponse{Success: true, Code: 200, Message: "Super to groups retrieved successfully", Data: superToGroups})
}

// GetSuperToGroupByID handles getting a super to group by ID
// @Summary Get super to group by ID
// @Description Get a super to group by its ID
// @Tags SuperToGroups
// @Produce json
// @Param id path int true "SuperToGroup ID"
// @Success 200 {object} schemas.SuccessResponse "Super to group details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid super to group ID"
// @Failure 404 {object} schemas.ErrorResponse "Super to group not found"
// @Router /api/super_to_groups/{id} [get]
func (h *SuperToGroupHandler) GetSuperToGroupByID(c *gin.Context) {
	superToGroupID := c.Param("superToGroupID")
	stgid, err := strconv.Atoi(superToGroupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid ID"})
		return
	}
	superToGroup, err := h.SuperToGroupUseCase.GetSuperToGroupByID(uint(stgid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Code: 500, Message: "Internal server error", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Super to group retrieved successfully",
		"data":    superToGroup,
	})
}

// UpdateSuperToGroup handles updating a super to group
// @Summary Update a super to group
// @Description Update a super to group by its ID
// @Tags SuperToGroups
// @Accept json
// @Produce json
// @Param id path int true "SuperToGroup ID"
// FIXME: UpdateSuperToGroupRequest does not exist. Use SuperToGroupRequest or define the struct in schemas.
// @Param super_to_group body schemas.SuperToGroupRequest true "SuperToGroup data"
// @Success 200 {object} schemas.SuccessResponse "Super to group updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid input"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/super_to_groups/{id} [patch]
func (h *SuperToGroupHandler) UpdateSuperToGroup(c *gin.Context) {
	superToGroupID := c.Param("superToGroupID")
	stgid, err := strconv.Atoi(superToGroupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid ID"})
		return
	}
	var superToGroup *entity.SuperToGroup
	if err := c.ShouldBindJSON(&superToGroup); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid request body", Details: err.Error()})
		return
	}
	superToGroup.ID = uint(stgid)
	err = h.SuperToGroupUseCase.UpdateSuperToGroup(superToGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Code: 500, Message: "Internal server error", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{Success: true, Code: 200, Message: "Super to group updated successfully", Data: superToGroup})
}

// DeleteSuperToGroup handles deleting a super to group
// @Summary Delete a super to group
// @Description Delete a super to group by its ID
// @Tags SuperToGroups
// @Produce json
// @Param id path int true "SuperToGroup ID"
// @Success 200 {object} schemas.SuccessResponse "Super to group deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid super to group ID"
// @Failure 404 {object} schemas.ErrorResponse "Super to group not found"
// @Router /api/super_to_groups/{id} [delete]
func (h *SuperToGroupHandler) DeleteSuperToGroup(c *gin.Context) {
	superToGroupID := c.Param("superToGroupID")
	stgid, err := strconv.Atoi(superToGroupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid ID"})
		return
	}
	err = h.SuperToGroupUseCase.DeleteSuperToGroup(uint(stgid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Code: 500, Message: "Internal server error", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{Success: true, Code: 200, Message: "Super to group deleted successfully"})
}
