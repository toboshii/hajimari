package util

import (
	"context"
	"fmt"
	"math"

	// v1 "k8s.io/api/apps/v1"
	"github.com/ullbergm/hajimari/internal/log"
	networkingV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
)

const (
	appNameLabelKey     = "app.kubernetes.io/name"
	serviceNameLabelKey = "kubernetes.io/service-name"
)

var logger = log.New()

// struct for the ReplicaStatusGetter object
type ReplicaStatusGetter struct {
	err               error
	replicas          int
	availableReplicas int
	kubeClient        kubernetes.Interface
}

// Initializes a ReplicaStatusGetter
func NewReplicaStatusGetter(kubeClient kubernetes.Interface) *ReplicaStatusGetter {
	return &ReplicaStatusGetter{
		kubeClient: kubeClient,
	}
}

// Gets replicaStatuses using the DiscoveryV1 api
func (rsg *ReplicaStatusGetter) GetEndpointStatuses(ingress networkingV1.Ingress) *ReplicaStatusGetter {
	namespace := ingress.ObjectMeta.GetNamespace()
	serviceNames := getServiceNames(ingress)
	var labelOptions metav1.ListOptions

	// Set LabelOptions to get labels for service names that the ingress references
	labelRequirements, err := labels.NewRequirement(serviceNameLabelKey, selection.In, serviceNames)
	if err != nil {
		logger.Error("Error setting labelSelector Requirements", err)
	}
	labelOptions.LabelSelector = labels.NewSelector().Add(*labelRequirements).String()

	epslices, err := rsg.kubeClient.DiscoveryV1().EndpointSlices(namespace).List(context.Background(), labelOptions)

	if err != nil {
		logger.Error("Error Getting EndpointSlices: ", err)
		rsg.err = err
	}

	if len(epslices.Items) > 1 {
		// This scenario can happen if the metrics endpointslices are included in the ingress
		logger.Debug(ingress.Name, " Multiple EndpointSlices found. Will try using all of them.")
	}

	if len(epslices.Items) == 0 {
		// This is indication that labels are mismatched somewhere
		logger.Debug(ingress.Name, " No EndpointSlice Found")
	}

	replicas := 0
	availableReplicas := 0

	for _, epslice := range epslices.Items {
		logger.Debug("Checking EndpointSlice: ", epslice.Name)
		replicas = replicas + len(epslice.Endpoints)
		for _, ep := range epslice.Endpoints {
			if *ep.Conditions.Ready == true {
				availableReplicas = availableReplicas + 1
			}
		}
	}

	if replicas == 0 {
		rsg.err = fmt.Errorf("No endpoints found for %s", ingress.Name)
	} else {
		rsg.replicas = replicas
		rsg.availableReplicas = availableReplicas
	}

	return rsg
}

// Gets the current value of replicas
func (rsg *ReplicaStatusGetter) GetReplicas() int {
	if rsg.err != nil {
		logger.Warn(rsg.err)
		return 0
	}
	return rsg.replicas
}

// Gets the current value of replicas
func (rsg *ReplicaStatusGetter) GetAvailableReplicas() int {
	if rsg.err != nil {
		logger.Warn(rsg.err)
		return 0
	}
	return rsg.availableReplicas
}

// Gets the current ratio of availableReplicas to replicas
// math.Round only works with float64
func (rsg *ReplicaStatusGetter) GetRatio() float64 {
	if rsg.err != nil {
		logger.Warn(rsg.err)
		return 0
	}
	return math.Round(float64(rsg.availableReplicas) / float64(rsg.replicas))
}

// Gets Service Names that the Ingress is actually meant for
func getServiceNames(ingress networkingV1.Ingress) []string {
	serviceNames := []string{}

	if ingress.Spec.DefaultBackend != nil {
		serviceNames = append(serviceNames, ingress.Spec.DefaultBackend.Service.Name)
	}
	if len(ingress.Spec.Rules) > 0 {
		for _, rule := range ingress.Spec.Rules {
			for _, path := range rule.HTTP.Paths {
				serviceNames = append(serviceNames, path.Backend.Service.Name)
			}
		}
	}

	return serviceNames
}
