package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (s service) Verify(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if code == "" {
		http.Error(w, "missing verification code", http.StatusBadRequest)
	}

	// if the code exists, check the date and validate/invalidate

	// at the end, we should return a valid token
}

func (s service) generate(accountID int) (string, error) {
	panic("")
	//
	//record, err := s.repo.InsertAuthenticationDetails(account.VerificationRecord{
	//	AccountID: uint(accountID),
	//	Code:      s.generateCode(),
	//})
	//if err != nil {
	//	return "", err
	//}
	//
	//return record.Code, err
}
