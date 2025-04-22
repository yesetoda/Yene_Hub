package handlers

import (
	"strconv"
	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type StippendHandler struct {
	StippendUsecase usecases.StipendUsecase
}

func NewStippendHandler(stippendUsecase usecases.StipendUsecase) *StippendHandler {
	return &StippendHandler{
		StippendUsecase: stippendUsecase,
	}
}

// CreateStipend handles creating a new stipend
// @Summary Create a new stipend
// @Description Create a new stipend entry
// @Tags Stipends
// @Accept json
// @Produce json
// @Param stipend body schemas.CreateStipendRequest true "Stipend data"
// @Success 201 {object} schemas.StipendResponse "Stipend created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Router /api/stipends [post]
func (h *StippendHandler) CreateStipend(c *gin.Context) {
	// Get the request body
	var stippend entity.Stipend
	if err := c.ShouldBindJSON(&stippend); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	// Call the use case to create the stippend
	if err := h.StippendUsecase.CreateStipend(&stippend); err != nil {
		c.JSON(400, gin.H{"error": "Failed to create stippend"})
		return
	}
	// Return a success response
	c.JSON(201, gin.H{"message": "Stippend created successfully"})
}

// ListStippends handles listing all stipends
// @Summary List stipends
// @Description Get a list of all stipends
// @Tags Stipends
// @Produce json
// @Success 200 {array} []*schemas.StipendResponse "List of stipends"
// @Failure 400 {object} map[string]interface{} "Failed to fetch stippends"
// @Router /api/stipends [get]
func (h *StippendHandler) ListStippends(c *gin.Context) {
	// Call the use case to get the list of stippends
	stippends, err := h.StippendUsecase.ListStipend()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch stippends"})
		return
	}
	// Return the list of stippends
	c.JSON(200, stippends)
}

// GetStippendByID handles getting a stipend by ID
// @Summary Get stipend by ID
// @Description Get a stipend by its ID
// @Tags Stipends
// @Produce json
// @Param stippend_id path int true "Stipend ID"
// @Success 200 {object} schemas.StipendResponse "Stipend details"
// @Failure 400 {object} map[string]interface{} "Invalid stipend ID"
// @Router /api/stipends/{stippend_id} [get]
func (h *StippendHandler) GetStippendByID(c *gin.Context) {
	// Get the stippend ID from the URL parameter
	sid := c.Param("stippend_id")
	if sid == "" {
		c.JSON(400, gin.H{"error": "Stippend ID is required"})
		return
	}
	// Convert stippend ID to uint
	stippendID, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid stippend ID"})
		return
	}
	// Call the use case to get the stippend by ID
	stippend, err := h.StippendUsecase.GetStipendByID(uint(stippendID))
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch stippend"})
		return
	}
	// Return the stippend details
	c.JSON(200, stippend)
}

// UpdateStipend handles updating a stipend
// @Summary Update a stipend
// @Description Update a stipend by its ID
// @Tags Stipends
// @Accept json
// @Produce json
// @Param stippend_id path int true "Stipend ID"
// @Param stipend body schemas.UpdateStipendRequest true "Stipend data"
// @Success 200 {object} schemas.StipendResponse "Stipend updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Router /api/stipends/{stippend_id} [patch]
func (h *StippendHandler) UpdateStipend(c *gin.Context) {
	var stipend entity.Stipend
	if err := c.ShouldBindJSON(&stipend); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.StippendUsecase.UpdateStipend(&stipend)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to update stippend"})
		return
	}
	c.JSON(200, gin.H{"message": "Stippend updated successfully"})
}

// DeleteStipend handles deleting a stipend
// @Summary Delete a stipend
// @Description Delete a stipend by its ID
// @Tags Stipends
// @Produce json
// @Param stippend_id path int true "Stipend ID"
// @Success 200 {object} schemas.StipendResponse "Stipend deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid stipend ID"
// @Router /api/stipends/{stippend_id} [delete]
func (h *StippendHandler) DeleteStipend(c *gin.Context) {
	// Get the stippend ID from the URL parameter
	sid := c.Param("stippend_id")
	if sid == "" {
		c.JSON(400, gin.H{"error": "Stippend ID is required"})
		return
	}
	usid, err := strconv.Atoi(sid)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid stippend ID"})
		return
	}
	// Call the use case to delete the stippend by ID
	err = h.StippendUsecase.DeleteStipend(uint(usid))
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete stippend"})
		return
	}
	c.JSON(200, gin.H{"message": "Stippend deleted successfully"})
}

// ForceSwaggoParse is a dummy function to ensure Swaggo parses this file.
func ForceSwaggoParseStippendHandler() {}
