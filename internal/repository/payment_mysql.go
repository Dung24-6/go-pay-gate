package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Dung24-6/go-pay-gate/internal/kafka"
	"github.com/Dung24-6/go-pay-gate/internal/model"
	"gorm.io/gorm"
)

// MySQLPaymentRepository implements PaymentRepository using MySQL
type MySQLPaymentRepository struct {
	db            *gorm.DB
	kafkaProducer *kafka.KafkaProducer
}

// NewMySQLPaymentRepository creates a new instance
func NewMySQLPaymentRepository(db *gorm.DB, kafkaProducer *kafka.KafkaProducer) *MySQLPaymentRepository {
	return &MySQLPaymentRepository{db: db, kafkaProducer: kafkaProducer}
}

// CreatePayment inserts a new payment record
func (r *MySQLPaymentRepository) CreatePayment(ctx context.Context, payment *model.Payment) error {
	if err := r.db.WithContext(ctx).Create(payment).Error; err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}

	// Publish sự kiện Kafka sau khi tạo payment thành công
	err := r.kafkaProducer.ProduceMessage("payment_created", payment.ID)
	if err != nil {
		log.Printf("Failed to publish Kafka message: %v", err)
	}

	return nil
}

// GetPaymentByID retrieves a payment by its ID
func (r *MySQLPaymentRepository) GetPaymentByID(ctx context.Context, id string) (*model.Payment, error) {
	var payment model.Payment
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}
	return &payment, nil
}

// UpdatePaymentStatus updates the status of a payment
func (r *MySQLPaymentRepository) UpdatePaymentStatus(ctx context.Context, id string, status string) error {
	if err := r.db.WithContext(ctx).Model(&model.Payment{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}
	return nil
}
