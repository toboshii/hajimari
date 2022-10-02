package services

import (
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/toboshii/hajimari/internal/config"
	"github.com/toboshii/hajimari/internal/hajimari/ingressapps"
	"github.com/toboshii/hajimari/internal/kube"
	"github.com/toboshii/hajimari/internal/kube/util"
	"github.com/toboshii/hajimari/internal/log"
	"github.com/toboshii/hajimari/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	logger          = log.New()
	mutex           sync.RWMutex
	ingressAppCache []models.AppGroup
)

type AppService interface {
	GetCachedIngressApps() []models.AppGroup
}

type appService struct {
	logger *logrus.Logger
}

func NewAppService(logger *logrus.Logger) *appService {
	// todo: make time configurable
	ticker := time.NewTicker(60 * time.Second)
	updaterChan := make(chan struct{})
	mutex = sync.RWMutex{}

	ingressAppCache = getIngressApps()

	go startIngressAppCacheUpdater(ticker, updaterChan)

	return &appService{logger: logger}
}

func (as *appService) GetCachedIngressApps() []models.AppGroup {
	return ingressAppCache
}

func getIngressApps() []models.AppGroup {
	appConfig, err := config.GetConfig()
	if err != nil {
		logger.Error("Failed to read configuration for hajimari: ", err)
		return nil
	}

	kubeClient := kube.GetClient()

	ingressAppsList := ingressapps.NewList(kubeClient, *appConfig)

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

	logger.Debug("Looking for hajimari apps in the following namespaces: ", namespacesString)

	ingressApps, err := ingressAppsList.Populate(namespaces...).Get()
	if err != nil {
		logger.Error("An error occurred while looking for hajimari apps", err)
		return nil
	}

	return ingressApps
}

func startIngressAppCacheUpdater(ticker *time.Ticker, updaterChan chan struct{}) {
	logger.Info("IngressApp cache daemon started")
	for {
		select {
		case <-ticker.C:
			// update cache
			mutex.Lock() // lock the cache before writing to it
			ingressAppCache = getIngressApps()
			mutex.Unlock() // unlock the cache after writing to it
		case <-updaterChan:
			// stop the daemon
			return
		}
	}
}
