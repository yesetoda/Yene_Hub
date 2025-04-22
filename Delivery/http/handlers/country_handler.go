package handlers

import (
	"net/http"
	"strconv"

	"a2sv.org/hub/Delivery/http/schemas"
	"a2sv.org/hub/usecases"
	"github.com/gin-gonic/gin"
)

// CountryHandler handles HTTP requests for country operations
type CountryHandler struct {
	countryUseCase usecases.CountryUseCase
}

func NewCountryHandler(countryUseCase usecases.CountryUseCase) *CountryHandler {
	return &CountryHandler{countryUseCase: countryUseCase}
}

// CreateCountry handles creating a new country
// @Summary Create a new country
// @Description Create a new country entry
// @Tags Countries
// @Accept json
// @Produce json
// @Param country body schemas.CreateCountryRequest true "Country data"
// @Success 201 {object} schemas.CountryResponse "Country created successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/countries [post]
func (h *CountryHandler) CreateCountry(c *gin.Context) {
	var input schemas.CreateCountryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}
	createdCountry, err := h.countryUseCase.Create(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Could not create country",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, createdCountry)
}

// GetCountryByID handles getting a country by ID
// @Summary Get country by ID
// @Description Get a country by its ID
// @Tags Countries
// @Produce json
// @Param id path int true "Country ID"
// @Success 200 {object} schemas.CountryResponse "Country details"
// @Failure 400 {object} schemas.ErrorResponse "Invalid country ID"
// @Failure 404 {object} schemas.ErrorResponse "Country not found"
// @Router /api/countries/{id} [get]
func (h *CountryHandler) GetCountryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid country ID",
			Details: err.Error(),
		})
		return
	}
	country, err := h.countryUseCase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Country not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, country)
}

// UpdateCountry handles updating a country
// @Summary Update country
// @Description Update a country by its ID
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "Country ID"
// @Param country body schemas.UpdateCountryRequest true "Country data"
// @Success 200 {object} schemas.CountryResponse "Country updated successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid request body or country ID"
// @Failure 404 {object} schemas.ErrorResponse "Country not found"
// @Failure 500 {object} schemas.ErrorResponse "Internal server error"
// @Router /api/countries/{id} [patch]
func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid country ID",
			Details: err.Error(),
		})
		return
	}
	var input schemas.UpdateCountryRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}
	updatedCountry, err := h.countryUseCase.Update(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Country not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updatedCountry)
}

// DeleteCountry handles deleting a country
// @Summary Delete a country
// @Description Delete a country by its ID
// @Tags Countries
// @Produce json
// @Param id path int true "Country ID"
// @Success 200 {object} schemas.SuccessResponse "Country deleted successfully"
// @Failure 400 {object} schemas.ErrorResponse "Invalid country ID"
// @Failure 404 {object} schemas.ErrorResponse "Country not found"
// @Router /api/countries/{id} [delete]
func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "Invalid country ID",
			Details: err.Error(),
		})
		return
	}
	err = h.countryUseCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{
			Code: http.StatusNotFound,
			Message: "Country not found",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schemas.SuccessResponse{Message: "Country deleted successfully"})
}

// ListCountries handles listing all countries
// @Summary List countries
// @Description Get a list of all countries
// @Tags Countries
// @Produce json
// @Success 200 {object} schemas.CountryListResponse "List of countries"
// @Router /api/countries [get]
func (h *CountryHandler) ListCountries(c *gin.Context) {
	countries, meta, err := h.countryUseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{
			Code: http.StatusInternalServerError,
			Message: "Failed to list countries",
			Details: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schemas.CountryListResponse{Data: countries, Meta: *meta})
}
