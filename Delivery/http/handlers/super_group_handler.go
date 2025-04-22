package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type SuperGroupHandler struct {
	superGroupUseCase usecases.SuperGroupUseCaseInterface
}

func NewSuperGroupHandler(superGroupUseCase usecases.SuperGroupUseCaseInterface) *SuperGroupHandler {
	return &SuperGroupHandler{
		superGroupUseCase: superGroupUseCase,
	}
}

// CreateSuperGroup handles creating a new super group
// @Summary Create a new super group
// @Description Create a new super group entry
// @Tags SuperGroups
// @Accept json
// @Produce json
// @Param super_group body schemas.CreateSuperGroupRequest true "SuperGroup data"
// @Success 201 {object} schemas.SuperGroupResponse "Super group created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/super_groups [post]
func (h *SuperGroupHandler) CreateSuperGroup(c *gin.Context) {
	var superGroup *entity.SuperGroup

	if err := c.ShouldBindJSON(&superGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	superGroup, err := h.superGroupUseCase.Create(superGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super group created successfully",
		"data":    superGroup,
	})
}

// GetSuperGroup handles getting a super group by ID
// @Summary Get super group by ID
// @Description Get a super group by its ID
// @Tags SuperGroups
// @Produce json
// @Param id path int true "SuperGroup ID"
// @Success 200 {object} schemas.SuperGroupResponse "Super group details"
// @Failure 400 {object} map[string]string "Invalid super group ID"
// @Failure 404 {object} map[string]string "Super group not found"
// @Router /api/super_groups/{id} [get]
func (h *SuperGroupHandler) GetSuperGroup(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	superGroup, err := h.superGroupUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super group retrieved successfully",
		"data":    superGroup,
	})
}

// UpdateSuperGroup handles updating a super group
// @Summary Update a super group
// @Description Update a super group by its ID
// @Tags SuperGroups
// @Accept json
// @Produce json
// @Param id path int true "SuperGroup ID"
// @Param super_group body schemas.UpdateSuperGroupRequest true "SuperGroup data"
// @Success 200 {object} schemas.SuperGroupResponse "Super group updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/super_groups/{id} [patch]
func (h *SuperGroupHandler) UpdateSuperGroup(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var superGroup *entity.SuperGroup

	if err := c.ShouldBindJSON(&superGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.superGroupUseCase.Update(superGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super group updated successfully",
		"data":    superGroup,
	})
}

// DeleteSuperGroup handles deleting a super group
// @Summary Delete a super group
// @Description Delete a super group by its ID
// @Tags SuperGroups
// @Produce json
// @Param id path int true "SuperGroup ID"
// @Success 200 {object} schemas.SuperGroupResponse "Super group deleted successfully"
// @Failure 400 {object} map[string]string "Invalid super group ID"
// @Failure 404 {object} map[string]string "Super group not found"
// @Router /api/super_groups/{id} [delete]
func (h *SuperGroupHandler) DeleteSuperGroup(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.superGroupUseCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super group deleted successfully",
	})
}

// ListSuperGroups handles listing all super groups
// @Summary List super groups
// @Description Get a list of all super groups
// @Tags SuperGroups
// @Produce json
// @Success 200 {array} []*schemas.SuperGroupResponse "List of super groups"
// @Router /api/super_groups [get]
func (h *SuperGroupHandler) ListSuperGroups(c *gin.Context) {
	superGroups, err := h.superGroupUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super groups retrieved successfully",
		"data":    superGroups,
	})
}
