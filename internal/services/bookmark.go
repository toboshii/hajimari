package services

import (
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/hajimari/crdbookmarks"
	"github.com/ullbergm/hajimari/internal/kube"
	"github.com/ullbergm/hajimari/internal/kube/util"
	"github.com/ullbergm/hajimari/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	kubeBookmarkCache []models.BookmarkGroup
)

type BookmarkService interface {
	GetCachedKubeBookmarks() []models.BookmarkGroup
}

type bookmarkService struct {
	logger *logrus.Logger
}

func NewBookmarkService(logger *logrus.Logger) *bookmarkService {
	// todo: make time configurable
	ticker := time.NewTicker(60 * time.Second)
	updaterChan := make(chan struct{})
	mutex = sync.RWMutex{}

	kubeBookmarkCache = getKubeBookmarks()
	go startKubeBookmarkCacheUpdater(ticker, updaterChan)

	return &bookmarkService{logger: logger}
}

func (as *bookmarkService) GetCachedKubeBookmarks() []models.BookmarkGroup {
	return kubeBookmarkCache
}

func getKubeBookmarks() []models.BookmarkGroup {
	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for hajimari: ", err)
		return nil
	}

	kubeClient := kube.GetClient()
	dynClient := kube.GetDynamicClient()

	namespaces, err := util.PopulateNamespaceList(kubeClient, appConfig.NamespaceSelector)
	if err != nil {
		logger.Error("An error occurred while populating namespaces: ", err)
		return nil
	}

	var namespacesString string
	// All Namespaces are selected
	if len(namespaces) == 1 && namespaces[0] == metav1.NamespaceAll {
		namespacesString = "* (All Namespaces)"
	} else {
		namespacesString = strings.Join(namespaces, ", ")
	}

	logger.Debug("Looking for Hajimari objects in the following namespaces: ", namespacesString)

	// Collect Custom Resource bookmarks
	crdBookmarksList := crdbookmarks.NewList(dynClient, *appConfig)

	crdBookmarks, err := crdBookmarksList.Populate(namespaces...).Get()
	if err != nil {
		logger.Error("An error occurred while looking for hajimari Custom Resource bookmarks", err)
		crdBookmarks = make([]models.BookmarkGroup, 0)
	}

	return crdBookmarks
}

func startKubeBookmarkCacheUpdater(ticker *time.Ticker, updaterChan chan struct{}) {
	logger.Info("KubeBookmark cache daemon started")
	for {
		select {
		case <-ticker.C:
			// update cache
			mutex.Lock() // lock the cache before writing to it
			kubeBookmarkCache = getKubeBookmarks()
			mutex.Unlock() // unlock the cache after writing to it
		case <-updaterChan:
			// stop the daemon
			return
		}
	}
}
