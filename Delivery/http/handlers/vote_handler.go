package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type VoteHandler struct {
	VoteUsecase usecases.VoteUsecase
}

func NewVoteHandler(voteUsecase usecases.VoteUsecase) *VoteHandler {
	return &VoteHandler{
		VoteUsecase: voteUsecase,
	}
}

// CreateVote handles creating a new vote
// @Summary Create a new vote
// @Description Create a new vote entry
// @Tags Votes
// @Accept json
// @Produce json
// @Param vote body schemas.CreateVoteRequest true "Vote data"
// @Success 201 {object} schemas.VoteResponse "Vote created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/votes [post]
func (h *VoteHandler) CreateVote(c *gin.Context) {
	var vote entity.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create the vote using the use case
	if err := h.VoteUsecase.CreateVote(&vote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the created vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// ListVote handles listing votes
// @Summary List votes
// @Description Get a list of votes
// @Tags Votes
// @Produce json
// @Success 200 {array} entity.Vote "List of votes"
// @Router /api/votes [get]
func (h *VoteHandler) ListVote(c *gin.Context) {
	// Get the vote data from the request body
	var vote entity.Vote
	// Create the vote using the use case
	if err := h.VoteUsecase.CreateVote(&vote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the created vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// GetVoteByID handles getting a vote by ID
// @Summary Get vote by ID
// @Description Get a vote by its ID
// @Tags Votes
// @Produce json
// @Param id path int true "Vote ID"
// @Success 200 {object} entity.Vote "Vote details"
// @Failure 400 {object} map[string]string "Invalid vote ID"
// @Failure 404 {object} map[string]string "Vote not found"
// @Router /api/votes/{id} [get]
func (h *VoteHandler) GetVoteByID(c *gin.Context) {
	// Get the vote ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	voteID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the vote using the use case
	vote, err := h.VoteUsecase.GetVoteByID(uint(voteID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// GetVoteByCommentID handles getting votes by comment ID
// @Summary Get votes by comment ID
// @Description Get votes for a specific comment
// @Tags Votes
// @Produce json
// @Param comment_id path int true "Comment ID"
// @Success 200 {array} entity.Vote "Votes for comment"
// @Failure 400 {object} map[string]string "Invalid comment ID"
// @Router /api/votes/comment/{comment_id} [get]
func (h *VoteHandler) GetVoteByCommentID(c *gin.Context) {
	// Get the comment ID from the URL parameter
	id := c.Param("commentID")
	// Convert the ID to uint
	commentID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the vote using the use case
	vote, err := h.VoteUsecase.GetVoteByCommentID(uint(commentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// GetVoteByPostID handles getting votes by post ID
// @Summary Get votes by post ID
// @Description Get votes for a specific post
// @Tags Votes
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {array} entity.Vote "Votes for post"
// @Failure 400 {object} map[string]string "Invalid post ID"
// @Router /api/votes/post/{post_id} [get]
func (h *VoteHandler) GetVoteByPostID(c *gin.Context) {
	// Get the Post ID from the URL parameter
	id := c.Param("PostID")
	// Convert the ID to uint
	PostID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the vote using the use case
	vote, err := h.VoteUsecase.GetVoteByPostID(uint(PostID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// GetVoteByUserID handles getting votes by user ID
// @Summary Get votes by user ID
// @Description Get votes for a specific user
// @Tags Votes
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} entity.Vote "Votes for user"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Router /api/votes/user/{user_id} [get]
func (h *VoteHandler) GetVoteByUserID(c *gin.Context) {
	// Get the User ID from the URL parameter
	id := c.Param("UserID")
	// Convert the ID to uint
	UserID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the vote using the use case
	vote, err := h.VoteUsecase.GetVoteByUserID(uint(UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// GetVoteByTrackID handles getting votes by track ID
// @Summary Get votes by track ID
// @Description Get votes for a specific track
// @Tags Votes
// @Produce json
// @Param track_id path int true "Track ID"
// @Success 200 {array} entity.Vote "Votes for track"
// @Failure 400 {object} map[string]string "Invalid track ID"
// @Router /api/votes/track/{track_id} [get]
func (h *VoteHandler) GetVoteByTrackID(c *gin.Context) {
	// Get the Track ID from the URL parameter
	id := c.Param("TrackID")
	// Convert the ID to uint
	TrackID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the vote using the use case
	vote, err := h.VoteUsecase.GetVoteByTrackID(uint(TrackID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// GetVoteBySubmissionID handles getting votes by submission ID
// @Summary Get votes by submission ID
// @Description Get votes for a specific submission
// @Tags Votes
// @Produce json
// @Param submission_id path int true "Submission ID"
// @Success 200 {array} entity.Vote "Votes for submission"
// @Failure 400 {object} map[string]string "Invalid submission ID"
// @Router /api/votes/submission/{submission_id} [get]
func (h *VoteHandler) GetVoteBySubmissionID(c *gin.Context) {
	// Get the Submission ID from the URL parameter
	id := c.Param("SubmissionID")
	// Convert the ID to uint
	SubmissionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the vote using the use case
	vote, err := h.VoteUsecase.GetVoteBySubmissionID(uint(SubmissionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// GetVoteByProblemID handles getting votes by problem ID
// @Summary Get votes by problem ID
// @Description Get votes for a specific problem
// @Tags Votes
// @Produce json
// @Param problem_id path int true "Problem ID"
// @Success 200 {array} entity.Vote "Votes for problem"
// @Failure 400 {object} map[string]string "Invalid problem ID"
// @Router /api/votes/problem/{problem_id} [get]
func (h *VoteHandler) GetVoteByProblemID(c *gin.Context) {
	// Get the Problem ID from the URL parameter
	id := c.Param("ProblemID")
	// Convert the ID to uint
	ProblemID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the vote using the use case
	vote, err := h.VoteUsecase.GetVoteByProblemID(uint(ProblemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// UpdateVote handles updating a vote
// @Summary Update a vote
// @Description Update a vote by its ID
// @Tags Votes
// @Accept json
// @Produce json
// @Param id path int true "Vote ID"
// @Param vote body schemas.UpdateVoteRequest true "Vote data"
// @Success 200 {object} schemas.VoteResponse "Vote updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/votes/{id} [patch]
func (h *VoteHandler) UpdateVote(c *gin.Context) {
	// Get the vote ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	voteID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Get the vote data from the request body
	var vote entity.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vote.ID = uint(voteID)
	// Update the vote using the use case
	if err := h.VoteUsecase.UpdateVote(&vote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the updated vote
	c.JSON(http.StatusOK, gin.H{"vote": vote})
}

// DeleteVote handles deleting a vote
// @Summary Delete a vote
// @Description Delete a vote by its ID
// @Tags Votes
// @Produce json
// @Param id path int true "Vote ID"
// @Success 200 {object} map[string]string "Vote deleted successfully"
// @Failure 400 {object} map[string]string "Invalid vote ID"
// @Failure 404 {object} map[string]string "Vote not found"
// @Router /api/votes/{id} [delete]
func (h *VoteHandler) DeleteVote(c *gin.Context) {
	// Get the vote ID from the URL parameter
	id := c.Param("id")
	// Convert the ID to uint
	voteID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// Delete the vote using the use case
	if err := h.VoteUsecase.DeleteVote(uint(voteID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Vote deleted successfully"})
}

// ForceSwaggoParse is a dummy function to ensure Swaggo parses this file.
func ForceSwaggoParseVoteHandler() {}
