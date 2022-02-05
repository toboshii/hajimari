package services

import (
	"context"

	"github.com/toboshii/hajimari/internal/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// List struct is used to list services
type List struct {
	appConfig  config.Config
	err        error // Used for bubbling errors
	items      []v1.Service
	kubeClient kubernetes.Interface
}

// FilterFunc defined for creating functions that comply with the filtering ingresses
type FilterFunc func(v1.Service, config.Config) bool

// NewList creates an List object that you can use to query services
func NewList(kubeClient kubernetes.Interface, appConfig config.Config, items ...v1.Service) *List {
	return &List{
		kubeClient: kubeClient,
		appConfig:  appConfig,
		items:      items,
	}
}

// Populate function returns a list of servicees
func (il *List) Populate(namespaces ...string) *List {
	for _, namespace := range namespaces {
		services, err := il.kubeClient.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			il.err = err
		}
		il.items = append(il.items, services.Items...)
	}

	return il
}

// Filter function applies a filter func that is passed as a parameter to the list of services
func (il *List) Filter(filterFunc FilterFunc) *List {

	var filtered []v1.Service

	for _, ingress := range il.items {
		if filterFunc(ingress, il.appConfig) {
			filtered = append(filtered, ingress)
		}
	}

	// Replace original services with filtered
	il.items = filtered
	return il
}

// Get function returns the services currently present in List
func (il *List) Get() ([]v1.Service, error) {
	return il.items, il.err
}
