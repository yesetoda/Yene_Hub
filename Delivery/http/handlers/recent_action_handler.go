package handlers

import (
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type RecentActionHandler struct {
	RecentActionUsecase usecases.RecentActionUsecase
}

func NewRecentActionHandler(recentActionUsecase usecases.RecentActionUsecase) *RecentActionHandler {
	return &RecentActionHandler{
		RecentActionUsecase: recentActionUsecase,
	}
}

// CreateRecentAction creates a new recent action
// @Summary Create recent action
// @Description Create a new recent user action
// @Tags RecentActions
// @Accept json
// @Produce json
// @Param recentAction body schemas.CreateRecentActionRequest true "Recent action details"
// @Success 201 {object} schemas.SuccessResponse "Recent action created"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/recent_actions [post]
func (h *RecentActionHandler) CreateRecentAction(c *gin.Context) {
	// Get the request body
	var recentAction schemas.CreateRecentActionRequest
	if err := c.ShouldBindJSON(&recentAction); err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid request body"})
		return
	}
	// Call the use case to create the recent action	
	rec_ac := entity.RecentAction{
		UserID: recentAction.UserID,
		Type: recentAction.ActionType,
		Description: recentAction.Description,
	}
	if err := h.RecentActionUsecase.CreateRecentAction(&rec_ac); err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		return
	}
	// Return a success response
	resp := schemas.SuccessResponse{
		Message: "Recent action created successfully",
		Data:    recentAction,
	}
	c.JSON(201, resp)
}

// ListRecentActions lists all recent actions
// @Summary List recent actions
// @Description Get a list of all recent user actions
// @Tags RecentActions
// @Produce json
// @Success 200 {object} schemas.SuccessResponse "List of recent actions"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/recent_actions [get]
func (h *RecentActionHandler) ListRecentActions(c *gin.Context) {
	// Call the use case to get the list of recent actions
	recentActions, err := h.RecentActionUsecase.ListRecentAction()
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		return
	}
	resp := schemas.SuccessResponse{
		Message: "Recent actions fetched successfully",
		Data:    recentActions,
	}
	// Return the list of recent actions
	c.JSON(200, resp)
}

// GetRecentActionByUserID gets recent actions by user ID
// @Summary Get recent actions by user ID
// @Description Get a list of recent actions for a specific user
// @Tags RecentActions
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} schemas.SuccessResponse "List of recent actions for the user"
// @Failure 400 {object} schemas.ErrorResponse "Invalid user ID"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/recent_actions/user/{user_id} [get]
func (h *RecentActionHandler) GetRecentActionByUserID(c *gin.Context) {
	// Get the user ID from the URL parameter
	uid := c.Param("user_id")
	if uid == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "User ID is required"})
		return
	}
	// Convert user ID to uint
	userID, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid user ID"})
		return
	}
	// Call the use case to get recent actions by user ID
	recentActions, err := h.RecentActionUsecase.GetRecentActionByUserID(uint(userID))
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		return
	}
	resp := schemas.SuccessResponse{
		Message: "Fetched recent actions for user:" + uid,
		Data:    recentActions,
	}
	// Return the list of recent actions for the user
	c.JSON(200, resp)
}

// GetRecentActionByID gets a recent action by ID
// @Summary Get recent action by ID
// @Description Get details of a recent action by its ID
// @Tags RecentActions
// @Produce json
// @Param id path int true "Recent Action ID"
// @Success 200 {object} schemas.SuccessResponse "Recent action details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid ID"
// @Failure 404 {object} schemas.ErrorResponse "Not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/recent_actions/{id} [get]
func (h *RecentActionHandler) GetRecentActionByID(c *gin.Context) {
	// Get the recent action ID from the URL parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Recent action ID is required"})
		return
	}
	// Convert ID to uint
	recentActionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid recent action ID"})
		return
	}
	// Call the use case to get recent actions by ID
	recentActions, err := h.RecentActionUsecase.GetRecentActionByID(uint(recentActionID))
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Not found"})
		} else {
			c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		}
		return
	}
	resp := schemas.SuccessResponse{
		Message: "Fetched recent actions by ID:" + id,
		Data:    recentActions,
	}
	// Return the list of recent actions
	c.JSON(200, resp)
}

