package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/todanni/authentication/pkg/account"
)

func (s *service) Register(w http.ResponseWriter, r *http.Request) {
	acc, err := s.validateRegisterRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	// Create account record in the DB
	createdAcc, err := s.repo.InsertAccount(acc)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// Create verification record in the DB
	_, err = s.repo.InsertVerificationRecord(account.VerificationRecord{
		AccountID: createdAcc.ID,
		Code:      s.generateCode(),
	})

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *service) validateRegisterRequest(r *http.Request) (account.Account, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return account.Account{}, err
	}

	if reqBody == nil {
		return account.Account{}, errors.New("body was empty")
	}

	var registerRequest account.Request
	err = json.Unmarshal(reqBody, &registerRequest)
	if err != nil {
		return account.Account{}, err
	}

	if registerRequest.Password == "" {
		return account.Account{}, errors.New("password was empty")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 14)
	if err != nil {
		return account.Account{}, err
	}

	return account.Account{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		AuthDetails: account.AuthDetails{
			Email:    registerRequest.Email,
			Password: string(pass),
		},
	}, err
}

func (s service) generateCode() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
