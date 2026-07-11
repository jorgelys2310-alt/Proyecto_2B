package dto

import "github.com/shopspring/decimal"

type ProductRequest struct {
	Name        string          `json:"name" binding:"required"`
	Price       decimal.Decimal `json:"price" binding:"required"`
	Description string          `json:"description"`
	Amount      int             `json:"amount" binding:"gte=0"`
	ImageURL    string          `json:"imageUrl"`
}

type ProductResponse struct {
	ProductID   uint            `json:"productId"`
	Name        string          `json:"name"`
	Price       decimal.Decimal `json:"price"`
	Description string          `json:"description"`
	Amount      int             `json:"amount"`
	ImageURL    string          `json:"imageUrl"`
}
