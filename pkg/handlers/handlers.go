package handlers

import (
	"net/http"

	"github.com/xvbnm48/go-bookings/pkg/config"
	"github.com/xvbnm48/go-bookings/pkg/models"
	"github.com/xvbnm48/go-bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	// fmt.Fprintf(w, "Hello, World!")
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello aruno!"
	remoteIP := m.App.Session.Get(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP.(string)

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
