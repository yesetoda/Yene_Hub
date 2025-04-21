package handlers

import (
	"strconv"

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
// @Param recentAction body entity.RecentAction true "Recent action details"
// @Success 201 {object} map[string]string "Recent action created"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Router /api/recent-actions [post]
func (h *RecentActionHandler) CreateRecentAction(c *gin.Context) {
	// Get the request body
	var recentAction entity.RecentAction
	if err := c.ShouldBindJSON(&recentAction); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Call the use case to create the recent action
	if err := h.RecentActionUsecase.CreateRecentAction(&recentAction); err != nil {
		c.JSON(400, gin.H{"error": "Failed to create recent action"})
		return
	}
	// Return a success response
	c.JSON(201, gin.H{"message": "Recent action created successfully"})
}

// ListRecentActions lists all recent actions
// @Summary List recent actions
// @Description Get a list of all recent user actions
// @Tags RecentActions
// @Produce json
// @Success 200 {array} entity.RecentAction "List of recent actions"
// @Router /api/recent-actions [get]
func (h *RecentActionHandler) ListRecentActions(c *gin.Context) {
	// Call the use case to get the list of recent actions
	recentActions, err := h.RecentActionUsecase.ListRecentAction()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch recent actions"})
		return
	}
	// Return the list of recent actions
	c.JSON(200, recentActions)
}

// GetRecentActionByUserID gets recent actions by user ID
// @Summary Get recent actions by user ID
// @Description Get a list of recent actions for a specific user
// @Tags RecentActions
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} entity.RecentAction "List of recent actions for the user"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Router /api/recent-actions/user/{user_id} [get]
func (h *RecentActionHandler) GetRecentActionByUserID(c *gin.Context) {
	// Get the user ID from the URL parameter
	uid := c.Param("user_id")
	if uid == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}
	// Convert user ID to uint
	userID, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	// Call the use case to get recent actions by user ID
	recentActions, err := h.RecentActionUsecase.GetRecentActionByUserID(uint(userID))
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch recent actions"})
		return
	}
	// Return the list of recent actions for the user
	c.JSON(200, recentActions)
}

// GetRecentActionByID gets a recent action by ID
// @Summary Get recent action by ID
// @Description Get details of a recent action by its ID
// @Tags RecentActions
// @Produce json
// @Param id path int true "Recent Action ID"
// @Success 200 {object} entity.RecentAction "Recent action details"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "Not found"
// @Router /api/recent-actions/{id} [get]
func (h *RecentActionHandler) GetRecentActionByID(c *gin.Context) {
	// Get the recent action ID from the URL parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Recent action ID is required"})
		return
	}
	// Convert ID to uint
	recentActionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid recent action ID"})
		return
	}
	// Call the use case to get recent actions by ID
	recentActions, err := h.RecentActionUsecase.GetRecentActionByID(uint(recentActionID))
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch recent actions"})
		return
	}
	// Return the list of recent actions for the ID
	c.JSON(200, recentActions)
}

// GetRecentActionByType gets recent actions by type
// @Summary Get recent actions by type
// @Description Get a list of recent actions for a specific type
// @Tags RecentActions
// @Produce json
// @Param action_type path string true "Action Type"
// @Success 200 {array} entity.RecentAction "List of recent actions for the type"
// @Router /api/recent-actions/type/{action_type} [get]
func (h *RecentActionHandler) GetRecentActionByType(c *gin.Context) {
	// Get the action type from the URL parameter
	actionType := c.Param("action_type")
	if actionType == "" {
		c.JSON(400, gin.H{"error": "Action type is required"})
		return
	}
	// Call the use case to get recent actions by action type
	recentActions, err := h.RecentActionUsecase.GetRecentActionByType(actionType)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch recent actions"})
		return
	}
	// Return the list of recent actions for the action type
	c.JSON(200, recentActions)
}

// UpdateRecentAction updates a recent action
// @Summary Update recent action
// @Description Update a recent user action
// @Tags RecentActions
// @Accept json
// @Produce json
// @Param id path int true "Recent Action ID"
// @Param recentAction body entity.RecentAction true "Recent action details"
// @Success 200 {object} map[string]string "Recent action updated"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Router /api/recent-actions/{id} [put]
func (h *RecentActionHandler) UpdateRecentAction(c *gin.Context) {
	// Get the recent action ID from the URL parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Recent action ID is required"})
		return
	}
	// Convert ID to uint
	recentActionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid recent action ID"})
		return
	}
	// Get the request body
	var recentAction entity.RecentAction
	if err := c.ShouldBindJSON(&recentAction); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	recentAction.ID = uint(recentActionID)
	// Call the use case to update the recent action
	if err := h.RecentActionUsecase.UpdateRecentAction(&recentAction); err != nil {
		c.JSON(400, gin.H{"error": "Failed to update recent action"})
		return
	}
	// Return a success response
	c.JSON(200, gin.H{"message": "Recent action updated successfully"})
}

// DeleteRecentAction deletes a recent action
// @Summary Delete recent action
// @Description Delete a recent user action
// @Tags RecentActions
// @Produce json
// @Param id path int true "Recent Action ID"
// @Success 200 {object} map[string]string "Recent action deleted"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Router /api/recent-actions/{id} [delete]
func (h *RecentActionHandler) DeleteRecentAction(c *gin.Context) {
	// Get the recent action ID from the URL parameter
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Recent action ID is required"})
		return
	}
	// Convert ID to uint
	recentActionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid recent action ID"})
		return
	}
	// Call the use case to delete the recent action
	if err := h.RecentActionUsecase.DeleteRecentAction(uint(recentActionID)); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete recent action"})
		return
	}
	// Return a success response
	c.JSON(200, gin.H{"message": "Recent action deleted successfully"})	
}