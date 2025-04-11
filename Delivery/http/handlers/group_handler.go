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
func (h *GroupHandler) ListGroups(c *gin.Context) {
	groups, err := h.groupUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetGroupsByCountryID handles listing groups by country ID
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
