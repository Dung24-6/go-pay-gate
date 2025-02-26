package service

import (
	"context"
	"fmt"
	"log"

	"github.com/Dung24-6/go-pay-gate/internal/model"
	"github.com/Dung24-6/go-pay-gate/internal/repository"
)

type PaymentService struct {
	repo      repository.PaymentRepository
	s3Service *S3Service
}

func NewPaymentService(repo repository.PaymentRepository, s3Service *S3Service) *PaymentService {
	return &PaymentService{
		repo:      repo,
		s3Service: s3Service,
	}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, payment *model.Payment) error {
	if err := s.repo.CreatePayment(ctx, payment); err != nil {
		return fmt.Errorf("failed to save payment: %w", err)
	}

	log.Printf("Payment created: %+v", payment)
	return nil
}

func (s *PaymentService) GetPayment(ctx context.Context, id string) (*model.Payment, error) {
	payment, err := s.repo.GetPaymentByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch payment: %w", err)
	}
	return payment, nil
}

func (s *PaymentService) UpdatePaymentStatus(ctx context.Context, id string, status string) error {
	if err := s.repo.UpdatePaymentStatus(ctx, id, status); err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	payment, err := s.repo.GetPaymentByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to fetch payment after update: %w", err)
	}

	err = s.s3Service.UploadReceipt(id, payment)
	if err != nil {
		log.Printf("Failed to upload receipt to S3: %v", err)
	}

	log.Printf("Payment status updated: %s -> %s", id, status)
	return nil
}
