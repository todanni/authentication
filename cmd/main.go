package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/alerts"
	"github.com/todanni/authentication/internal/config"
	"github.com/todanni/authentication/internal/database"
	"github.com/todanni/authentication/internal/repository"
	authentication "github.com/todanni/authentication/internal/server"
	"github.com/todanni/authentication/pkg/auth"
	"github.com/todanni/email"
)

func main() {
	// Read config
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Open database connection
	db, err := database.Open(cfg)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Perform any migrations needed to run the service
	err = db.AutoMigrate(&auth.AuthenticationDetails{})
	if err != nil {
		log.Error(err)
	}

	// Initialise router
	router := mux.NewRouter()

	// Initialise HTTP client
	c := &http.Client{}

	// Create the service
	authentication.NewAuthService(
		repository.NewAuthRepository(db),
		router,
		email.NewEmailService(cfg.SendGridKey),
		c,
		alerts.NewDiscordAlerter(c, cfg.RegisterWebhook),
	)

	// Start the servers and listen
	log.Fatal(http.ListenAndServe(":8083", router))
}
