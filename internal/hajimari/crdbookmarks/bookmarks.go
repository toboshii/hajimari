package crdbookmarks

import (
	"github.com/mitchellh/mapstructure"
	"github.com/ullbergm/hajimari/internal/annotations"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/kube/lists/crdbookmarks"
	"github.com/ullbergm/hajimari/internal/kube/types/v1alpha1"
	"github.com/ullbergm/hajimari/internal/kube/wrappers"
	"github.com/ullbergm/hajimari/internal/log"
	"github.com/ullbergm/hajimari/internal/models"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
)

var (
	logger = log.New()
)

// List struct is used for listing hajimari bookmarks
type List struct {
	appConfig config.Config
	err       error // Used for forwarding errors
	items     []models.BookmarkGroup
	dynClient dynamic.Interface
}

// NewList func creates a new instance of bookmarks lister
func NewList(dynClient dynamic.Interface, appConfig config.Config) *List {
	return &List{
		appConfig: appConfig,
		dynClient: dynClient,
	}
}

// Populate function that populates a list of hajimari bookmarks from ingresses in selected namespaces
func (al *List) Populate(namespaces ...string) *List {
	bookmarksList, err := crdbookmarks.NewList(al.dynClient, al.appConfig).
		Populate(namespaces...).
		Get()

	// Apply Instance filter
	if len(al.appConfig.InstanceName) != 0 {
		bookmarksList, err = crdbookmarks.NewList(al.dynClient, al.appConfig, bookmarksList...).
			Get()
	}

	if err != nil {
		al.err = err
	}

	al.items = bookmarksToHajimariBookmarks(bookmarksList)

	return al
}

// Get function returns the bookmarks currently present in List
func (al *List) Get() ([]models.BookmarkGroup, error) {
	return al.items, al.err
}

func bookmarksToHajimariBookmarks(bookmarks []unstructured.Unstructured) (bookmarkGroups []models.BookmarkGroup) {
	for _, bookmark := range bookmarks {
		name := bookmark.GetName()
		namespace := bookmark.GetNamespace()

		logger.Debugf("Found bookmarks with Name '%v' in Namespace '%v'", name, namespace)

		bookmarkObj := v1alpha1.Bookmark{}
		err := mapstructure.Decode(bookmark.UnstructuredContent(), &bookmarkObj)
		if err != nil {
			logger.Error("Could not unmarshall object: %s/", name, namespace)
		}

		wrapper := wrappers.NewBookmarkWrapper(&bookmarkObj)

		groupMap := make(map[string]int, len(bookmarkGroups))
		for i, v := range bookmarkGroups {
			groupMap[v.Group] = i
		}

		if _, ok := groupMap[wrapper.GetGroup()]; !ok {
			bookmarkGroups = append(bookmarkGroups, models.BookmarkGroup{
				Group: wrapper.GetGroup(),
			})
		}

		bookmarkMap := make(map[string]int, len(bookmarkGroups))
		for i, v := range bookmarkGroups {
			bookmarkMap[v.Group] = i
		}

		if i, ok := bookmarkMap[wrapper.GetGroup()]; ok {
			location := 1000
			if wrapper.GetLocation() != 0 {
				location = wrapper.GetLocation()
			}

			bookmarkGroups[i].Bookmarks = append(bookmarkGroups[i].Bookmarks, models.Bookmark{
				Name:        wrapper.GetName(),
				Icon:        wrapper.GetAnnotationValue(annotations.HajimariIconAnnotation),
				URL:         wrapper.GetURL(),
				TargetBlank: wrapper.GetTargetBlank(),
				Location:    location,
			})
		}

	}
	return
}
