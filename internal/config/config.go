package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Env         string
	DatabaseUrl string
}

func MustLoad() Config {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to load .env: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is required.")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required.")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DATABASE_URL is required.")
	}

	return Config{
		Port:        port,
		Env:         env,
		DatabaseUrl: dbUrl,
	}
}
