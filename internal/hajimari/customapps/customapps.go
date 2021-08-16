package customapps

import (
	"github.com/toboshii/hajimari/internal/config"
	"github.com/toboshii/hajimari/internal/hajimari"
	"github.com/toboshii/hajimari/internal/log"
)

var (
	logger = log.New()
)

// List struct is used for listing hajimari apps
type List struct {
	appConfig config.Config
	err       error // Used for forwarding errors
	items     []hajimari.App
}

// NewList func creates a new instance of apps lister
func NewList(appConfig config.Config) *List {
	return &List{
		appConfig: appConfig,
	}
}

// Populate function that populates a list of custom apps
func (al *List) Populate() *List {
	al.items = convertCustomAppsToHajimariApps(al.appConfig.CustomApps)

	return al
}

// Get function returns the apps currently present in List
func (al *List) Get() ([]hajimari.App, error) {
	return al.items, al.err
}

func convertCustomAppsToHajimariApps(customApps []config.CustomApp) (apps []hajimari.App) {
	for _, customApp := range customApps {
		logger.Debugf("Found custom app with Name '%v'", customApp.Name)

		apps = append(apps, hajimari.App{
			Name:  customApp.Name,
			URL:   customApp.URL,
			Icon:  customApp.Icon,
			Group: customApp.Group,
		})
	}

	return apps
}
