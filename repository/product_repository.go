package repository

import (
	"Proyecto_2B/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product

	err := r.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product

	err := r.DB.Order("product_id ASC").Find(&products).Error
	return products, err
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) Delete(product *models.Product) error {
	return r.DB.Delete(product).Error
}
