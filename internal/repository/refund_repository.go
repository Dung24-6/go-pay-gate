package repository

import (
	"context"

	"github.com/Dung24-6/go-pay-gate/internal/model"
)

// RefundRepository defines methods to interact with refund data.
type RefundRepository interface {
	CreateRefund(ctx context.Context, refund *model.Refund) error
	GetRefundByID(ctx context.Context, id string) (*model.Refund, error)
	UpdateRefundStatus(ctx context.Context, id string, status string) error
}
