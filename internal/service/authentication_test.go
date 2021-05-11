package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/todanni/alerts"
	"github.com/todanni/authentication/internal/repository"
	"github.com/todanni/authentication/pkg/account"
	"github.com/todanni/email"
	"golang.org/x/crypto/bcrypt"
)

func Test_service_Login(t *testing.T) {
	repo := &repository.RepositoryMock{}

	pass, err := bcrypt.GenerateFromPassword([]byte("password"), 14)
	assert.NoError(t, err)

	repo.On("GetAuthDetails", "test@mail.com").Return(account.AuthDetails{
		AccountID: 1,
		Email:     "test@mail.com",
		Password:  string(pass),
		Verified:  true,
	}, nil)

	alerter := alerts.MockAlerter{}
	alerter.On("SendLoginAlert", alerts.LoginRequest{Email: "test@mail.com"}).Return(nil)
	emailer := email.MockEmail{}
	s := NewService(repo, mux.NewRouter(), &emailer, &http.Client{}, &alerter)

	// Create request
	lr := account.LoginRequest{
		Email:    "test@mail.com",
		Password: "password",
	}
	postBody, err := json.Marshal(lr)
	assert.NoError(t, err)

	b := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", "/api/account/auth", b)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.Login)
	handler.ServeHTTP(rr, req)

	bodyBytes, err := ioutil.ReadAll(rr.Body)
	assert.NoError(t, err)
	log.Print(string(bodyBytes))

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code didn't match")
}
