package util

import (
	"context"
	// "fmt"
	"strings"

	"github.com/toboshii/hajimari/internal/log"
	v1 "k8s.io/api/apps/v1"
	networkingV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Initializes logger
var (
	logger = log.New()
)

// struct for the StatusGetter object
type StatusGetter struct {
	err        error
	status     string
	kubeClient kubernetes.Interface
}

// Initializes a StatusGetter
func NewStatusGetter(kubeClient kubernetes.Interface) *StatusGetter {
	return &StatusGetter{
		kubeClient: kubeClient,
	}
}

// GetDeploymentStatus gets the status conditions of the deployment and sets the status to match the Available status reported
func (ssg *StatusGetter) GetDeploymentStatus(ingress networkingV1.Ingress) *StatusGetter {

	if ssg.err != nil {
		ssg.err = nil
	}

	deployment, err := ssg.kubeClient.AppsV1().Deployments(ingress.ObjectMeta.GetNamespace()).Get(context.Background(), ingress.ObjectMeta.GetName(), metav1.GetOptions{})

	if err != nil {
		ssg.err = err
	} else {

		conditions := deployment.Status.Conditions

		for _, c := range conditions {
			logger.Debug(c.String())
			if (c.Type==v1.DeploymentAvailable) {
				ssg.status = strings.ToLower(string(c.Status))
			}
		}
	}

	return ssg
}

// Gets the current value of status
func (ssg *StatusGetter) Get() string {
	if ssg.err != nil {
		logger.Warn(ssg.err)
		return ""
	}

	return ssg.status
}
