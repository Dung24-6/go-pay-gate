package repository

import (
	"context"
	"fmt"

	"github.com/Dung24-6/go-pay-gate/internal/model"
	"gorm.io/gorm"
)

type MySQLTransactionRepository struct {
	db *gorm.DB
}

func NewMySQLTransactionRepository(db *gorm.DB) *MySQLTransactionRepository {
	return &MySQLTransactionRepository{db: db}
}

func (r *MySQLTransactionRepository) CreateTransaction(ctx context.Context, tx *model.Transaction) error {
	if err := r.db.WithContext(ctx).Create(tx).Error; err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}

func (r *MySQLTransactionRepository) GetTransactionByID(ctx context.Context, id string) (*model.Transaction, error) {
	var tx model.Transaction
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&tx).Error; err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}
	return &tx, nil
}

func (r *MySQLTransactionRepository) UpdateTransactionStatus(ctx context.Context, id string, status string) error {
	if err := r.db.WithContext(ctx).Model(&model.Transaction{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return fmt.Errorf("failed to update transaction status: %w", err)
	}
	return nil
}
