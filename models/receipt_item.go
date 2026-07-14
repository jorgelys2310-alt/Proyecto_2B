package models

import "github.com/shopspring/decimal"

type ReceiptItem struct {
	ReceiptItemID uint `gorm:"primaryKey;column:receipt_item_id" json:"receiptItemId"`

	ReceiptID uint    `gorm:"column:receipt_id;not null;index" json:"receiptId"`
	Receipt   Receipt `gorm:"foreignKey:ReceiptID;references:ReceiptID" json:"-"`

	ProductID uint    `gorm:"column:product_id;not null;index" json:"productId"`
	Product   Product `gorm:"foreignKey:ProductID;references:ProductID" json:"product"`

	Quantity  int             `gorm:"not null" json:"quantity"`
	UnitPrice decimal.Decimal `gorm:"type:numeric(10,2);column:unit_price;not null" json:"unitPrice"`
	Subtotal  decimal.Decimal `gorm:"type:numeric(10,2);not null" json:"subtotal"`
}

func (ReceiptItem) TableName() string {
	return "receipt_items"
}
