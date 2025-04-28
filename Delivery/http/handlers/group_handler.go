package handlers

import (
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
// @Success 201 {object} schemas.SuccessResponse "Group created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/groups [post]
func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var input schemas.CreateGroupRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}
	createdGroup, err := h.groupUseCase.Create(&input)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Could not create group",
			Details: err.Error(),
		})
		return
	}
	c.JSON(201, schemas.SuccessResponse{
		Success: true,
		Code:    201,
		Message: "Group created successfully",
		Data:    createdGroup})
}

// GetGroupByID handles getting a group by ID
// @Summary Get group by ID
// @Description Get a group by its ID
// @Tags Groups
// @Produce json
// @Param id path int true "Group ID"
// @Success 200 {object} schemas.SuccessResponse "Group details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid group ID"
// @Failure 404 {object} schemas.ErrorResponse "Group not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/groups/{id} [get]
func (h *GroupHandler) GetGroupByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid group ID",
			Details: err.Error(),
		})
		return
	}
	group, err := h.groupUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{
			Code:    404,
			Message: "Group not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Group details",
		Data:    group})
}

// UpdateGroup handles updating a group
// @Summary Update group
// @Description Update a group by its ID
// @Tags Groups
// @Accept json
// @Produce json
// @Param id path int true "Group ID"
// @Param group body schemas.UpdateGroupRequest true "Group data"
// @Success 200 {object} schemas.SuccessResponse "Group updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body or group ID"
// @Failure 404 {object} schemas.ErrorResponse "Group not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/groups/{id} [patch]
func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid group ID",
			Details: err.Error(),
		})
		return
	}
	var input schemas.UpdateGroupRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}
	updatedGroup, err := h.groupUseCase.Update(uint(id), &input)
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{
			Code:    404,
			Message: "Group not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Group updated successfully",
		Data:    updatedGroup})
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
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/groups/{id} [delete]
func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid group ID",
			Details: err.Error(),
		})
		return
	}
	err = h.groupUseCase.Delete(uint(id))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{
			Code:    404,
			Message: "Group not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Group deleted successfully",
		Data:    nil})
}

// ListGroups handles listing all groups
// @Summary List groups
// @Description Get a list of all groups
// @Tags Groups
// @Produce json
// @Success 200 {object} schemas.SuccessResponse "List of groups"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/groups [get]
func (h *GroupHandler) ListGroups(c *gin.Context) {
	groups, _, err := h.groupUseCase.List()
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to list groups",
			Details: err.Error(),
		})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of groups",
		Data:    groups})
}

// GetGroupsByCountryID handles listing groups by country ID
// @Summary List groups by country ID
// @Description Get a list of groups by country ID
// @Tags Groups
// @Produce json
// @Param country_id path int true "Country ID"
// @Success 200 {object} schemas.SuccessResponse "List of groups"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/groups/country/{country_id} [get]
func (h *GroupHandler) GetGroupsByCountryID(c *gin.Context) {
	idStr := c.Param("country_id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{
			Code:    400,
			Message: "Invalid country ID",
			Details: err.Error(),
		})
		return
	}
	groups, _, err := h.groupUseCase.GetByCountryID(uint(id))
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{
			Code:    500,
			Message: "Failed to list groups",
			Details: err.Error(),
		})
		return
	}
	c.JSON(200, schemas.SuccessResponse{	
		Success: true,
		Code:    200,
		Message: "List of groups",
		Data:    groups})
}
