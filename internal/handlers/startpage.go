package handlers

import (
	"context"
	"errors"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/toboshii/hajimari/internal/config"
	"github.com/toboshii/hajimari/internal/hajimari"
	"github.com/toboshii/hajimari/internal/hajimari/customapps"
	"github.com/toboshii/hajimari/internal/hajimari/ingressapps"
	"github.com/toboshii/hajimari/internal/kube"
	"github.com/toboshii/hajimari/internal/kube/util"
	"github.com/toboshii/hajimari/internal/models"
	"github.com/toboshii/hajimari/internal/services"
	"github.com/toboshii/hajimari/internal/util/tplutil"
)

type startpageResource struct {
	service services.StartpageService
	tpl     *template.Template
}

func NewStartpageResource(service services.StartpageService, tpl *template.Template) *startpageResource {
	return &startpageResource{service: service, tpl: tpl}
}

func (rs *startpageResource) StartpageRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", rs.GetStartpage)
	router.Get("/config", rs.GetDefaultConfig)
	router.Post("/", rs.CreateStartpage)

	router.Route("/{startpageID}", func(r chi.Router) {
		r.Use(rs.StartpageCtx)
		r.Get("/", rs.GetStartpage)
		r.Get("/config", rs.GetStartpageConfig)
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
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "startpage", startpage)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (sr *startpageResource) GetStartpage(w http.ResponseWriter, r *http.Request) {
	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for hajimari", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Context().Value("startpage") != nil {
		startpage := r.Context().Value("startpage").(*models.Startpage)

		sr.service.ConvertStartpageToConfig(appConfig, startpage)
	}

	kubeClient := kube.GetClient()

	var hajimariApps []hajimari.App

	ingressAppsList := ingressapps.NewList(kubeClient, *appConfig)

	namespaces, err := util.PopulateNamespaceList(kubeClient, appConfig.NamespaceSelector)

	if err != nil {
		logger.Error("An error occurred while populating namespaces", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Debug("Namespaces to look for hajimari apps: ", namespaces)
	hajimariApps, err = ingressAppsList.Populate(namespaces...).Get()

	if err != nil {
		logger.Error("An error occurred while looking for hajimari apps", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customAppsList := customapps.NewList(*appConfig)

	customHajimariApps, err := customAppsList.Populate().Get()
	if err != nil {
		logger.Error("An error occured while populating custom hajimari apps", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Append both generated and custom apps
	hajimariApps = append(hajimariApps, customHajimariApps...)

	w.Header().Add("Content-Type", "text/html")

	err = sr.tpl.Execute(w, struct {
		Title     string
		Greeting  string
		Date      string
		Apps      []hajimari.App
		Groups    []config.Group
		Providers []config.Provider
		Modules   []config.Module
	}{
		Title:     appConfig.Title,
		Greeting:  tplutil.Greet(appConfig.Name, time.Now().Hour()),
		Date:      time.Now().Format("Mon, Jan 02"),
		Apps:      hajimariApps,
		Groups:    appConfig.Groups,
		Providers: appConfig.Providers,
		Modules:   appConfig.Modules,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (sr *startpageResource) GetDefaultConfig(w http.ResponseWriter, r *http.Request) {
	startpage := models.Startpage{}

	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for hajimari", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sr.service.ConvertConfigToStartpage(appConfig, &startpage)

	if err := render.Render(w, r, NewStartpageResponse(&startpage)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (sr *startpageResource) GetStartpageConfig(w http.ResponseWriter, r *http.Request) {
	startpage := r.Context().Value("startpage").(*models.Startpage)

	if err := render.Render(w, r, NewStartpageResponse(startpage)); err != nil {
		render.Render(w, r, ErrRender(err))
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
	_, _ = sr.service.NewStartpage(startpage)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewStartpageResponse(startpage))
}

func (sr *startpageResource) UpdateStartpage(w http.ResponseWriter, r *http.Request) {
	startpage := r.Context().Value("startpage").(*models.Startpage)

	data := &StartpageRequest{Startpage: startpage}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	startpage = data.Startpage
	_, _ = sr.service.UpdateStartpage(startpage.ID, startpage)

	render.Render(w, r, NewStartpageResponse(startpage))
}

func (sr *startpageResource) DeleteStartpage(w http.ResponseWriter, r *http.Request) {
	var err error

	startpage := r.Context().Value("startpage").(*models.Startpage)

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

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
