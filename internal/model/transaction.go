package model

import (
	"time"
)

// Transaction represents a payment transaction record
type Transaction struct {
	ID            string    `gorm:"primaryKey"`
	PaymentID     string    `gorm:"index;not null"` // Liên kết với Payment
	Payment       Payment   `gorm:"foreignKey:PaymentID;constraint:OnDelete:CASCADE"`
	Amount        float64   `gorm:"not null"`
	Currency      string    `gorm:"not null"`
	Status        string    `gorm:"not null"` // pending, success, failed
	TransactionID string    `gorm:"uniqueIndex;not null"`
	Provider      string    `gorm:"not null"` // VNPay, Momo, ZaloPay
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
