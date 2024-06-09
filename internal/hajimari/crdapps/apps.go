package crdapps

import (
	"github.com/mitchellh/mapstructure"
	"github.com/ullbergm/hajimari/internal/annotations"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/kube/lists/crdapps"
	"github.com/ullbergm/hajimari/internal/kube/types/v1alpha1"
	"github.com/ullbergm/hajimari/internal/kube/wrappers"
	"github.com/ullbergm/hajimari/internal/log"
	"github.com/ullbergm/hajimari/internal/models"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
)

var (
	logger = log.New()
)

// List struct is used for listing hajimari apps
type List struct {
	appConfig config.Config
	err       error // Used for forwarding errors
	items     []models.AppGroup
	dynClient dynamic.Interface
}

// NewList func creates a new instance of apps lister
func NewList(dynClient dynamic.Interface, appConfig config.Config) *List {
	return &List{
		appConfig: appConfig,
		dynClient: dynClient,
	}
}

// Populate function that populates a list of hajimari apps from ingresses in selected namespaces
func (al *List) Populate(namespaces ...string) *List {
	appsList, err := crdapps.NewList(al.dynClient, al.appConfig).
		Populate(namespaces...).
		Get()

	// Apply Instance filter
	if len(al.appConfig.InstanceName) != 0 {
		appsList, err = crdapps.NewList(al.dynClient, al.appConfig, appsList...).
			Get()
	}

	if err != nil {
		al.err = err
	}

	al.items = appsToHajimariApps(appsList)

	return al
}

// Get function returns the apps currently present in List
func (al *List) Get() ([]models.AppGroup, error) {
	return al.items, al.err
}

func appsToHajimariApps(apps []unstructured.Unstructured) (appGroups []models.AppGroup) {
	for _, app := range apps {
		name := app.GetName()
		namespace := app.GetNamespace()

		logger.Debugf("Found apps with Name '%v' in Namespace '%v'", name, namespace)

		appObj := v1alpha1.Application{}
		err := mapstructure.Decode(app.UnstructuredContent(), &appObj)
		if err != nil {
			logger.Error("Could not unmarshall object: %s/", name, namespace)
		}

		wrapper := wrappers.NewAppWrapper(&appObj)

		groupMap := make(map[string]int, len(appGroups))
		for i, v := range appGroups {
			groupMap[v.Group] = i
		}

		if _, ok := groupMap[wrapper.GetGroup()]; !ok {
			appGroups = append(appGroups, models.AppGroup{
				Group: wrapper.GetGroup(),
			})
		}

		appMap := make(map[string]int, len(appGroups))
		for i, v := range appGroups {
			appMap[v.Group] = i
		}

		if i, ok := appMap[wrapper.GetGroup()]; ok {
			location := 1000
			if wrapper.GetLocation() != 0 {
				location = wrapper.GetLocation()
			}

			appGroups[i].Apps = append(appGroups[i].Apps, models.App{
				Name:        wrapper.GetName(),
				Icon:        wrapper.GetAnnotationValue(annotations.HajimariIconAnnotation),
				IconColor:   wrapper.GetIconColor(),
				URL:         wrapper.GetURL(),
				TargetBlank: wrapper.GetTargetBlank(),
				Info:        wrapper.GetInfo(),
				Location:    location,
			})
		}

	}
	return
}
