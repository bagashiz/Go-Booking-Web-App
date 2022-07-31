package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bagashiz/Go-Booking-Web-App/internal/config"
	"github.com/bagashiz/Go-Booking-Web-App/internal/models"
	"github.com/justinas/nosurf"
)

// functions is a variable that holds the FuncMap for the templates
var functions = template.FuncMap{}

// app is a variable that holds the application configuration
var app *config.AppConfig

// NewTemplates is a function that sets the application configuration to the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultFunctions is a function that adds the default data from TemplateData to the pages
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate is a function that renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	h, ok := tc[tmpl]
	if !ok {
		log.Fatal("Cannot get template: ", tmpl)
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	h.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Printf("Error writing template %v: %v\n", tmpl, err)
	}

	return nil
}

// CreateTemplateCache is a function that creates a template cache to store the templates in memory for faster rendering
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)

		templateSet, err := template.New(pageName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

			myCache[pageName] = templateSet
		}
	}
	return myCache, nil
}
