package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/todanni/authentication/pkg/account"

	log "github.com/sirupsen/logrus"
	"github.com/todanni/alerts"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return
	}

	var registerRequest account.Request
	err = json.Unmarshal(reqBody, &registerRequest)
	if err != nil {
		log.Error("couldn't unmarshal", err)
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 14)
	if err != nil {
		log.Error(err)
		return
	}

	// Save details in DB
	returnedAuthDetails, err := s.repo.InsertAuthDetails(account.AuthDetails{
		Email:    registerRequest.Email,
		Password: string(pass),
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req := alerts.RegisterRequest{Email: registerRequest.Email}
	err = s.alerter.SendRegisterAlert(req)
	if err != nil {
		log.Error(err)
	}

	// TODO: send create request to verification service
	// needs AccountID

	marshalled, err := json.Marshal(returnedAuthDetails)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(marshalled)
	if err != nil {
		log.Error(err)
	}
}
