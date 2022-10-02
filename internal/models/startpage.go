package models

type Startpage struct {
	ID                    string           `json:"id,omitempty"`
	Name                  string           `json:"name"`
	Title                 string           `json:"title"`
	LightTheme            string           `json:"lightTheme"`
	DarkTheme             string           `json:"darkTheme"`
	CustomThemes          []Theme          `json:"customThemes"`
	ShowSearch            *bool            `json:"showSearch"`
	ShowGreeting          *bool            `json:"showGreeting"`
	ShowApps              *bool            `json:"showApps"`
	ShowAppGroups         *bool            `json:"showAppGroups"`
	ShowAppUrls           *bool            `json:"showAppUrls"`
	ShowAppInfo           *bool            `json:"showAppInfo"`
	ShowAppStatus         *bool            `json:"showAppStatus"`
	DefaultAppIcon        string           `json:"defaultAppIcon"`
	ShowBookmarks         *bool            `json:"showBookmarks"`
	ShowGlobalBookmarks   *bool            `json:"showGlobalBookmarks"`
	ShowBookmarkGroups    *bool            `json:"showBookmarkGroups"`
	AlwaysTargetBlank     *bool            `json:"alwaysTargetBlank"`
	DefaultSearchProvider string           `json:"defaultSearchProvider"`
	SearchProviders       []SearchProvider `json:"searchProviders"`
	Bookmarks             []BookmarkGroup  `json:"bookmarks"`
}
