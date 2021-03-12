package repository

import (
	"fmt"

	"github.com/todanni/authentication/pkg/account"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) account.Repository {
	r := &repo{
		db: db,
	}
	r.runMigrations()
	return r
}

func (r repo) runMigrations() {
	err := r.db.AutoMigrate(&account.Account{}, &account.AuthDetails{}, &account.VerificationRecord{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
