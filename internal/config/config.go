package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
  Port string
  Env string
  DatabaseUrl string
}

func MustLoad() Config {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Failed to load environment variables.")
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

  return Config {
    Port: port,
    Env: env,
    DatabaseUrl: dbUrl,
  }
}