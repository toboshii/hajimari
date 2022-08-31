package models

type BookmarkGroup struct {
	Name      string     `json:"name"`
	Bookmarks []Bookmark `json:"bookmarks"`
}
