package handlers

import (
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
// @Success 201 {object} schemas.ProblemResponse "Problem created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/problems [post]
func (h *ProblemHandler) CreateProblem(c *gin.Context) {
	problem := &entity.Problem{}
	err := c.BindJSON(problem)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err = h.ProblemUsecase.CreateProblem(problem)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(201, problem)
}

// ListProblems handles listing all problems
// @Summary List problems
// @Description Get a list of all problems
// @Tags Problems
// @Produce json
// @Success 200 {array} schemas.ProblemResponse "List of problems"
// @Router /api/problems [get]
func (h *ProblemHandler) ListProblems(c *gin.Context) {
	problems, err := h.ProblemUsecase.ListProblem()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, problems)
}

// GetProblemByName handles getting a problem by name
// @Summary Get problem by name
// @Description Get a problem by its name
// @Tags Problems
// @Produce json
// @Param name path string true "Problem name"
// @Success 200 {object} schemas.ProblemResponse "Problem details"
// @Failure 400 {object} map[string]string "Invalid problem name"
// @Failure 404 {object} map[string]string "Problem not found"
// @Router /api/problems/name/{name} [get]
func (h *ProblemHandler) GetProblemByName(c *gin.Context) {
	name := c.Param("name")
	problem, err := h.ProblemUsecase.GetProblemByName(name)
	if err != nil {
		c.JSON(404, gin.H{"error": "Problem not found"})
		return
	}
	c.JSON(200, problem)
}

// GetProblemByID handles getting a problem by ID
// @Summary Get problem by ID
// @Description Get a problem by its ID
// @Tags Problems
// @Produce json
// @Param id path int true "Problem ID"
// @Success 200 {object} schemas.ProblemResponse "Problem details"
// @Failure 400 {object} map[string]string "Invalid problem ID"
// @Failure 404 {object} map[string]string "Problem not found"
// @Router /api/problems/{id} [get]
func (h *ProblemHandler) GetProblemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid problem ID"})
		return
	}
	problem, err := h.ProblemUsecase.GetProblemByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Problem not found"})
		return
	}
	c.JSON(200, problem)
}

// UpdateProblem handles updating a problem
// @Summary Update a problem
// @Description Update a problem by its ID
// @Tags Problems
// @Accept json
// @Produce json
// @Param id path int true "Problem ID"
// @Param problem body schemas.UpdateProblemRequest true "Problem data"
// @Success 200 {object} schemas.ProblemResponse "Problem updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/problems/{id} [patch]
func (h *ProblemHandler) UpdateProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid problem ID"})
		return
	}
	problem := &schemas.UpdateProblemRequest{}
	err = c.BindJSON(problem)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	problem.ID = uint(id)
	err = h.ProblemUsecase.UpdateProblem(problem)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, problem)
}

// DeleteProblem handles deleting a problem
// @Summary Delete a problem
// @Description Delete a problem by its ID
// @Tags Problems
// @Produce json
// @Param id path int true "Problem ID"
// @Success 200 {object} map[string]string "Problem deleted successfully"
// @Failure 400 {object} map[string]string "Invalid problem ID"
// @Failure 404 {object} map[string]string "Problem not found"
// @Router /api/problems/{id} [delete]
func (h *ProblemHandler) DeleteProblem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid problem ID"})
		return
	}
	err = h.ProblemUsecase.DeleteProblem(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Problem not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Problem deleted successfully"})
}

func (h *ProblemHandler) GetProblemByDifficulty(difficulty string) ([]*entity.Problem, error) {
	problems, err := h.ProblemUsecase.GetProblemByDifficulty(difficulty)
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (h *ProblemHandler) GetProblemByTag(tag string) ([]*entity.Problem, error) {
	problems, err := h.ProblemUsecase.GetProblemByTag(tag)
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (h *ProblemHandler) GetProblemByPlatform(platform string) ([]*entity.Problem, error) {
	problems, err := h.ProblemUsecase.GetProblemByPlatform(platform)
	if err != nil {
		return nil, err
	}
	return problems, nil
}
