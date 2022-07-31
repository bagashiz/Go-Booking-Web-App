package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig is a struct that holds the application configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
	UseCache      bool
	InProduction  bool
}
