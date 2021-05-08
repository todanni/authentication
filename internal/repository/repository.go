package repository

import (
	log "github.com/sirupsen/logrus"

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

func (r *repo) runMigrations() {
	err := r.db.AutoMigrate(&account.VerificationRecord{})
	if err != nil {
		log.Error(err)
	}

	err = r.db.AutoMigrate(&account.AuthDetails{})
	if err != nil {
		log.Error(err)
	}

	err = r.db.AutoMigrate(&account.Account{})
	if err != nil {
		log.Error(err)
	}
}
