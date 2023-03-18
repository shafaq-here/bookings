package handlers

import (
	"fmt"
	"net/http"

	"github.com/shafaq-here/bookings/pkg/config"
	"github.com/shafaq-here/bookings/pkg/models"
	render "github.com/shafaq-here/bookings/pkg/render"
)

// Type repository to interact with the data
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}

}

var Repo *Repository

func NewHandlers(r *Repository) {
	Repo = r
}

// this will handle request for localhost:8080/
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	//wrie the logic
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})

}

// this will handle request for localhost:8080/about
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	myMap := make(map[string]string)
	myMap["test"] = "Hello there."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	myMap["remote_ip"] = remoteIP

	//m.App.Session.GetInt()
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: myMap,
	})
}

func (m *Repository) SectionsHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "sections.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start Date entered is %s and End date entered is %s", start, end)))
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservations.page.tmpl", &models.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
