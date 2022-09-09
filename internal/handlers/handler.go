package handlers

import (
	"net/http"
	"time"

	loggerMiddleware "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/toboshii/hajimari/internal/log"
	"github.com/toboshii/hajimari/internal/services"
	"github.com/toboshii/hajimari/internal/stores"
)

var (
	logger = log.New()
)

func NewHandler() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(loggerMiddleware.Logger("router", logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	startpageService := services.NewStartpageService(stores.NewMemoryStore(), logger)
	appService := services.NewAppService(logger)

	// router.Mount("/config", NewConfigResource().ConfigRoutes())
	router.Mount("/apps", NewAppResource(appService).AppRoutes())
	router.Mount("/bookmarks", NewBookmarkResource().BookmarkRoutes())
	router.Mount("/startpage", NewStartpageResource(startpageService).StartpageRoutes())

	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(404)
	render.Render(w, r, ErrNotFound)
}
