package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/todanni/alerts"
	"github.com/todanni/authentication/internal/repository"
	"github.com/todanni/authentication/pkg/account"
	"github.com/todanni/email"
	"golang.org/x/crypto/bcrypt"
)

func Test_Login(t *testing.T) {
	// Setup
	pass, err := bcrypt.GenerateFromPassword([]byte("password"), 14)
	assert.NoError(t, err)

	repo := &repository.RepositoryMock{}
	repo.On("GetAuthDetails", "test@mail.com").Return(account.AuthDetails{
		AccountID: 1,
		Email:     "test@mail.com",
		Password:  string(pass),
		Verified:  true,
	}, nil)

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("eyJhbGciOiJSUzI1NiIsImtpZCI6IjU2OWEyNjdhMzNkMGRiZjAyYWVkIiwidHlwIjoiSldUIn0.eyJleHAiOjE2MjExNjg0MDYsImlzcyI6InRvZGFubmktYWNjb3VudC1zZXJ2aWNlIiwidWlkIjoxfQ.uCy1-w6lNMbhufcNmeYaORe18wfEhSspaaXTEheiTm9tmuEFw0KPnVluxJc1JnxlIX_3IwXMAdV002cu0rFp2M4fILjFLGT0d38WccJhMlPmIcxEEvi9mWiGyi48goeErNMs0o9OHicEc4hjYcrNtNa9vYYoQk32SDZ_4l359c2SaPcEgKninr6O3XjhsbZ45r26T4GnarEWo2l6TR47t1ba_f3QtI9q4dBdFZ44MgLoxMxxsoRKplBZ2OYsNYeAYvgCGOYBC7W6dimeuYtDoxFYyM4FMf3XrFNC189rd-slEQAt8zrpREsHsi8DIOUsw9qxdA9KKK6d1XK8Y_p8dQ"))
	}))
	defer ts.Close()
	client := ts.Client()

	_, err = client.Get(ts.URL)
	assert.NoError(t, err)

	err = os.Setenv("TKN_ISSUER_URL", ts.URL)
	assert.NoError(t, err)

	alerter := alerts.MockAlerter{}
	alerter.On("SendLoginAlert", alerts.LoginRequest{Email: "test@mail.com"}).Return(nil)
	emailer := email.MockEmail{}
	s := NewService(repo, mux.NewRouter(), &emailer, client, &alerter)

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

	// Run
	handler.ServeHTTP(rr, req)

	//Verify
	bodyBytes, err := ioutil.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code didn't match")
	log.Print(string(bodyBytes))
}
