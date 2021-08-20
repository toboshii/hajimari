package services

import (
	"github.com/sirupsen/logrus"
	"github.com/toboshii/hajimari/internal/config"
	"github.com/toboshii/hajimari/internal/models"
	"github.com/toboshii/hajimari/internal/stores"
)

type StartpageService interface {
	NewStartpage(startpage *models.Startpage) (string, error)
	GetStartpage(id string) (*models.Startpage, error)
	UpdateStartpage(id string, startpage *models.Startpage) (*models.Startpage, error)
	RemoveStartpage(id string) (*models.Startpage, error)
	ConvertStartpageToConfig(appConfig *config.Config, startpage *models.Startpage)
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
	log := s.logger.WithFields(logrus.Fields{
		"startpageId": id,
	})
	log.Info("GetStartpage")
	return s.store.GetStartpage(id)
}

func (s *startpageService) UpdateStartpage(id string, startpage *models.Startpage) (*models.Startpage, error) {
	return s.store.UpdateStartpage(id, startpage)
}

func (s *startpageService) RemoveStartpage(id string) (*models.Startpage, error) {
	return s.store.RemoveStartpage(id)
}

func (s *startpageService) ConvertStartpageToConfig(appConfig *config.Config, startpage *models.Startpage) {
	appConfig.Name = startpage.Name
	appConfig.Groups = []config.Group{}

	for _, g := range startpage.Groups {
		links := []config.Link{}

		for _, l := range g.Links {
			links = append(links, config.Link(l))
		}

		appConfig.Groups = append(appConfig.Groups, config.Group{
			Name:  g.Name,
			Links: links,
		})
	}
}

func (s *startpageService) ConvertConfigToStartpage(appConfig *config.Config, startpage *models.Startpage) {
	startpage.Name = appConfig.Name
	startpage.Groups = []models.Group{}

	for _, g := range appConfig.Groups {
		links := []models.Link{}

		for _, l := range g.Links {
			links = append(links, models.Link(l))
		}

		startpage.Groups = append(startpage.Groups, models.Group{
			Name:  g.Name,
			Links: links,
		})
	}
}
