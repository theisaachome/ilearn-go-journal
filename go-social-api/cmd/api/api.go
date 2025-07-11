package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/theisaachome/social/internal/store"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// Middleware for logging, recovering from panics, etc.
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))


	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler) 
	})
	return r
}

func (app *application) run(mux http.Handler) error{
	server :=&http.Server{
		Addr:app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute * 1,
	}
	fmt.Printf("Starting server on %s\n", app.config.addr)
	return server.ListenAndServe()
}