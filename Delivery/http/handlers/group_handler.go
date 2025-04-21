package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

// GroupHandler handles HTTP requests for group operations
type GroupHandler struct {
	groupUseCase usecases.GroupUseCase
}

// NewGroupHandler creates a new GroupHandler instance
func NewGroupHandler(groupUseCase usecases.GroupUseCase) *GroupHandler {
	return &GroupHandler{
		groupUseCase: groupUseCase,
	}
}

// CreateGroup handles creating a new group
// @Summary Create a new group
// @Description Create a new group entry
// @Tags Groups
// @Accept json
// @Produce json
// @Param group body entity.Group true "Group data"
// @Success 201 {object} entity.Group "Group created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/groups [post]
func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var group entity.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdGroup, err := h.groupUseCase.Create(&group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Group created successfully",
		"group":   createdGroup,
	})
}

// GetGroupByID handles getting a group by ID
// @Summary Get group by ID
// @Description Get a group by its ID
// @Tags Groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} entity.Group "Group details"
// @Failure 400 {object} map[string]string "Invalid group ID"
// @Failure 404 {object} map[string]string "Group not found"
// @Router /api/groups/{id} [get]
func (h *GroupHandler) GetGroupByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	group, err := h.groupUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}

// UpdateGroup handles updating a group
// @Summary Update a group
// @Description Update a group by its ID
// @Tags Groups
// @Accept json
// @Produce json
// @Param id path int true "Group ID"
// @Param group body entity.Group true "Group data"
// @Success 200 {object} entity.Group "Group updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/groups/{id} [patch]
func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	var group entity.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	group.ID = uint(id)
	updatedGroup, err := h.groupUseCase.Update(&group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Group updated successfully",
		"group":   updatedGroup,
	})
}

// DeleteGroup handles deleting a group
// @Summary Delete a group
// @Description Delete a group by its ID
// @Tags Groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} map[string]string "Group deleted successfully"
// @Failure 400 {object} map[string]string "Invalid group ID"
// @Failure 404 {object} map[string]string "Group not found"
// @Router /api/groups/{id} [delete]
func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	if err := h.groupUseCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Group deleted successfully",
	})
}

// ListGroups handles listing all groups
// @Summary List groups
// @Description Get a list of all groups
// @Tags Groups
// @Produce json
// @Success 200 {array} entity.Group "List of groups"
// @Router /api/groups [get]
func (h *GroupHandler) ListGroups(c *gin.Context) {
	groups, err := h.groupUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetGroupsByCountryID handles listing groups by country ID
// @Summary List groups by country ID
// @Description Get a list of groups by country ID
// @Tags Groups
// @Produce json
// @Param country_id path int true "Country ID"
// @Success 200 {array} entity.Group "List of groups"
// @Router /api/groups/country/{country_id} [get]
func (h *GroupHandler) GetGroupsByCountryID(c *gin.Context) {
	idStr := c.Param("country_id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	groups, err := h.groupUseCase.GetGroupsByCountryID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}
