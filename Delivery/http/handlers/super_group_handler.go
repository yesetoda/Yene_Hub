package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

type SuperGroupHandler struct {
	superGroupUseCase usecases.SuperGroupUseCaseInterface
}

func NewSuperGroupHandler(superGroupUseCase usecases.SuperGroupUseCaseInterface) *SuperGroupHandler {
	return &SuperGroupHandler{
		superGroupUseCase: superGroupUseCase,
	}
}

func (h *SuperGroupHandler) CreateSuperGroup(c *gin.Context) {
	var superGroup *entity.SuperGroup

	if err := c.ShouldBindJSON(&superGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	superGroup, err := h.superGroupUseCase.Create(superGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super group created successfully",
		"data":    superGroup,
	})
}

func (h *SuperGroupHandler) GetSuperGroup(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	superGroup, err := h.superGroupUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super group retrieved successfully",
		"data":    superGroup,
	})
}

func (h *SuperGroupHandler) UpdateSuperGroup(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var superGroup *entity.SuperGroup

	if err := c.ShouldBindJSON(&superGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.superGroupUseCase.Update(superGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super group updated successfully",
		"data":    superGroup,
	})
}

func (h *SuperGroupHandler) DeleteSuperGroup(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.superGroupUseCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super group deleted successfully",
	})
}

func (h *SuperGroupHandler) ListSuperGroups(c *gin.Context) {
	superGroups, err := h.superGroupUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Super groups retrieved successfully",
		"data":    superGroups,
	})
}
