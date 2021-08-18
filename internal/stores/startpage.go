package stores

import (
	"github.com/toboshii/hajimari/internal/models"
)

type StartpageStore interface {
	NewStartpage(startpage *models.Startpage) (string, error)
	GetStartpage(id string) (*models.Startpage, error)
	UpdateStartpage(id string, startpage *models.Startpage) (*models.Startpage, error)
	RemoveStartpage(id string) (*models.Startpage, error)
}
