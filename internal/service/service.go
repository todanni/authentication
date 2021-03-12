package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/todanni/alerts"
	"github.com/todanni/authentication/pkg/account"
	"github.com/todanni/email"
)

type service struct {
	repo    account.Repository
	email   email.Service
	alerter alerts.Alerter
	router  *mux.Router
	client  *http.Client
}

func (s *service) ResetPassword(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NewService(repo account.Repository, router *mux.Router, email email.Service, client *http.Client, alerter alerts.Alerter) account.Service {
	s := &service{
		repo:    repo,
		router:  router,
		email:   email,
		client:  client,
		alerter: alerter,
	}
	//s.routes()
	return s
}

func (s *service) routes() {
	// Register endpoint
	s.router.HandleFunc("/api/authentication/", s.Register).Methods(http.MethodPost)

	// Account verification endpoint
	s.router.HandleFunc("/verify/:code", s.Verify).Methods(http.MethodGet)
}
