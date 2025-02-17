package repository

import (
	"context"

	"github.com/Dung24-6/go-pay-gate/internal/model"
)

// TransactionRepository defines methods to interact with transaction data.
type TransactionRepository interface {
	CreateTransaction(ctx context.Context, tx *model.Transaction) error
	GetTransactionByID(ctx context.Context, id string) (*model.Transaction, error)
	UpdateTransactionStatus(ctx context.Context, id string, status string) error
}
