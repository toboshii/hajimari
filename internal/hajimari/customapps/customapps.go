package customapps

import (
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/models"
	utilStrings "github.com/ullbergm/hajimari/internal/util/strings"
)

// List struct is used for listing hajimari apps
type List struct {
	appConfig config.Config
	err       error // Used for forwarding errors
	items     []models.AppGroup
}

// NewList func creates a new instance of apps lister
func NewList(appConfig config.Config) *List {
	return &List{
		appConfig: appConfig,
	}
}

// Populate function that populates a list of custom apps
func (al *List) Populate() *List {

	var customApps []models.AppGroup

	for _, group := range al.appConfig.CustomApps {
		group.Group = utilStrings.NormalizeString(group.Group)
		customApps = append(customApps, group)
	}

	al.items = customApps

	return al
}

// Get function returns the apps currently present in List
func (al *List) Get() ([]models.AppGroup, error) {
	return al.items, al.err
}
