package models

import (
	"time"
)

type UserBalance struct {
	UserID    uint    `gorm:"primaryKey"`
	Balance   float64 `gorm:"type:decimal(20,8);not null;default:0"`
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
}
