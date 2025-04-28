package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type ProblemTrackHandler struct {
	Usecase usecases.ProblemTracksUsecaseInterface
}

func NewProblemTrackHandler(usecase usecases.ProblemTracksUsecaseInterface) *ProblemTrackHandler {
	return &ProblemTrackHandler{Usecase: usecase}
}

// AddProblemToTrack godoc
// @Summary Add a problem to a track
// @Description Add a problem to a specific track
// @Tags ProblemTracks
// @Accept json
// @Produce json
// @Param track_id path int true "Track ID"
// @Param problem body schemas.AddProblemToTrackRequest true "Problem"
// @Success 201 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /api/tracks/tid/{track_id}/problems [post]
func (h *ProblemTrackHandler) AddProblemToTrack(c *gin.Context) {
	trackID, err := strconv.ParseUint(c.Param("track_id"), 10, 64)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid track ID"})
		return
	}
	var Pt schemas.AddProblemToTrackRequest
	if err := c.ShouldBindJSON(&Pt); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	if err := h.Usecase.AddProblemToTrack(uint(trackID), Pt.ProblemID); err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, schemas.SuccessResponse{
		Success: true,
		Code:    201,
		Message: "Problem added to track",
		Data:    "Problem added to track",
	})
}

// ListProblemsInTrack godoc
// @Summary List problems in a track
// @Description Get all problems for a specific track
// @Tags ProblemTracks
// @Produce json
// @Param track_id path int true "Track ID"
// @Success 200 {array} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /api/tracks/tid/{track_id}/problems [get]
func (h *ProblemTrackHandler) ListProblemsInTrack(c *gin.Context) {
	trackID, err := strconv.ParseUint(c.Param("track_id"), 10, 64)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid track ID"})
		return
	}
	problems, err := h.Usecase.ListProblemsInTrack(uint(trackID))
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	var problemsResp []schemas.ProblemResponse
	for _, problem := range problems {
		problemsResp = append(problemsResp, schemas.ProblemResponse{
			ID:        problem.ID,
			Name:      problem.Name,
			Difficulty: schemas.Difficulty(problem.Difficulty),
			Tag:       []string{problem.Tag},
			Platform:  problem.Platform,
			CreatedAt: problem.CreatedAt,
			UpdatedAt: problem.UpdatedAt,
		})
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of problems in track",
		Data:    problemsResp})
}

// GetProblemInTracksByName godoc
// @Summary Get problem in track by name
// @Description Get a problem in a track by its name
// @Tags ProblemTracks
// @Produce json
// @Param track_id path int true "Track ID"
// @Param name query string true "Problem name"
// @Success 200 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 404 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /api/tracks/tid/{track_id}/problems/by-name [get]
func (h *ProblemTrackHandler) GetProblemInTracksByName(c *gin.Context) {
	trackID, err := strconv.ParseUint(c.Param("track_id"), 10, 64)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid track ID"})
		return
	}
	name := c.Query("name")
	if name == "" {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "name is required"})
		return
	}
	pt, err := h.Usecase.GetProblemInTracksByName(uint(trackID), name)
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: http.StatusNotFound, Message: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Problem in track by name",
		Data: schemas.ProblemResponse{
			ID:        pt.ID,
			Name:      pt.Name,
			Difficulty: schemas.Difficulty(pt.Difficulty),
			Tag:       []string{pt.Tag},
			Platform:  pt.Platform,
				CreatedAt: pt.CreatedAt,
			UpdatedAt: pt.UpdatedAt,
		}})
}

