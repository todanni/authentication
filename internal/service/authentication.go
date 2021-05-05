package service

import (
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/todanni/alerts"
	"github.com/todanni/authentication/pkg/account"
	"github.com/todanni/token"
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

	jwt, err := token.Generate(int(authDetails.AccountID), *s.client)
	log.Info(string(jwt))

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Error(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    string(jwt),
		Secure:   true,
		SameSite: 3,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jwt)
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
