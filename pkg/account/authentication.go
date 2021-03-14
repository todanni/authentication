package account

import (
	"time"

	"gorm.io/gorm"
)

type AuthDetails struct {
	AccountID          uint               `json:"accountID" gorm:"primaryKey;autoIncrement;unique;not null"`
	Email              string             `json:"email" gorm:"not null,unique"`
	Password           string             `json:"-" gorm:"not null"`
	Verified           bool               `json:"verified" gorm:"default:false"`
	CreatedAt          time.Time          `json:"-" gorm:"autoCreateTime"`
	UpdatedAt          time.Time          `json:"-" gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt     `json:"-" gorm:"index"`
	VerificationRecord VerificationRecord `json:"-" gorm:"foreignKey:account_id"`
}
