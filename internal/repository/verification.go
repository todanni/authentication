package repository

import (
	"github.com/todanni/authentication/pkg/account"
	"gorm.io/gorm"
)

type verificationRepository struct {
	db *gorm.DB
}

func (r verificationRepository) Insert(record account.VerificationRecord) (account.VerificationRecord, error) {
	err := r.db.Create(&record).Error
	return record, err
}

func (r verificationRepository) Update(record account.VerificationRecord) (account.VerificationRecord, error) {
	panic("implement me")
}

func (r verificationRepository) Select(record account.VerificationRecord) (account.VerificationRecord, error) {
	panic("")
}

func NewVerificationRepository(db *gorm.DB) account.Repository {
	return &verificationRepository{
		db: db,
	}
}
