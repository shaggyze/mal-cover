package ping

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shaggyze/mal-cover/internal/utils"
)

// Ping contains basic routes.
type Ping struct{}

// New to create new ping and other base routes.
func New() *Ping {
	return &Ping{}
}

// Register to register common routes.
func (p Ping) Register(r chi.Router) {
	r.Get("/", p.handleRoot)
	r.Get("/ping", p.handlePing)
	r.Get("/robots.txt", p.handleRobots)
	r.Get("/favicon.ico", p.handleFavIcon)
	r.NotFound(http.HandlerFunc(p.handleNotFound))
	r.MethodNotAllowed(http.HandlerFunc(p.handleMethodNotAllowed))
}

func (p Ping) handleRoot(w http.ResponseWriter, _ *http.Request) {
	utils.RespondWithCSS(w, http.StatusOK, "it's working\n\nfor more info: https://github.com/shaggyze/mal-cover", nil)
}

func (p Ping) handlePing(w http.ResponseWriter, _ *http.Request) {
	utils.RespondWithCSS(w, http.StatusOK, "pong", nil)
}

func (p Ping) handleNotFound(w http.ResponseWriter, _ *http.Request) {
	utils.RespondWithCSS(w, http.StatusNotFound, "page not found\njust like your future", nil)
}

func (p Ping) handleMethodNotAllowed(w http.ResponseWriter, _ *http.Request) {
	utils.RespondWithCSS(w, http.StatusMethodNotAllowed, "wrong method", nil)
}

func (p Ping) handleFavIcon(w http.ResponseWriter, _ *http.Request) {
	utils.RespondWithCSS(w, http.StatusOK, "ok", nil)
}

func (p Ping) handleRobots(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("User-agent: *\nDisallow: /"))
}
