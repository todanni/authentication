package account

import (
	"net/http"

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

type Service interface {
	// Login handles the requests coming to the authenticate endpoint
	Login(w http.ResponseWriter, r *http.Request)

	// Register handles the requests coming to the register endpoint
	Register(w http.ResponseWriter, r *http.Request)

	// Verify handles the requests coming to the verify endpoint
	Verify(w http.ResponseWriter, r *http.Request)

	// Login handles the requests coming to the reset password endpoint
	ResetPassword(w http.ResponseWriter, r *http.Request)
}

type Repository interface {
	// InsertAccount
	InsertAccount(account Account) (Account, error)

	// InsertAuthDetails
	InsertAuthDetails(details AuthDetails) (AuthDetails, error)

	// UpdateAuthDetails
	UpdateAuthDetails(details AuthDetails) (AuthDetails, error)

	// GetAuthDetails
	GetAuthDetails(email string) (AuthDetails, error)

	// InsertVerificationRecord
	InsertVerificationRecord(record VerificationRecord) (VerificationRecord, error)

	// GetVerificationRecord
	GetVerificationRecord(accountID int) (VerificationRecord, error)

	// UpdateVerificationRecord
	UpdateVerificationRecord(code string) (VerificationRecord, error)
}
