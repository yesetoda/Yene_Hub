//:TODO update it from the domain entity
package schemas

import "time"

// PaymentStatus represents the status of a stipend payment
type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentApproved  PaymentStatus = "approved"
	PaymentRejected  PaymentStatus = "rejected"
	PaymentPaid      PaymentStatus = "paid"
	PaymentFailed    PaymentStatus = "failed"
)

// PaymentMethod represents the method of stipend payment
type PaymentMethod string

const (
	BankTransfer PaymentMethod = "bank_transfer"
	MobileMoney  PaymentMethod = "mobile_money"
	Cash         PaymentMethod = "cash"
)

// CreateStipendRequest represents the request body for creating a new stipend
// swagger:model
type CreateStipendRequest struct {
	// Required: true
	UserID        uint          `json:"user_id" binding:"required" example:"1"`
	Amount        float64       `json:"amount" binding:"required" example:"500.00"`
	Currency      string        `json:"currency" binding:"required" example:"USD"`
	PaymentMethod PaymentMethod `json:"payment_method" binding:"required" example:"bank_transfer"`
	Month         time.Time     `json:"month" binding:"required"`
	
	// Optional fields
	BankName      string `json:"bank_name,omitempty" example:"Example Bank"`
	AccountNumber string `json:"account_number,omitempty" example:"1234567890"`
	AccountName   string `json:"account_name,omitempty" example:"John Doe"`
	PhoneNumber   string `json:"phone_number,omitempty" example:"+1234567890"`
	Description   string `json:"description,omitempty" example:"Monthly stipend for April 2025"`
}

// UpdateStipendRequest represents the request body for updating a stipend
// swagger:model
type UpdateStipendRequest struct {
	Amount        *float64       `json:"amount,omitempty" example:"550.00"`
	Currency      *string        `json:"currency,omitempty" example:"USD"`
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty" example:"mobile_money"`
	Status        *PaymentStatus `json:"status,omitempty" example:"payment_approved"`
	
	BankName      *string `json:"bank_name,omitempty" example:"New Bank"`
	AccountNumber *string `json:"account_number,omitempty" example:"0987654321"`
	AccountName   *string `json:"account_name,omitempty" example:"John Doe"`
	PhoneNumber   *string `json:"phone_number,omitempty" example:"+1234567890"`
	Description   *string `json:"description,omitempty" example:"Updated monthly stipend"`
	
	// Payment processing fields
	TransactionID *string    `json:"transaction_id,omitempty" example:"TXN123456"`
	PaidAt       *time.Time `json:"paid_at,omitempty"`
	Notes        *string    `json:"notes,omitempty" example:"Payment processed successfully"`
}

// StipendResponse represents a stipend in responses
// swagger:model
type StipendResponse struct {
	ID            uint          `json:"id" example:"1"`
	UserID        uint          `json:"user_id" example:"1"`
	Amount        float64       `json:"amount" example:"500.00"`
	Currency      string        `json:"currency" example:"USD"`
	PaymentMethod PaymentMethod `json:"payment_method" example:"bank_transfer"`
	Status        PaymentStatus `json:"status" example:"payment_pending"`
	Month         time.Time     `json:"month"`
	
	// Payment details
	BankName      string `json:"bank_name,omitempty" example:"Example Bank"`
	AccountNumber string `json:"account_number,omitempty" example:"1234567890"`
	AccountName   string `json:"account_name,omitempty" example:"John Doe"`
	PhoneNumber   string `json:"phone_number,omitempty" example:"+1234567890"`
	Description   string `json:"description,omitempty" example:"Monthly stipend for April 2025"`
	
	// Processing information
	TransactionID string     `json:"transaction_id,omitempty" example:"TXN123456"`
	PaidAt       *time.Time  `json:"paid_at,omitempty"`
	Notes        string      `json:"notes,omitempty" example:"Payment processed successfully"`
	
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

// StipendListQuery represents query parameters for listing stipends
// swagger:model
type StipendListQuery struct {
	Page          int            `form:"page,default=1" example:"1"`
	PageSize      int            `form:"page_size,default=10" example:"10"`
	UserID        *uint          `form:"user_id" example:"1"`
	Status        *PaymentStatus `form:"status" example:"payment_pending"`
	PaymentMethod *PaymentMethod `form:"payment_method" example:"bank_transfer"`
	StartDate     *time.Time     `form:"start_date"`
	EndDate       *time.Time     `form:"end_date"`
	MinAmount     *float64       `form:"min_amount" example:"400.00"`
	MaxAmount     *float64       `form:"max_amount" example:"600.00"`
}

// StipendListResponse represents paginated stipend results
// swagger:model
type StipendListResponse struct {
	Data []*StipendResponse `json:"data"`
	Meta PaginationMeta     `json:"meta"`
}

// StipendSummary represents summary statistics for stipends
// swagger:model
type StipendSummary struct {
	TotalAmount      float64 `json:"total_amount" example:"5000.00"`
	PaidAmount       float64 `json:"paid_amount" example:"4500.00"`
	PendingAmount    float64 `json:"pending_amount" example:"500.00"`
	TotalStipends    int     `json:"total_stipends" example:"10"`
	ProcessedStipends int     `json:"processed_stipends" example:"9"`
	PendingStipends  int     `json:"pending_stipends" example:"1"`
	Currency         string  `json:"currency" example:"USD"`
	Month            time.Time `json:"month"`
}
