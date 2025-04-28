package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type SubmissionHandler struct {
	SubmissionUsecase usecases.SubmissionUsecase
}

func NewSubmissionHandeler(submissionUsecase usecases.SubmissionUsecase) *SubmissionHandler {
	return &SubmissionHandler{
		SubmissionUsecase: submissionUsecase,
	}
}

// CreateSubmission handles creating a new submission
// @Summary Create a new submission
// @Description Create a new submission entry
// @Tags Submissions
// @Accept json
// @Produce json
// @Param submission body schemas.CreateSubmissionRequest true "Submission data"
// @Success 201 {object} schemas.SuccessResponse "Submission created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/submissions [post]
func (s *SubmissionHandler) CreateSubmission(c *gin.Context) {
	// Get the request body
	var submission entity.Submission
	if err := c.BindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Code: 400, Message: "Invalid request body", Details: err.Error()})
		return
	}
	err := s.SubmissionUsecase.CreateSubmission(&submission)
	// Handle the error
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Code: 500, Message: "Internal server error", Details: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, schemas.SuccessResponse{Success: true, Code: 201, Message: "Submission created successfully", Data: submission})
}

// ListSubmission handles listing all submissions
// @Summary List submissions
// @Description Get a list of all submissions
// @Tags Submissions
// @Produce json
// @Success 200 {object} schemas.SuccessResponse "List of submissions"
// @Router /api/submissions [get]
func (s *SubmissionHandler) ListSubmission(c *gin.Context) {
	submission, err := s.SubmissionUsecase.ListSubmission()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Code: 500, Message: "Internal server error", Details: err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{Success: true, Code: 200, Message: "List of submissions", Data: submission})
}

// GetSubmissionByID handles getting a submission by ID
// @Summary Get submission by ID
// @Description Get a submission by its ID
// @Tags Submissions
// @Produce json
// @Param id path int true "Submission ID"
// @Success 200 {object} schemas.SuccessResponse "Submission details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid submission ID"
// @Failure 404 {object} schemas.ErrorResponse "Submission not found"
// @Router /api/submissions/{id} [get]
func (s *SubmissionHandler) GetSubmissionByID(c *gin.Context) {
	submissionID := c.Param("id")
	sid, err := strconv.Atoi(submissionID)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid submission ID", Details: err.Error()})
		return
	}
	// Get the submission by ID
	submission, err := s.SubmissionUsecase.GetSubmissionByID(uint(sid))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Submission not found", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Success: true, Code: 200, Message: "Submission details", Data: submission})
}

// GetSubmissionByProblemID handles getting submissions by problem ID
// @Summary Get submissions by problem ID
// @Description Get submissions for a specific problem
// @Tags Submissions
// @Produce json
// @Param problem_id path int true "Problem ID"
// @Success 200 {object} schemas.SuccessResponse "Submissions for problem"
// @Failure 400 {object} schemas.ErrorResponse "Invalid problem ID"
// @Router /api/submissions/problem/{problem_id} [get]
func (s *SubmissionHandler) GetSubmissionByProblemID(c *gin.Context) {
	problemID := c.Param("id")
	
	pid, err := strconv.Atoi(problemID)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid problem ID", Details: err.Error()})
		return
	}
	// Get the submission by problem ID
	submission, err := s.SubmissionUsecase.GetSubmissionByProblemID(uint(pid))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Submission not found", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Success: true, Code: 200, Message: "Submissions for problem", Data: submission})
}

// GetSubmissionByUserID handles getting submissions by user ID
// @Summary Get submissions by user ID
// @Description Get submissions for a specific user
// @Tags Submissions
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} schemas.SuccessResponse "Submissions for user"
// @Failure 400 {object} schemas.ErrorResponse "Invalid user ID"
// @Router /api/submissions/user/{user_id} [get]
func (s *SubmissionHandler) GetSubmissionByUserID(c *gin.Context) {
	userID := c.Param("id")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(400, schemas.ErrorResponse{Code: 400, Message: "Invalid user ID", Details: err.Error()})
		return
	}

	// Get the submission by user ID
	submission, err := s.SubmissionUsecase.GetSubmissionByUserID(uint(uid))
	if err != nil {
		c.JSON(404, schemas.ErrorResponse{Code: 404, Message: "Submission not found", Details: err.Error()})
		return
	}
	c.JSON(200, schemas.SuccessResponse{Success: true, Code: 200, Message: "Submissions for user", Data: submission})
}
