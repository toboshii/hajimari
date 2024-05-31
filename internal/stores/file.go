package stores

import (
	"fmt"
	"os"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ullbergm/hajimari/internal/models"
	"gopkg.in/yaml.v3"
)

const filePath = "/data/%s.yaml"

type fileStore struct {
	StartpageStore
}

func NewFileStore() StartpageStore {
	return &fileStore{}
}

func (s *fileStore) NewStartpage(startpage *models.Startpage) (string, error) {
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}

	startpage.ID = id

	if data, err := yaml.Marshal(&startpage); err != nil {
		return "", err
	} else if err := os.WriteFile(fmt.Sprintf(filePath, id), data, 0600); err != nil {
		return "", err
	}

	return startpage.ID, nil
}

func (s *fileStore) GetStartpage(id string) (*models.Startpage, error) {
	file, err := os.ReadFile(fmt.Sprintf(filePath, id))
	if err != nil {
		return nil, fmt.Errorf("startpage %s doesn't exist", id)
	}
	startpage := models.Startpage{}

	err2 := yaml.Unmarshal(file, &startpage)
	if err2 != nil {
		return nil, err2
	}

	return &startpage, nil
}

func (s *fileStore) UpdateStartpage(id string, startpage *models.Startpage) (*models.Startpage, error) {
	_, err := os.ReadFile(fmt.Sprintf(filePath, id))
	if err != nil {
		return nil, fmt.Errorf("startpage %s doesn't exist", id)
	}

	if data, err := yaml.Marshal(&startpage); err != nil {
		return nil, err
	} else if err := os.WriteFile(fmt.Sprintf(filePath, id), data, 0600); err != nil {
		return nil, err
	}

	return startpage, nil
}

func (s *fileStore) RemoveStartpage(id string) (*models.Startpage, error) {
	if _, err := os.Stat(fmt.Sprintf(filePath, id)); err != nil {
		return nil, fmt.Errorf("startpage %s doesn't exist", id)
	}

	if err := os.Remove(fmt.Sprintf(filePath, id)); err != nil {
		return nil, err
	}

	return nil, nil
}
