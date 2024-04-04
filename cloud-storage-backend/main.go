package main

import (
	"cloud-storage-backend/router"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"
)

var (
	ShutdownTimeout = 60
	Port            = ":5000"
)

func main() {
	r := router.NewGinRouter()

	srv := &http.Server{
		Addr:    Port,
		Handler: r,
	}

	startServer(srv)
}

func startServer(srv *http.Server) {
	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		log.Logger.Info().Err(err).Msg("Server shutdown or closed")
	} else if err != nil {
		log.Logger.Err(err).Msg("Unexpected shutdown or close")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Logger.Fatal().Msgf("%s: %s", msg, err)
	}
}
