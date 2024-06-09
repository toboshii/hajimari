package wrappers

import (
	// "fmt"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"github.com/ullbergm/hajimari/internal/kube/types/v1alpha1"
	// "k8s.io/apimachinery/pkg/api/meta"
)

// BookmarkWrapper struct wraps a kubernetes ingress object
type BookmarkWrapper struct {
	bookmark *v1alpha1.Bookmark
}

// NewBookmarkWrapper func creates an instance of BookmarkWrapper
func NewBookmarkWrapper(bookmark *v1alpha1.Bookmark) *BookmarkWrapper {
	return &BookmarkWrapper{
		bookmark: bookmark,
	}
}

// GetName func extracts name of the bookmark wrapped by the object
func (aw *BookmarkWrapper) GetName() string {
	return aw.bookmark.Spec.Name
}

// GetNamespace func extracts namespace of the bookmark wrapped by the object
func (aw *BookmarkWrapper) GetNamespace() string {
	return aw.bookmark.ObjectMeta.Namespace
}

// GetGroup func extracts group name from the bookmark
func (aw *BookmarkWrapper) GetGroup() string {
	return aw.bookmark.Spec.Group
}

// GetGroup func extracts group name from the bookmark
func (aw *BookmarkWrapper) GetInfo() string {
	return aw.bookmark.Spec.Info
}

// GetGroup func extracts group name from the bookmark
func (aw *BookmarkWrapper) GetAnnotationValue(annotationKey string) string {
	return aw.bookmark.Spec.Icon
}

// GetTargetBlank func extracts open in new window feature gate from the bookmark
// @default false
func (aw *BookmarkWrapper) GetTargetBlank() bool {
	return aw.bookmark.Spec.TargetBlank
}

// GetURL func extracts url of the bookmark wrapped by the object
func (aw *BookmarkWrapper) GetURL() string {
	return aw.bookmark.Spec.URL
}

// GetLocation func extracts sorting location of the bookmark wrapped by the object
func (aw *BookmarkWrapper) GetLocation() int {
	return aw.bookmark.Spec.Location
}
