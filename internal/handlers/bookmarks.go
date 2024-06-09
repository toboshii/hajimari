package handlers

import (
	"cmp"
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/models"
	"github.com/ullbergm/hajimari/internal/services"
	"github.com/ullbergm/hajimari/internal/util/strings"
	utilStrings "github.com/ullbergm/hajimari/internal/util/strings"
)

type bookmarkResource struct {
	service services.BookmarkService
}

func NewBookmarkResource(service services.BookmarkService) *bookmarkResource {
	return &bookmarkResource{service: service}
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

	// Collect kubernetes bookmarks
	cachedKubeBookmarks := rs.service.GetCachedKubeBookmarks()
	logger.Debug("CachedKubeBookmarks: ", cachedKubeBookmarks)

	var kubeBookmarks = make([]models.BookmarkGroup, len(cachedKubeBookmarks))

	copy(kubeBookmarks, cachedKubeBookmarks)

	globalBookmarks := appConfig.GlobalBookmarks

	for i, bookmarkGroup := range globalBookmarks {
		for j, bookmark := range bookmarkGroup.Bookmarks {
			bookmark.Location = 0
			if bookmark.Location == 0 {
				bookmark.Location = 1000
			}
			bookmarkGroup.Bookmarks[j] = bookmark
		}
		globalBookmarks[i] = bookmarkGroup
	}

	// Merge bookmarks together
	var bookmarks []models.BookmarkGroup

	for i, kubeBookmarkGroup := range kubeBookmarks {
		for x, customBookmarkGroup := range globalBookmarks {
			if strings.NormalizeString(customBookmarkGroup.Group) == strings.NormalizeString(kubeBookmarkGroup.Group) {
				kubeBookmarks[i].Bookmarks = append(kubeBookmarks[i].Bookmarks, customBookmarkGroup.Bookmarks...)
				globalBookmarks = append(globalBookmarks[:x], globalBookmarks[x+1:]...)
			}
		}

		// Sort by Location then by Name
		slices.SortFunc(kubeBookmarks[i].Bookmarks, func(a, b models.Bookmark) int {
			logger.Debug("Comparing: ", a.Name, " (", a.Location, ") and ", b.Name, " (", b.Location, ")")
			return cmp.Or(
				cmp.Compare(a.Location, b.Location),
				cmp.Compare(utilStrings.NormalizeString(a.Name), utilStrings.NormalizeString(b.Name)),
			)
		})

	}

	bookmarks = append(kubeBookmarks, globalBookmarks...)

	// Sort Bookmark Groups alphabetically
	slices.SortFunc(bookmarks, func(a, b models.BookmarkGroup) int {
		return utilStrings.CompareNormalized(a.Group, b.Group)
	})

	logger.Debug("Bookmarks: ", bookmarks)

	if err := render.RenderList(w, r, NewBookmarkListResponse(bookmarks)); err != nil {
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
