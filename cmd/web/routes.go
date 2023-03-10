package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shafaq-here/bookings/pkg/config"
	"github.com/shafaq-here/bookings/pkg/handlers"
)

func routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.HomeHandler))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutHandler))
	// mux.Get("/sections", http.HandlerFunc(handlers.Repo.SectionsHandler))
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(LoadSession)
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/sections", handlers.Repo.SectionsHandler)

	return mux
}
