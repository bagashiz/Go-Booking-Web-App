package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

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
