package models

import "github.com/shafaq-here/bookings/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Data      map[string]interface{}
	CSRFToken string
	Form      *forms.Form
	Flash     string
	Warning   string
	Error     string
}
