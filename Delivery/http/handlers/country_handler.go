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
// @Summary Create a new country
// @Description Create a new country entry
// @Tags Countries
// @Accept json
// @Produce json
// @Param country body entity.Country true "Country data"
// @Success 201 {object} entity.Country "Country created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/countries [post]
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
// @Summary Get country by ID
// @Description Get a country by its ID
// @Tags Countries
// @Produce json
// @Param id path int true "Country ID"
// @Success 200 {object} entity.Country "Country details"
// @Failure 400 {object} map[string]string "Invalid country ID"
// @Failure 404 {object} map[string]string "Country not found"
// @Router /api/countries/{id} [get]
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
// @Summary Update a country
// @Description Update a country by its ID
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "Country ID"
// @Param country body entity.Country true "Country data"
// @Success 200 {object} entity.Country "Country updated successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/countries/{id} [patch]
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
// @Summary Delete a country
// @Description Delete a country by its ID
// @Tags Countries
// @Produce json
// @Param id path int true "Country ID"
// @Success 200 {object} map[string]string "Country deleted successfully"
// @Failure 400 {object} map[string]string "Invalid country ID"
// @Failure 404 {object} map[string]string "Country not found"
// @Router /api/countries/{id} [delete]
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
// @Summary List countries
// @Description Get a list of all countries
// @Tags Countries
// @Produce json
// @Success 200 {array} entity.Country "List of countries"
// @Router /api/countries [get]
func (h *CountryHandler) ListCountries(c *gin.Context) {
	countries, err := h.countryUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, countries)
}
