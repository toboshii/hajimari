package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/models"
	"github.com/ullbergm/hajimari/internal/services"
)

type contextKey int

const contextKeyStartpage contextKey = iota

type startpageResource struct {
	service services.StartpageService
}

func NewStartpageResource(service services.StartpageService) *startpageResource {
	return &startpageResource{service: service}
}

func (rs *startpageResource) StartpageRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", rs.GetDefaultStartpage)
	router.Post("/", rs.CreateStartpage)

	router.Route("/{startpageID}", func(r chi.Router) {
		r.Use(rs.StartpageCtx)
		r.Get("/", rs.GetStartpage)
		r.Put("/", rs.UpdateStartpage)
		r.Delete("/", rs.DeleteStartpage)
	})

	return router
}

// StartpageCtx middleware is used to load a Startpage object from
// the URL parameters passed through as the request. In case
// the Startpage could not be found, we stop here and return a 404.
func (sr *startpageResource) StartpageCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var startpage *models.Startpage
		var err error

		if startpageID := chi.URLParam(r, "startpageID"); startpageID != "" {
			startpage, err = sr.service.GetStartpage(startpageID)
		} else {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyStartpage, startpage)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (sr *startpageResource) GetStartpage(w http.ResponseWriter, r *http.Request) {
	startpage := r.Context().Value(contextKeyStartpage).(*models.Startpage)

	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for hajimari: ", err)
		render.Render(w, r, ErrServerError(err))
		return
	}

	sr.service.ConvertConfigToStartpage(appConfig, startpage)

	if err := render.Render(w, r, NewStartpageResponse(startpage)); err != nil {
		render.Render(w, r, ErrServerError(err))
		return
	}
}

func (sr *startpageResource) GetDefaultStartpage(w http.ResponseWriter, r *http.Request) {
	startpage := models.Startpage{}

	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for hajimari: ", err)
		render.Render(w, r, ErrServerError(err))
		return
	}

	sr.service.ConvertConfigToStartpage(appConfig, &startpage)

	startpage.Bookmarks = appConfig.GlobalBookmarks

	if err := render.Render(w, r, NewStartpageResponse(&startpage)); err != nil {
		render.Render(w, r, ErrServerError(err))
		return
	}
}

func (sr *startpageResource) CreateStartpage(w http.ResponseWriter, r *http.Request) {
	data := &StartpageRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	startpage := data.Startpage

	if _, err := sr.service.NewStartpage(startpage); err != nil {
		logger.Error("Error writing startpage data: ", err)
		render.Render(w, r, ErrServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewStartpageResponse(startpage))
}

func (sr *startpageResource) UpdateStartpage(w http.ResponseWriter, r *http.Request) {
	startpage := r.Context().Value(contextKeyStartpage).(*models.Startpage)

	data := &StartpageRequest{Startpage: startpage}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	startpage = data.Startpage

	if _, err := sr.service.UpdateStartpage(startpage.ID, startpage); err != nil {
		logger.Error("Error writing startpage data: ", err)
		render.Render(w, r, ErrServerError(err))
		return
	}

	render.Render(w, r, NewStartpageResponse(startpage))
}

func (sr *startpageResource) DeleteStartpage(w http.ResponseWriter, r *http.Request) {
	var err error

	startpage := r.Context().Value(contextKeyStartpage).(*models.Startpage)

	startpage, err = sr.service.RemoveStartpage(startpage.ID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewStartpageResponse(startpage))
}

type StartpageRequest struct {
	*models.Startpage

	ProtectedID string `json:"id"`
}

func (s *StartpageRequest) Bind(r *http.Request) error {
	if s.Startpage == nil {
		return errors.New("missing required Startpage fields")
	}

	return nil
}

type StartpageResponse struct {
	*models.Startpage
}

func NewStartpageResponse(startpage *models.Startpage) *StartpageResponse {
	resp := &StartpageResponse{Startpage: startpage}

	return resp
}

func (rd *StartpageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
