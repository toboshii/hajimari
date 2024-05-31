package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/models"
)

type bookmarkResource struct{}

func NewBookmarkResource() *bookmarkResource {
	return &bookmarkResource{}
}

func (rs *bookmarkResource) BookmarkRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", rs.ListBookmarks)

	return router
}

func (rs *bookmarkResource) ListBookmarks(w http.ResponseWriter, r *http.Request) {
	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for hajimari", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	globalBookmarks := appConfig.GlobalBookmarks

	if err := render.RenderList(w, r, NewBookmarkListResponse(globalBookmarks)); err != nil {
		render.Render(w, r, ErrServerError(err))
		return
	}

}

type BookmarkResponse struct {
	models.BookmarkGroup
}

func NewBookmarkResponse(bookmarkGroup models.BookmarkGroup) *BookmarkResponse {
	resp := &BookmarkResponse{BookmarkGroup: bookmarkGroup}

	return resp
}

func (rd *BookmarkResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewBookmarkListResponse(bookmarkGroups []models.BookmarkGroup) []render.Renderer {
	list := []render.Renderer{}
	for _, bookmark := range bookmarkGroups {
		list = append(list, NewBookmarkResponse(bookmark))
	}
	return list
}
