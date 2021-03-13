package service

import (
	"encoding/json"
	"github.com/todanni/authentication/pkg/account"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strings"
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
	}

	err = bcrypt.CompareHashAndPassword([]byte(authDetails.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(w, "wrong password", http.StatusForbidden)
	}

	// TODO: generate JWT token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *service) validateLoginRequest(r *http.Request) (account.AuthDetails, error){
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return account.AuthDetails{}, err
	}

	var loginRequest account.LoginRequest
	err = json.Unmarshal(reqBody, &loginRequest)
	if err != nil{
		return account.AuthDetails{}, err
	}

	if loginRequest.Email == ""  || loginRequest.Password == ""{
		return account.AuthDetails{}, err
	}

	return account.AuthDetails{
		Email:    strings.ToLower(loginRequest.Email),
		Password:  loginRequest.Password,
	}, err
}
