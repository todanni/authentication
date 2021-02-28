package server

import (
	"github.com/gorilla/mux"
	"github.com/todanni/template-repository/pkg/example"
	"net/http"
)

func NewExampleService(repo example.Repository, router *mux.Router) example.Service {
	server := &exampleService{
		repo: repo,
		router: router,
	}
	server.routes()
	return server
}

type exampleService struct {
	repo example.Repository
	router *mux.Router
}

func (s *exampleService) ExampleMethod(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

