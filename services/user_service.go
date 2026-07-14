package services

import (
	"errors"

	"Proyecto_2B/dto"
	"Proyecto_2B/exceptions"
	"Proyecto_2B/models"
	"Proyecto_2B/repository"
	"Proyecto_2B/utils"

	"gorm.io/gorm"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(
	userRepository *repository.UserRepository,
) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) Register(
	request dto.RegisterUserRequest,
) (*dto.UserResponse, error) {

	exists, err := s.UserRepository.ExistsByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, exceptions.ErrEmailAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		Password:    hashedPassword,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}

	if err := s.UserRepository.Create(&user); err != nil {
		return nil, err
	}

	response := s.toResponse(&user)
	return &response, nil
}

func (s *UserService) Login(
	request dto.LoginRequest,
) (*dto.UserResponse, error) {

	user, err := s.UserRepository.FindByEmail(request.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrInvalidCredentials
		}

		return nil, err
	}

	if !utils.CheckPassword(request.Password, user.Password) {
		return nil, exceptions.ErrInvalidCredentials
	}

	response := s.toResponse(user)
	return &response, nil
}

func (s *UserService) FindByID(
	id uint,
) (*dto.UserResponse, error) {

	user, err := s.UserRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, err
	}

	response := s.toResponse(user)
	return &response, nil
}

func (s *UserService) Update(
	id uint,
	request dto.UpdateUserRequest,
) (*dto.UserResponse, error) {

	user, err := s.UserRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, err
	}

	if request.FirstName != nil {
		user.FirstName = *request.FirstName
	}

	if request.LastName != nil {
		user.LastName = *request.LastName
	}

	if request.Email != nil && *request.Email != user.Email {
		exists, err := s.UserRepository.ExistsByEmail(*request.Email)
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, exceptions.ErrEmailAlreadyExists
		}

		user.Email = *request.Email
	}

	if request.Password != nil && *request.Password != "" {
		hashedPassword, err := utils.HashPassword(*request.Password)
		if err != nil {
			return nil, err
		}

		user.Password = hashedPassword
	}

	if request.Address != nil {
		user.Address = *request.Address
	}

	if request.PhoneNumber != nil {
		user.PhoneNumber = *request.PhoneNumber
	}

	if err := s.UserRepository.Update(user); err != nil {
		return nil, err
	}

	response := s.toResponse(user)
	return &response, nil
}

func (s *UserService) Delete(id uint) error {
	user, err := s.UserRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.ErrNotFound
		}

		return err
	}

	return s.UserRepository.Delete(user)
}

func (s *UserService) FindEntityByID(
	id uint,
) (*models.User, error) {

	user, err := s.UserRepository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *UserService) toResponse(
	user *models.User,
) dto.UserResponse {

	return dto.UserResponse{
		UserID:      user.UserID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
	}
}
