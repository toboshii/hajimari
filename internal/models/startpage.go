package models

type Startpage struct {
	ID                    string           `json:"id,omitempty"`
	Name                  string           `json:"name"`
	Title                 string           `json:"title"`
	LightTheme            string           `json:"lightTheme"`
	DarkTheme             string           `json:"darkTheme"`
	CustomThemes          []Theme          `json:"customThemes"`
	TargetBlank           *bool            `json:"targetBlank"`
	ShowGreeting          *bool            `json:"showGreeting"`
	ShowApps              *bool            `json:"showApps"`
	ShowAppGroups         *bool            `json:"showAppGroups"`
	ShowAppUrls           *bool            `json:"showAppUrls"`
	ShowAppInfo           *bool            `json:"showAppInfo"`
	ShowAppStatus         *bool            `json:"showAppStatus"`
	ShowBookmarks         *bool            `json:"showBookmarks"`
	ShowGlobalBookmarks   *bool            `json:"showGlobalBookmarks"`
	ShowBookmarkGroups    *bool            `json:"showBookmarkGroups"`
	DefaultSearchProvider string           `json:"defaultSearchProvider"`
	SearchProviders       []SearchProvider `json:"searchProviders"`
	Bookmarks             []BookmarkGroup  `json:"bookmarks"`
}
