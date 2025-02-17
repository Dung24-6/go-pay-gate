package model

import (
	"time"
)

// Refund represents a refund transaction
type Refund struct {
	ID            string      `gorm:"primaryKey"`
	PaymentID     string      `gorm:"index;not null"`
	Payment       Payment     `gorm:"foreignKey:PaymentID;constraint:OnDelete:CASCADE"`
	TransactionID string      `gorm:"index;not null"` // Liên kết với Transaction
	Transaction   Transaction `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE"`
	Amount        float64     `gorm:"not null"`
	Currency      string      `gorm:"not null"`
	Status        string      `gorm:"not null"` // pending, success, failed
	Reason        string      `gorm:"not null"`
	ProcessedAt   time.Time   `gorm:"autoCreateTime"`
	CreatedAt     time.Time   `gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime"`
}
