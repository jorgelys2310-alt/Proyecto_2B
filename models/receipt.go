package models

import "time"

type Receipt struct {
	ReceiptID     uint      `gorm:"primaryKey;column:receipt_id" json:"receiptId"`
	UserID        uint      `gorm:"column:created_by;not null" json:"userId"`
	User          User      `gorm:"foreignKey:UserID;references:UserID" json:"user"`
	Total         float64   `gorm:"type:numeric(10,2);not null" json:"total"`
	CreatedAt     time.Time `gorm:"column:created_at;not null" json:"createdAt"`
	AmountOfItems int       `gorm:"column:amount_of_items;not null" json:"amountOfItems"`
}

func (Receipt) TableName() string {
	return "receipts"
}