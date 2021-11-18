package ingresses

import (
	"context"

	"github.com/toboshii/hajimari/internal/config"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// List struct is used to list ingresses
type List struct {
	appConfig  config.Config
	err        error // Used for forwarding errors
	items      []v1.Ingress
	kubeClient kubernetes.Interface
}

// FilterFunc defined for creating functions that comply with the filtering ingresses
type FilterFunc func(v1.Ingress, config.Config) bool

// NewList creates an List object that you can use to query ingresses
func NewList(kubeClient kubernetes.Interface, appConfig config.Config, items ...v1.Ingress) *List {
	return &List{
		kubeClient: kubeClient,
		appConfig:  appConfig,
		items:      items,
	}
}

// Populate function returns a list of ingresses
func (il *List) Populate(namespaces ...string) *List {
	for _, namespace := range namespaces {
		ingresses, err := il.kubeClient.NetworkingV1().Ingresses(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			il.err = err
		}
		il.items = append(il.items, ingresses.Items...)
	}

	return il
}

// Filter function applies a filter func that is passed as a parameter to the list of ingresses
func (il *List) Filter(filterFunc FilterFunc) *List {

	var filtered []v1.Ingress

	for _, ingress := range il.items {
		if filterFunc(ingress, il.appConfig) {
			filtered = append(filtered, ingress)
		}
	}

	// Replace original ingresses with filtered
	il.items = filtered
	return il
}

// Get function returns the ingresses currently present in List
func (il *List) Get() ([]v1.Ingress, error) {
	return il.items, il.err
}
