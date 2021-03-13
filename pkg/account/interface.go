package account

import "net/http"

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
	// InsertAccount - used
	InsertAccount(account Account) (Account, error)

	// GetAuthDetails - used
	GetAuthDetails(email string) (AuthDetails, error)

	// InsertVerificationRecord - used
	InsertVerificationRecord(record VerificationRecord) (VerificationRecord, error)

	// UpdateVerificationRecord - used
	GetVerificationRecordByCode(code string) (VerificationRecord, error)

	// SetAuthDetailsValid - sets the auth details record to verified - used
	SetAuthDetailsValid(accountID uint) error
}
