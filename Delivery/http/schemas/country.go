package schemas

import "time"

// CreateCountryRequest represents the request body for creating a new country
// swagger:model CreateCountryRequest
type CreateCountryRequest struct {
	Name string `json:"name" binding:"required" example:"Ethiopia"`
	Code string `json:"code" binding:"required" example:"ET"`
}

// UpdateCountryRequest represents the request body for updating a country
// swagger:model UpdateCountryRequest
type UpdateCountryRequest struct {
	Name *string `json:"name,omitempty" example:"Ethiopia"`
	Code *string `json:"code,omitempty" example:"ET"`
}

// CountryResponse represents a country in responses
// swagger:model CountryResponse
type CountryResponse struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"Ethiopia"`
	Code      string    `json:"code" example:"ET"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CountryListResponse represents paginated country results
// swagger:model CountryListResponse
type CountryListResponse struct {
	Data []*CountryResponse `json:"data"`
	Meta PaginationMeta     `json:"meta"`
}
