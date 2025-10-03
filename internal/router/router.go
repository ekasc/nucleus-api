package router

import (
	"net/http"
	"time"

	"github.com/ekasc/nucleus-api/internal/config"
	"github.com/ekasc/nucleus-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
)

func New(c config.Config, log *zerolog.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   c.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/healthz", handlers.Healthz)
	r.Get("/metrics", handlers.Metrics)

	return r
}
