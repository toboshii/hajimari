package ingressapps

import (
	"github.com/toboshii/hajimari/internal/annotations"
	"github.com/toboshii/hajimari/internal/config"
	"github.com/toboshii/hajimari/internal/kube/lists/ingresses"
	"github.com/toboshii/hajimari/internal/kube/util"
	"github.com/toboshii/hajimari/internal/kube/wrappers"
	"github.com/toboshii/hajimari/internal/log"
<<<<<<< HEAD
	v1 "k8s.io/api/networking/v1"
=======
	"github.com/toboshii/hajimari/internal/models"
	"k8s.io/api/extensions/v1beta1"
>>>>>>> 97467966a4db33a45732f6f5d967b7ec5cb7a754
	"k8s.io/client-go/kubernetes"
)

var (
	logger = log.New()
)

// List struct is used for listing hajimari apps
type List struct {
	appConfig  config.Config
	err        error // Used for forwarding errors
	items      []models.AppGroup
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

	al.items = convertIngressesToHajimariApps(ingressList, *util.NewStatusGetter(al.kubeClient))

	return al
}

// Get function returns the apps currently present in List
func (al *List) Get() ([]models.AppGroup, error) {
	return al.items, al.err
}

<<<<<<< HEAD
func convertIngressesToHajimariApps(ingresses []v1.Ingress, ssg util.StatusGetter) (apps []hajimari.App) {
=======
func convertIngressesToHajimariApps(ingresses []v1beta1.Ingress) (appGroups []models.AppGroup) {
>>>>>>> 97467966a4db33a45732f6f5d967b7ec5cb7a754
	for _, ingress := range ingresses {
		logger.Debugf("Found ingress with Name '%v' in Namespace '%v'", ingress.Name, ingress.Namespace)
		status := ssg.GetDeploymentStatus(ingress).Get()
		var emptyStatus string = "undefined"

		wrapper := wrappers.NewIngressWrapper(&ingress)
<<<<<<< HEAD
		if wrapper.GetStatusCheckEnabled() && len(status)>0 {
			apps = append(apps, hajimari.App{
				Name:   wrapper.GetName(),
				Group:  wrapper.GetGroup(),
				Icon:   wrapper.GetAnnotationValue(annotations.HajimariIconAnnotation),
				URL:    wrapper.GetURL(),
				Status: status,
			})
		} else {
			apps = append(apps, hajimari.App{
				Name:  wrapper.GetName(),
				Group: wrapper.GetGroup(),
				Icon:  wrapper.GetAnnotationValue(annotations.HajimariIconAnnotation),
				URL:   wrapper.GetURL(),
				Status: emptyStatus,
=======

		groupMap := make(map[string]int, len(appGroups))
		for i, v := range appGroups {
			groupMap[v.Name] = i
		}

		if _, ok := groupMap[wrapper.GetGroup()]; !ok {
			appGroups = append(appGroups, models.AppGroup{
				Name: wrapper.GetGroup(),
			})
		}

		appMap := make(map[string]int, len(appGroups))
		for i, v := range appGroups {
			appMap[v.Name] = i
		}

		if i, ok := appMap[wrapper.GetGroup()]; ok {
			appGroups[i].Apps = append(appGroups[i].Apps, models.App{
				Name: wrapper.GetName(),
				Icon: wrapper.GetAnnotationValue(annotations.HajimariIconAnnotation),
				URL:  wrapper.GetURL(),
>>>>>>> 97467966a4db33a45732f6f5d967b7ec5cb7a754
			})
		}

	}
	return
}
