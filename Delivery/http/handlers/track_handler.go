package handlers

import (
	"net/http"
	"strconv"

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
// @Param track body entity.Track true "Track data"
// @Success 201 {object} entity.Track "Track created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/tracks [post]
func (h *TrackHandler) CreateTrack(c *gin.Context) {
	// Get the track data from the request body
	var track entity.Track
	if err := c.ShouldBindJSON(&track); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create the track using the use case
	if err := h.TrackUsecase.CreateTrack(&track); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the created track
	c.JSON(http.StatusOK, gin.H{"track": track})
}

// ListTrack handles listing all tracks
// @Summary List tracks
// @Description Get a list of all tracks
// @Tags Tracks
// @Produce json
// @Success 200 {array} entity.Track "List of tracks"
// @Router /api/tracks [get]
func (h *TrackHandler) ListTrack(c *gin.Context) {
	// Get the list of tracks using the use case
	tracks, err := h.TrackUsecase.ListTrack()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the list of tracks
	c.JSON(http.StatusOK, gin.H{"tracks": tracks})
}

// GetTrackByID handles getting a track by ID
// @Summary Get track by ID
// @Description Get a track by its ID
// @Tags Tracks
// @Produce json
// @Param id path int true "Track ID"
// @Success 200 {object} entity.Track "Track details"
// @Failure 400 {object} map[string]string "Invalid track ID"
// @Failure 404 {object} map[string]string "Track not found"
// @Router /api/tracks/{id} [get]
func (h *TrackHandler) GetTrackByID(c *gin.Context) {
	// Get the track ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	trackID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the track using the use case
	track, err := h.TrackUsecase.GetTrackByID(uint(trackID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the track
	c.JSON(http.StatusOK, gin.H{"track": track})
}

// GetTrackByName handles getting a track by name
// @Summary Get track by name
// @Description Get a track by its name
// @Tags Tracks
// @Produce json
// @Param name path string true "Track Name"
// @Success 200 {object} entity.Track "Track details"
// @Failure 400 {object} map[string]string "Invalid track name"
// @Failure 404 {object} map[string]string "Track not found"
// @Router /api/tracks/name/{name} [get]
func (h *TrackHandler) GetTrackByName(c *gin.Context) {
	// Get the track name from the URL parameter
	name := c.Param("name")
	// Get the track using the use case
	track, err := h.TrackUsecase.GetTrackByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the track
	c.JSON(http.StatusOK, gin.H{"track": track})
}

// UpdateTrack handles updating a track
// @Summary Update a track
// @Description Update a track by its ID
// @Tags Tracks
// @Accept json
// @Produce json
// @Param id path int true "Track ID"
// @Param track body entity.Track true "Track data"
// @Success 200 {object} entity.Track "Track updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/tracks/{id} [patch]
func (h *TrackHandler) UpdateTrack(c *gin.Context) {
	// Get the track ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	trackID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the updated track data from the request body
	var track entity.Track
	if err := c.ShouldBindJSON(&track); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Set the ID of the track to update
	track.ID = uint(trackID)
	// Update the track using the use case
	if err := h.TrackUsecase.UpdateTrack(&track); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the updated track
	c.JSON(http.StatusOK, gin.H{"track": track})
}

// DeleteTrack handles deleting a track
// @Summary Delete a track
// @Description Delete a track by its ID
// @Tags Tracks
// @Produce json
// @Param id path int true "Track ID"
// @Success 200 {object} map[string]string "Track deleted successfully"
// @Failure 400 {object} map[string]string "Invalid track ID"
// @Failure 404 {object} map[string]string "Track not found"
// @Router /api/tracks/{id} [delete]
func (h *TrackHandler) DeleteTrack(c *gin.Context) {
	// Get the track ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	trackID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Delete the track using the use case
	if err := h.TrackUsecase.DeleteTrack(uint(trackID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Track deleted successfully"})
}

// ForceSwaggoParse is a dummy function to ensure Swaggo parses this file.
func ForceSwaggoParseTrackHandler() {}