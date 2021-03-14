package account

import (
	"gorm.io/gorm"
)

type VerificationRecord struct {
	AccountID uint
	Code      string `gorm:"not null;unique"`
	gorm.Model
}
