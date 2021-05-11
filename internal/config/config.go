package config

import (
	"github.com/caarlos0/env/v6"
)

// Config contains the env variables needed to run the servers
type Config struct {
	DBHost     string `env:"POSTGRES_HOST,required"`
	DBPort     int    `env:"POSTGRES_PORT,required"`
	DBUser     string `env:"POSTGRES_USER,required"`
	DBPassword string `env:"POSTGRES_PASSWORD,required"`
	DBName     string `env:"POSTGRES_NAME,required"`

	LoginWebhook    string `env:"DISCORD_LOGIN_WEBHOOK,required"`
	RegisterWebhook string `env:"DISCORD_REGISTER_WEBHOOK,required"`

	SendGridKey string `env:"SENDGRID_API_KEY,required"`

	RMQUser     string `env:"RMQ_USER,required"`
	RMQPassword string `env:"RMQ_PASSWORD,required"`
	RMQHost     string `env:"RMQ_HOST,required"`
}

func NewFromEnv() (Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
