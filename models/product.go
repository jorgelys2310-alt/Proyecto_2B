package models

type Product struct {
	ProductID   uint    `gorm:"primaryKey;column:product_id" json:"productId"`
	Name        string  `gorm:"not null;column:product_name" json:"name"`
	Price       float64 `gorm:"type:numeric(10,2);not null" json:"price"`
	Description string  `json:"description"`
	Amount      int     `gorm:"not null" json:"amount"`
	ImageURL    string  `gorm:"column:image_url" json:"imageUrl"`
}

func(Product) TableName() string {
	return "products"
}