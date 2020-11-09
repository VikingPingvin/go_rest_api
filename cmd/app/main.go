package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"vikingpingvin/restpractice/article"
	"vikingpingvin/restpractice/router"

	"github.com/rs/zerolog/log"
)

func main() {

	// READ CONF
	log.Info().Msg("Beginning configuration parsing")
	// INITIALIZE LOGGER

	// INITIALIZE ROUTER
	router := router.InitRouter()
	log.Info().Msg("Routers initialized.")

	// INIT DATABASE
	articlesDb := article.InitializeArticles() // Dummy db
	log.Info().Msg(fmt.Sprintf("Database initialized with size #%d", len(*articlesDb)))

	// START SERVER
	srvPort := 8080
	log.Info().Msgf("Starting server on PORT %d", srvPort)
	address := fmt.Sprintf(":%d", 8080)

	s := http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().
				Err(errors.New("Server startup error")).
				Msg("Error bringing up server!")
		}
	}()

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt)
	<-osSignal

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	log.Info().Msg("Shutting down server...")
	s.Shutdown(ctx)
	os.Exit(0)
}
