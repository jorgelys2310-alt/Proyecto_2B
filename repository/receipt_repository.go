package repository

import (
	"Proyecto_2B/models"

	"gorm.io/gorm"
)

type ReceiptRepository struct {
	DB *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) *ReceiptRepository {
	return &ReceiptRepository{
		DB: db,
	}
}

func (r *ReceiptRepository) Create(receipt *models.Receipt) error {
	return r.DB.Create(receipt).Error
}

func (r *ReceiptRepository) CreateItem(item *models.ReceiptItem) error {
	return r.DB.Create(item).Error
}

func (r *ReceiptRepository) FindByID(id uint) (*models.Receipt, error) {
	var receipt models.Receipt

	err := r.DB.
		Preload("User").
		First(&receipt, id).Error

	if err != nil {
		return nil, err
	}

	return &receipt, nil
}

func (r *ReceiptRepository) FindAll() ([]models.Receipt, error) {
	var receipts []models.Receipt

	err := r.DB.
		Preload("User").
		Order("receipt_id ASC").
		Find(&receipts).Error

	return receipts, err
}

func (r *ReceiptRepository) FindByUserID(userID uint) ([]models.Receipt, error) {
	var receipts []models.Receipt

	err := r.DB.
		Preload("User").
		Where("created_by = ?", userID).
		Order("receipt_id ASC").
		Find(&receipts).Error

	return receipts, err
}

func (r *ReceiptRepository) FindItemsByReceiptID(receiptID uint) ([]models.ReceiptItem, error) {
	var items []models.ReceiptItem

	err := r.DB.
		Preload("Product").
		Where("receipt_id = ?", receiptID).
		Find(&items).Error

	return items, err
}

func (r *ReceiptRepository) Delete(receipt *models.Receipt) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("receipt_id = ?", receipt.ReceiptID).
			Delete(&models.ReceiptItem{}).Error; err != nil {
			return err
		}

		return tx.Delete(receipt).Error
	})
}

func (r *ReceiptRepository) Transaction(
	callback func(tx *gorm.DB) error,
) error {
	return r.DB.Transaction(callback)
}
