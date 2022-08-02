package models

import "github.com/bagashiz/Go-Booking-Web-App/internal/forms"

// TemplateData is a struct that holds data sent from the handlers
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
