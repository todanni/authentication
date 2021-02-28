package server

import "net/http"

func (s *exampleService) routes() {
	// GET an example
	s.router.HandleFunc("/api/example", s.ExampleMethod).Methods(http.MethodGet)
}