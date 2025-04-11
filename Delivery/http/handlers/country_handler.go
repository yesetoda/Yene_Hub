package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Domain/entity"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

// CountryHandler handles HTTP requests for country operations
type CountryHandler struct {
	countryUseCase usecases.CountryUseCase
}

// NewCountryHandler creates a new CountryHandler instance
func NewCountryHandler(countryUseCase usecases.CountryUseCase) *CountryHandler {
	return &CountryHandler{
		countryUseCase: countryUseCase,
	}
}

// CreateCountry handles creating a new country
func (h *CountryHandler) CreateCountry(c *gin.Context) {
	var country entity.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	createdCountry, err := h.countryUseCase.Create(&country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Country created successfully",
		"country": createdCountry,
	})
}

// GetCountryByID handles getting a country by ID
func (h *CountryHandler) GetCountryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	country, err := h.countryUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}

	c.JSON(http.StatusOK, country)
}

// UpdateCountry handles updating a country
func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	var country entity.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	country.ID = uint(id)
	updatedCountry, err := h.countryUseCase.Update(&country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Country updated successfully",
		"country": updatedCountry,
	})
}

// DeleteCountry handles deleting a country
func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	if err := h.countryUseCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Country deleted successfully",
	})
}

// ListCountries handles listing all countries
func (h *CountryHandler) ListCountries(c *gin.Context) {
	countries, err := h.countryUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, countries)
}
