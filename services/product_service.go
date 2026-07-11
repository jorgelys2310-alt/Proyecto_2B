package services

import (
	"errors"

	"Proyecto_2B/dto"
	"Proyecto_2B/exceptions"
	"Proyecto_2B/models"
	"Proyecto_2B/repository"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

func NewProductService(
	productRepository *repository.ProductRepository,
) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (s *ProductService) Create(
	request dto.ProductRequest,
) (*dto.ProductResponse, error) {

	if request.Price.LessThanOrEqual(decimal.Zero) {
		return nil, exceptions.ErrBadRequest
	}

	product := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
		Amount:      request.Amount,
		ImageURL:    request.ImageURL,
	}

	if err := s.ProductRepository.Create(&product); err != nil {
		return nil, err
	}

	response := s.toResponse(&product)
	return &response, nil
}

func (s *ProductService) FindAll(
) ([]dto.ProductResponse, error) {

	products, err := s.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.ProductResponse, 0, len(products))

	for _, product := range products {
		responses = append(
			responses,
			s.toResponse(&product),
		)
	}

	return responses, nil
}

func (s *ProductService) FindByID(
	id uint,
) (*dto.ProductResponse, error) {

	product, err := s.ProductRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, err
	}

	response := s.toResponse(product)
	return &response, nil
}

func (s *ProductService) Update(
	id uint,
	request dto.ProductRequest,
) (*dto.ProductResponse, error) {

	product, err := s.ProductRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, err
	}

	if request.Price.LessThanOrEqual(decimal.Zero) {
		return nil, exceptions.ErrBadRequest
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Description = request.Description
	product.Amount = request.Amount
	product.ImageURL = request.ImageURL

	if err := s.ProductRepository.Update(product); err != nil {
		return nil, err
	}

	response := s.toResponse(product)
	return &response, nil
}

func (s *ProductService) Delete(id uint) error {
	product, err := s.ProductRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.ErrNotFound
		}

		return err
	}

	return s.ProductRepository.Delete(product)
}

func (s *ProductService) FindEntityByID(
	id uint,
) (*models.Product, error) {

	product, err := s.ProductRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, err
	}

	return product, nil
}

func (s *ProductService) toResponse(
	product *models.Product,
) dto.ProductResponse {

	return dto.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Amount:      product.Amount,
		ImageURL:    product.ImageURL,
	}
}