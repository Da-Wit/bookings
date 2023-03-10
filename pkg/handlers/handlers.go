package handlers

import (
	"net/http"

	"github.com/da-wit/bookings/pkg/config"
	"github.com/da-wit/bookings/pkg/models"
	"github.com/da-wit/bookings/pkg/render"
)

// Repo Repository variable used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo returns a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetRepo sets Repo as r
func SetRepo(r *Repository) {
	Repo = r
}

// Home is the handler function of / page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler function of /about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	var StringMap = map[string]string{}
	StringMap["random_index"] = "The sun will come out."
	StringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: StringMap,
	})
}
