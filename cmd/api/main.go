package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/md-asharaf/olx-go/internal/config"
	"github.com/md-asharaf/olx-go/internal/db"
	"github.com/md-asharaf/olx-go/internal/handlers"
)

func main(){
  cfg := config.MustLoad()

  _, err := db.Connect(cfg.DatabaseUrl)
  if err != nil {
    log.Fatalf("db.Connect: %v", err)
  }

  fmt.Println("Database Connected.")
  fmt.Println("Starting Olx Server...")
  mux := http.NewServeMux()

  mux.HandleFunc("GET /healthz",handlers.Health)

  srv := http.Server{
    Addr: ":" + cfg.Port,
    Handler: mux,
    ReadTimeout: time.Second * 10,
    WriteTimeout: time.Second * 30,
    IdleTimeout: time.Second * 60,
  }

  log.Printf("Server is running on port %v", srv.Addr)
  if err := srv.ListenAndServe(); err != nil {
    log.Fatalf("Server failed: %v", err)
  }
}