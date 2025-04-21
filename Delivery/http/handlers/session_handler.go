package handlers

import (
	"strconv"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	SessionUsecase usecases.SessionUsecase
}

func NewSessionHandler(sessionUsecase usecases.SessionUsecase) *SessionHandler {
	return &SessionHandler{
		SessionUsecase: sessionUsecase,
	}
}

// CreateSession handles creating a new session
// @Summary Create a new session
// @Description Create a new session for a user
// @Tags Sessions
// @Accept json
// @Produce json
// @Param session body entity.Session true "Session data"
// @Success 201 {object} map[string]interface{} "Session created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/sessions [post]
func (h *SessionHandler) CreateSession(c *gin.Context) {
	// Create a new session
	var session entity.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.SessionUsecase.CreateSession(&session); err != nil {
		c.JSON(400, gin.H{"error": "Failed to create session"})
		return
	}
	// Return a success response
	c.JSON(201, gin.H{"message": "Session created successfully"})
}

// ListSessions handles listing all sessions
// @Summary List sessions
// @Description Get a list of all sessions
// @Tags Sessions
// @Produce json
// @Success 200 {array} entity.Session "List of sessions"
// @Router /api/sessions [get]
func (h *SessionHandler) ListSessions(c *gin.Context) {
	// List all sessions
	sessions, err := h.SessionUsecase.ListSession()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch sessions"})
		return
	}
	// Return the list of sessions
	c.JSON(200, sessions)
}

// GetSessionByID handles getting a session by ID
// @Summary Get session by ID
// @Description Get a session by its ID
// @Tags Sessions
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {object} entity.Session "Session details"
// @Failure 400 {object} map[string]interface{} "Invalid session ID"
// @Failure 404 {object} map[string]interface{} "Session not found"
// @Router /api/sessions/{id} [get]
func (h *SessionHandler) GetSessionByID(c *gin.Context) {
	// Get a session by ID
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Session ID is required"})
		return
	}
	sid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid session ID"})
		return
	}
	session, err := h.SessionUsecase.GetSessionByID(uint(sid))
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch session"})
		return
	}
	// Return the session
	c.JSON(200, session)
}

// GetSessionByName handles getting a session by name
// @Summary Get session by name
// @Description Get a session by its name
// @Tags Sessions
// @Produce json
// @Param name path string true "Session name"
// @Success 200 {array} entity.Session "Session details"
// @Router /api/sessions/name/{name} [get]
func (h *SessionHandler) GetSessionByName(c *gin.Context) {
	// Get a session by name
	name := c.Param("name")
	if name == "" {
		c.JSON(400, gin.H{"error": "Session name is required"})
		return
	}
	sessions, err := h.SessionUsecase.GetSessionByName(name)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch session"})
		return
	}
	// Return the session
	c.JSON(200, sessions)
}

// GetSessionByStartTime handles getting a session by start time
// @Summary Get session by start time
// @Description Get a session by its start time
// @Tags Sessions
// @Produce json
// @Param start_time path string true "Session start time"
// @Success 200 {array} entity.Session "Session details"
// @Router /api/sessions/start-time/{start_time} [get]
func (h *SessionHandler) GetSessionByStartTime(c *gin.Context) {
	// Get a session by start time
	startTime := c.Param("start_time")
	if startTime == "" {
		c.JSON(400, gin.H{"error": "Session start time is required"})
		return
	}
	sessions, err := h.SessionUsecase.GetSessionByStartTime(startTime)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch session"})
		return
	}
	// Return the session
	c.JSON(200, sessions)
}

// UpdateSession handles updating a session
// @Summary Update a session
// @Description Update a session by its ID
// @Tags Sessions
// @Accept json
// @Produce json
// @Param id path int true "Session ID"
// @Param session body entity.Session true "Session data"
// @Success 200 {object} map[string]interface{} "Session updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 404 {object} map[string]interface{} "Session not found"
// @Router /api/sessions/{id} [patch]
func (h *SessionHandler) UpdateSession(c *gin.Context) {
	// Update a session
	var session entity.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.SessionUsecase.UpdateSession(&session); err != nil {
		c.JSON(400, gin.H{"error": "Failed to update session"})
		return
	}
	// Return a success response
	c.JSON(200, gin.H{"message": "Session updated successfully"})
}

// DeleteSession handles deleting a session
// @Summary Delete a session
// @Description Delete a session by its ID
// @Tags Sessions
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {object} map[string]interface{} "Session deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid session ID"
// @Failure 404 {object} map[string]interface{} "Session not found"
// @Router /api/sessions/{id} [delete]
func (h *SessionHandler) DeleteSession(c *gin.Context) {
	// Delete a session
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "Session ID is required"})
		return
	}
	sid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid session ID"})
		return
	}
	if err := h.SessionUsecase.DeleteSession(uint(sid)); err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete session"})
		return
	}
	// Return a success response
	c.JSON(200, gin.H{"message": "Session deleted successfully"})
}

// ForceSwaggoParse is a dummy function to ensure Swaggo parses this file.
func ForceSwaggoParseSessionHandler() {}
