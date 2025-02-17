package repository

import (
	"context"

	"github.com/Dung24-6/go-pay-gate/internal/model"
	"gorm.io/gorm"
)

// PaymentRepository defines methods to interact with payment data.
type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment *model.Payment) error
	GetPaymentByID(ctx context.Context, id string) (*model.Payment, error)
	UpdatePaymentStatus(ctx context.Context, id string, status string) error
}

// paymentRepository implements PaymentRepository
type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository creates a new instance of PaymentRepository
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// CreatePayment inserts a new payment record into the database
func (r *paymentRepository) CreatePayment(ctx context.Context, payment *model.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

// GetPaymentByID retrieves a payment record by ID
func (r *paymentRepository) GetPaymentByID(ctx context.Context, id string) (*model.Payment, error) {
	var payment model.Payment
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// UpdatePaymentStatus updates the status of a payment record
func (r *paymentRepository) UpdatePaymentStatus(ctx context.Context, id string, status string) error {
	return r.db.WithContext(ctx).Model(&model.Payment{}).Where("id = ?", id).Update("status", status).Error
}
