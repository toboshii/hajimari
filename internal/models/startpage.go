package models

type Startpage struct {
	ID                  string          `json:"id,omitempty"`
	Name                string          `json:"name"`
	Title               string          `json:"title"`
	Theme               string          `json:"theme"`
	LightTheme          string          `json:"lightTheme"`
	DarkTheme           string          `json:"darkTheme"`
	ShowGreeting        *bool           `json:"showGreeting"`
	ShowAppGroups       *bool           `json:"showAppGroups"`
	ShowAppUrls         *bool           `json:"showAppUrls"`
	ShowBookmarkGroups  *bool           `json:"showBookmarkGroups"`
	ShowGlobalBookmarks *bool           `json:"showGlobalBookmarks"`
	Bookmarks           []BookmarkGroup `json:"bookmarks"`
}
