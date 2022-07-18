package infrastructure

import (
	"log"

	"github.com/daniel5268/go-meye/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewGormPostgresClient returns a gorm instance for the PostgreSQL DB
func NewGormPostgresClient() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection was not correctly established")
	}

	return db
}
