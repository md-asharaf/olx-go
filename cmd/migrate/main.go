package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/md-asharaf/olx-go/internal/config"
)

func main() {
	args := os.Args
	fmt.Println(args)

	if len(args) < 2 {
		log.Fatal("usage: migrate <up | down>")
	}

	cfg := config.MustLoad()

	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl,
	)
	if err != nil {
		log.Fatalf("migrate.New: %v", err)
	}

	switch args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatalf("m.Up: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatalf("m.Down: %v", err)
		}
	default:
		log.Fatalf("unknown argument: %s", args[1])
	}

	fmt.Println("Running migrations...")
}
