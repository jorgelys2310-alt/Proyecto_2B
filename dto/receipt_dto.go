package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type ReceiptItemRequest struct {
	ProductID uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type CreateReceiptRequest struct {
	UserID uint                 `json:"userId" binding:"required"`
	Items  []ReceiptItemRequest `json:"items" binding:"required,min=1,dive"`
}

type ReceiptItemResponse struct {
	ProductID   uint            `json:"productId"`
	ProductName string          `json:"productName"`
	Quantity    int             `json:"quantity"`
	UnitPrice   decimal.Decimal `json:"unitPrice"`
	Subtotal    decimal.Decimal `json:"subtotal"`
}

type ReceiptResponse struct {
	ReceiptID     uint                  `json:"receiptId"`
	UserID        uint                  `json:"userId"`
	UserEmail     string                `json:"userEmail"`
	Total         decimal.Decimal       `json:"total"`
	AmountOfItems int                   `json:"amountOfItems"`
	CreatedAt     time.Time             `json:"createdAt"`
	Items         []ReceiptItemResponse `json:"items"`
}
