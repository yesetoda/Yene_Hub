package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type ProblemHandler struct {
	ProblemUsecase usecases.ProblemUseCaseInterface
}

func NewProblemHandler(problemUsecase usecases.ProblemUseCaseInterface) *ProblemHandler {
	return &ProblemHandler{
		ProblemUsecase: problemUsecase,
	}
}

// CreateProblem handles creating a new problem
// @Summary Create a new problem
// @Description Create a new problem entry
// @Tags Problems
// @Accept json
// @Produce json
// @Param problem body schemas.CreateProblemRequest true "Problem data"
// @Success 201 {object} schemas.SuccessResponse "Problem created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/problems [post]
func (h *ProblemHandler) CreateProblem(c *gin.Context) {
	problem := &entity.Problem{}
	err := c.BindJSON(problem)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid request body"})
		return
	}
	err = h.ProblemUsecase.CreateProblem(problem)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error"})
		return
	}
	c.JSON(201, schemas.SuccessResponse{
		Success: true,
		Code:    201,
		Message: "Problem created successfully",
		Data:    problem})
}

// ListProblems handles listing all problems
// @Summary List problems
// @Description Get a list of all problems
// @Tags Problems
// @Produce json
// @Success 200 {object} schemas.SuccessResponse "List of problems"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/problems [get]
func (h *ProblemHandler) ListProblems(c *gin.Context) {
	problems, err := h.ProblemUsecase.ListProblem()
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error"})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of problems",
		Data:    problems})
}

// GetProblemByName handles getting a problem by name
// @Summary Get problem by name
// @Description Get a problem by its name
// @Tags Problems
// @Produce json
// @Param name path string true "Problem name"
// @Success 200 {object} schemas.SuccessResponse "Problem details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid problem name"
// @Failure 404 {object} schemas.ErrorResponse "Problem not found"
// @Router /api/problems/name/{name} [get]
func (h *ProblemHandler) GetProblemByName(c *gin.Context) {
	name := c.Param("name")
	problem, err := h.ProblemUsecase.GetProblemByName(name)
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: http.StatusNotFound, Message: "Problem not found"})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Problem details",
		Data:    problem})
}

// GetProblemByID handles getting a problem by ID
// @Summary Get problem by ID
// @Description Get a problem by its ID
// @Tags Problems
// @Produce json
// @Param id path int true "Problem ID"
// @Success 200 {object} schemas.SuccessResponse "Problem details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid problem ID"
// @Failure 404 {object} schemas.ErrorResponse "Problem not found"
// @Router /api/problems/{id} [get]
func (h *ProblemHandler) GetProblemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid problem ID"})
		return
	}
	problem, err := h.ProblemUsecase.GetProblemByID(uint(id))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: http.StatusNotFound, Message: "Problem not found"})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Problem details",
		Data:    problem})
}

// UpdateProblem handles updating a problem
// @Summary Update a problem
// @Description Update a problem by its ID
// @Tags Problems
// @Accept json
// @Produce json
// @Param id path int true "Problem ID"
// @Param problem body schemas.UpdateProblemRequest true "Problem data"
// @Success 200 {object} schemas.SuccessResponse "Problem updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body or problem ID"
// @Failure 404 {object} schemas.ErrorResponse "Problem not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/problems/{id} [patch]
func (h *ProblemHandler) UpdateProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid problem ID"})
		return
	}
	problem := &schemas.UpdateProblemRequest{}
	err = c.BindJSON(problem)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid request body"})
		return
	}
	problem.ID = uint(id)
	err = h.ProblemUsecase.UpdateProblem(problem)
	if err != nil {
		c.JSON(500, schemas.ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal server error"})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Problem updated successfully",
		Data:    problem})
}

// DeleteProblem handles deleting a problem
// @Summary Delete a problem
// @Description Delete a problem by its ID
// @Tags Problems
// @Produce json
// @Param id path int true "Problem ID"
// @Success 200 {object} schemas.SuccessResponse "Problem deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid problem ID"
// @Failure 404 {object} schemas.ErrorResponse "Problem not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/problems/{id} [delete]
func (h *ProblemHandler) DeleteProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid problem ID"})
		return
	}
	err = h.ProblemUsecase.DeleteProblem(uint(id))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: http.StatusNotFound, Message: "Problem not found"})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "Problem deleted successfully",
		Data:    nil})
}

func (h *ProblemHandler) GetProblemByDifficulty(c *gin.Context) {
	problems, err := h.ProblemUsecase.GetProblemByDifficulty(c.Param("difficulty"))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: http.StatusNotFound, Message: "Problem not found"})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of problems by difficulty",
		Data:    problems})
}

func (h *ProblemHandler) GetProblemByTag(c *gin.Context) {
	tag := c.Param("tag")
	problems, err := h.ProblemUsecase.GetProblemByTag(tag)
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: http.StatusNotFound, Message: "Problem not found"})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of problems by tag",
		Data:    problems})
}

func (h *ProblemHandler) GetProblemByPlatform(	c *gin.Context) {
	platform := c.Param("platform")
	problems, err := h.ProblemUsecase.GetProblemByPlatform(platform)
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: http.StatusNotFound, Message: "Problem not found"})
		return
	}
	c.JSON(200, schemas.SuccessResponse{
		Success: true,
		Code:    200,
		Message: "List of problems by platform",
		Data:    problems})
}
