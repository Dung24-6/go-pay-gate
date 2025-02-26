package repository

import (
	"context"
	"fmt"

	"github.com/Dung24-6/go-pay-gate/internal/model"
	"gorm.io/gorm"
)

type MySQLRefundRepository struct {
	db *gorm.DB
}

func NewMySQLRefundRepository(db *gorm.DB) *MySQLRefundRepository {
	return &MySQLRefundRepository{db: db}
}

func (r *MySQLRefundRepository) CreateRefund(ctx context.Context, refund *model.Refund) error {
	if err := r.db.WithContext(ctx).Create(refund).Error; err != nil {
		return fmt.Errorf("failed to create refund: %w", err)
	}
	return nil
}

func (r *MySQLRefundRepository) GetRefundByID(ctx context.Context, id string) (*model.Refund, error) {
	var refund model.Refund
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&refund).Error; err != nil {
		return nil, fmt.Errorf("failed to get refund: %w", err)
	}
	return &refund, nil
}

func (r *MySQLRefundRepository) UpdateRefundStatus(ctx context.Context, id string, status string) error {
	if err := r.db.WithContext(ctx).Model(&model.Refund{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return fmt.Errorf("failed to update refund status: %w", err)
	}
	return nil
}
