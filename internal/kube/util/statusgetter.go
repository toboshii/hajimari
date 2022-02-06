package util

import (
	"context"

	"github.com/toboshii/hajimari/internal/log"
	v1 "k8s.io/api/apps/v1"
	networkingV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	logger = log.New()
)

type StatusGetter struct {
	err        error
	status     string
	kubeClient kubernetes.Interface
}

func NewStatusGetter(kubeClient kubernetes.Interface) *StatusGetter {
	return &StatusGetter{
		kubeClient: kubeClient,
	}
}

func (ssg *StatusGetter) GetDeploymentStatus(ingress networkingV1.Ingress) *StatusGetter {
	deployment, err := ssg.kubeClient.AppsV1().Deployments(ingress.ObjectMeta.GetNamespace()).Get(context.Background(), ingress.ObjectMeta.GetName(), metav1.GetOptions{})

	if err != nil {
		ssg.err = err
	}

	conditions := deployment.Status.Conditions

	for _, c := range conditions {
		if (c.Type==v1.DeploymentAvailable) {
			ssg.status = string(c.Status)
		}
	}

	return ssg
}

func (ssg *StatusGetter) GetServiceStatus(ingress networkingV1.Ingress) *StatusGetter {
	service, err := ssg.kubeClient.CoreV1().Services(ingress.Namespace).Get(context.Background(), ingress.Name, metav1.GetOptions{})

	if err != nil {
		ssg.err = err
	}

	ssg.status = service.Status.String()

	return ssg
}

func (ssg *StatusGetter) Get() string {
	if ssg.err != nil {
		logger.Warn(ssg.err)
		return ""
	}

	return ssg.status
}
