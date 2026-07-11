package dto

type RegisterUserRequest struct {
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName   *string `json:"firstName"`
	LastName    *string `json:"lastName"`
	Email       *string `json:"email" binding:"omitempty,email"`
	Password    *string `json:"password" binding:"omitempty,min=6"`
	Address     *string `json:"address"`
	PhoneNumber *string `json:"phoneNumber"`
}

type UserResponse struct {
	UserID      uint   `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}
