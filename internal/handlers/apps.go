package handlers

import (
	"cmp"
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/hajimari/customapps"
	"github.com/ullbergm/hajimari/internal/models"
	"github.com/ullbergm/hajimari/internal/services"
	"github.com/ullbergm/hajimari/internal/util/strings"
	utilStrings "github.com/ullbergm/hajimari/internal/util/strings"
)

type appResource struct {
	service services.AppService
}

func NewAppResource(service services.AppService) *appResource {
	return &appResource{service: service}
}

func (rs *appResource) AppRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", rs.ListApps)

	return router
}

func (rs *appResource) ListApps(w http.ResponseWriter, r *http.Request) {
	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for hajimari: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Collect Kubernetes apps

	cachedKubeApps := rs.service.GetCachedKubeApps()
	logger.Debug("CachedKubeApps: ", cachedKubeApps)

	var kubeApps = make([]models.AppGroup, len(cachedKubeApps))

	copy(kubeApps, cachedKubeApps)

	// Collect Custom apps

	customAppsList := customapps.NewList(*appConfig)

	customApps, err := customAppsList.Populate().Get()
	if err != nil {
		logger.Error("An error occured while populating custom hajimari apps: ", err)
		render.Render(w, r, ErrServerError(err))
	}

	// Merge apps together

	var apps []models.AppGroup

	for i, kubeAppGroup := range kubeApps {
		for x, customAppGroup := range customApps {
			if strings.NormalizeString(customAppGroup.Group) == strings.NormalizeString(kubeAppGroup.Group) {
				kubeApps[i].Apps = append(kubeApps[i].Apps, customAppGroup.Apps...)
				customApps = append(customApps[:x], customApps[x+1:]...)
			}
		}

		// Sort by Location then by Name
		slices.SortFunc(kubeApps[i].Apps, func(a, b models.App) int {
			logger.Debug("Comparing: ", a.Name, " (", a.Location, ") and ", b.Name, " (", b.Location, ")")
			return cmp.Or(
				cmp.Compare(a.Location, b.Location),
				cmp.Compare(utilStrings.NormalizeString(a.Name), utilStrings.NormalizeString(b.Name)),
			)
		})

	}

	apps = append(kubeApps, customApps...)

	// Sort App Groups alphabetically
	slices.SortFunc(apps, func(a, b models.AppGroup) int {
		return utilStrings.CompareNormalized(a.Group, b.Group)
	})

	if err := render.RenderList(w, r, NewAppListResponse(apps)); err != nil {
		render.Render(w, r, ErrServerError(err))
		return
	}

}

type AppResponse struct {
	models.AppGroup
}

func NewAppResponse(appGroup models.AppGroup) *AppResponse {
	resp := &AppResponse{AppGroup: appGroup}

	return resp
}

func (rd *AppResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewAppListResponse(appGroups []models.AppGroup) []render.Renderer {
	list := []render.Renderer{}
	for _, app := range appGroups {
		list = append(list, NewAppResponse(app))
	}
	return list
}
