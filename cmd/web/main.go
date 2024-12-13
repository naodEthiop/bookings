package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/naodEthiop/bookings/pkg/config"
	"github.com/naodEthiop/bookings/pkg/handlers"
	"github.com/naodEthiop/bookings/pkg/render"
	"log"
	"net/http"
	"time"
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
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application click: http://127.0.0.1%s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app), //  using route module to handle requests

	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
