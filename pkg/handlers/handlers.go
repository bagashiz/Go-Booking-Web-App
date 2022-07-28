package handlers

import (
	"net/http"

	"github.com/bagashiz/Go-Booking-Web-App/pkg/config"
	"github.com/bagashiz/Go-Booking-Web-App/pkg/models"
	"github.com/bagashiz/Go-Booking-Web-App/pkg/render"
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

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Executive is the executive page handler function
func (m *Repository) Executive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "executive.page.tmpl", &models.TemplateData{})
}

// Deluxe is the deluxe page handler function
func (m *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "deluxe.page.tmpl", &models.TemplateData{})
}

// Reservation is the reservation page handler function
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.page.tmpl", &models.TemplateData{})
}

// Contact is the contact page handler function
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// Contact is the contact page handler function
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}
