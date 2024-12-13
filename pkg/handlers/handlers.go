package handlers

import (
	"github.com/naodEthiop/bookings/pkg/config"
	"github.com/naodEthiop/bookings/pkg/models"
	"github.com/naodEthiop/bookings/pkg/render"
	"net/http"
)

// TemplateData holds data sent from handlers to template

// Repo the repository used by the handlers
var Repo *Repository

//repository type

type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		a,
	}
}

// newHandler  sets the repository for the handlers

func NewHandler(r *Repository) {
	Repo = r

}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Developer", "NaodEthiop")
	intm := make(map[string]int)
	intm["year"] = 2032
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		IntMap: intm,
	})

}

// About  is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	//intmap := make(map[string]int)
	//intmap["age"] = 19
	stringMap["test"] = "Hello, again."
	//send the data to the template
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
