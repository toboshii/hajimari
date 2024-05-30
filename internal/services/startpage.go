package services

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/ullbergm/hajimari/internal/config"
	"github.com/ullbergm/hajimari/internal/models"
	"github.com/ullbergm/hajimari/internal/stores"
)

type StartpageService interface {
	NewStartpage(startpage *models.Startpage) (string, error)
	GetStartpage(id string) (*models.Startpage, error)
	UpdateStartpage(id string, startpage *models.Startpage) (*models.Startpage, error)
	RemoveStartpage(id string) (*models.Startpage, error)
	ConvertConfigToStartpage(appConfig *config.Config, startpage *models.Startpage)
}

type startpageService struct {
	store  stores.StartpageStore
	logger *logrus.Logger
}

func NewStartpageService(store stores.StartpageStore, logger *logrus.Logger) *startpageService {
	return &startpageService{store: store, logger: logger}
}

func (s *startpageService) NewStartpage(startpage *models.Startpage) (string, error) {
	return s.store.NewStartpage(startpage)
}

func (s *startpageService) GetStartpage(id string) (*models.Startpage, error) {
	startpage, err := s.store.GetStartpage(id)
	if err != nil {
		return startpage, err
	}

	if len(startpage.Bookmarks) == 0 {
		startpage.Bookmarks = []models.BookmarkGroup{}
	}

	if len(startpage.Groups) > 0 {
		logger.Warnf("Startpage %s contains deprecated option `groups`, please convert it to `bookmarks`", startpage.ID)
	}

	return startpage, err
}

func (s *startpageService) UpdateStartpage(id string, startpage *models.Startpage) (*models.Startpage, error) {
	return s.store.UpdateStartpage(id, startpage)
}

func (s *startpageService) RemoveStartpage(id string) (*models.Startpage, error) {
	return s.store.RemoveStartpage(id)
}

func (s *startpageService) ConvertConfigToStartpage(appConfig *config.Config, startpage *models.Startpage) {
	var configInterface map[string]interface{}
	var startpageInterface map[string]interface{}

	configJson, _ := json.Marshal(appConfig)
	json.Unmarshal(configJson, &configInterface)

	startpageJson, _ := json.Marshal(startpage)
	json.Unmarshal(startpageJson, &startpageInterface)

	for k, v := range startpageInterface {
		switch v := v.(type) {
		case string:
			if v == "" && startpageInterface[k] != configInterface[k] {
				startpageInterface[k] = configInterface[k]
			}
		case int:
			if v == 0 && startpageInterface[k] != configInterface[k] {
				startpageInterface[k] = configInterface[k]
			}
		case nil:
			startpageInterface[k] = configInterface[k]
		}
	}

	mergedJson, _ := json.Marshal(startpageInterface)
	json.Unmarshal(mergedJson, &startpage)
}
