package handlers

import (
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type TrackHandler struct {
	TrackUsecase usecases.TrackUsecase
}

func NewTrackHandler(trackUsecase usecases.TrackUsecase) *TrackHandler {
	return &TrackHandler{
		TrackUsecase: trackUsecase,
	}
}

// CreateTrack handles creating a new track
// @Summary Create a new track
// @Description Create a new track entry
// @Tags Tracks
// @Accept json
// @Produce json
// @Param track body schemas.CreateTrackRequest true "Track data"
// @Success 201 {object} schemas.SuccessResponse "Track created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/tracks [post]
func (h *TrackHandler) CreateTrack(c *gin.Context) {
	// Get the track data from the request body
	var track entity.Track
	if err := c.ShouldBindJSON(&track); err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid request body", Details: err.Error()})
		return
	}
	// Create the track using the use case
	if err := h.TrackUsecase.CreateTrack(&track); err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Failed to create track", Details: err.Error()})
		return
	}
	// Return the created track
	c.JSON(201, schemas.SuccessResponse{Success: true, Code: 201, Message: "Track created successfully", Data: track})
}

// ListTrack handles listing all tracks
// @Summary List tracks
// @Description Get a list of all tracks
// @Tags Tracks
// @Produce json
// @Success 200 {object} schemas.SuccessResponse "List of tracks"
// @Router /api/tracks [get]
func (h *TrackHandler) ListTrack(c *gin.Context) {
	// Get the list of tracks using the use case
	tracks, err := h.TrackUsecase.ListTrack()
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Failed to list tracks", Details: err.Error()})
		return
	}
	// Return the list of tracks
	c.JSON(200, schemas.SuccessResponse{Success: true, Code: 200, Message: "List of tracks", Data: tracks})
}

// GetTrackByID handles getting a track by ID
// @Summary Get track by ID
// @Description Get a track by its ID
// @Tags Tracks
// @Produce json
// @Param id path int true "Track ID"
// @Success 200 {object} schemas.SuccessResponse "Track details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid track ID"
// @Failure 404 {object} schemas.ErrorResponse "Track not found"
// @Router /api/tracks/{id} [get]
func (h *TrackHandler) GetTrackByID(c *gin.Context) {
	// Get the track ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	trackID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid track ID", Details: err.Error()})
		return
	}
	// Get the track using the use case
	track, err := h.TrackUsecase.GetTrackByID(uint(trackID))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Track not found", Details: err.Error()})
		return
	}
	// Return the track
	c.JSON(200, schemas.SuccessResponse{Success: true, Code: 200, Message: "Track details", Data: track})
}

// GetTrackByName handles getting a track by name
// @Summary Get track by name
// @Description Get a track by its name
// @Tags Tracks
// @Produce json
// @Param name path string true "Track Name"
// @Success 200 {object} schemas.SuccessResponse "Track details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid track name"
// @Failure 404 {object} schemas.ErrorResponse "Track not found"
// @Router /api/tracks/name/{name} [get]
func (h *TrackHandler) GetTrackByName(c *gin.Context) {
	// Get the track name from the URL parameter
	name := c.Param("name")
	// Get the track using the use case
	track, err := h.TrackUsecase.GetTrackByName(name)
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Track not found", Details: err.Error()})
		return
	}
	// Return the track
	c.JSON(200, schemas.SuccessResponse{Success: true, Code: 200, Message: "Track details", Data: track})
}

// UpdateTrack handles updating a track
// @Summary Update a track
// @Description Update a track by its ID
// @Tags Tracks
// @Accept json
// @Produce json
// @Param id path int true "Track ID"
// @Param track body schemas.UpdateTrackRequest true "Track data"
// @Success 200 {object} schemas.SuccessResponse "Track updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid input"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/tracks/{id} [patch]
func (h *TrackHandler) UpdateTrack(c *gin.Context) {
	// Get the track ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	trackID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid track ID", Details: err.Error()})
		return
	}
	// Get the updated track data from the request body
	var track entity.Track
	if err := c.ShouldBindJSON(&track); err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid input", Details: err.Error()})
		return
	}
	// Set the ID of the track to update
	track.ID = uint(trackID)
	// Update the track using the use case
	if err := h.TrackUsecase.UpdateTrack(&track); err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: 500, Message: "Failed to update track", Details: err.Error()})
		return
	}
	// Return the updated track
	c.JSON(200, schemas.SuccessResponse{Success: true, Code: 200, Message: "Track updated successfully", Data: track})
}

// DeleteTrack handles deleting a track
// @Summary Delete a track
// @Description Delete a track by its ID
// @Tags Tracks
// @Produce json
// @Param id path int true "Track ID"
// @Success 200 {object} schemas.SuccessResponse "Track deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid track ID"
// @Failure 404 {object} schemas.ErrorResponse "Track not found"
// @Router /api/tracks/{id} [delete]
func (h *TrackHandler) DeleteTrack(c *gin.Context) {
	// Get the track ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	trackID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid track ID", Details: err.Error()})
		return
	}
	// Delete the track using the use case
	if err := h.TrackUsecase.DeleteTrack(uint(trackID)); err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Track not found", Details: err.Error()})
		return
	}
	// Return success message
	c.JSON(200, schemas.SuccessResponse{Success: true, Code: 200, Message: "Track deleted successfully"})
}

// ForceSwaggoParse is a dummy function to ensure Swaggo parses this file.
func ForceSwaggoParseTrackHandler() {}
