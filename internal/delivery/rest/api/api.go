package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rl404/fairy/log"
	"github.com/rl404/mal-cover/internal/service"
	"github.com/rl404/mal-cover/internal/utils"
)

// API contains all functions for api endpoints.
type API struct {
	service service.Service
}

// New to create new api endpoints.
func New(service service.Service) *API {
	return &API{
		service: service,
	}
}

// Register to register api routes.
func (api *API) Register(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Use(log.MiddlewareWithLog(utils.GetLogger(), log.MiddlewareConfig{
			QueryParam: true,
			Error:      true,
		}))

		r.Get("/{user}/{type}", api.handleGetCover)
	})
}

func (api *API) handleGetCover(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user")
	mainType := chi.URLParam(r, "type")
	style := r.URL.Query().Get("style")

	css, code, err := api.service.GenerateCover(r.Context(), service.GenerateCoverRequest{
		Username: username,
		Type:     mainType,
		Style:    style,
	})

	utils.RespondWithCSS(w, code, css, err)
}