package models

type BookmarkGroup struct {
	Group     string     `json:"group"`
	Bookmarks []Bookmark `json:"bookmarks"`
}
