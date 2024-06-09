package crdbookmarks

import (
	"context"

	"github.com/ullbergm/hajimari/internal/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// List struct is used to list ingresses
type List struct {
	appConfig config.Config
	err       error // Used for forwarding errors
	items     []unstructured.Unstructured
	dynClient dynamic.Interface
}

var bookmarkResource = schema.GroupVersionResource{
	Group:    "hajimari.io",
	Version:  "v1alpha1",
	Resource: "bookmarks",
}

// NewList creates an List object that you can use to query ingresses
func NewList(dynClient dynamic.Interface, appConfig config.Config, items ...unstructured.Unstructured) *List {
	return &List{
		dynClient: dynClient,
		appConfig: appConfig,
		items:     items,
	}
}

// Populate function returns a list of ingresses
func (il *List) Populate(namespaces ...string) *List {
	for _, namespace := range namespaces {
		bookmarks, err := il.dynClient.
			Resource(bookmarkResource).
			Namespace(namespace).
			List(context.Background(), metav1.ListOptions{})

		if err != nil {
			il.err = err
		} else {
			il.items = append(il.items, bookmarks.Items...)
		}
	}

	return il
}

// Get function returns the ingresses currently present in List
func (il *List) Get() ([]unstructured.Unstructured, error) {
	return il.items, il.err
}
