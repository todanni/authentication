package repository

import (
	"github.com/todanni/authentication/pkg/account"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) account.Repository {
	return &repo{
		db: db,
	}
}
