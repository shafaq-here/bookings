package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shafaq-here/bookings/internal/config"
	"github.com/shafaq-here/bookings/internal/forms"
	"github.com/shafaq-here/bookings/internal/models"
	render "github.com/shafaq-here/bookings/internal/render"
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

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "make-reservations.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)

	//form.Has("first_name", r)
	form.Required("first_name", "last_name", "email")
	//form.Atleast3Chars("first_name")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservations.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("could not get data from the session")
		m.App.Session.Put(r.Context(), "error", "Cant get reservation from session")
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
