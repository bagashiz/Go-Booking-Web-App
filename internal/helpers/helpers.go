package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/bagashiz/Go-Booking-Web-App/internal/config"
)

var app *config.AppConfig

// NewHelpers is a function that sets the config for the helpers package
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError is a function that sends a client error response to the user
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of ", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError is a function that sends a server error response to the user
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
