package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bagashiz/Go-Booking-Web-App/internal/config"
	"github.com/bagashiz/Go-Booking-Web-App/internal/forms"
	"github.com/bagashiz/Go-Booking-Web-App/internal/helpers"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Executive is the executive page handler function
func (m *Repository) Executive(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "executive.page.tmpl", &models.TemplateData{})
}

// Contact is the contact page handler function
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
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
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Reservation is the make-reservation page handler function
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation is the POST request of make-reservation page handler function
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// ReservationSummary is the reservation-summary page handler function
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Cannot get error from session")
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from session.")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
