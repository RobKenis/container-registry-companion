package main

import (
	"net/http"
	"os"
	"time"

	"github.com/robkenis/container-registry-companion/internal/ports/http/health_handler"
	"github.com/robkenis/container-registry-companion/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.elastic.co/ecszerolog"
)

func main() {
	mode := utils.GetEnv("MODE", "production")
	switch mode {
	case "development":
		log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly}).With().Timestamp().Logger()
	case "production":
		log.Logger = ecszerolog.New(os.Stdout)
	}

	r := http.NewServeMux()

	r.Handle("GET /health", health_handler.Handler{})

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info().Msg("Starting server on port 8080...")
	log.Fatal().Err(srv.ListenAndServe())
}
