package repository

import (
	"github.com/todanni/authentication/pkg/auth"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func (r authRepository) Insert(details auth.AuthenticationDetails) (auth.AuthenticationDetails, error) {
	err := r.db.Create(&details).Error
	return details, err
}

func (r authRepository) Update(details auth.AuthenticationDetails) (auth.AuthenticationDetails, error) {
	// TODO: I left myself a note that this doesn't work, so test
	err := r.db.Updates(&details).Error
	return details, err
}

func (r authRepository) Get(userID int) (auth.AuthenticationDetails, error) {
	var details auth.AuthenticationDetails
	err := r.db.First(&details, userID).Error
	return details, err
}

func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepository{
		db: db,
	}
}
