package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bagashiz/Go-Booking-Web-App/internal/config"
	"github.com/bagashiz/Go-Booking-Web-App/internal/models"
	"github.com/bagashiz/Go-Booking-Web-App/internal/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func getRoutes() http.Handler {
	// Things to store in the session
	gob.Register(models.Reservation{})

	//* change this to true when in production
	app.InProduction = false

	// session configuration
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache: ", err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)
	render.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/executive", Repo.Executive)
	mux.Get("/deluxe", Repo.Deluxe)
	mux.Get("/contact", Repo.Contact)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)

	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

// NoSurf is a function that adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad is a function that loads and save the session from current request to the cookie
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// CreateTestTemplateCache is a function that creates a template cache to store the templates in memory for faster rendering
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)

		templateSet, err := template.New(pageName).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}

			myCache[pageName] = templateSet
		}
	}
	return myCache, nil
}
