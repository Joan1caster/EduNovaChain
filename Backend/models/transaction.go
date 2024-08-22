package models

import (
	"time"
)

type Transaction struct {
	ID          uint   `gorm:"primaryKey"`
	OrderID     uint   `gorm:"not null"`
	TxHash      string `gorm:"unique;not null"`
	Amount      string `gorm:"type:varchar(78);not null"` // 使用 string 存储
	GasFeeEther string `gorm:"type:varchar(78);not null"` // 使用 string 存储
	Status      string `gorm:"type:enum('PENDING','COMPLETED','FAILED');not null"`
	CreatedAt   time.Time
	Order       Order `gorm:"foreignKey:OrderID"`
}

type PurchaseInfo struct {
	OrderID            uint    `json:"orderID"`
	NFTContractAddress string  `json:"nftContractAddress"`
	TokenID            string  `json:"tokenID"`
	Price              float64 `json:"price"`
	SellerAddress      string  `json:"sellerAddress"`
}
