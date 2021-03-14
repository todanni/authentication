package account

import (
	"gorm.io/gorm"
)

type Account struct {
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	ProfilePicture string      `json:"profilePicture"`
	JobTitle       string      `json:"jobTitle"`
	AuthDetails    AuthDetails `json:"-" gorm:"foreignKey:account_id"`
	gorm.Model
}
