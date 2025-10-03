package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ekasc/nucleus-api/internal/config"
	"github.com/ekasc/nucleus-api/internal/logger"
	"github.com/ekasc/nucleus-api/internal/router"
)

func main() {
	cfg := config.MustLoad()
	logg := logger.New(cfg)
	r := router.New(cfg, logg)
	addr := fmt.Sprintf(":%d", cfg.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logg.Info().Str("addr", addr).Msg("server listening")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logg.Fatal().Err(err).Msg("server failed")
		}

	}()

	<-done
	logg.Info().Msg("server shutting down")

	if err := server.Close(); err != nil {
		logg.Error().Err(err).Msg("error shutting down server")
	}

}
