package wrappers

import (
	// "fmt"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"github.com/ullbergm/hajimari/internal/kube/types/v1alpha1"
	// "k8s.io/apimachinery/pkg/api/meta"
)

// AppWrapper struct wraps a kubernetes ingress object
type AppWrapper struct {
	app *v1alpha1.Application
}

// NewAppWrapper func creates an instance of AppWrapper
func NewAppWrapper(app *v1alpha1.Application) *AppWrapper {
	return &AppWrapper{
		app: app,
	}
}

// GetName func extracts name of the app wrapped by the object
func (aw *AppWrapper) GetName() string {
	return aw.app.Spec.Name
}

// GetNamespace func extracts namespace of the app wrapped by the object
func (aw *AppWrapper) GetNamespace() string {
	return aw.app.ObjectMeta.Namespace
}

// GetGroup func extracts group name from the app
func (aw *AppWrapper) GetGroup() string {
	return aw.app.Spec.Group
}

// GetGroup func extracts group name from the app
func (aw *AppWrapper) GetInfo() string {
	return aw.app.Spec.Info
}

// GetGroup func extracts group name from the app
func (aw *AppWrapper) GetAnnotationValue(annotationKey string) string {
	return aw.app.Spec.Icon
}

// GetTargetBlank func extracts open in new window feature gate from the app
// @default false
func (aw *AppWrapper) GetTargetBlank() bool {
	return aw.app.Spec.TargetBlank
}

// GetURL func extracts url of the app wrapped by the object
func (aw *AppWrapper) GetURL() string {
	return aw.app.Spec.URL
}

// GetLocation func extracts sorting location of the app wrapped by the object
func (aw *AppWrapper) GetLocation() int {
	return aw.app.Spec.Location
}
