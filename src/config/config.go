package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env int8

const (
	Test env = iota
	Prod
	Dev
)

var (
	JwtIssuer   string
	JwtSecret   string
	DatabaseURL string
	Port        string
)

func LoadConfig(e env) {
	var envFile string
	switch e {
	case Test:
		envFile = "../../test.env"
	case Dev:
		envFile = "dev.env"
	case Prod:
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}
	JwtIssuer = os.Getenv("JWT_ISSUER")
	JwtSecret = os.Getenv("JWT_SECRET")
	DatabaseURL = os.Getenv("DATABASE_URL")
	Port = os.Getenv("PORT")
}
