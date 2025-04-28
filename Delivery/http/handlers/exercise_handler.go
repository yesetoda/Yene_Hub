package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type ExerciseHandler struct {
	Usecase *usecases.ExerciseUseCase
}

func NewExerciseHandler(usecase *usecases.ExerciseUseCase) *ExerciseHandler {
	return &ExerciseHandler{Usecase: usecase}
}

// @Summary Create a new exercise
// @Tags Exercises
// @Accept json
// @Produce json
// @Param exercise body schemas.CreateExerciseRequest true "Exercise"
// @Success 201 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Router /exercises [post]
func (h *ExerciseHandler) CreateExercise(c *gin.Context) {
	var exercise entity.Exercise
	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	if err := h.Usecase.Create(&exercise); err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	exerciseresp := schemas.ExerciseResponse{
		ID:        exercise.ID,
		TrackID:   exercise.TrackID,
		ProblemID: exercise.ProblemID,
		GroupID:   exercise.GroupID,
		CreatedAt: exercise.CreatedAt,
		UpdatedAt: exercise.UpdatedAt,
	}
	c.JSON(http.StatusCreated, schemas.SuccessResponse{
		Success: true,
		Code:    201,
		Message: "Exercise created successfully",
		Data:    exerciseresp})
}

// @Summary Get an exercise by ID
// @Tags Exercises
// @Produce json
// @Param id path int true "Exercise ID"
// @Success 200 {object} schemas.SuccessResponse
// @Failure 404 {object} schemas.ErrorResponse
// @Router /exercises/{id} [get]
func (h *ExerciseHandler) GetExerciseByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "invalid id"})
		return
	}
	exercise, err := h.Usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Exercise details",
		Data: schemas.ExerciseResponse{
			ID:        exercise.ID,
			TrackID:   exercise.TrackID,
			ProblemID: exercise.ProblemID,
			GroupID:   exercise.GroupID,
			CreatedAt: exercise.CreatedAt,
			UpdatedAt: exercise.UpdatedAt,
		},
	})
}

// @Summary List all exercises
// @Tags Exercises
// @Produce json
// @Success 200 {array} schemas.SuccessResponse
// @Router /exercises [get]
func (h *ExerciseHandler) ListExercises(c *gin.Context) {
	exercises, err := h.Usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	var exerciseResponses []schemas.ExerciseResponse
	for _, exercise := range exercises {
		exerciseResponses = append(exerciseResponses, schemas.ExerciseResponse{
			ID:        exercise.ID,
			TrackID:   exercise.TrackID,
			ProblemID: exercise.ProblemID,
			GroupID:   exercise.GroupID,
			CreatedAt: exercise.CreatedAt,
			UpdatedAt: exercise.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of exercises",
		Data:    exerciseResponses,
	})
}

// @Summary Update an exercise
// @Tags Exercises
// @Accept json
// @Produce json
// @Param id path int true "Exercise ID"
// @Param exercise body schemas.UpdateExerciseRequest true "Exercise"
// @Success 200 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 404 {object} schemas.ErrorResponse
// @Router /exercises/{id} [patch]
func (h *ExerciseHandler) UpdateExercise(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "invalid id"})
		return
	}
	var exercise entity.Exercise
	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	exercise.ID = uint(id)
	if err := h.Usecase.Update(&exercise); err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Exercise updated successfully",
		Data:    exercise})
}

// @Summary Delete an exercise
// @Tags Exercises
// @Param id path int true "Exercise ID"
// @Success 204 {object} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 404 {object} schemas.ErrorResponse
// @Router /exercises/{id} [delete]
func (h *ExerciseHandler) DeleteExercise(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "invalid id"})
		return
	}
	if err := h.Usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, schemas.SuccessResponse{
		Success: true,
		Code:    204,
		Message: "Exercise deleted successfully",
		Data:    ""})
}

// @Summary Get exercises by group ID
// @Tags Exercises
// @Produce json
// @Param group_id path int true "Group ID"
// @Success 200 {array} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Router /groups/gid/:group_id/exercises [get]	
func (h *ExerciseHandler) GetExercisesByGroupID(c *gin.Context) {
	groupID, err := strconv.ParseUint(c.Param("group_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "invalid group id"})
		return
	}
	exercises, err := h.Usecase.GetByGroupID(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	var exerciseResponses []schemas.ExerciseResponse
	for _, exercise := range exercises {
		exerciseResponses = append(exerciseResponses, schemas.ExerciseResponse{
			ID:        exercise.ID,
			TrackID:   exercise.TrackID,
			ProblemID: exercise.ProblemID,
			GroupID:   exercise.GroupID,
			CreatedAt: exercise.CreatedAt,
			UpdatedAt: exercise.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of exercises",
		Data:    exerciseResponses})
}

// @Summary Get exercises by track ID
// @Tags Exercises
// @Produce json
// @Param track_id path int true "Track ID"
// @Success 200 {array} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Router /tracks/tid/:track_id/exercises [get]
func (h *ExerciseHandler) GetExercisesByTrackID(c *gin.Context) {
	trackID, err := strconv.ParseUint(c.Param("track_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "invalid track id"})
		return
	}
	exercises, err := h.Usecase.GetByTrackID(uint(trackID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	var exerciseResponses []schemas.ExerciseResponse
	for _, exercise := range exercises {
		exerciseResponses = append(exerciseResponses, schemas.ExerciseResponse{
			ID:        exercise.ID,
			TrackID:   exercise.TrackID,
			ProblemID: exercise.ProblemID,
			GroupID:   exercise.GroupID,
			CreatedAt: exercise.CreatedAt,
			UpdatedAt: exercise.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of exercises",
		Data:    exerciseResponses})
}

// @Summary Get exercises by problem ID
// @Tags Exercises
// @Produce json
// @Param problem_id path int true "Problem ID"
// @Success 200 {array} schemas.SuccessResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Router /problems/pid/:problem_id/exercises [get]
func (h *ExerciseHandler) GetExercisesByProblemID(c *gin.Context) {
	problemID, err := strconv.ParseUint(c.Param("problem_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "invalid problem id"})
		return
	}
	exercises, err := h.Usecase.GetByProblemID(uint(problemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	var exerciseResponses []schemas.ExerciseResponse
	for _, exercise := range exercises {
		exerciseResponses = append(exerciseResponses, schemas.ExerciseResponse{
			ID:        exercise.ID,
			TrackID:   exercise.TrackID,
			ProblemID: exercise.ProblemID,
			GroupID:   exercise.GroupID,
			CreatedAt: exercise.CreatedAt,
			UpdatedAt: exercise.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of exercises",
		Data:    exerciseResponses})
}