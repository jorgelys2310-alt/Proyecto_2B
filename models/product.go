package models

import "github.com/shopspring/decimal"

type Product struct {
	ProductID   uint            `gorm:"primaryKey;column:product_id" json:"productId"`
	Name        string          `gorm:"not null;column:product_name" json:"name"`
	Price       decimal.Decimal `gorm:"type:numeric(10,2);not null" json:"price"`
	Description string          `json:"description"`
	Amount      int             `gorm:"not null" json:"amount"`
	ImageURL    string          `gorm:"column:image_url" json:"imageUrl"`
}

func (Product) TableName() string {
	return "products"
}
