package schemas

// ErrorResponse represents a standardized error response
// swagger:model
// (Already refined, used for all error endpoints)
type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Invalid request parameters"`
	Details string `json:"details,omitempty" example:"Email field is required"`
}

// PaginationMeta contains pagination information
// swagger:model
// (Already refined, used for all paginated responses)
type PaginationMeta struct {
	Total      int `json:"total" example:"100"`
	Page       int `json:"page" example:"1"`
	PageSize   int `json:"page_size" example:"10"`
	TotalPages int `json:"total_pages" example:"10"`
}

// SuccessResponse represents a generic success response
// swagger:model
// (Already refined)
type SuccessResponse struct {
	Message string `json:"message" example:"Operation successful"`
	Data    any    `json:"data,omitempty"`
}

// BulkOperationResponse represents the response for bulk operations
// swagger:model
// (Already refined)
type BulkOperationResponse struct {
	Message     string `json:"message" example:"Bulk operation completed"`
	Successful  int    `json:"successful" example:"95"`
	Failed      int    `json:"failed" example:"5"`
	TotalCount  int    `json:"total_count" example:"100"`
	FailedItems []any  `json:"failed_items,omitempty"`
}
