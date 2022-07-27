package main

import (
	"net/http"

	"github.com/bagashiz/Go-Booking-Web-App/pkg/config"
	"github.com/bagashiz/Go-Booking-Web-App/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes is a function that creates a new router and registers the handlers
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/executive", handlers.Repo.Executive)
	mux.Get("/deluxe", handlers.Repo.Deluxe)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
