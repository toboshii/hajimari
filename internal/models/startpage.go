package models

type Startpage struct {
	ID                    string           `json:"id,omitempty"`
	Name                  string           `json:"name"`
	Title                 string           `json:"title"`
	LightTheme            string           `json:"lightTheme"`
	DarkTheme             string           `json:"darkTheme"`
	CustomThemes          []Theme          `json:"customThemes"`
	ShowGreeting          *bool            `json:"showGreeting"`
	ShowAppGroups         *bool            `json:"showAppGroups"`
	ShowAppUrls           *bool            `json:"showAppUrls"`
	ShowAppInfo           *bool            `json:"showAppInfo"`
	ShowAppStatus         *bool            `json:"showAppStatus"`
	ShowBookmarkGroups    *bool            `json:"showBookmarkGroups"`
	ShowGlobalBookmarks   *bool            `json:"showGlobalBookmarks"`
	DefaultSearchProvider string           `json:"defaultSearchProvider"`
	SearchProviders       []SearchProvider `json:"searchProviders"`
	Bookmarks             []BookmarkGroup  `json:"bookmarks"`
}
