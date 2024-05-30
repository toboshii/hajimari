package stores

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ullbergm/hajimari/internal/models"
	"github.com/ullbergm/hajimari/internal/util/pointer"
)

var startpages = []*models.Startpage{
	{
		ID:                  "108ZTGB77K09P4QJfu4vG",
		Name:                "Adam",
		ShowAppGroups:       pointer.Of(true),
		ShowGlobalBookmarks: pointer.Of(true),
		ShowAppUrls:         pointer.Of(true),
		ShowAppInfo:         pointer.Of(true),
		AlwaysTargetBlank:   pointer.Of(false),
		CustomThemes: []models.Theme{
			{
				Name:            "danger",
				BackgroundColor: "#0c0e0c",
				PrimaryColor:    "#eaebea",
				AccentColor:     "#d8323c",
			},
		},
		Bookmarks: []models.BookmarkGroup{
			{
				Group: "Media",
				Bookmarks: []models.Bookmark{
					{
						Name: "Youtube",
						URL:  "https://youtube.com",
					},
					{
						Name: "Netflix",
						URL:  "https://netflix.com",
					},
				},
			},
			{
				Group: "Code",
				Bookmarks: []models.Bookmark{
					{
						Name: "Github",
						URL:  "https://github.com",
					},
				},
			},
		},
	},
	{
		ID:   "r1NqSpS1C0z3cMHyzX-Y5",
		Name: "Bob",
		Groups: []map[string]interface{}{
			{
				"name": "communicate",
				"links": []map[string]interface{}{
					{
						"name": "Discord",
						"url":  "https://discord.im",
					},
				},
			},
		},
	},
	{ID: "NmcvYTdRozUDNpbjAxFTO", Name: "Carol"},
	{ID: "3OG1X_GuRS_TrCcWb5YjK", Name: "Derrick"},
	{ID: "Qq4EF6MaSKaYZpEldtLTn", Name: "Emily"},
}

type memoryStore struct {
	StartpageStore
}

func NewMemoryStore() StartpageStore {
	return &memoryStore{}
}

func (s *memoryStore) NewStartpage(startpage *models.Startpage) (string, error) {
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	startpage.ID = id
	startpages = append(startpages, startpage)
	return startpage.ID, nil
}

func (s *memoryStore) GetStartpage(id string) (*models.Startpage, error) {
	for _, s := range startpages {
		if s.ID == id {
			return s, nil
		}
	}
	return nil, errors.New("startpage not found")
}

func (s *memoryStore) UpdateStartpage(id string, startpage *models.Startpage) (*models.Startpage, error) {
	for i, s := range startpages {
		if s.ID == id {
			startpages[i] = startpage
			return startpage, nil
		}
	}
	return nil, errors.New("startpage not found")
}

func (s *memoryStore) RemoveStartpage(id string) (*models.Startpage, error) {
	for i, s := range startpages {
		if s.ID == id {
			startpages = append((startpages)[:i], (startpages)[i+1:]...)
			return s, nil
		}
	}
	return nil, errors.New("startpage not found")
}
