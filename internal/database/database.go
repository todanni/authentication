package database

import (
	"fmt"

	"github.com/todanni/authentication/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(cfg config.Config) (*gorm.DB, error) {
	// Make connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBConf.DBHost, cfg.DBConf.DBPort, cfg.DBConf.DBUser, cfg.DBConf.DBPassword, cfg.DBConf.DBName)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return db, err
}
