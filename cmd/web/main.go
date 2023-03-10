package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	config "github.com/shafaq-here/bookings/pkg/config"
	handlers "github.com/shafaq-here/bookings/pkg/handlers"
	render "github.com/shafaq-here/bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	// http.HandleFunc("/", handlers.Repo.HomeHandler)
	// http.HandleFunc("/about", handlers.Repo.AboutHandler)
	// http.HandleFunc("/sections", handlers.Repo.SectionsHandler)

	fmt.Printf("Starting server at port %s", portNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	//start the server

	// http.ListenAndServe(portNumber, nil)
	err = server.ListenAndServe()
	log.Fatal(err)

}
