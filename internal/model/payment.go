package model

import (
	"time"
)

// Payment represents a payment transaction
type Payment struct {
	ID            string        `gorm:"primaryKey"`
	OrderID       string        `gorm:"index"`
	Amount        float64       `gorm:"not null"`
	Currency      string        `gorm:"not null"`
	Status        string        `gorm:"not null"`
	PaymentMethod string        `gorm:"not null"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime"`
	Transactions  []Transaction `gorm:"foreignKey:PaymentID;constraint:OnDelete:CASCADE"`
	Refunds       []Refund      `gorm:"foreignKey:PaymentID;constraint:OnDelete:CASCADE"`
}
