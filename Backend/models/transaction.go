package models

import (
	"time"
)

type Transaction struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"not null"`
	TxHash    string    `gorm:"unique;not null"`
	Amount    float64   `gorm:"type:decimal(20,8);not null"`
	GasFee    float64   `gorm:"type:decimal(20,8);not null"`
	Status    string    `gorm:"type:enum('PENDING','COMPLETED','FAILED');not null"`
	CreatedAt time.Time
	Order     Order     `gorm:"foreignKey:OrderID"`
}
