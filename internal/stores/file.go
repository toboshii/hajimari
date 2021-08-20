package stores

import (
	"errors"
	"io/ioutil"
	"os"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/toboshii/hajimari/internal/models"
	"gopkg.in/yaml.v3"
)

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

	data, err2 := yaml.Marshal(&startpage)
	if err != nil {
		return "", err2
	}

	err3 := ioutil.WriteFile("/data/"+id+".yaml", data, 0600)
	if err3 != nil {
		return "", err
	}

	return startpage.ID, nil
}

func (s *fileStore) GetStartpage(id string) (*models.Startpage, error) {
	file, err := ioutil.ReadFile("/data/" + id + ".yaml")
	if err != nil {
		return nil, errors.New("startpage not found")
	}
	startpage := models.Startpage{}

	err2 := yaml.Unmarshal(file, &startpage)
	if err2 != nil {
		return nil, err2
	}

	return &startpage, nil
}

func (s *fileStore) UpdateStartpage(id string, startpage *models.Startpage) (*models.Startpage, error) {
	_, err := ioutil.ReadFile("/data/" + id + ".yaml")
	if err != nil {
		return nil, errors.New("startpage not found")
	}

	data, err2 := yaml.Marshal(&startpage)
	if err != nil {
		return nil, err2
	}

	err3 := ioutil.WriteFile("/data/"+id+".yaml", data, 0600)
	if err3 != nil {
		return nil, err
	}

	return startpage, nil
}

func (s *fileStore) RemoveStartpage(id string) (*models.Startpage, error) {
	file, err := ioutil.ReadFile("/data/" + id + ".yaml")
	if err != nil {
		return nil, errors.New("startpage not found")
	}
	startpage := models.Startpage{}

	err2 := yaml.Unmarshal(file, &startpage)
	if err2 != nil {
		return nil, err2
	}

	err3 := os.Remove("/data/" + id + ".yaml")
	if err3 != nil {
		return nil, err3
	}

	return &startpage, nil
}
