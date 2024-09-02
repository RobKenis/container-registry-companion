package main

import (
	"net/http"
	"os"
	"time"

	"github.com/robkenis/container-registry-companion/internal/catalog"
	catalog_handler "github.com/robkenis/container-registry-companion/internal/ports/http/catalog_handler"
	"github.com/robkenis/container-registry-companion/internal/ports/http/health_handler"
	"github.com/robkenis/container-registry-companion/internal/ports/http/index_handler"
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

	webDirectory := utils.GetEnv("STATIC_WEB_DIRECTORY", "./web/static")
	log.Debug().Msg("Using web directory: " + webDirectory)

	r := http.NewServeMux()

	r.Handle("GET /health", health_handler.Handler{})

	r.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir(webDirectory))))
	r.Handle("GET /", index_handler.Handler{})

	r.Handle("GET /repositories", catalog_handler.Handler{
		Catalog: catalog.NewCatalog(utils.GetEnv("REGISTRY_URL", "http://localhost:5000")),
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info().Msg("Starting server on port 8080...")
	log.Fatal().Err(srv.ListenAndServe())
}
