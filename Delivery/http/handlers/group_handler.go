package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

// GroupHandler handles HTTP requests for group operations
type GroupHandler struct {
	groupUseCase usecases.GroupUseCase
}

func NewGroupHandler(groupUseCase usecases.GroupUseCase) *GroupHandler {
	return &GroupHandler{groupUseCase: groupUseCase}
}

// CreateGroup handles creating a new group
// @Summary Create a new group
// @Description Create a new group entry
// @Tags Groups
// @Accept json
// @Produce json
// @Param group body schemas.CreateGroupRequest true "Group data"
// @Success 201 {object} schemas.GroupResponse "Group created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/groups [post]
func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var input schemas.CreateGroupRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}
	createdGroup, err := h.groupUseCase.Create(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Could not create group",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, createdGroup)
}

// GetGroupByID handles getting a group by ID
// @Summary Get group by ID
// @Description Get a group by its ID
// @Tags Groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} schemas.GroupResponse "Group details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid group ID"
// @Failure 404 {object} schemas.ErrorResponse "Group not found"
// @Router /api/groups/{id} [get]
func (h *GroupHandler) GetGroupByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid group ID",
			Details: err.Error(),
		})
		return
	}
	group, err := h.groupUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Group not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, group)
}

// UpdateGroup handles updating a group
// @Summary Update group
// @Description Update a group by its ID
// @Tags Groups
// @Accept json
// @Produce json
// @Param id path int true "Group ID"
// @Param group body schemas.UpdateGroupRequest true "Group data"
// @Success 200 {object} schemas.GroupResponse "Group updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body or group ID"
// @Failure 404 {object} schemas.ErrorResponse "Group not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/groups/{id} [patch]
func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid group ID",
			Details: err.Error(),
		})
		return
	}
	var input schemas.UpdateGroupRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}
	updatedGroup, err := h.groupUseCase.Update(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Group not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updatedGroup)
}

// DeleteGroup handles deleting a group
// @Summary Delete a group
// @Description Delete a group by its ID
// @Tags Groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} schemas.SuccessResponse "Group deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid group ID"
// @Failure 404 {object} schemas.ErrorResponse "Group not found"
// @Router /api/groups/{id} [delete]
func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid group ID",
			Details: err.Error(),
		})
		return
	}
	err = h.groupUseCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Group not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{Message: "Group deleted successfully"})
}

// ListGroups handles listing all groups
// @Summary List groups
// @Description Get a list of all groups
// @Tags Groups
// @Produce json
// @Success 200 {object} schemas.GroupListResponse "List of groups"
// @Router /api/groups [get]
func (h *GroupHandler) ListGroups(c *gin.Context) {
	groups, meta, err := h.groupUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Failed to list groups",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schemas.GroupListResponse{Data: groups, Meta: *meta})
}

// GetGroupsByCountryID handles listing groups by country ID
// @Summary List groups by country ID
// @Description Get a list of groups by country ID
// @Tags Groups
// @Produce json
// @Param country_id path int true "Country ID"
// @Success 200 {object} schemas.GroupListResponse "List of groups"
// @Router /api/groups/country/{country_id} [get]
func (h *GroupHandler) GetGroupsByCountryID(c *gin.Context) {
	idStr := c.Param("country_id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid country ID",
			Details: err.Error(),
		})
		return
	}
	groups, meta, err := h.groupUseCase.GetByCountryID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Failed to list groups",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schemas.GroupListResponse{Data: groups, Meta: *meta})
}
