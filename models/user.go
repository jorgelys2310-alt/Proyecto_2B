package models

type User struct {
	UserID      uint   `gorm:"primaryKey;column:user_id" json:"userId"`
	FirstName   string `gorm:"not null;column:first_name" json:"firstName"`
	LastName    string `gorm:"not null;column:last_name" json:"lastName"`
	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"-"`
	Address     string `json:"address"`
	PhoneNumber string `gorm:"column:phone_number" json:"phoneNumber"`
}

func (User) TableName() string {
	return "users"
}
