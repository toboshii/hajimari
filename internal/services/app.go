package services

import (
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/hajimari/crdapps"
	"github.com/ullbergm/hajimari/internal/hajimari/ingressapps"
	"github.com/ullbergm/hajimari/internal/kube"
	"github.com/ullbergm/hajimari/internal/kube/util"
	"github.com/ullbergm/hajimari/internal/log"
	"github.com/ullbergm/hajimari/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	logger       = log.New()
	mutex        sync.RWMutex
	kubeAppCache []models.AppGroup
)

type AppService interface {
	GetCachedKubeApps() []models.AppGroup
}

type appService struct {
	logger *logrus.Logger
}

func NewAppService(logger *logrus.Logger) *appService {
	// todo: make time configurable
	ticker := time.NewTicker(60 * time.Second)
	updaterChan := make(chan struct{})
	mutex = sync.RWMutex{}

	kubeAppCache = getKubeApps()
	go startKubeAppCacheUpdater(ticker, updaterChan)

	return &appService{logger: logger}
}

func (as *appService) GetCachedKubeApps() []models.AppGroup {
	return kubeAppCache
}

func getKubeApps() []models.AppGroup {
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

	// Collect Ingress apps

	ingressAppsList := ingressapps.NewList(kubeClient, *appConfig)

	ingressApps, err := ingressAppsList.Populate(namespaces...).Get()
	if err != nil {
		logger.Error("An error occurred while looking for hajimari Ingress apps", err)
		return nil
	}

	// Collect Custom Resource apps

	crdAppsList := crdapps.NewList(dynClient, *appConfig)

	crdApps, err := crdAppsList.Populate(namespaces...).Get()
	if err != nil {
		logger.Error("An error occurred while looking for hajimari Custom Resource apps", err)
		return nil
	}

	// Merge together

	ingressApps = append(ingressApps, crdApps...)

	return ingressApps
}

func startKubeAppCacheUpdater(ticker *time.Ticker, updaterChan chan struct{}) {
	logger.Info("KubeApp cache daemon started")
	for {
		select {
		case <-ticker.C:
			// update cache
			mutex.Lock() // lock the cache before writing to it
			kubeAppCache = getKubeApps()
			mutex.Unlock() // unlock the cache after writing to it
		case <-updaterChan:
			// stop the daemon
			return
		}
	}
}
