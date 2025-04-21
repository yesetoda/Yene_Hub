package handlers

import (
	"net/http"
	"strconv"

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

// CreateSuperToGroup handles creating a new super to group
// @Summary Create a new super to group
// @Description Create a new super to group entry
// @Tags SuperToGroups
// @Accept json
// @Produce json
// @Param super_to_group body entity.SuperToGroup true "SuperToGroup data"
// @Success 201 {object} entity.SuperToGroup "Super to group created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/super_to_groups [post]
func (h *SuperToGroupHandler) CreateSuperToGroup(c *gin.Context) {
	var superToGroup *entity.SuperToGroup

	if err := c.ShouldBindJSON(&superToGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.SuperToGroupUseCase.CreateSuperToGroup(superToGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super to group created successfully",
		"data":    superToGroup,
	})
}

// ListSuperToGroup handles listing all super to groups
// @Summary List super to groups
// @Description Get a list of all super to groups
// @Tags SuperToGroups
// @Produce json
// @Success 200 {array} entity.SuperToGroup "List of super to groups"
// @Router /api/super_to_groups [get]
func (h *SuperToGroupHandler) ListSuperToGroup(c *gin.Context) {
	superToGroups, err := h.SuperToGroupUseCase.ListSuperToGroup()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super to groups retrieved successfully",
		"data":    superToGroups,
	})
}

// GetSuperToGroupByID handles getting a super to group by ID
// @Summary Get super to group by ID
// @Description Get a super to group by its ID
// @Tags SuperToGroups
// @Produce json
// @Param id path int true "SuperToGroup ID"
// @Success 200 {object} entity.SuperToGroup "Super to group details"
// @Failure 400 {object} map[string]string "Invalid super to group ID"
// @Failure 404 {object} map[string]string "Super to group not found"
// @Router /api/super_to_groups/{id} [get]
func (h *SuperToGroupHandler) GetSuperToGroupByID(c *gin.Context) {
	superToGroupID := c.Param("superToGroupID")
	stgid, err := strconv.Atoi(superToGroupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	superToGroup, err := h.SuperToGroupUseCase.GetSuperToGroupByID(uint(stgid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Param super_to_group body entity.SuperToGroup true "SuperToGroup data"
// @Success 200 {object} entity.SuperToGroup "Super to group updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/super_to_groups/{id} [patch]
func (h *SuperToGroupHandler) UpdateSuperToGroup(c *gin.Context) {
	superToGroupID := c.Param("superToGroupID")
	stgid, err := strconv.Atoi(superToGroupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var superToGroup *entity.SuperToGroup
	if err := c.ShouldBindJSON(&superToGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	superToGroup.ID = uint(stgid)
	err = h.SuperToGroupUseCase.UpdateSuperToGroup(superToGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Super to group updated successfully",
		"data":    superToGroup,
	})
}

// DeleteSuperToGroup handles deleting a super to group
// @Summary Delete a super to group
// @Description Delete a super to group by its ID
// @Tags SuperToGroups
// @Produce json
// @Param id path int true "SuperToGroup ID"
// @Success 200 {object} map[string]string "Super to group deleted successfully"
// @Failure 400 {object} map[string]string "Invalid super to group ID"
// @Failure 404 {object} map[string]string "Super to group not found"
// @Router /api/super_to_groups/{id} [delete]
func (h *SuperToGroupHandler) DeleteSuperToGroup(c *gin.Context) {
	superToGroupID := c.Param("superToGroupID")
	stgid, err := strconv.Atoi(superToGroupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.SuperToGroupUseCase.DeleteSuperToGroup(uint(stgid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Super to group deleted successfully",
	})
}