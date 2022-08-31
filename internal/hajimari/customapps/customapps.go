package customapps

import (
	"strings"

	"github.com/toboshii/hajimari/internal/config"
	"github.com/toboshii/hajimari/internal/models"
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

	for _, v := range al.appConfig.CustomApps {
		v.Group = strings.ToLower(v.Group)
		customApps = append(customApps, v)
	}

	al.items = customApps

	return al
}

// Get function returns the apps currently present in List
func (al *List) Get() ([]models.AppGroup, error) {
	return al.items, al.err
}
