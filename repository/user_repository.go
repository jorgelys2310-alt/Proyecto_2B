package repository

import (
	"Proyecto_2B/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User

	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64

	err := r.DB.Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).Error

	return count > 0, err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(user *models.User) error {
	return r.DB.Delete(user).Error
}
