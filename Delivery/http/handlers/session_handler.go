package handlers

import (
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
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
// @Param session body schemas.CreateSessionRequest true "Session data"
// @Success 201 {object} schemas.SuccessResponse "Session created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/sessions [post]
func (h *SessionHandler) CreateSession(c *gin.Context) {
	var session entity.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid request body", Details: err.Error()})
		return
	}
	if err := h.SessionUsecase.CreateSession(&session); err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Failed to create session", Details: err.Error()})
		return
	}
	c.JSON(201, schemas.SuccessResponse{Message: "Session created successfully", Data: session})
}

// ListSessions handles listing all sessions
// @Summary List sessions
// @Description Get a list of all sessions
// @Tags Sessions
// @Produce json
// @Success 200 {object} schemas.SuccessResponse "List of sessions"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/sessions [get]
func (h *SessionHandler) ListSessions(c *gin.Context) {
	sessions, err := h.SessionUsecase.ListSession()
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Failed to fetch sessions", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Message: "List of sessions", Data: sessions})
}

// GetSessionByID handles getting a session by ID
// @Summary Get session by ID
// @Description Get a session by its ID
// @Tags Sessions
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {object} schemas.SuccessResponse "Session details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid session ID"
// @Failure 404 {object} schemas.ErrorResponse "Session not found"
// @Router /api/sessions/{id} [get]
func (h *SessionHandler) GetSessionByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Session ID is required"})
		return
	}
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid session ID", Details: err.Error()})
		return
	}
	session, err := h.SessionUsecase.GetSessionByID(uint(sessionID))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Session not found", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Message: "Session details", Data: session})
}

// GetSessionByName handles getting a session by name
// @Summary Get session by name
// @Description Get a session by its name
// @Tags Sessions
// @Produce json
// @Param name path string true "Session name"
// @Success 200 {object} schemas.SuccessResponse "Session details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid session name"
// @Failure 404 {object} schemas.ErrorResponse "Session not found"
// @Router /api/sessions/name/{name} [get]
func (h *SessionHandler) GetSessionByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Session name is required"})
		return
	}
	sessions, err := h.SessionUsecase.GetSessionByName(name)
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Session not found", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Message: "Session details", Data: sessions})
}

// GetSessionByStartTime handles getting a session by start time
// @Summary Get session by start time
// @Description Get a session by its start time
// @Tags Sessions
// @Produce json
// @Param start_time path string true "Session start time"
// @Success 200 {object} schemas.SuccessResponse "Session details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid session start time"
// @Failure 404 {object} schemas.ErrorResponse "Session not found"
// @Router /api/sessions/start-time/{start_time} [get]
func (h *SessionHandler) GetSessionByStartTime(c *gin.Context) {
	startTime := c.Param("start_time")
	if startTime == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Session start time is required"})
		return
	}
	sessions, err := h.SessionUsecase.GetSessionByStartTime(startTime)
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Session not found", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Message: "Session details", Data: sessions})
}

// UpdateSession handles updating a session
// @Summary Update a session
// @Description Update a session by its ID
// @Tags Sessions
// @Accept json
// @Produce json
// @Param id path int true "Session ID"
// @Param session body schemas.UpdateSessionRequest true "Session data"
// @Success 200 {object} schemas.SuccessResponse "Session updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 404 {object} schemas.ErrorResponse "Session not found"
// @Router /api/sessions/{id} [patch]
func (h *SessionHandler) UpdateSession(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Session ID is required"})
		return
	}
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid session ID", Details: err.Error()})
		return
	}
	var session entity.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid request body", Details: err.Error()})
		return
	}
	session.ID = uint(sessionID)
	if err := h.SessionUsecase.UpdateSession(&session); err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Session not found", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Message: "Session updated successfully", Data: session})
}

// DeleteSession handles deleting a session
// @Summary Delete a session
// @Description Delete a session by its ID
// @Tags Sessions
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {object} schemas.SuccessResponse "Session deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid session ID"
// @Failure 404 {object} schemas.ErrorResponse "Session not found"
// @Router /api/sessions/{id} [delete]
func (h *SessionHandler) DeleteSession(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Session ID is required"})
		return
	}
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid session ID", Details: err.Error()})
		return
	}
	if err := h.SessionUsecase.DeleteSession(uint(sessionID)); err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Session not found", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Message: "Session deleted successfully"})
}

// ForceSwaggoParseSessionHandler is a dummy function to ensure Swaggo parses this file.
func ForceSwaggoParseSessionHandler() {}
