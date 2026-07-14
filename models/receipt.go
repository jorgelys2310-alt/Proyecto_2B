package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Receipt struct {
	ReceiptID uint `gorm:"primaryKey;column:receipt_id" json:"receiptId"`

	UserID uint `gorm:"column:created_by;not null;index" json:"userId"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`

	Total         decimal.Decimal `gorm:"type:numeric(10,2);not null" json:"total"`
	CreatedAt     time.Time       `gorm:"column:created_at;not null" json:"createdAt"`
	AmountOfItems int             `gorm:"column:amount_of_items;not null" json:"amountOfItems"`
}

func (Receipt) TableName() string {
	return "receipts"
}
