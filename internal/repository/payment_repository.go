package repository

import (
	"context"

	"github.com/Dung24-6/go-pay-gate/internal/model"
)

// PaymentRepository defines methods to interact with payment data.
type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *model.Payment) error
	GetPaymentByID(ctx context.Context, id string) (*model.Payment, error)
	UpdatePaymentStatus(ctx context.Context, id string, status string) error
}
