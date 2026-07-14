package services

import (
	"errors"
	"fmt"
	"time"

	"Proyecto_2B/dto"
	"Proyecto_2B/exceptions"
	"Proyecto_2B/models"
	"Proyecto_2B/repository"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ReceiptService struct {
	DB                *gorm.DB
	ReceiptRepository *repository.ReceiptRepository
	UserRepository    *repository.UserRepository
	ProductRepository *repository.ProductRepository
}

func NewReceiptService(
	db *gorm.DB,
	receiptRepository *repository.ReceiptRepository,
	userRepository *repository.UserRepository,
	productRepository *repository.ProductRepository,
) *ReceiptService {
	return &ReceiptService{
		DB:                db,
		ReceiptRepository: receiptRepository,
		UserRepository:    userRepository,
		ProductRepository: productRepository,
	}
}

func (s *ReceiptService) Create(
	request dto.CreateReceiptRequest,
) (*dto.ReceiptResponse, error) {

	user, err := s.UserRepository.FindByID(request.UserID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf(
				"%w: usuario con id %d",
				exceptions.ErrNotFound,
				request.UserID,
			)
		}

		return nil, err
	}

	var createdReceipt models.Receipt

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		total := decimal.Zero
		amountOfItems := 0

		receipt := models.Receipt{
			UserID:        user.UserID,
			Total:         decimal.Zero,
			CreatedAt:     time.Now(),
			AmountOfItems: 0,
		}

		if err := tx.Create(&receipt).Error; err != nil {
			return err
		}

		for _, itemRequest := range request.Items {
			var product models.Product

			err := tx.First(
				&product,
				itemRequest.ProductID,
			).Error

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf(
						"%w: producto con id %d",
						exceptions.ErrNotFound,
						itemRequest.ProductID,
					)
				}

				return err
			}

			if product.Amount < itemRequest.Quantity {
				return fmt.Errorf(
					"%w para el producto: %s",
					exceptions.ErrInsufficientStock,
					product.Name,
				)
			}

			unitPrice := product.Price

			subtotal := unitPrice.Mul(
				decimal.NewFromInt(
					int64(itemRequest.Quantity),
				),
			)

			item := models.ReceiptItem{
				ReceiptID: receipt.ReceiptID,
				ProductID: product.ProductID,
				Quantity:  itemRequest.Quantity,
				UnitPrice: unitPrice,
				Subtotal:  subtotal,
			}

			if err := tx.Create(&item).Error; err != nil {
				return err
			}

			product.Amount -= itemRequest.Quantity

			if err := tx.Save(&product).Error; err != nil {
				return err
			}

			total = total.Add(subtotal)
			amountOfItems += itemRequest.Quantity
		}

		receipt.Total = total
		receipt.AmountOfItems = amountOfItems

		if err := tx.Save(&receipt).Error; err != nil {
			return err
		}

		createdReceipt = receipt
		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.buildResponse(&createdReceipt)
}

func (s *ReceiptService) FindByID(
	id uint,
) (*dto.ReceiptResponse, error) {

	receipt, err := s.ReceiptRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, err
	}

	return s.buildResponse(receipt)
}

func (s *ReceiptService) FindAll() ([]dto.ReceiptResponse, error) {

	receipts, err := s.ReceiptRepository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ReceiptResponse, 0, len(receipts))

	for i := range receipts {
		response, err := s.buildResponse(&receipts[i])
		if err != nil {
			return nil, err
		}

		responses = append(responses, *response)
	}

	return responses, nil
}

func (s *ReceiptService) FindByUserID(
	userID uint,
) ([]dto.ReceiptResponse, error) {

	_, err := s.UserRepository.FindByID(userID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, err
	}

	receipts, err :=
		s.ReceiptRepository.FindByUserID(userID)

	if err != nil {
		return nil, err
	}

	responses := make([]dto.ReceiptResponse, 0, len(receipts))

	for i := range receipts {
		response, err := s.buildResponse(&receipts[i])
		if err != nil {
			return nil, err
		}

		responses = append(responses, *response)
	}

	return responses, nil
}

func (s *ReceiptService) Delete(id uint) error {
	receipt, err := s.ReceiptRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.ErrNotFound
		}

		return err
	}

	return s.ReceiptRepository.Delete(receipt)
}

func (s *ReceiptService) buildResponse(
	receipt *models.Receipt,
) (*dto.ReceiptResponse, error) {

	items, err :=
		s.ReceiptRepository.FindItemsByReceiptID(
			receipt.ReceiptID,
		)

	if err != nil {
		return nil, err
	}

	itemResponses := make(
		[]dto.ReceiptItemResponse,
		0,
		len(items),
	)

	for _, item := range items {
		itemResponses = append(
			itemResponses,
			dto.ReceiptItemResponse{
				ProductID:   item.ProductID,
				ProductName: item.Product.Name,
				Quantity:    item.Quantity,
				UnitPrice:   item.UnitPrice,
				Subtotal:    item.Subtotal,
			},
		)
	}

	response := dto.ReceiptResponse{
		ReceiptID:     receipt.ReceiptID,
		UserID:        receipt.UserID,
		UserEmail:     receipt.User.Email,
		Total:         receipt.Total,
		AmountOfItems: receipt.AmountOfItems,
		CreatedAt:     receipt.CreatedAt,
		Items:         itemResponses,
	}

	return &response, nil
}
