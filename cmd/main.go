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
	"github.com/todanni/authentication/internal/service"
	"github.com/todanni/authentication/pkg/account"
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
	err = db.AutoMigrate(&account.AuthDetails{}, &account.VerificationRecord{})
	if err != nil {
		log.Error(err)
	}

	// Initialise router
	router := mux.NewRouter()

	// Initialise HTTP client
	c := &http.Client{}

	// Create services
	service.NewService(
		repository.NewRepository(db),
		router,
		email.NewEmailService(cfg.SendGridKey),
		c,
		alerts.NewDiscordAlerter(c, cfg.RegisterWebhook),
	)

	// Start the servers and listen
	log.Fatal(http.ListenAndServe(":8083", router))
}
