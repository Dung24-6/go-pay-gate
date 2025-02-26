package gateway

import (
	"context"
	"time"
)

type PaymentGateway interface {
	CreatePayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error)
	QueryStatus(ctx context.Context, id string) (*PaymentStatus, error)
	ProcessCallback(data []byte) (*CallbackResponse, error)
}

type PaymentRequest struct {
	Amount         float64
	Currency       string
	OrderID        string
	Description    string
	CustomerID     string
	CustomerName   string
	CustomerEmail  string
	CustomerPhone  string
	RedirectURL    string
	WebhookURL     string
	ExpiryDuration time.Duration
	PaymentMethod  string
	Metadata       map[string]string
}

type PaymentResponse struct {
	TransactionID string
	PaymentURL    string
	Amount        float64
	Currency      string
	Status        string
	OrderID       string
	CreatedAt     time.Time
	ExpiresAt     time.Time
	QRCode        string // Optional: For QR-based payments
	DeepLink      string // Optional: For mobile payments
	PaymentMethod string
	RawResponse   map[string]interface{} // Store provider-specific response
}

// PaymentStatus represents the current status of a payment
type PaymentStatus struct {
	TransactionID  string
	OrderID        string
	Status         string // pending, success, failed, expired, refunded
	Amount         float64
	Currency       string
	PaymentMethod  string
	PaidAmount     float64   // Actual amount paid
	RefundedAmount float64   // Amount refunded (if any)
	PaidAt         time.Time // When payment was completed
	LastUpdated    time.Time
	ErrorCode      string                 // If payment failed
	ErrorMessage   string                 // Error description
	RawStatus      map[string]interface{} // Store provider-specific status
}

// CallbackResponse represents the webhook/callback response
type CallbackResponse struct {
	TransactionID  string
	OrderID        string
	Status         string
	Amount         float64
	Currency       string
	PaymentMethod  string
	PaidAmount     float64
	PaidAt         time.Time
	SignatureValid bool                   // Indicates if callback signature is valid
	RawData        map[string]interface{} // Original callback data
}

// PaymentError represents a payment processing error
type PaymentError struct {
	Code    string
	Message string
	Source  string      // Which payment provider returned the error
	Raw     interface{} // Raw error from provider
}

// Error implements the error interface
func (e *PaymentError) Error() string {
	return e.Message
}

// Config represents base configuration for payment gateways
type Config struct {
	MerchantID    string
	MerchantName  string
	ApiKey        string
	ApiSecret     string
	Environment   string
	ApiEndpoint   string
	WebhookSecret string
	Version       string
	Timeout       time.Duration
	RetryAttempts int
}

// Constants for payment statuses
const (
	StatusPending  = "pending"
	StatusSuccess  = "success"
	StatusFailed   = "failed"
	StatusExpired  = "expired"
	StatusRefunded = "refunded"
)

// Constants for payment methods
const (
	MethodCard   = "card"
	MethodQR     = "qr"
	MethodWallet = "wallet"
	MethodBank   = "bank"
)

// Constants for environments
const (
	EnvSandbox    = "sandbox"
	EnvProduction = "production"
)