// GetRecentActionByType gets recent actions by type
// @Summary Get recent actions by type
// @Description Get a list of recent actions for a specific type
// @Tags RecentActions
// @Produce json
// @Param action_type path string true "Action Type"
// @Success 200 {object} schemas.SuccessResponse "List of recent actions for the type"
// @Failure 400 {object} schemas.ErrorResponse "Invalid action type"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/recent_actions/type/{action_type} [get]
func (h *RecentActionHandler) GetRecentActionByType(c *gin.Context) {
	// Get the action type from the URL parameter
	actionType := c.Param("action_type")
	if actionType == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Action type is required"})
		return
	}
	// Call the use case to get recent actions by action type
	recentActions, err := h.RecentActionUsecase.GetRecentActionByType(actionType)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		return
	}
	resp := schemas.SuccessResponse{
		Message: "Fetched recent actions by type:" + actionType,
		Data:    recentActions,
	}
	// Return the list of recent actions for the action type
	c.JSON(200, resp)
}

// UpdateRecentAction updates a recent action
// @Summary Update recent action
// @Description Update a recent user action
// @Tags RecentActions
// @Accept json
// @Produce json
// @Param id path int true "Recent Action ID"
// @Param recentAction body schemas.UpdateRecentActionRequest true "Recent action details"
// @Success 200 {object} schemas.SuccessResponse "Recent action updated"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 404 {object} schemas.ErrorResponse "Recent action not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/recent_actions/{id} [patch]
func (h *RecentActionHandler) UpdateRecentAction(c *gin.Context) {
	// Get the recent action ID from the URL parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Recent action ID is required"})
		return
	}
	// Convert ID to uint
	recentActionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid recent action ID"})
		return
	}
	// Get the request body
	var recentAction schemas.UpdateRecentActionRequest
	if err := c.ShouldBindJSON(&recentAction); err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid request body"})
		return
	}
	recentAction.ID = uint(recentActionID)
	// Call the use case to update the recent action
	if err := h.RecentActionUsecase.UpdateRecentAction(&entity.RecentAction{ID: recentAction.ID}); err != nil {
		if err.Error() == "not found" {
			c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Recent action not found"})
		} else {
			c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		}
		return
	}
	resp := schemas.SuccessResponse{
		Message: "Recent action updated successfully",
		Data:    recentAction,
	}
	// Return a success response
	c.JSON(200, resp)
}

// DeleteRecentAction deletes a recent action
// @Summary Delete recent action
// @Description Delete a recent user action
// @Tags RecentActions
// @Produce json
// @Param id path int true "Recent Action ID"
// @Success 200 {object} schemas.SuccessResponse "Recent action deleted"
// @Failure 400 {object} schemas.ErrorResponse "Invalid ID"
// @Failure 404 {object} schemas.ErrorResponse "Recent action not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/recent_actions/{id} [delete]
func (h *RecentActionHandler) DeleteRecentAction(c *gin.Context) {
	// Get the recent action ID from the URL parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Recent action ID is required"})
		return
	}
	// Convert ID to uint
	recentActionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid recent action ID"})
		return
	}
	// Call the use case to delete the recent action
	if err := h.RecentActionUsecase.DeleteRecentAction(uint(recentActionID)); err != nil {
		if err.Error() == "not found" {
			c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Recent action not found"})
		} else {
			c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Internal server error"})
		}
		return
	}
	// Return a success response
	resp := schemas.SuccessResponse{
		Message: "Recent action deleted successfully",
		Data:    nil,
	}
	c.JSON(200, resp)
}
