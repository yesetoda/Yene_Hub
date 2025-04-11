package schemas

import "a2sv.org/hub/Domain/entity"

// ErrorResponse represents a standardized error response
// swagger:model
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// PaginationMeta contains pagination information
// swagger:model
type PaginationMeta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
}

// PaginatedUsers represents paginated user results
// swagger:model
type PaginatedUsers struct {
	Data []*entity.User `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

// AuthTokenResponse represents authentication token response
// swagger:model
type AuthTokenResponse struct {
	Token string       `json:"token"`
	User  *entity.User `json:"user"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type LoginResponse struct {
	Message string       `json:"message"`
	Token   string       `json:"token"`
	User    *entity.User `json:"user"`
}
