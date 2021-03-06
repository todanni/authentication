package authentication

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/alerts"
	"github.com/todanni/authentication/pkg/auth"
	"github.com/todanni/email"
)

func NewAuthService(repo auth.Repository, router *mux.Router, email email.Service, client *http.Client, alerter alerts.Alerter) auth.Service {
	server := &authService{
		repo:    repo,
		router:  router,
		email:   email,
		client:  client,
		alerter: alerter,
	}
	server.routes()
	return server
}

type authService struct {
	repo    auth.Repository
	email   email.Service
	alerter alerts.Alerter
	router  *mux.Router
	client  *http.Client
}

func (s *authService) Register(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return
	}

	var authDetails auth.AuthenticationDetails
	err = json.Unmarshal(reqBody, &authDetails)
	if err != nil {
		log.Error("couldn't unmarshal", err)
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
		return
	}

	req := alerts.RegisterRequest{Email: authDetails.Email}
	err = s.alerter.SendRegisterAlert(req)
	if err != nil {
		log.Error(err)
	}

	// Save details in DB
	returnedAuthDetails, err := s.repo.Insert(authDetails)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send verification email
	//err = s.email.SendVerificationEmail("someMadeupCode", email.Recipient{Email: authDetails.Email})
	//if err != nil {
	//	log.Error(err)
	//}

	marshalled, err := json.Marshal(returnedAuthDetails)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(marshalled)
	if err != nil {
		log.Error(err)
	}

}

func (s *authService) Login(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *authService) ResetPassword(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
