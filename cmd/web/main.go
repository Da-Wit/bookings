package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/da-wit/bookings/pkg/config"
	"github.com/da-wit/bookings/pkg/handlers"
	"github.com/da-wit/bookings/pkg/render"
)

const portNumber string = ":8080"

// app is a application-wide configuration variable.
// it's properties are accessible throughout the project.
var app config.AppConfig

func main() {
	// change this value to ture when in production
	app.InProduction = false

	session := scs.New()
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

	render.UpdateConfig(&app)

	repo := handlers.NewRepo(&app)
	handlers.SetRepo(repo)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Println("portNumber is set to", portNumber)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
