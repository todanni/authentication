package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/todanni/alerts"
	"github.com/todanni/authentication/pkg/account"
	"github.com/todanni/email"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(w http.ResponseWriter, r *http.Request) {
	acc, err := s.validateRegisterRequest(r)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Create account record in the DB
	createdAcc, err := s.repo.InsertAccount(acc)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send Discord alert
	err = s.alerter.SendRegisterAlert(alerts.RegisterRequest{
		FullName: createdAcc.FirstName + " " + createdAcc.LastName,
		Email:    createdAcc.AuthDetails.Email,
	})
	if err != nil {
		log.Error(err)
	}

	// Create verification record in the DB
	rec, err := s.repo.InsertVerificationRecord(account.VerificationRecord{
		AccountID: createdAcc.ID,
		Code:      s.generateCode(),
	})
	if err != nil {
		log.Error("failed to create verification record", err)
		return
	}

	// Send verification email
	err = s.email.SendVerificationEmail(rec.Code, email.Recipient{
		Email:    createdAcc.AuthDetails.Email,
		FullName: createdAcc.FirstName + " " + createdAcc.LastName,
	})
	if err != nil {
		log.Error(err)
	}

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

	var registerRequest account.RegisterRequest
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
			Email:    strings.ToLower(registerRequest.Email),
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