// GetProblemInTracksByDifficulty godoc
// @Summary Get problems in track by difficulty
// @Description Get problems in a track by difficulty
// @Tags ProblemTracks
// @Produce json
// @Param track_id path int true "Track ID"
// @Param difficulty query string true "Difficulty"
// @Success 200 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /api/tracks/tid/{track_id}/problems/by-difficulty [get]
func (h *ProblemTrackHandler) GetProblemInTracksByDifficulty(c *gin.Context) {
	trackID, err := strconv.ParseUint(c.Param("track_id"), 10, 64)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid track ID"})
		return
	}
	difficulty := c.Query("difficulty")
	if difficulty == "" {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "difficulty is required"})
		return
	}
	pts, err := h.Usecase.GetProblemInTracksByDifficulty(uint(trackID), difficulty)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	var problemsResp []schemas.ProblemResponse
	for _, pt := range pts {
		problemsResp = append(problemsResp, schemas.ProblemResponse{
			ID:        pt.ID,
			Name:      pt.Name,
			Difficulty: schemas.Difficulty(pt.Difficulty),
			Tag:       []string{pt.Tag},
			Platform:  pt.Platform,
			CreatedAt: pt.CreatedAt,
			UpdatedAt: pt.UpdatedAt,
		})
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of problems in track by difficulty",
		Data:    problemsResp})
}

// GetProblemInTracksByTag godoc
// @Summary Get problems in track by tag
// @Description Get problems in a track by tag
// @Tags ProblemTracks
// @Produce json
// @Param track_id path int true "Track ID"
// @Param tag query string true "Tag"
// @Success 200 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /api/tracks/tid/{track_id}/problems/by-tag [get]
func (h *ProblemTrackHandler) GetProblemInTracksByTag(c *gin.Context) {
	trackID, err := strconv.ParseUint(c.Param("track_id"), 10, 64)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid track ID"})
		return
	}
	tag := c.Query("tag")
	if tag == "" {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "tag is required"})
		return
	}
	pts, err := h.Usecase.GetProblemInTracksByTag(uint(trackID), tag)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	var problemsResp []schemas.ProblemResponse
	for _, pt := range pts {
		problemsResp = append(problemsResp, schemas.ProblemResponse{
			ID:        pt.ID,
			Name:      pt.Name,
			Difficulty: schemas.Difficulty(pt.Difficulty),
			Tag:       []string{pt.Tag},
			Platform:  pt.Platform,
			CreatedAt: pt.CreatedAt,
			UpdatedAt: pt.UpdatedAt,
		})
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of problems in track by tag",
		Data:    problemsResp})
}

// GetProblemInTracksByPlatform godoc
// @Summary Get problems in track by platform
// @Description Get problems in a track by platform
// @Tags ProblemTracks
// @Produce json
// @Param track_id path int true "Track ID"
// @Param platform query string true "Platform"
// @Success 200 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /api/tracks/tid/{track_id}/problems/by-platform [get]
func (h *ProblemTrackHandler) GetProblemInTracksByPlatform(c *gin.Context) {
	trackID, err := strconv.ParseUint(c.Param("track_id"), 10, 64)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid track ID"})
		return
	}
	platform := c.Query("platform")
	if platform == "" {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "platform is required"})
		return
	}
	pts, err := h.Usecase.GetProblemInTracksByPlatform(uint(trackID), platform)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	var problemsResp []schemas.ProblemResponse
	for _, pt := range pts {
		problemsResp = append(problemsResp, schemas.ProblemResponse{
			ID:        pt.ID,
			Name:      pt.Name,
			Difficulty: schemas.Difficulty(pt.Difficulty),
			Tag:       []string{pt.Tag},
			Platform:  pt.Platform,
			CreatedAt: pt.CreatedAt,
			UpdatedAt: pt.UpdatedAt,
		})
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of problems in track by platform",
		Data:    problemsResp})
}

// RemoveProblemFromTrack godoc
// @Summary Remove a problem from a track
// @Description Remove a problem from a track by ProblemTrack ID
// @Tags ProblemTracks
// @Param id path int true "ProblemTrack ID"
// @Success 204 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /api/problem-tracks/{id} [delete]
func (h *ProblemTrackHandler) RemoveProblemFromTrack(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid id"})
		return
	}
	if err := h.Usecase.RemoveProblemFromTrack(uint(id)); err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(204, schemas.SuccessResponse{
		Success: true,
		Code:    204,
		Message: "Problem removed from track",
		Data:    ""})
}