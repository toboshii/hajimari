package ingressapps

import (
	"math"

	"github.com/ullbergm/hajimari/internal/annotations"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/kube/lists/ingresses"
	"github.com/ullbergm/hajimari/internal/kube/util"
	"github.com/ullbergm/hajimari/internal/kube/wrappers"
	"github.com/ullbergm/hajimari/internal/log"
	"github.com/ullbergm/hajimari/internal/models"
	v1 "k8s.io/api/networking/v1"
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

	al.items = convertIngressesToHajimariApps(ingressList, *util.NewReplicaStatusGetter(al.kubeClient))

	return al
}

// Get function returns the apps currently present in List
func (al *List) Get() ([]models.AppGroup, error) {
	return al.items, al.err
}

func convertIngressesToHajimariApps(ingresses []v1.Ingress, rsg util.ReplicaStatusGetter) (appGroups []models.AppGroup) {
	for _, ingress := range ingresses {
		logger.Debugf("Found ingress with Name '%v' in Namespace '%v'", ingress.Name, ingress.Namespace)
		replicaStatus := rsg.GetEndpointStatuses(ingress)

		wrapper := wrappers.NewIngressWrapper(&ingress)

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
			if wrapper.GetStatusCheckEnabled() && (replicaStatus.GetReplicas() != 0) {
				appGroups[i].Apps = append(appGroups[i].Apps, models.App{
					Name:        wrapper.GetName(),
					Icon:        wrapper.GetAnnotationValue(annotations.HajimariIconAnnotation),
					IconColor:   wrapper.GetIconColor(),
					URL:         wrapper.GetURL(),
					Info:        wrapper.GetInfo(),
					TargetBlank: wrapper.GetTargetBlank(),
					Replicas: models.ReplicaInfo{
						Total:     replicaStatus.GetReplicas(),
						Available: replicaStatus.GetAvailableReplicas(),
						PctReady:  math.Round(replicaStatus.GetRatio() * 100),
					},
					Location: wrapper.GetLocation(),
				})
			} else {
				appGroups[i].Apps = append(appGroups[i].Apps, models.App{
					Name:        wrapper.GetName(),
					Icon:        wrapper.GetAnnotationValue(annotations.HajimariIconAnnotation),
					IconColor:   wrapper.GetIconColor(),
					URL:         wrapper.GetURL(),
					TargetBlank: wrapper.GetTargetBlank(),
					Info:        wrapper.GetInfo(),
				})
			}
		}

	}
	return
}
