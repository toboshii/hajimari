package util

import (
	"context"
	"math"

	// v1 "k8s.io/api/apps/v1"
	networkingV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

// struct for the ReplicaStatusGetter object
type ReplicaStatusGetter struct {
	err                error
	replicas           int32
	availableReplicas  int32
	status             string
	kubeClient kubernetes.Interface
}

// Initializes a ReplicaStatusGetter
func NewReplicaStatusGetter(kubeClient kubernetes.Interface) *ReplicaStatusGetter {
	return &ReplicaStatusGetter{
		kubeClient: kubeClient,
	}
}

// GetDeploymentStatus gets the status conditions of the deployment and sets the status to match the Available status reported
func (rsg *ReplicaStatusGetter) GetDeploymentStatus(ingress networkingV1.Ingress) *ReplicaStatusGetter {

	if rsg.err != nil {
		rsg.err = nil
	}

	// labelMap := ingress.GetLabels()

	deployment, err := rsg.kubeClient.AppsV1().Deployments(ingress.ObjectMeta.GetNamespace()).Get(context.Background(), ingress.ObjectMeta.GetName(), metav1.GetOptions{})

	if err != nil {
		rsg.err = err
	} else {
		// Non-terminated pods targeted by the deployment
		rsg.replicas = deployment.Status.Replicas
		// Using Available instead of Ready as it would pass the minReadySeconds threshold
		rsg.availableReplicas = deployment.Status.AvailableReplicas
	}

	return rsg
}

// GetDeploymentStatus gets the status conditions of the deployment and sets the status to match the Available status reported
func (rsg *ReplicaStatusGetter) GetDaemonSetStatus(ingress networkingV1.Ingress) *ReplicaStatusGetter {

	if rsg.err != nil {
		rsg.err = nil
	}

	daemonSet, err := rsg.kubeClient.AppsV1().DaemonSets(ingress.ObjectMeta.GetNamespace()).Get(context.Background(), ingress.ObjectMeta.GetName(), metav1.GetOptions{})

	if err != nil {
		rsg.err = err
	} else {
		// Non-terminated pods targeted by the deployment
		rsg.replicas = daemonSet.Status.DesiredNumberScheduled
		// Using Available instead of Ready as it would pass the minReadySeconds threshold
		rsg.availableReplicas = daemonSet.Status.NumberAvailable
	}

	return rsg
}
// Gets the current value of replicas
func (rsg *ReplicaStatusGetter) GetReplicas() int32 {
	if rsg.err != nil {
		logger.Warn(rsg.err)
		return 0
	}
	return rsg.replicas
}

// Gets the current value of replicas
func (rsg *ReplicaStatusGetter) GetAvailableReplicas() int32 {
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
	return math.Round(float64(rsg.availableReplicas)/float64(rsg.replicas))
}
