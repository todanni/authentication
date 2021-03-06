package authentication

import "net/http"

func (s *authService) routes() {
	s.router.HandleFunc("/api/authentication/", s.Register).Methods(http.MethodPost)
}
