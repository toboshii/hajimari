package ingressapps

import (
	"github.com/toboshii/hajimari/pkg/annotations"
	"github.com/toboshii/hajimari/pkg/config"
	"github.com/toboshii/hajimari/pkg/hajimari"
	"github.com/toboshii/hajimari/pkg/kube/lists/ingresses"
	"github.com/toboshii/hajimari/pkg/kube/wrappers"
	"github.com/toboshii/hajimari/pkg/log"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/kubernetes"
)

var (
	logger = log.New()
)

// List struct is used for listing hajimari apps
type List struct {
	appConfig  config.Config
	err        error // Used for forwarding errors
	items      []hajimari.App
	kubeClient kubernetes.Interface
}

// NewList func creates a new instance of apps lister
func NewList(kubeClient kubernetes.Interface, appConfig config.Config) *List {
	return &List{
		appConfig:  appConfig,
		kubeClient: kubeClient,
	}
}

// Populate function that populates a list of hajimari apps from ingresses in selected namespaces
func (al *List) Populate(namespaces ...string) *List {
	ingressList, err := ingresses.NewList(al.kubeClient, al.appConfig).
		Populate(namespaces...).
		Filter(byHajimariEnableAnnotation).Get()

	// Apply Instance filter
	if len(al.appConfig.InstanceName) != 0 {
		ingressList, err = ingresses.NewList(al.kubeClient, al.appConfig, ingressList...).
			Filter(byHajimariInstanceAnnotation).Get()
	}

	if err != nil {
		al.err = err
	}

	al.items = convertIngressesToHajimariApps(ingressList)

	return al
}

// Get function returns the apps currently present in List
func (al *List) Get() ([]hajimari.App, error) {
	return al.items, al.err
}

func convertIngressesToHajimariApps(ingresses []v1beta1.Ingress) (apps []hajimari.App) {
	for _, ingress := range ingresses {
		logger.Infof("Found ingress with Name '%v' in Namespace '%v'", ingress.Name, ingress.Namespace)

		wrapper := wrappers.NewIngressWrapper(&ingress)
		apps = append(apps, hajimari.App{
			Name:  wrapper.GetName(),
			Group: wrapper.GetGroup(),
			Icon:  wrapper.GetAnnotationValue(annotations.HajimariIconAnnotation),
			URL:   wrapper.GetURL(),
		})
	}
	return
}
