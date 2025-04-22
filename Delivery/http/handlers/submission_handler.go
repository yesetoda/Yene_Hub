package handlers

import (
	"net/http"
	"strconv"

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
// @Success 201 {object} schemas.SubmissionResponse "Submission created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/submissions [post]
func (s *SubmissionHandler) CreateSubmission(c *gin.Context) {
	// Get the request body
	var submission entity.Submission
	if err := c.BindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	err := s.SubmissionUsecase.CreateSubmission(&submission)
	// Handle the error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "submission created successfully",
	})
}

// ListSubmission handles listing all submissions
// @Summary List submissions
// @Description Get a list of all submissions
// @Tags Submissions
// @Produce json
// @Success 200 {array} []*schemas.SubmissionResponse "List of submissions"
// @Router /api/submissions [get]
func (s *SubmissionHandler) ListSubmission(c *gin.Context) {
	submission, err := s.SubmissionUsecase.ListSubmission()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"submissions": submission,
	})
}

// GetSubmissionByID handles getting a submission by ID
// @Summary Get submission by ID
// @Description Get a submission by its ID
// @Tags Submissions
// @Produce json
// @Param id path int true "Submission ID"
// @Success 200 {object} schemas.SubmissionResponse "Submission details"
// @Failure 400 {object} map[string]string "Invalid submission ID"
// @Failure 404 {object} map[string]string "Submission not found"
// @Router /api/submissions/{id} [get]
func (s *SubmissionHandler) GetSubmissionByID(c *gin.Context) {
	submissionID := c.Param("id")
	sid, err := strconv.Atoi(submissionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	// Get the submission by ID
	submission, err := s.SubmissionUsecase.GetSubmissionByID(uint(sid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "submission not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"submission": submission,
	})
}

// GetSubmissionByProblemID handles getting submissions by problem ID
// @Summary Get submissions by problem ID
// @Description Get submissions for a specific problem
// @Tags Submissions
// @Produce json
// @Param problem_id path int true "Problem ID"
// @Success 200 {array} []*schemas.SubmissionResponse "Submissions for problem"
// @Failure 400 {object} map[string]string "Invalid problem ID"
// @Router /api/submissions/problem/{problem_id} [get]
func (s *SubmissionHandler) GetSubmissionByProblemID(c *gin.Context) {
	problemID := c.Param("id")
	pid, err := strconv.Atoi(problemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	// Get the submission by problem ID
	submission, err := s.SubmissionUsecase.GetSubmissionByProblemID(uint(pid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "submission not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"submissions": submission,
	})
}

// GetSubmissionByUserID handles getting submissions by user ID
// @Summary Get submissions by user ID
// @Description Get submissions for a specific user
// @Tags Submissions
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} []*schemas.SubmissionResponse "Submissions for user"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Router /api/submissions/user/{user_id} [get]
func (s *SubmissionHandler) GetSubmissionByUserID(c *gin.Context) {
	userID := c.Param("id")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	// Get the submission by user ID
	submission, err := s.SubmissionUsecase.GetSubmissionByUserID(uint(uid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "submission not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"submissions": submission,
	})
}