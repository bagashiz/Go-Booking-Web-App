package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bagashiz/Go-Booking-Web-App/internal/config"
	"github.com/bagashiz/Go-Booking-Web-App/internal/handlers"
	"github.com/bagashiz/Go-Booking-Web-App/internal/helpers"
	"github.com/bagashiz/Go-Booking-Web-App/internal/models"
	"github.com/bagashiz/Go-Booking-Web-App/internal/render"
)

// app is a variable that holds the application configuration from config.go
var app config.AppConfig

// session is a variable that holds the session used by the handlers
var session *scs.SessionManager

// infoLog is a variable that holds the info log for the application
var infoLog *log.Logger

// errorLog is a variable that holds the error log for the application
var errorLog *log.Logger

// portNumber is a constant that holds the port number for the application locally
const portNumber = ":8080"

// main is the main application function
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Starting server on http://localhost%v/\n", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}

func run() error {
	// Things to store in the session
	gob.Register(models.Reservation{})

	//* change this to true when in production
	app.InProduction = false

	// set infoLog to output to stdout
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	// set errorLog to output to stderr
	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// session configuration
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache: ", err)
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return nil
}
