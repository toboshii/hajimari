package util

import (
	"context"

	"github.com/toboshii/hajimari/internal/log"
	networkingV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	logger = log.New()
)

type ServiceStatusGetter struct {
	err        error
	status     string
	kubeClient kubernetes.Interface
}

func NewServiceStatusGetter(kubeClient kubernetes.Interface) *ServiceStatusGetter {
	return &ServiceStatusGetter{
		kubeClient: kubeClient,
	}
}

func (ssg *ServiceStatusGetter) GetServiceStatus(ingress networkingV1.Ingress) *ServiceStatusGetter {
	service, err := ssg.kubeClient.CoreV1().Services(ingress.Namespace).Get(context.Background(), ingress.Name, metav1.GetOptions{})

	if err != nil {
		ssg.err = err
	}

	ssg.status = service.Status.String()

	return ssg
}

func (ssg *ServiceStatusGetter) Get() string {
	if ssg.err != nil {
		logger.Warn(ssg.err)
		return ""
	}

	return ssg.status
}
