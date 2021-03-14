package service

import (
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/todanni/alerts"

	"github.com/todanni/authentication/pkg/account"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := s.validateLoginRequest(r)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	authDetails, err := s.repo.GetAuthDetails(loginRequest.Email)
	if err != nil {
		http.Error(w, "invalid request", http.StatusNotFound)
		return
	}

	if !authDetails.Verified {
		http.Error(w, "please verify your email first", http.StatusForbidden)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(authDetails.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(w, "wrong password", http.StatusForbidden)
		return
	}

	// Send login alert
	err = s.alerter.SendLoginAlert(alerts.LoginRequest{Email: authDetails.Email})
	if err != nil {
		log.Error(err)
	}

	// TODO: generate JWT token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *service) validateLoginRequest(r *http.Request) (account.AuthDetails, error) {
	var loginRequest account.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		return account.AuthDetails{}, err
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		return account.AuthDetails{}, err
	}

	return account.AuthDetails{
		Email:    strings.ToLower(loginRequest.Email),
		Password: loginRequest.Password,
	}, err
}
