package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bagashiz/Go-Booking-Web-App/internal/config"
	"github.com/bagashiz/Go-Booking-Web-App/internal/forms"
	"github.com/bagashiz/Go-Booking-Web-App/internal/models"
	"github.com/bagashiz/Go-Booking-Web-App/internal/render"
)

// Repo is a variable that holds the repository used by the handlers
var Repo *Repository

// Repository is a struct that holds the application configuration
type Repository struct {
	App *config.AppConfig
}

// NewRepo is a function that creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers is a function that sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler function
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Executive is the executive page handler function
func (m *Repository) Executive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "executive.page.tmpl", &models.TemplateData{})
}

// Deluxe is the deluxe page handler function
func (m *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "deluxe.page.tmpl", &models.TemplateData{})
}

// Availability is the search for available room page handler function
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability is the search for available room page handler function
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	checkInDate := r.Form.Get("checkInDate")
	checkOutDate := r.Form.Get("checkOutDate")

	w.Write([]byte(fmt.Sprintf("Check-In Date: %s\nCheck-Out Date: %s\n", checkInDate, checkOutDate)))
}

// jsonResponse is a struct to store JSON
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON is handler function to handle request for availibility room and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact is the contact page handler function
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is the make-reservation page handler function
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation is the POST request of make-reservation page handler function
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
}
